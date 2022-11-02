package routes

import (
	"hr-dapp/srv/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRoutes(r *gin.Engine) {
	var worker = r.Group("/api/worker")
	{
		con1 := &controllers.WorkerSummary{}
		con2 := &controllers.WorkerDetail{}
		worker.POST("/", con1.CreateWorker)
		worker.GET("/all", con1.GetWorkerSummary)
		worker.GET("/:workerId", con2.GetWorkerDetail)
	}
	var cert = r.Group("/api/certificate")
	{
		con := &controllers.WorkerDetail{}
		cert.POST("/", con.CreateCertificate)
	}
	var career = r.Group("/api/career")
	{
		con := &controllers.WorkerDetail{}
		career.POST("/new", con.CreateCareer)
		career.PUT("/:careerId/finish", con.FinishCareer)
		career.POST("/validate", con.ValidateWorker)
	}
}
