package main

import (
	"agenda/db"
	"agenda/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Connect()
	routes.CrudLembretes(r)
	r.Run(":8080")
}
