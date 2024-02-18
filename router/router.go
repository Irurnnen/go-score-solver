package router

import (
	"github.com/Irurnnen/go-score/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	mathController := new(controllers.MathController)

	v1 := router.Group("v1")
	{
		v1.GET("/addition", mathController.SolveAddition)
		v1.GET("/subtraction", mathController.SolveSubtraction)
		v1.GET("/multiplication", mathController.SolveMultiplication)
		v1.GET("/division", mathController.SolveDivision)
	}
	return router
}
