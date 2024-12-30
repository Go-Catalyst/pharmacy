package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pharmacy/internal/categories/models"
	"pharmacy/internal/categories/repository"
)

var categoryRepository = repository.NewCategoryRepository()

// GetCategories godoc
// @Summary Get all categories
// @Description Get all categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Category
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	categories := categoryRepository.GetAllCategories()
	c.JSON(http.StatusOK, categories)
}

// GetCategory godoc
// @Summary Get a category by ID
// @Description Get a category by ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Router /categories/{id} [get]
func GetCategory(c *gin.Context) {
	id := c.Param("id")
	category, err := categoryRepository.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body models.Category true "Category"
// @Success 201 {object} models.Category
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var category models.Category
	log.Print("inja")
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categoryRepository.CreateCategory(&category)
	c.JSON(http.StatusCreated, category)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Param category body models.Category true "Category"
// @Success 200 {object} models.Category
// @Router /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedCategory, err := categoryRepository.UpdateCategory(id, category)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, updatedCategory)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 204
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := categoryRepository.DeleteCategory(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
