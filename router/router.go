package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Iniciallize router
	router := gin.Default()
	// Iniciallize routes
	iniciallizeRoutes(router)
	// Run server
	router.Run("localhost:8080") // listen and serve on 0.0.0.0:8080

}
