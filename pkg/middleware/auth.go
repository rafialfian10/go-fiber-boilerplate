package middleware

import (
	"fmt"
	"go-restapi-boilerplate/dto"
	jwtToken "go-restapi-boilerplate/pkg/jwt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UserAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from the Authorization header
		token := c.Get("Authorization")
		if token == "" {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		// Remove "Bearer " prefix
		token = strings.Replace(token, "Bearer ", "", 1)

		// Log the token
		// fmt.Println("Token received:", token)

		// Validate the token and extract claims
		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		// Log the claims
		// fmt.Println("Claims decoded:", claims)

		// Set the context value and proceed to the next handler
		c.Locals("userData", claims)
		return c.Next()
	}
}

func AdminAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from the Authorization header
		token := c.Get("Authorization")
		if token == "" {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		// Remove "Bearer " prefix
		token = strings.Replace(token, "Bearer ", "", 1)

		// Validate the token and extract claims
		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		// Log the claims
		// fmt.Println("Claims decoded in AdminAuth:", claims)

		// Validate whether the user is an admin
		roleId, err := strconv.Atoi(fmt.Sprintf("%v", claims["roleId"]))
		if err != nil || (roleId != 1 && roleId != 2) {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized, you're not an administrator",
			}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		// Set the context value and proceed to the next handler
		c.Locals("userData", claims)
		return c.Next()
	}
}

func SuperAdminAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from the Authorization header
		token := c.Get("Authorization")
		if token == "" {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		// Remove "Bearer " prefix
		token = strings.Replace(token, "Bearer ", "", 1)

		// Validate the token and extract claims
		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		// Log the claims
		// fmt.Println("Claims decoded in SuperAdminAuth:", claims)

		// Validate whether the user is a super admin
		roleId, err := strconv.Atoi(fmt.Sprintf("%v", claims["roleId"]))
		if err != nil || roleId != 1 {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized, you're not a Super Administrator",
			}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		// Set the context value and proceed to the next handler
		c.Locals("userData", claims)
		return c.Next()
	}
}
