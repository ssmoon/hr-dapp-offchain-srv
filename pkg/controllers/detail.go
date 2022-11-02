package controllers

import (
	models "hr-dapp/srv/pkg/models/db"
	"hr-dapp/srv/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
)

type WorkerDetail struct {
}

func (controller *WorkerDetail) GetWorkerDetail(context *gin.Context) {
	workerId, _ := strconv.Atoi(context.Query("workerId"))
	context.JSON(http.StatusOK, services.GetWorkerDetail(uint32(workerId)))
}

func (controller *WorkerDetail) CreateCertificate(context *gin.Context) {
	workerId, _ := strconv.Atoi(context.Query("workerId"))
	certificateId, _ := strconv.Atoi(context.Query("certificateId"))
	acquiredAt := carbon.Parse(context.Query("acquireAt"))
	services.CreateCertificate(uint32(workerId), uint32(certificateId), acquiredAt.Carbon2Time())
	context.Status(http.StatusNoContent)
}

func (controller *WorkerDetail) CreateCareer(context *gin.Context) {
	var json *models.WorkerCareer
	if err := context.ShouldBindJSON(&json); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.CreateCareer(json)
	context.Status(http.StatusNoContent)
}

func (controller *WorkerDetail) FinishCareer(context *gin.Context) {
	workerId, _ := strconv.Atoi(context.Query("workerId"))
	endAt := carbon.Parse(context.Query("endAt"))
	services.FinishCareer(uint32(workerId), endAt.Carbon2Time())
	context.Status(http.StatusNoContent)
}

func (controller *WorkerDetail) ValidateWorker(context *gin.Context) {
	workerId, _ := strconv.Atoi(context.Query("workerId"))
	result, err := services.ValidateOnChain(uint32(workerId))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, result)
}
