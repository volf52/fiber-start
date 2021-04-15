package main

import (
	"embed"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/volf52/fiber-start/routing"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed templates
var embeddedFiles embed.FS

func main() {
	fsys, err := fs.Sub(embeddedFiles, "templates")
	if err != nil {
		panic(err)
	}

	engine := html.NewFileSystem(http.FS(fsys), ".gohtml")

	app := fiber.New(fiber.Config{Views: engine})
	app.Use(logger.New())

	routing.SetupRoutes(app)

	port := getEnvDef("PORT", "9003")
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}

func getEnvDef(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
