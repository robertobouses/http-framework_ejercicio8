package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/http-framework_ejercicio8/app"
	"github.com/robertobouses/http-framework_ejercicio8/entity"
)

type HTTP interface {
	PostMeasurement(ctx *gin.Context)
}

type http struct {
	service app.APP
}

func NewHTTP(service app.APP) HTTP {
	return &http{
		service: service,
	}
}

func (h *http) PostMeasurement(ctx *gin.Context) {
	var newMeasurement entity.Measurement
	err := ctx.BindJSON(&newMeasurement)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error al hacer BindJSON": err.Error()})
		return
	}

	err = h.service.CreateMeasurement(newMeasurement)
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	ctx.JSON(nethttp.StatusOK, newMeasurement)

}
