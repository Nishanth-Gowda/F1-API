package controllers

import (
	"f1/model"
	"f1/service"
	"strconv"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type F1TeamController interface{
	FindAll() []model.F1Team
	Save(ctx *gin.Context) model.F1Team
	FindById(ctx *gin.Context) model.F1Team
	DeleteById(ctx *gin.Context) model.F1Team
	UpdateDriverById(ctx *gin.Context) model.F1Team
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


func (c *f1TeamController) UpdateDriverById(ctx *gin.Context) model.F1Team {
    // Get the team ID from the request URL parameter
    idStr := ctx.Param("id")

    // Convert the ID string to an integer
    id, err := strconv.Atoi(idStr)
    if err != nil {
        // Handle the error (e.g., return an error response to the client)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
        return model.F1Team{} // Return an empty F1Team
    }

    // Parse the JSON request body to get the new drivers
    var newDrivers []string
    if err := ctx.BindJSON(&newDrivers); err != nil {
        // Handle the error (e.g., return an error response to the client)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
        return model.F1Team{} // Return an empty F1Team
    }

    // Call your service's UpdateDrivers function to update the team's drivers
    updatedTeam := c.service.UpdateDrivers(id, newDrivers)

    // Check if the team was found and updated successfully
    if updatedTeam.TeamID == 0 {
        // Handle the case where the team was not found
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
        return model.F1Team{} // Return an empty F1Team
    }

    // Return the updated team as a response to the client
    return updatedTeam
}
