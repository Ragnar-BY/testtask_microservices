package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ragnar-BY/testtask_microservices/proto"
	"github.com/gorilla/mux"
)

// api is for manager.
type api struct {
	player proto.PlayerServiceClient
}

func newGateway(player proto.PlayerServiceClient, router *mux.Router) *mux.Router {
	mngrRouter := api{player}
	router.HandleFunc("/add", mngrRouter.addPlayerHandler).
		Methods(http.MethodPost).
		Queries("name", "{name}")
	router.HandleFunc("/balance/{playerId:[0-9]+}", mngrRouter.balancePlayerHandler).
		Methods(http.MethodGet)
	router.HandleFunc("/fund/{playerId:[0-9]+}", mngrRouter.fundPointsHandler).
		Methods(http.MethodPut).
		Queries("points", "{points:[0-9]+\\.?[0-9]{0,2}}")
	router.HandleFunc("/take/{playerId:[0-9]+}", mngrRouter.takePointsHandler).
		Methods(http.MethodPut).
		Queries("points", "{points:[0-9]+\\.?[0-9]{0,2}}")
	router.HandleFunc("/remove/{playerId:[0-9]+}", mngrRouter.removePlayerHandler).
		Methods(http.MethodDelete)
	return router
}

// addPlayerHandler creates new player, returns id.
func (c *api) addPlayerHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	if name == "" {
		Error(w, http.StatusBadRequest, "wrong name")
		return
	}
	reply, err := c.player.CreateNewPlayer(r.Context(), &proto.CreatePlayerRequest{Name: name})
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot create new player: %v", err))
		return
	}
	JSON(w, http.StatusCreated, reply.Id)
}

// balancePlayerHandler returns player balance.
func (c *api) balancePlayerHandler(w http.ResponseWriter, r *http.Request) {
	playerID, err := getIntValue(r, "playerId")
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot get playerId: %v", err))
		return
	}
	reply, err := c.player.GetPlayerPoints(r.Context(), &proto.PlayerIDRequest{PlayerID: int32(playerID)})
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot get player points: %v", err))
		return
	}
	JSON(w, http.StatusOK, reply.Balance)
}

// fundPointsHandler gives points to player, returns new balance.
func (c *api) fundPointsHandler(w http.ResponseWriter, r *http.Request) {
	playerID, points, err := getPlayerIDAndPoints(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	reply, err := c.player.FundPointsToPlayer(r.Context(), &proto.PlayerIDPointRequest{PlayerID: int32(playerID), Points: points})
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot fund points to player: %v", err))
		return
	}
	JSON(w, http.StatusOK, reply.Balance)
}

///takePointsHandler takes points if possible from player, returns new balance.
func (c *api) takePointsHandler(w http.ResponseWriter, r *http.Request) {
	playerID, points, err := getPlayerIDAndPoints(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	reply, err := c.player.TakePointsFromPlayer(r.Context(), &proto.PlayerIDPointRequest{PlayerID: int32(playerID), Points: points})
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot take points from player: %v", err))
		return
	}
	JSON(w, http.StatusOK, reply.Balance)
}

func (c *api) removePlayerHandler(w http.ResponseWriter, r *http.Request) {
	playerID, err := getIntValue(r, "playerId")
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot get playerId: %v", err))
		return
	}
	_, err = c.player.RemovePlayer(r.Context(), &proto.PlayerIDRequest{PlayerID: int32(playerID)})
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot remove player: %v", err))
		return
	}
	JSON(w, http.StatusOK, playerID)
}
func getPlayerIDAndPoints(r *http.Request) (int, float32, error) {
	playerID, err := getIntValue(r, "playerId")
	if err != nil {
		return 0, 0, fmt.Errorf("cannot get playerId: %v", err)
	}
	points, err := getFloatValue(r, "points")
	if err != nil {
		return 0, 0, fmt.Errorf("cannot get points: %v", err)
	}
	return playerID, points, nil
}

// TODO it is better move this functions to other file.
func getIntValue(r *http.Request, key string) (int, error) {
	val := mux.Vars(r)[key]
	return strconv.Atoi(val)
}

func getFloatValue(r *http.Request, key string) (float32, error) {
	val := mux.Vars(r)[key]
	f64, err := strconv.ParseFloat(val, 32)
	return float32(f64), err
}
