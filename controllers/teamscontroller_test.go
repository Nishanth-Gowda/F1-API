package controllers

import (
	"f1/model"
	"f1/service"
	"f1/utils"
	"testing"
)



func TestSave(t *testing.T){

	team := model.F1Team{
        TeamID:           1,
        TeamName:         "Mercedes-AMG Petronas Formula One Team",
        Country:          "Germany",
        YearFounded:      1954,
        ChampionshipsWon: 8,
        Drivers:          []string{"Lewis Hamilton", "Valtteri Bottas"},
    }

	addedTeam := service.New().Save(team)
	// Write test assertions to verify that the team was added correctly.
    // You can use the testing package's functions like t.Errorf or t.Fatalf to report test failures.
    if addedTeam.TeamID != team.TeamID {
        t.Errorf("Expected TeamID to be %d, but got %d", team.TeamID, addedTeam.TeamID)
    }
}

func TestFindAll(t *testing.T) {
    // Generate mock Formula 1 teams
    mockTeams := utils.GenerateMockF1Teams()

    // Create a new instance of the service
    teamService := service.New()

    // Variables to keep track of successful and failed saves
    successfulSaves := 0
    failedSaves := 0

    // Iterate through the mock teams and test the Save operation
    for _, team := range mockTeams {
        savedTeam := teamService.Save(team)

        // Perform validations on the saved team
        if savedTeam.TeamID != team.TeamID {
            t.Errorf("Expected TeamID %d, but got %d", team.TeamID, savedTeam.TeamID)
            failedSaves++
        } else {
            successfulSaves++
        }

        // Example validation:
        if savedTeam.TeamName != team.TeamName {
            t.Errorf("Expected TeamName %s, but got %s", team.TeamName, savedTeam.TeamName)
            failedSaves++
        } else {
            successfulSaves++
        }
    }


    // Optionally, you can assert that no saves failed
    if failedSaves > 0 {
        t.Errorf("%d saves failed", failedSaves)
    }
	
}


func TestFindById(t *testing.T) {
    // Generate mock Formula 1 teams
    mockTeams := utils.GenerateMockF1Teams()
    teamService := service.New()
    for _, team := range mockTeams {
        teamService.Save(team)
        foundTeam := teamService.FindById(team.TeamID)
        if foundTeam.TeamID != team.TeamID {
            t.Errorf("Expected TeamID %d, but got %d", team.TeamID, foundTeam.TeamID)
            return
        }
    }
}

func TestDeleteById(t *testing.T) {
    // Generate mock Formula 1 teams
    mockTeams := utils.GenerateMockF1Teams()   
    teamService := service.New()
    for _, team := range mockTeams {
        teamService.Save(team)
        teamService.DeleteById(team.TeamID)
        foundTeam := teamService.FindById(team.TeamID)
        if foundTeam.TeamID != 0 {
            t.Errorf("Expected TeamID %d, but got %d", team.TeamID, foundTeam.TeamID)
            return
        }
    }
}