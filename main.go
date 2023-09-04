package main

import (
	"f1/controllers"
	"f1/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	f1teamservice service.F1TeamService = service.New()
	f1TeamController controllers.F1TeamController = controllers.New(f1teamservice)
)

func main() {
	router := gin.Default()

	router.GET("/teams", func(ctx *gin.Context) {
		ctx.JSON(200, f1TeamController.FindAll())
	})

	router.POST("/teams", func(ctx *gin.Context) {
		ctx.JSON(200, f1TeamController.Save(ctx))
	})

	router.GET("/teams/:id", func(ctx *gin.Context) {
		ctx.JSON(200, f1TeamController.FindById(ctx))
	})

	router.DELETE("/teams/:id", func(ctx *gin.Context) {
		ctx.JSON(200, f1TeamController.DeleteById(ctx))
	})

	router.PUT("/teams/:id", func(ctx *gin.Context) {
		ctx.JSON(200, f1TeamController.UpdateDriverById(ctx))
	})

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "UP & Running",
		})
	})

	log.Fatal(router.Run(":3000"))
}