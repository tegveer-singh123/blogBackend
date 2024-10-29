package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tegveer-singh123/blog/util"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message":"Unauthorized",
		})
	}

	return c.Next()
}
