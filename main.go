package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kriangkrai/GoFiber/Controller"
	"github.com/kriangkrai/GoFiber/Router"
)

func main() {

	if err := Controller.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	api := app.Group("/api") // /api

	v1 := api.Group("/v1") // /api/v1
	v1.Get("/get/:device", Router.Get)
	v1.Get("/gets", Router.Gets)
	v1.Post("/insert", Router.Insert)
	v1.Put("/update", Router.Update)
	v1.Delete("/delete/:device", Router.Delete)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Listen(":" + port)
	//log.Fatal(app.Listen(port))
}
