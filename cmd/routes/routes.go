package routes

import (
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"

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

	app.Use(cache.New(cache.Config{
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.OriginalURL()
		},
	}))
	app.Use(compress.New())
	app.Use(cors.New())
	// app.Use(limiter.New())

	api := app.Group("/api")
	api.Get("plz/all", func(c *fiber.Ctx) error {
		return handlers.GetAllPostalCodes(c, &dbInstance)
	})

	api.Get("plz/find", func(c *fiber.Ctx) error {
		return handlers.FindPostalCodes(c, &dbInstance)
	})
	
	api.Get("plz/search", func(c *fiber.Ctx) error {
		return handlers.SearchPostalCodes(c, &dbInstance)
	})

	app.Get("/docs/*", swagger.HandlerDefault)

	return app
}