package handlers

import (
	"net/http"
	"pharmacy/internal/drugs/models"
	"pharmacy/internal/drugs/repository"

	"github.com/gin-gonic/gin"
)

type DrugHandler struct {
	Repo *repository.DrugRepository
}

func NewDrugHandler(repo *repository.DrugRepository) *DrugHandler {
	return &DrugHandler{Repo: repo}
}

func (h *DrugHandler) AddDrug(c *gin.Context) {
	var drug models.Drug

	if err := c.ShouldBindJSON(&drug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.Repo.CreateDrug(&drug); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add drug"})
		return
	}

	c.JSON(http.StatusCreated, drug)
}
