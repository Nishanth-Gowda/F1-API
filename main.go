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

	// Configure Gin to load HTML templates
	router.LoadHTMLGlob("public/*.html")


	// Route to serve HTML page
	router.GET("/home", func(ctx *gin.Context) {
    	ctx.HTML(http.StatusOK, "index.html", nil)
	})

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

	router.PUT("/teams/:id", f1TeamController.UpdateDriversById)
	

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "UP & Running",
		})
	})

	log.Fatal(router.Run(":3000"))
}