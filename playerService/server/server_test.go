package server

import (
	"context"
	"errors"
	"testing"

	"github.com/Ragnar-BY/testtask_microservices/player"
	"github.com/Ragnar-BY/testtask_microservices/playerService/manager"
	"github.com/Ragnar-BY/testtask_microservices/proto"
	"github.com/stretchr/testify/require"
)

var errSome = errors.New("some error")

func TestServer_CreateNewPlayer(t *testing.T) {
	db := &manager.MockDB{}
	s := NewServer(manager.NewManager(db))

	type dbArguments struct {
		playerName  string
		returnID    int
		returnError error
	}
	tt := []struct {
		name          string
		dbArgs        *dbArguments
		expectedError error
		expectedValue int32
	}{
		{
			name:          "Success",
			dbArgs:        &dbArguments{playerName: "player1", returnID: 1, returnError: nil},
			expectedError: nil,
			expectedValue: 1,
		},
		{
			name:          "DBError",
			dbArgs:        &dbArguments{playerName: "player2", returnID: 0, returnError: errSome},
			expectedError: errSome,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			playerName := ""
			if tc.dbArgs != nil {
				db.On("AddPlayer", context.Background(), tc.dbArgs.playerName).Return(tc.dbArgs.returnID, tc.dbArgs.returnError)
				playerName = tc.dbArgs.playerName
			}
			reply, err := s.CreateNewPlayer(context.Background(), &proto.CreatePlayerRequest{Name: playerName})
			if tc.expectedError != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedValue, reply.Id)
			}
		})
	}
	db.AssertExpectations(t)
}

func TestServer_GetPlayerPoints(t *testing.T) {
	db := &manager.MockDB{}
	s := NewServer(manager.NewManager(db))

	type dbArguments struct {
		playerID     int
		returnPlayer *player.Player
		returnError  error
	}
	tt := []struct {
		name            string
		playerID        int32
		dbArgs          *dbArguments
		expectedError   error
		expectedBalance float32
	}{
		{
			name:     "Success",
			playerID: 1,
			dbArgs: &dbArguments{
				playerID:     1,
				returnPlayer: &player.Player{ID: 1, Balance: 1.5},
				returnError:  nil},
			expectedError:   nil,
			expectedBalance: 1.5,
		},
		{
			name:          "DBError",
			playerID:      2,
			dbArgs:        &dbArguments{playerID: 2, returnPlayer: nil, returnError: errSome},
			expectedError: errSome,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.dbArgs != nil {
				db.On("PlayerByID", context.Background(), tc.dbArgs.playerID).Return(tc.dbArgs.returnPlayer, tc.dbArgs.returnError)
			}
			reply, err := s.GetPlayerPoints(context.Background(), &proto.PlayerIDRequest{PlayerID: tc.playerID})
			if tc.expectedError != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedBalance, reply.Balance)
			}
		})
	}
	db.AssertExpectations(t)
}
func TestServer_TakePointsFromPlayer(t *testing.T) {
	db := &manager.MockDB{}
	s := NewServer(manager.NewManager(db))
	type dbArguments struct {
		playerID          int
		returnPlayerByID  *player.Player
		returnErrorByID   error
		updatePlayer      *player.Player
		UpdatePlayerError error
	}

	tt := []struct {
		name            string
		playerID        int32
		points          float32
		dbArgs          dbArguments
		expectedBalance float32
		expectedError   error
	}{
		{
			name:     "Success",
			points:   2.5,
			playerID: 1,
			dbArgs: dbArguments{
				playerID: 1,
				returnPlayerByID: &player.Player{
					ID:      1,
					Balance: 4.0,
				},
				returnErrorByID: nil,
				updatePlayer: &player.Player{
					ID:      1,
					Balance: 1.5,
				},
				UpdatePlayerError: nil,
			},
			expectedBalance: 1.5,
			expectedError:   nil,
		},
		{
			name:     "DBError",
			points:   2.5,
			playerID: 2,
			dbArgs: dbArguments{
				playerID:        2,
				returnErrorByID: errSome,
			},
			expectedError: errSome,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			db.On("PlayerByID", context.Background(), tc.dbArgs.playerID).Return(tc.dbArgs.returnPlayerByID, tc.dbArgs.returnErrorByID)
			if tc.dbArgs.updatePlayer != nil {
				db.On("UpdatePlayer", context.Background(), tc.dbArgs.playerID, *tc.dbArgs.updatePlayer).Return(tc.dbArgs.UpdatePlayerError)
			}
			reply, err := s.TakePointsFromPlayer(context.Background(), &proto.PlayerIDPointRequest{PlayerID: tc.playerID, Points: tc.points})
			if tc.expectedError != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedBalance, reply.Balance)
			}
		})
	}
}

func TestServer_FundPointsToPlayer(t *testing.T) {
	db := &manager.MockDB{}
	s := NewServer(manager.NewManager(db))
	type dbArguments struct {
		playerID          int
		returnPlayerByID  *player.Player
		returnErrorByID   error
		updatePlayer      *player.Player
		UpdatePlayerError error
	}

	tt := []struct {
		name            string
		playerID        int32
		points          float32
		dbArgs          dbArguments
		expectedBalance float32
		expectedError   error
	}{
		{
			name:     "Success",
			points:   2.5,
			playerID: 1,
			dbArgs: dbArguments{
				playerID: 1,
				returnPlayerByID: &player.Player{
					ID:      1,
					Balance: 1.5,
				},
				returnErrorByID: nil,
				updatePlayer: &player.Player{
					ID:      1,
					Balance: 4.0,
				},
				UpdatePlayerError: nil,
			},
			expectedBalance: 4.0,
			expectedError:   nil,
		},
		{
			name:     "DBError",
			points:   2.5,
			playerID: 2,
			dbArgs: dbArguments{
				playerID:        2,
				returnErrorByID: errSome,
			},
			expectedError: errSome,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			db.On("PlayerByID", context.Background(), tc.dbArgs.playerID).Return(tc.dbArgs.returnPlayerByID, tc.dbArgs.returnErrorByID)
			if tc.dbArgs.updatePlayer != nil {
				db.On("UpdatePlayer", context.Background(), tc.dbArgs.playerID, *tc.dbArgs.updatePlayer).Return(tc.dbArgs.UpdatePlayerError)
			}
			reply, err := s.FundPointsToPlayer(context.Background(), &proto.PlayerIDPointRequest{PlayerID: tc.playerID, Points: tc.points})
			if tc.expectedError != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedBalance, reply.Balance)
			}
		})
	}
}

func TestManagerRouter_RemovePlayer(t *testing.T) {
	db := &manager.MockDB{}
	s := NewServer(manager.NewManager(db))

	type dbArguments struct {
		playerID          int
		returnErrorDelete error
	}
	tt := []struct {
		name          string
		dbArgs        dbArguments
		playerID      int32
		expectedError error
	}{
		{
			name:     "Success",
			playerID: 1,
			dbArgs:   dbArguments{playerID: 1},
		},
		{
			name:          "DBError",
			playerID:      2,
			dbArgs:        dbArguments{playerID: 2, returnErrorDelete: errSome},
			expectedError: errSome,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			db.On("DeletePlayer", context.Background(), tc.dbArgs.playerID).Return(tc.dbArgs.returnErrorDelete)
			_, err := s.RemovePlayer(context.Background(), &proto.PlayerIDRequest{PlayerID: tc.playerID})
			if tc.expectedError != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
