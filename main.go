package main

import (
	"log"

	"github.com/Sasfort/personal-backend/database"
	"github.com/Sasfort/personal-backend/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the backend side of Syabib's website!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", welcome)

	api := app.Group("/api")

	originApi := api.Group("/origins")
	originApi.Post("/", routes.CreateOrigin)
	originApi.Get("/", routes.ReadAllOrigins)
	originApi.Get("/:id", routes.ReadOrigin)
	originApi.Put("/:id", routes.UpdateOrigin)
	originApi.Delete("/:id", routes.DeleteOrigin)

	charactersApi := api.Group("/characters")
	charactersApi.Post("/", routes.CreateCharacter)
	charactersApi.Get("/", routes.ReadAllCharacters)
	charactersApi.Get("/:id", routes.ReadCharacter)
	charactersApi.Put("/:id", routes.UpdateCharacter)
	charactersApi.Delete("/:id", routes.DeleteCharacter)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	err := app.Listen("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}
