package service

import "f1/model"

type F1TeamService interface {
	Save(model.F1Team) model.F1Team
	FindAll() []model.F1Team
	FindById(id int) model.F1Team
	DeleteById(id int) model.F1Team
	UpdateDrivers(id int, drivers []string) model.F1Team
}

type f1TeamService struct {
	f1teams []model.F1Team
}

func New() F1TeamService{
	return &f1TeamService{}
}



func (f1 *f1TeamService) Save(team model.F1Team) model.F1Team {
	f1.f1teams = append(f1.f1teams, team)
	return team
}


func (f1 *f1TeamService) FindAll() []model.F1Team {
	return f1.f1teams
}

func (f1 *f1TeamService) FindById(id int) model.F1Team {
	for _, team := range f1.f1teams {
		if team.TeamID == id {
			return team
		}
		
	}
	return model.F1Team{}
}

func (f1 *f1TeamService) DeleteById(id int) model.F1Team {
	var team model.F1Team
	for i, t := range f1.f1teams {
		if t.TeamID == id {
			team = t
			f1.f1teams = append(f1.f1teams[:i], f1.f1teams[i+1:]...)
		}
	}
	return team
}

func (f1 *f1TeamService) UpdateDrivers(id int, drivers []string) model.F1Team{
	for i, team := range f1.f1teams {
		if team.TeamID == id {
			team.Drivers = drivers
			f1.f1teams[i] = team // Update the team in the slice
		}
	}
	return f1.FindById(id)
}
