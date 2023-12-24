package router

import (
	Handler "github.com/Kaebraza/gopportunities/handlers"
	"github.com/gin-gonic/gin"
)

func iniciallizeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", Handler.ShowOpeningHandler)
		v1.POST("/opening", Handler.CreateOpeningHandler)
		v1.DELETE("/opening", Handler.DeleteOpeningHandler)
		v1.PUT("/opening", Handler.UpdateOpeningHandler)
		v1.GET("/openings", Handler.ListOpeningsHandler)
	}

}
