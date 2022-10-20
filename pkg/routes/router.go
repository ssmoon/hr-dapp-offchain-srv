package routes

import (
	"hr-dapp/srv/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRoutes(r *gin.Engine) {
	var worker = r.Group("/api/worker")
	{
		con := &controllers.WorkerSummary{}
		worker.POST("/", con.CreateWorker)
		worker.GET("/all", con.GetWorkerSummary)
	}
	var cert = r.Group("/api/certificate")
	{
		con := &controllers.WorkerSummary{}
		cert.POST("/", con.CreateWorker)
		cert.GET("/all", con.GetWorkerSummary)
	}
	var career = r.Group("/api/career")
	{
		con := &controllers.WorkerSummary{}
		career.POST("/", con.CreateWorker)
		career.GET("/all", con.GetWorkerSummary)
	}
}
