package controller

import (
	"log"

	"github.com/JUNGHUNKIM-7/gpt"
	"github.com/JUNGHUNKIM-7/model"
	"github.com/gofiber/fiber/v2"
)

func GetCompletion(c *fiber.Ctx) error {
	var (
		body       model.Body
		completion string
	)

	err := c.BodyParser(&body)
	if err != nil {
		log.Fatal(err)
	}

	completion = gpt.GetCompletion(
		model.Roles{
			Role:    "system",
			Content: "you are a helpful assistant",
		},
		model.Roles{
			Role:    "user",
			Content: body.Q,
		},
		body.Config,
	)

	return c.JSON(model.Response{
		A: completion,
	})
}
