package main

import (
	"github.com/NoIdeaCoder/001/handlers/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)
func main(){
    engine := django.New("web/templates",".html")
    app := fiber.New(fiber.Config{
        Views: engine,
    })
    app.Static("/static","web/static")
    routes.HandleRoutes(app)
    app.Listen(":8080")
    
     
}
