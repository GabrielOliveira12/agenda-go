package main

import (
	"agenda/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	var lembrete Lembrete
	if err := c.ShouldBindJSON(&lembrete); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := Database.Create(&lembrete)
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
	var lembretes []Lembrete
	result := Database.Find(&lembretes)
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

	var lembrete Lembrete
	if err := Database.First(&lembrete, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Lembrete não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&lembrete); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := Database.Save(&lembrete)
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

	result := Database.Delete(&Lembrete{}, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Lembrete excluído com sucesso"})
}

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
	var usuario Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := Database.Create(&usuario)
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
	var usuario []Usuario
	result := Database.Preload("Lembretes").Find(&usuario)
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

	var usuario Usuario
	if err := Database.First(&usuario, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "usuario não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := Database.Save(&usuario)
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

	result := Database.Delete(&Usuario{}, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Usuário excluído com sucesso"})
}

func CrudLembretes(r *gin.Engine) {

	docs.SwaggerInfo.BasePath = ""

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	//LEMBRETE
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.POST("/lembretes", InsertLembretes)

	r.GET("/lembretes", ListLembretes)

	r.DELETE("/lembretes/:id", DeleteLembretes)

	//USUARIO
	r.POST("/usuario", InsertUsuario)

	r.GET("/usuario", ListUsuario)

	r.PUT("/usuario/:id", UpdateUsuario)

	r.DELETE("/usuario/:id", DeleteUsuario)
}
