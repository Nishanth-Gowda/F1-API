package model

type F1Team struct{
	TeamID int `json:"teamId"`
	TeamName string `json:"teamName"`
	Country string `json:"country"`
	YearFounded int `json:"yearFounded"`
	ChampionshipsWon int `json:"championshipsWon"`
	Drivers []string `json:"drivers"`
}