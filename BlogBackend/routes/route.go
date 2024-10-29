package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tegveer-singh123/blog/controller"
	"github.com/tegveer-singh123/blog/middleware"
)

func Setup(app *fiber.App) {
    app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticated)

	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPost)
	app.Get("/api/allpost/:id", controller.DetailPost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
	app.Get("/api/uniquepost", controller.UniquePost)
	app.Get("/api/deletepost", controller.DeletePost)



}
