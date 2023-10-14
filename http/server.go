package http

import (
	"fmt"
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/http-framework_ejercicio8/app"
	"github.com/robertobouses/http-framework_ejercicio8/entity"
)

type HTTP interface {
	PostMeasurement(ctx *gin.Context)
	GetMeasurement(ctx *gin.Context)
	DeleteMeasurement(ctx *gin.Context, id string)
	GetCubic(ctx *gin.Context, id string)
	GetScale(ctx *gin.Context, id string)
	PostCompare(ctx *gin.Context)
	DeleteAllMeasurement(ctx *gin.Context)
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

func (h *http) GetMeasurement(ctx *gin.Context) {
	measurement, err := h.service.PrintMeasurement()
	if err != nil {
		log.Printf("Error al obtener alimentos", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.Print("alimentos en cada capa http:", measurement)
	ctx.JSON(nethttp.StatusOK, measurement)
}

func (h *http) DeleteMeasurement(ctx *gin.Context, id string) {
	err := h.service.DeleteMeasurement(id)
	if err != nil {
		log.Print("Error al llamar de http a app. Función delete", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})

	}
	ctx.JSON(nethttp.StatusOK, "se ha borrado correctamente")
}

func (h *http) GetCubic(ctx *gin.Context, id string) {
	cubic, err := h.service.CalcCubic(id)
	if err != nil {
		log.Printf("Error", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(nethttp.StatusOK, cubic)

}

func (h *http) GetScale(ctx *gin.Context, id string) {
	cubic, err := h.service.CalcCubic(id)
	if err != nil {
		log.Printf("Error: %v", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var solution string

	if cubic < 10000 && cubic > 100 {
		solution = fmt.Sprintf("APTO. El valor cúbico %d se encuentra dentro de la escala subvencionable", cubic)
	} else {
		solution = fmt.Sprintf("NO APTO. El valor cúbico %d está fuera de la escala subvencionable", cubic)
	}

	ctx.JSON(nethttp.StatusOK, gin.H{"solution": solution})
}

func (h *http) PostCompare(ctx *gin.Context) {
	var measurementRequest entity.MeasurementRequest
	if err := ctx.ShouldBindJSON(&measurementRequest); err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error al hacer ShouldBindJSON": err.Error()})
		return
	}

	if len(measurementRequest.IDs) < 2 {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": "Se requieren al menos 2 IDs para comparar"})
		return
	}

	var cubicResults []int
	for _, id := range measurementRequest.IDs {
		cubic, err := h.service.CalcCubic(id)
		if err != nil {
			log.Printf("Error al calcular el cubo para ID %s: %v", id, err)
			ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		cubicResults = append(cubicResults, cubic)
	}

	minCubic, maxCubic := h.service.FindMinAndMaxCubic(cubicResults)

	ctx.JSON(nethttp.StatusOK, gin.H{
		"min_cubic": minCubic,
		"max_cubic": maxCubic,
	})
}

func (h *http) DeleteAllMeasurement(ctx *gin.Context) {
	err := h.service.DeleteAllMeasurement()
	if err != nil {
		log.Print("Error al llamar http a app. Función delete", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(nethttp.StatusOK, "se ha borrado correctamente")
}
