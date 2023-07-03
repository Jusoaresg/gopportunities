package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializa o router utilizando as config Default do gin
	router := gin.Default()
	//Definindo a Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
