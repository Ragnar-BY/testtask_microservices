package main

import (
	"log"

	"github.com/Ragnar-BY/testtask_microservices/playerService/manager"
	"github.com/Ragnar-BY/testtask_microservices/playerService/mongo"
	"github.com/Ragnar-BY/testtask_microservices/playerService/mysql"
	"github.com/Ragnar-BY/testtask_microservices/playerService/server"
)

func startWithMongo(opts settings) {
	var session mongo.Session
	err := session.Open(opts.Address)
	if err != nil {
		log.Fatalf("Cannot start MongoDB on the %s: %v", opts.Address, err)
	}
	players, err := session.Players(opts.DBName, opts.PlayerCollection)
	if err != nil {
		log.Fatalf("Cannot get player collection: %v", err)
	}
	mngr := manager.NewManager(players)
	s := server.NewServer(mngr)
	s.Start(opts.PlayerServerAddress)
}

func startWithMySQL(opts settings) {
	sql, err := mysql.Open(opts.User, opts.Password, opts.DBName)
	if err != nil {
		log.Fatal("Cannot start MySQL: ", err)
	}
	ps := mysql.PlayerService{DB: sql, Name: opts.PlayerCollection}

	mngr := manager.NewManager(ps)
	s := server.NewServer(mngr)
	s.Start(opts.PlayerServerAddress)

}
func main() {
	opts := new(settings)
	err := opts.Parse()
	if err != nil {
		log.Fatalf("Cannot parse settings: %v", err)
	}
	if opts.Type == "mysql" {
		startWithMySQL(*opts)
	} else {
		startWithMongo(*opts)
	}
}
