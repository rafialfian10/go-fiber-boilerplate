package handlerTransaction

import (
	"crypto/tls"
	"fmt"
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
)

func (h *handlerTransaction) CreateTransaction(c *fiber.Ctx) error {
	var request dto.CreateTransactionRequest

	err := c.BodyParser(&request)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	claims, ok := c.Locals("userData").(jwt.MapClaims)
	if !ok {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User data from jwt payload is not found",
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// Extract user data from JWT claims
	userId, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// Generate a unique transaction ID
	var transactionIsMatch = false
	var transactionId uint
	for !transactionIsMatch {
		transactionId = uint(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.GetTransactionByID(transactionId)
		if transactionData == nil || transactionData.ID == 0 {
			transactionIsMatch = true
		}
	}

	transaction := models.Transaction{
		ID:              transactionId,
		UserID:          userId,
		DisasterID:      request.DisasterID,
		TransactionDate: time.Now(),
		Status:          "pending",
	}

	addedTransaction, err := h.TransactionRepository.CreateTransaction(&transaction)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	newTransaction, err := h.TransactionRepository.GetTransactionByID(addedTransaction.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(int(newTransaction.ID)),
			GrossAmt: int64(newTransaction.Disaster.Donate),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: newTransaction.User.FullName,
			Email: newTransaction.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	updateTokenTransaction, _ := h.TransactionRepository.UpdateTokenTransaction(snapResp.Token, newTransaction.ID)

	transactionUpdated, _ := h.TransactionRepository.GetTransactionByID(updateTokenTransaction.ID)

	response := dto.Result{
		Status:  http.StatusCreated,
		Message: "Transaction successfully paid",
		Data: map[string]interface{}{
			"data":                  convertTransactionResponse(transactionUpdated),
			"midtrans_token":        snapResp.Token,
			"midtrans_redirect_url": snapResp.RedirectURL,
		},
	}
	return c.Status(http.StatusCreated).JSON(response)
}

func (h *handlerTransaction) Notification(c *fiber.Ctx) error {
	var notificationPayload map[string]interface{}

	if err := c.BodyParser(&notificationPayload); err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)
	order_id, _ := strconv.Atoi(orderId)

	transaction, _ := h.TransactionRepository.GetTransactionByID(uint(order_id))

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			transaction.Status = "pending"
			h.TransactionRepository.UpdateTransaction(transaction)
		} else if fraudStatus == "accept" {
			transaction.Status = "success"
			h.TransactionRepository.UpdateTransaction(transaction)
			SendMail("Transaction Success", *transaction)
		}
	} else if transactionStatus == "settlement" {
		transaction.Status = "success"
		h.TransactionRepository.UpdateTransaction(transaction)
		SendMail("Transaction Success", *transaction)
	} else if transactionStatus == "deny" {
		transaction.Status = "failed"
		h.TransactionRepository.UpdateTransaction(transaction)
		SendMail("Transaction Failed", *transaction)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		transaction.Status = "failed"
		h.TransactionRepository.UpdateTransaction(transaction)
		SendMail("Transaction Failed", *transaction)
	} else if transactionStatus == "pending" {
		transaction.Status = "pending"
		SendMail("Transaction Failed", *transaction)
		h.TransactionRepository.UpdateTransaction(transaction)
	}

	response := dto.Result{
		Status:  http.StatusOK,
		Message: "Notification processed successfully",
	}
	return c.Status(http.StatusOK).JSON(response)
}

func SendMail(status string, transaction models.Transaction) {
	var CONFIG_SMTP_HOST = os.Getenv("CONFIG_SMTP_HOST")
	var CONFIG_SMTP_PORT = os.Getenv("CONFIG_SMTP_PORT")
	var CONFIG_SENDER_NAME = os.Getenv("CONFIG_SENDER_NAME")
	var CONFIG_AUTH_EMAIL = os.Getenv("CONFIG_AUTH_EMAIL")
	var CONFIG_AUTH_PASSWORD = os.Getenv("CONFIG_AUTH_PASSWORD")

	smtpPort, err := strconv.Atoi(CONFIG_SMTP_PORT)
	if err != nil {
		log.Fatal(err.Error())
	}

	var title = transaction.Disaster.Title
	var donate = strconv.Itoa(transaction.Disaster.Donate)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", transaction.User.Email)
	mailer.SetHeader("Subject", "Transaction Status")
	mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8" />
				<meta http-equiv="X-UA-Compatible" content="IE=edge" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0" />
				<title>Document</title>
				<style>
					.container {
						text-align: center;
						padding: 30px;
						background-color: rgb(252, 163, 163);
					}
					.image {
						display: flex !important;
						justify-content: center !important;
						align-items: center !important;
					}
					h2 {
						text-align: center;
						font-weight: 800;
					}
					ul {
						display: inline-block;
						width: 25vw;
						margin-bottom: 50px;
						text-align: left;
						list-style-type: none;
						padding: 0px;
					}
					li {
						margin-bottom: 10px;
					}
					p {
						text-align: center;
						font-weight: 800;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<div class="image">
						<img src="https://res.cloudinary.com/dixxnrj9b/image/upload/v1721350769/holyways/holyways-icon_sxcgnz.png" alt="Holyways Icon" width="100px" height="70px" style="margin: auto" />
					</div>
					<h2>Donation Details:</h2>
					<ul>
						<li>Disaster: %s</li>
						<li>Total Donation: Rp.%s</li>
						<li>Status: <b>%s</b></li>
					</ul>
					<p>Thank you for your donation through Holyways!</p>
				</div>
			</body>
		</html>`, title, donate, status))

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		smtpPort,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Mail sent! to " + transaction.User.Email)
}
