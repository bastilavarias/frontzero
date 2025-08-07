// cmd/app/main.go

package main

import (
	"log"

	"github.com/bastilavarias/frontzero/internal/router"
)

func main() {
	app := router.New()

	log.Fatal(app.Listen(":3000"))
}
