package store

import "gorm.io/gorm"

type Team struct {
	Id int `gorm."primaryKey"`
	Name string
	Players []Player
	Coach string
	Record Record	

}

type Record struct {
	Wins int32 
	Loses int32
}

func GetTeamPlayers(teamId int) ([]Player, error){
	var players []Player
	result := DB.Where("Id = ?", teamId).Find(&players)
	return players, result.Error
}

func CreateTheMets(db *gorm.DB) error {
    team := Team{
        Name: "The Mets",
        Players: []Player{
            {Name: "Donny"},
            {Name: "Jaxx"},
            {Name: "Dez"},
            {Name: "Luke"},
            {Name: "Nathan"},
            {Name: "Lucas"},
            {Name: "Andrew"},
            {Name: "Lorenzo"},
            {Name: "Veevaan"},    
            {Name: "Calen"},
            {Name: "Caleb"},
            {Name: "Jonah"},
        },
    }

    return db.Create(&team).Error
}

