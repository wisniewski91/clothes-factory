package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/wisniewski91/clothes-factory/internal/clients"
	jobtypes "github.com/wisniewski91/clothes-factory/internal/jobTypes"
	"github.com/wisniewski91/clothes-factory/internal/parts"
	"github.com/wisniewski91/clothes-factory/internal/storage"
)

var (
	engine *html.Engine
)

func StartServer() {
	storage.InitDb()

	clients.Init()
	jobtypes.Init()
	parts.Init()

	engine = html.New("./views/admin", ".html")

	app := fiber.New(fiber.Config{
		ServerHeader: "Gooo!",
		AppName:      "Clothes factory",
		Views:        engine,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))

	clients.SetupRoutes(app)
	jobtypes.SetupRoutes(app)
	parts.SetupRoutes(app)

	app.Static("/static", "./static/admin")
	app.Listen("localhost:4000")
}
