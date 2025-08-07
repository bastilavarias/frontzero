package main

import (
	_ "fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type Entity struct {
	Name   string  `yaml:"name"`
	Fields []Field `yaml:"fields"`
}

type Manifest struct {
	Entities []Entity `yaml:"entities"`
}

func main() {
	data, err := os.ReadFile("frontzero.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var manifest Manifest
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	for _, entity := range manifest.Entities {
		app.Get(strings.ToLower(entity.Name), func(c *fiber.Ctx) error {
			return c.SendString(entity.Name)
		})
	}
	log.Fatal(app.Listen(":3000"))
}
