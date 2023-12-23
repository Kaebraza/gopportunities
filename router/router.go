package router

import "github.com/gin-gonic/gin"

// nome da função tem que ser maiúsculo para ser exportada
func Initialize() { 
	// incializa o router ultilizando as configs default do gin
	router := gin.Default()
	// go mod tidy limpa o go mod, instala o que falta e exclui o que não usa
	// Define uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// rodando a APIgo
	router.Run("8080") // listen and serve on 0.0.0.0:8080

}
