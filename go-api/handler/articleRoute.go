package routes

//import (
//	"strconv"
//
//	"github.com/gofiber/fiber/v2"
//	"github.com/swanhtetaungphyo/burmaguru/models"
//	"github.com/swanhtetaungphyo/burmaguru/services"
//	"github.com/swanhtetaungphyo/burmaguru/utils"
//)
//
//func CreateArticle(c *fiber.Ctx) error {
//	var article models.Article
//
//	if err := c.BodyParser(&article); err != nil {
//		return utils.ErrorResponse(c, "Invalid input", fiber.StatusBadRequest, err.Error())
//	}
//
//	createdArticle, err := services.CreateArticleService(article)
//	if err != nil {
//		return utils.ErrorResponse(c, "Failed to create article", fiber.StatusInternalServerError, err.Error())
//	}
//
//	return utils.SuccessResponse(c, "Article created successfully", fiber.StatusCreated, createdArticle)
//}
//
//func GetArticle(c *fiber.Ctx) error {
//	requestParam := c.Params("id")
//	articleId, err := strconv.Atoi(requestParam)
//	if err != nil {
//		return utils.ErrorResponse(c, "Invalid parameter", fiber.StatusBadRequest, err.Error())
//	}
//
//	article, err := services.GetArticleByIdService(articleId)
//	if err != nil {
//		return utils.ErrorResponse(c, "Article not found", fiber.StatusNotFound, err.Error())
//	}
//
//	return utils.SuccessResponse(c, "Article retrieval successful", fiber.StatusOK, article)
//}
//
//func UpdateArticle(c *fiber.Ctx) error {
//	var article models.Article
//	requestParam := c.Params("id")
//	articleId, err := strconv.Atoi(requestParam)
//	if err != nil {
//		return utils.ErrorResponse(c, "Invalid parameter", fiber.StatusBadRequest, err.Error())
//	}
//
//	if err := c.BodyParser(&article); err != nil {
//		return utils.ErrorResponse(c, "Invalid input", fiber.StatusBadRequest, err.Error())
//	}
//
//	article.ID = articleId
//
//	if err := services.UpdateArticleService(article); err != nil {
//		return utils.ErrorResponse(c, "Failed to update article", fiber.StatusInternalServerError, err.Error())
//	}
//
//	return utils.SuccessResponse(c, "Article updated successfully", fiber.StatusOK, nil)
//}
//
//func DeleteArticle(c *fiber.Ctx) error {
//	requestParam := c.Params("id")
//	articleId, err := strconv.Atoi(requestParam)
//	if err != nil {
//		return utils.ErrorResponse(c, "Invalid parameter", fiber.StatusBadRequest, err.Error())
//	}
//
//	if err := services.DeleteArticleService(articleId); err != nil {
//		return utils.ErrorResponse(c, "Failed to delete article", fiber.StatusInternalServerError, err.Error())
//	}
//
//	return utils.SuccessResponse(c, "Article deleted successfully", fiber.StatusOK, nil)
//}
//
//func GetArticles(c *fiber.Ctx) error {
//	articles, err := services.GetAllArticlesService()
//	if err != nil {
//		return utils.ErrorResponse(c, "Failed to retrieve articles", fiber.StatusInternalServerError, err.Error())
//	}
//
//	return utils.SuccessResponse(c, "Articles retrieval successful", fiber.StatusOK, articles)
//}
