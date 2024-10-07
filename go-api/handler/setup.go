package routes

import (
    "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
        app.Post("/auth/register", Register)        // User Registration
        app.Get("/auth/verify-email/:token", VerifyEmail) // Email Verification
        app.Post("/auth/login", Login)              // User Login

        app.Get("/users",GetAllUsers)
        app.Get("/users/:id", GetUser)             // Get User by ID
        app.Put("/users/:id", UpdateUser)          // Update User by ID
        app.Delete("/users/:id", DeleteUser)       // Delete User by ID

        app.Post("/articles", CreateArticle)        // Create Article
        app.Get("/articles/:id", GetArticle)        // Get Article by ID
        app.Put("/articles/:id", UpdateArticle)     // Update Article by ID
        app.Delete("/articles/:id", DeleteArticle)  // Delete Article by ID
        app.Get("/articles", GetArticles)           // Get All Articles

        app.Post("/categories", CreateCategory)      // Create Category
        app.Get("/categories/:id", GetCategory)      // Get Category by ID
        app.Put("/categories/:id", UpdateCategory)   // Update Category by ID
        app.Delete("/categories/:id", DeleteCategory) // Delete Category by ID
        app.Get("/categories", GetCategories)        // Get All Categories

        app.Post("/articles/:id/comments", CreateComment) // Create Comment
        app.Get("/articles/:id/comments", GetComments)     // Get Comments by Article ID
        app.Delete("/comments/:id", DeleteComment)         // Delete Comment by ID

        app.Post("/guides", CreateGuide)              // Create Guide
        app.Get("/guides/:id", GetGuide)              // Get Guide by ID
        app.Put("/guides/:id", UpdateGuide)           // Update Guide by ID
        app.Delete("/guides/:id", DeleteGuide)        // Delete Guide by ID
}
