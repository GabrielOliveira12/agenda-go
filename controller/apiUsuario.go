package controller

import (
	"agenda/db"
	"agenda/model"

	"github.com/gin-gonic/gin"
)

// USUARIO

// InsertUsuario godoc
// @Summary insertusuario
// @Schemes
// @Param Usuarios body Usuario true "struct"
// @Description Insere os usuarios
// @Tags Usuarios
// @Accept json
// @Produce json
// @Success 200 {string} InsertUsuario
// @Router /usuario [post]
func InsertUsuario(c *gin.Context) {
	var usuario model.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := db.Database.Create(&usuario)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, usuario)
}

// ListUsuario godoc
// @Summary listusuario
// @Schemes
// @Tags Usuarios
// @Accept json
// @Produce json
// @Success 200 {string} ListUsuario
// @Router /usuario [get]
func ListUsuario(c *gin.Context) {
	var usuario []model.Usuario
	result := db.Database.Preload("Lembretes").Find(&usuario)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, usuario)
}

// UpdateUsuario godoc
// @Summary updateusuario
// @Schemes
// @Param Usuarios body Usuario true "struct"
// @Param id path string true "Usuario ID"
// @Description Altera os usuarios
// @Tags Usuarios
// @Accept json
// @Produce json
// @Success 200 {string} UpdateUsuario
// @Router /usuario/{id} [put]
func UpdateUsuario(c *gin.Context) {
	id := c.Param("id")

	var usuario model.Usuario
	if err := db.Database.First(&usuario, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "usuario não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := db.Database.Save(&usuario)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, usuario)
}

// DeleteUsuario godoc
// @Summary deleteusuario\
// @Schemes
// @Param id path string true "Usuario ID"
// @Description Deleta os usuarios
// @Tags Usuarios
// @Accept json
// @Produce json
// @Success 200 {string} DeleteUsuario
// @Router /usuario/{id} [delete]
func DeleteUsuario(c *gin.Context) {
	id := c.Param("id")

	result := db.Database.Delete(&model.Usuario{}, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Usuário excluído com sucesso"})
}
