package store

type Player struct {
	Id int `gorm."primaryKey"`
	Name string
	Positions []Position 
}

type Position struct {
	Id int `gorm."primaryKey"`
	Number int32
	Name int32
}

func GetAllPlayers() ([]Player, error){
	var players []Player
	result := DB.Find(&players)
	return players, result.Error
}
