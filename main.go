package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/robertobouses/http-framework_ejercicio8/app"
	"github.com/robertobouses/http-framework_ejercicio8/http"
	"github.com/robertobouses/http-framework_ejercicio8/internal"
	"github.com/robertobouses/http-framework_ejercicio8/repository"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASS:", os.Getenv("DB_PASS"))
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("DB_DATABASE:", os.Getenv("DB_DATABASE"))

	db, err := internal.NewPostgres(internal.DBConfig{
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	})

	if err != nil {
		panic(err)
	}
	log.Println("la conexión se abrió correctamente")
	defer db.Close()

	var repo repository.REPOSITORY
	repo, err = repository.NewRepository(db)
	if err != nil {
		panic(err)
	}
	var appController app.APP
	var httpController http.HTTP

	appController = app.NewAPP(repo)
	httpController = http.NewHTTP(appController)

	server := gin.Default()

	server.POST("/measurement", func(ctx *gin.Context) {
		httpController.PostMeasurement(ctx)
	})

	server.GET("/measurements", func(ctx *gin.Context) {
		httpController.GetMeasurement(ctx)
	})

	server.DELETE("/measurement/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		httpController.DeleteMeasurement(ctx, id)

	})

	server.GET("/cubic/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		httpController.GetCubic(ctx, id)
	})

	server.GET("/scale/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		httpController.GetScale(ctx, id)
	})

	server.POST("/compare", func(ctx *gin.Context) {
		httpController.PostCompare(ctx)

	})

	//hacer un endpoint que permita borrar los que tienen vacío en el id

	server.DELETE("/empty", func(ctx *gin.Context) {
		httpController.DeleteEmptyMeasurement(ctx)
	})

	server.DELETE("/all", func(ctx *gin.Context) {
		httpController.DeleteAllMeasurement(ctx)
	})

	port := ":8080"
	log.Printf("Escuchando en el puerto%s\n", port)

	if err := server.Run(port); err != nil {
		log.Fatal(err)
	}
}
