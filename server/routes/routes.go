package routes

import (
	"iharacee/server"
	_ "iharacee/server/routes/api/giris"
	_ "iharacee/server/routes/api/sunucusaati"
	_ "iharacee/server/routes/api/telemetri_gonder"

	/* _ "iharacee/server/routes/session" */

	"github.com/gofiber/fiber/v2"
)

func init() {
	app := server.App

	router := app.Group("/")
	router.Get("/", mainRoute)
}

func mainRoute(c *fiber.Ctx) error {
	return c.SendString("Hello, Neostellar 👋!")
}
