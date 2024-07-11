package middleware

import (
	"fmt"
	"go-restapi-boilerplate/dto"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadSingleFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// parse multipart form with max memory size 8 Mb
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println("Request parse error: ", err)
			return err
		}

		// single file
		fileHeaders := form.File["image"]
		if len(fileHeaders) == 0 {
			// set up context value and send it to next handler
			c.Locals("image", "")
			return c.Next()
		}

		file, err := fileHeaders[0].Open()
		defer file.Close()
		if err != nil {
			// handle file open error
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
			return c.Status(http.StatusBadRequest).JSON(response)
		}

		log.Println(fileHeaders[0].Filename)

		// validation format file
		if filepath.Ext(fileHeaders[0].Filename) != ".jpg" && filepath.Ext(fileHeaders[0].Filename) != ".jpeg" && filepath.Ext(fileHeaders[0].Filename) != ".png" {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: "Invalid file type",
			}
			return c.Status(http.StatusBadRequest).JSON(response)
		}

		// generate randomized filename using timestamps that convert to milliseconds
		newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeaders[0].Filename))

		// get active directory
		dir, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}

		// set file location
		fileLocation := filepath.Join(dir, "uploads/photo", newFileName)

		// save file to specific destination
		err = c.SaveFile(fileHeaders[0], fileLocation)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
			return c.Status(http.StatusBadRequest).JSON(response)
		}

		var imgUrl string
		if strings.Contains(c.Hostname(), "localhost") || strings.Contains(c.Hostname(), "127.0.0.1") {
			imgUrl = fmt.Sprintf("http://%s/static/photo/%s", c.Hostname(), newFileName)
		} else {
			imgUrl = fmt.Sprintf("https://%s/static/photo/%s", c.Hostname(), newFileName)
		}

		// set up context value and send it to next handler
		c.Locals("image", imgUrl)
		return c.Next()
	}
}

func UploadMultipleFiles() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var arrImages []string

		// parse multipart form with max memory size 8 Mb
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println("Request parse error: ", err)
			return err
		}

		// parsing files
		fileHeaders := form.File["images"]
		if len(fileHeaders) == 0 {
			// set up context value and send it to next handler
			c.Locals("images", []string{})
			return c.Next()
		}

		for _, fileHeader := range fileHeaders {
			file, err := fileHeader.Open()
			defer file.Close()
			if err != nil {
				response := dto.Result{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
				}
				return c.Status(http.StatusBadRequest).JSON(response)
			}

			log.Println(fileHeader.Filename)

			// validation format file
			if filepath.Ext(fileHeader.Filename) != ".jpg" && filepath.Ext(fileHeader.Filename) != ".jpeg" && filepath.Ext(fileHeader.Filename) != ".png" {
				response := dto.Result{
					Status:  http.StatusBadRequest,
					Message: "Invalid file type",
				}
				return c.Status(http.StatusBadRequest).JSON(response)
			}

			// generate randomized filename using timestamps that convert to milliseconds
			newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))

			// get active directory
			dir, err := os.Getwd()
			if err != nil {
				panic(err.Error())
			}

			// set file location
			fileLocation := filepath.Join(dir, "uploads/photo", newFileName)

			// save file to specific destination
			err = c.SaveFile(fileHeader, fileLocation)
			if err != nil {
				response := dto.Result{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
				}
				return c.Status(http.StatusBadRequest).JSON(response)
			}

			var imgUrl string
			if strings.Contains(c.Hostname(), "localhost") || strings.Contains(c.Hostname(), "127.0.0.1") {
				imgUrl = fmt.Sprintf("http://%s/static/photo/%s", c.Hostname(), newFileName)
			} else {
				imgUrl = fmt.Sprintf("https://%s/static/photo/%s", c.Hostname(), newFileName)
			}

			arrImages = append(arrImages, imgUrl)
		}

		// set up context value and send it to next handler
		c.Locals("images", arrImages)
		return c.Next()
	}
}
