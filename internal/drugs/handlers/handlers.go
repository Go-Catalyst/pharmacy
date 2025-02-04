package handlers

import (
	"net/http"
	"pharmacy/internal/drugs/models"
	"pharmacy/internal/drugs/repository"
	"strconv"

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

func (h *DrugHandler) GetAllDrugs(c *gin.Context) {
	drugs, err := h.Repo.GetAllDrugs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve drugs"})
		return
	}

	c.JSON(http.StatusOK, drugs)
}

func (h *DrugHandler) GetDrugByID(c *gin.Context) {
	idParam := c.Param("id")
	drugID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid drug ID"})
		return
	}

	drug, err := h.Repo.GetDrugByID(uint(drugID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Drug not found"})
		return
	}

	c.JSON(http.StatusOK, drug)
}