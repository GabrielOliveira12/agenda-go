package controller

import (
	"agenda/db"
	"agenda/model"

	"github.com/gin-gonic/gin"
)

// InsertLembretes godoc
// @Summary insertlembretes
// @Schemes
// @Param Lembretes body Lembrete true "struct"
// @Description Insere os lembretes
// @Tags Lembretes
// @Accept json
// @Produce json
// @Success 200 {string} InsertLembretes
// @Router /lembretes [post]
func InsertLembretes(c *gin.Context) {
	var lembrete model.Lembrete
	if err := c.ShouldBindJSON(&lembrete); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := db.Database.Create(&lembrete)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, lembrete)
}

// ListLembretes godoc
// @Summary listlembretes
// @Schemes
// @Tags Lembretes
// @Accept json
// @Produce json
// @Success 200 {string} ListLembretes
// @Router /lembretes [get]
func ListLembretes(c *gin.Context) {
	var lembretes []model.Lembrete
	result := db.Database.Find(&lembretes)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, lembretes)
}

// UpdateLembretes godoc
// @Summary updatelembretes
// @Schemes
// @Param Lembretes body Lembrete true "struct"
// @Param id path string true "Lembrete ID"
// @Description Altera os lembretes
// @Tags Lembretes
// @Accept json
// @Produce json
// @Success 200 {string} UpdateLembretes
// @Router /lembretes/{id} [put]
func UpadateLembretes(c *gin.Context) {
	id := c.Param("id")

	var lembrete model.Lembrete
	if err := db.Database.First(&lembrete, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Lembrete não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&lembrete); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := db.Database.Save(&lembrete)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, lembrete)
}

// DeleteLembretes godoc
// @Summary deletelembretes
// @Schemes
// @Param id path string true "Lembrete ID"
// @Description Exclui os lembretes
// @Tags Lembretes
// @Accept json
// @Produce json
// @Success 200 {string} DeleteLembretes
// @Router /lembretes/{id} [delete]
func DeleteLembretes(c *gin.Context) {
	id := c.Param("id")

	result := db.Database.Delete(&model.Lembrete{}, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Lembrete excluído com sucesso"})
}
