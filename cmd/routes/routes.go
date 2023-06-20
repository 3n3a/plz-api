package routes

import (
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"

	// "github.com/gofiber/fiber/v2/middleware/limiter"
	swagger "github.com/arsmn/fiber-swagger/v2"

	"github.com/3n3a/plz-api/cmd/db"
	"github.com/3n3a/plz-api/cmd/handlers"
	_ "github.com/3n3a/plz-api/docs"
)

func New(dbInstance db.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(cache.New())
	app.Use(compress.New())
	// app.Use(limiter.New())

	api := app.Group("/api")
	api.Get("postal-codes", func(c *fiber.Ctx) error {
		return handlers.GetAllPostalCodes(c, &dbInstance)
	})
	

	app.Get("/docs/*", swagger.HandlerDefault)

	return app
}