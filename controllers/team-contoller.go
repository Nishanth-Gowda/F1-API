package controllers

import (
	"encoding/json"
	"f1/model"
	"f1/service"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type F1TeamController interface{
	FindAll() []model.F1Team
	Save(ctx *gin.Context) model.F1Team
	FindById(ctx *gin.Context) model.F1Team
	DeleteById(ctx *gin.Context) model.F1Team
	UpdateDriversById(ctx *gin.Context)
} 

type f1TeamController struct{
	service service.F1TeamService
}

func New(service service.F1TeamService) F1TeamController{
	return &f1TeamController{
		service: service,
	}
}

func (c *f1TeamController) FindAll() []model.F1Team{
	return c.service.FindAll()
}

func (c *f1TeamController) Save(ctx *gin.Context) model.F1Team{
	var team model.F1Team
	ctx.BindJSON(&team)
	c.service.Save(team)
	return team
}

func (c *f1TeamController) FindById(ctx *gin.Context) model.F1Team {
    var team model.F1Team
    idStr := ctx.Param("id")

    // Convert the idStr to an integer
    id, err := strconv.Atoi(idStr)
    if err != nil {
        // Handle the error, e.g., return an error response
        // or set a default value for id.
        // For now, we'll just return an empty team and log the error.
        log.Println("Error converting id to int:", err)
        return team
    }

    // Now, you can call the FindById method with the integer id
    team = c.service.FindById(id)
    return team
}

func (c *f1TeamController) DeleteById(ctx *gin.Context) model.F1Team {
	var team model.F1Team
	idStr := ctx.Param("id")

    // Convert the idStr to an integer
    id, err := strconv.Atoi(idStr)
    if err != nil {
        // Handle the error, e.g., return an error response
        // or set a default value for id.
        // For now, we'll just return an empty team and log the error.
        log.Println("Error converting id to int:", err)
        return team
    }

	team = c.service.DeleteById(id)
	return team
}


// UpdateDriversById handles updating drivers for a team
func (c *f1TeamController) UpdateDriversById(ctx *gin.Context) {

    // Get team ID from path parameter
    teamID, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
      ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
      return
    }
  
    // Read request body into raw byte slice
    reqBody, err := io.ReadAll(ctx.Request.Body)
    if err != nil {
      ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error reading request body"})
      return
    }
  
    // Log raw request body
    log.Printf("Raw request body: %s", string(reqBody))
  
    // Unmarshal request body into drivers array
    var drivers []string
    if err := json.Unmarshal(reqBody, &drivers); err != nil {
      ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error unmarshalling JSON"})
      log.Printf("Unmarshal error: %+v", err)
      return
    }
  
    // Call service to update drivers
    updatedTeam := c.service.UpdateDrivers(teamID, drivers)
  
    // Check for 404 case
    if updatedTeam.TeamID == 0 {
      ctx.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
      return 
    }
  
    // Return updated team
    ctx.JSON(http.StatusOK, updatedTeam)
  }