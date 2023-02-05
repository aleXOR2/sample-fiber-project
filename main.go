// https://tutorialedge.net/golang/basic-rest-api-go-fiber/
package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func validateParamsV1() {
	log.Print("validateParamsV1 called")
}

func validateParamsV2() {
	log.Print("validateParamsV2 called")
}

func addToDatabase() {
	log.Print("writing to database")
}

func v1(c *fiber.Ctx) {
	c.Send("V1 triggered.")
	validateParamsV1()
	addToDatabase()
}

func v2(c *fiber.Ctx) {
	c.Send("V2 triggered.")
	validateParamsV2()
	addToDatabase()
}

func rootRoute(c *fiber.Ctx) {
	log.Print("redirecting to V1")
	v1(c)
}

func setupRoutes(app *fiber.App) {
	app.Get("/v1", v1)
	app.Get("/v2", v2)
	app.Get("/", rootRoute)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	log.Fatal(app.Listen(3000))
}
