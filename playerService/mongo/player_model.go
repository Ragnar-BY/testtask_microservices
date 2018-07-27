package mongo

import (
	"github.com/Ragnar-BY/testtask_microservices/player"
	"github.com/globalsign/mgo/bson"
)

// PlayerModel represents Player model for MongoDB.
type PlayerModel struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	PlayerID int           `bson:"playerId"`
	Name     string        `bson:"name"`
	Balance  float32       `bson:"balance"`
}

// NewPlayerModel creates new playerModel from player.Player.
func NewPlayerModel(pl player.Player) PlayerModel {
	return PlayerModel{PlayerID: pl.ID, Name: pl.Name, Balance: pl.Balance}
}

// ToPlayer converts PlayerModel to player.Player.
func (pm PlayerModel) ToPlayer() *player.Player {
	return &player.Player{ID: pm.PlayerID, Name: pm.Name, Balance: pm.Balance}
}
