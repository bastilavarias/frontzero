// internal/router/router.go

package router

import (
	"log"
	_ "os"
	"strings"

	"github.com/bastilavarias/frontzero/internal/config"
	"github.com/gofiber/fiber/v2"
)

// New initializes and returns a new Fiber app with dynamically configured routes.
func New() *fiber.App {
	// 1. Read and parse the manifest file
	manifest, err := config.LoadManifest("frontzero.yaml")
	if err != nil {
		log.Fatalf("failed to load manifest: %v", err)
	}
	app := fiber.New()

	for _, entity := range manifest.Entities {
		// Create a copy of the entity for the closure
		entity := entity
		app.Get("/"+strings.ToLower(entity.Name), func(c *fiber.Ctx) error {
			return c.SendString(entity.Name)
		})
	}

	return app
}
