package routes

import (

	"github.com/NoIdeaCoder/001/handlers/utilities"
	"github.com/gofiber/fiber/v2"
)

func HandleRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", nil)
	})
	app.Get("/explore", func(c *fiber.Ctx) error {
		return c.Render("explore", nil)
	})
	app.Get("/navigate/:foldername?", func(c *fiber.Ctx) error {
		return utilities.NavigateDirectories(c, c.Params("foldername"))
	})
	app.Get("/back", func(c *fiber.Ctx) error {
		return utilities.NavigateBackwards(c)
	})
	app.Get("/download/:filename", func(c *fiber.Ctx) error {
		return utilities.DownloadFile(c, c.Params("filename"))
	})

}
