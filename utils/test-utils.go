package utils

import "f1/model"

// GenerateMockF1Teams generates mock data for a list of Formula 1 teams.
func GenerateMockF1Teams() []model.F1Team {
    teams := []model.F1Team{
        {TeamID: 1, TeamName: "Mercedes-AMG Petronas Formula One Team", Country: "Germany", YearFounded: 1954, ChampionshipsWon: 8, Drivers: []string{"Lewis Hamilton", "Valtteri Bottas"}},
        {TeamID: 2, TeamName: "Red Bull Racing", Country: "Austria", YearFounded: 2005, ChampionshipsWon: 4, Drivers: []string{"Max Verstappen", "Sergio Perez"}},
        {TeamID: 3, TeamName: "Scuderia Ferrari", Country: "Italy", YearFounded: 1929, ChampionshipsWon: 16, Drivers: []string{"Charles Leclerc", "Carlos Sainz"}},
        {TeamID: 4, TeamName: "McLaren F1 Team", Country: "United Kingdom", YearFounded: 1963, ChampionshipsWon: 8, Drivers: []string{"Lando Norris", "Daniel Ricciardo"}},
        {TeamID: 5, TeamName: "Aston Martin Cognizant Formula One Team", Country: "United Kingdom", YearFounded: 1952, ChampionshipsWon: 0, Drivers: []string{"Sebastian Vettel", "Lance Stroll"}},
        // Add more Formula 1 teams here
    }
    return teams
}