package category_handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pharmacy/internal/categories/models"
	"pharmacy/internal/categories/repository"
)

var categoryRepository = repository.NewCategoryRepository()

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param user body models.Category true "Category"
// @Success 201 {object} models.Category
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categoryRepository.CreateCategory(&category)
	c.JSON(http.StatusCreated, category)
}
