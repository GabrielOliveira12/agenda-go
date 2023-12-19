package routes

import (
	"agenda/controller"
	"agenda/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CrudLembretes(r *gin.Engine) {

	docs.SwaggerInfo.BasePath = ""

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.POST("/lembretes", controller.InsertLembretes)

	r.GET("/lembretes", controller.ListLembretes)

	r.DELETE("/lembretes/:id", controller.DeleteLembretes)

	r.POST("/usuario", controller.InsertUsuario)

	r.GET("/usuario", controller.ListUsuario)

	r.PUT("/usuario/:id", controller.UpdateUsuario)

	r.DELETE("/usuario/:id", controller.DeleteUsuario)
}
