package dashboard

import (
	"github.com/gofiber/fiber/v2"
)

const viewPath = "modules/dashboard/"

func dashboardHandler(c *fiber.Ctx) error {
	return c.Render(viewPath+"index", nil, "layout/main")
}
