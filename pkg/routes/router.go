package routes

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRoutes(r *gin.Engine) {

	createOpRouter(r)
}
