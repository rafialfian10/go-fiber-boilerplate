package middleware

import (
	"strings"
	"time"

	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/pkg/mysql"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := mysql.DB

		// Read entire request body
		reqBody := c.Request().Body()

		logData := models.Log{
			Date:      time.Now(),
			IPAddress: c.IP(),
			Host:      c.Hostname(),
			Path:      c.Path(),
			Method:    c.Method(),
		}

		// Check for sensitive data in request body
		if strings.Contains(string(reqBody), "password") {
			logData.Body = "this data is encrypted, because it contains credentials"
		} else if len(reqBody) > 0 {
			// Split request body if it contains multiple parts (e.g., multipart/form-data)
			body := strings.Split(string(reqBody), "----------------------------")

			var (
				textBody string
				fileBody string
			)

			for _, b := range body {
				if strings.Contains(b, "image") {
					fileBody = "----------------------------" + strings.Split(b, "\r\n\r\n")[0]
				} else {
					if len(b) >= 1 && b[0] == '{' {
						textBody += b
					} else if b != "" {
						textBody += "----------------------------" + b
					}
				}
			}

			logData.Body = textBody
			logData.File = fileBody
		}

		// Proceed with the request
		if err := c.Next(); err != nil {
			return err
		}

		// Log response details
		logData.ResposeTime = time.Since(logData.Date).Seconds()
		logData.StatusCode = c.Response().StatusCode()

		// Persist log data to database
		if err := db.Create(&logData).Error; err != nil {
			return err
		}

		return nil
	}
}
