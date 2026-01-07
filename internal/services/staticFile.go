package services

import "github.com/gofiber/fiber/v2"

func StaticFile(c *fiber.Ctx) error {

	params := c.Params("path")
	return c.SendFile("./public/uploads/courses/" + params)
}
