package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	Connect()
	CrudLembretes(r)
	r.Run(":8080")
}
