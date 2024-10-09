package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// @title BurmaGuru
// @version 1.0
// @description This is the Prototype server for the BurmaGuru
// @host localhost:8080
// @BasePath /api

func Setup(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/auth/register", limiter.New(), Register)
	//api.Get("/auth/verify-email/:token", VerifyEmail)
	//api.Post("/auth/login", Login)

	api.Get("/users", limiter.New(), GetAllUsers)
	api.Get("/users/:id", GetUser)
	api.Put("/users/:email", UpdateUser)
	api.Delete("/users/:id", DeleteUser)

	//api.Post("/articles", CreateArticle)
	//api.Get("/articles/:id", GetArticle)
	//api.Put("/articles/:id", UpdateArticle)
	//api.Delete("/articles/:id", DeleteArticle)
	//api.Get("/articles", GetArticles)
	//
	//api.Post("/categories", CreateCategory)
	//api.Get("/categories/:id", GetCategory)
	//api.Put("/categories/:id", UpdateCategory)
	//api.Delete("/categories/:id", DeleteCategory)
	//api.Get("/categories", GetCategories)

	api.Post("/articles/:id/comments", CreateComment)
	api.Get("/articles/:id/comments", GetComments)
	api.Delete("/comments/:id", DeleteComment)

}
