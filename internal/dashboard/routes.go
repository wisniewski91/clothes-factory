package dashboard

import "github.com/gofiber/fiber/v2"

func SetRoutes(app *fiber.App) {

	// WWW ROUTES
	app.Get("/", dashboardHandler)

}
