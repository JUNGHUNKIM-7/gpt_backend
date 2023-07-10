package main

import (
	"log"
	"os"

	"github.com/JUNGHUNKIM-7/controller"
	"github.com/JUNGHUNKIM-7/model"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	token := os.Getenv("GPT_TOKEN")
	model.Environment = &model.Env{
		GptToken: token,
	}
}

func main() {
	app := fiber.New()

	chat := app.Group("/chat")
	chat.Post("/", controller.GetCompletion)

	app.Listen(":3000")
}
