package controllers

import (
	models "hr-dapp/srv/pkg/models/db"
	"hr-dapp/srv/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WorkerSummary struct {
}

func (controller *WorkerSummary) GetWorkerSummary(context *gin.Context) {

}

func (controller *WorkerSummary) CreateWorker(context *gin.Context) {
	var json *models.Worker
	if err := context.ShouldBindJSON(&json); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.CreateWorker(json)
	context.Status(http.StatusNoContent)
}
