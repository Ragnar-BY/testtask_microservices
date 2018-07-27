package manager

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Ragnar-BY/testtask_microservices/player"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestManager_CreateNewPlayer(t *testing.T) {
	db := &MockDB{}
	m := NewManager(db)
	ctx := context.Background()

	tests := []struct {
		name          string
		playerName    string
		expectedID    int
		expectedError error
	}{
		{
			name:          "CreateNewPlayer Success",
			playerName:    "player1",
			expectedID:    1,
			expectedError: nil,
		},
		{
			name:          "CreateNewPlayer Error",
			playerName:    "player2",
			expectedError: errors.New("wrong id"),
		},
	}

	for _, tt := range tests {
		db.On("AddPlayer", ctx, tt.playerName).Return(tt.expectedID, tt.expectedError)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := m.CreateNewPlayer(ctx, tt.playerName)
			if tt.expectedError != nil {
				assert.Error(t, err, tt.expectedError.Error())

			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedID, id)
			}
		})
	}
	db.AssertExpectations(t)
}

func TestManager_GetPlayerPoints(t *testing.T) {
	db := &MockDB{}
	m := NewManager(db)
	ctx := context.Background()

	tests := []struct {
		name            string
		playerID        int
		expectedPlayer  *player.Player
		expectedDBError error
	}{
		{
			name:     "GetPlayerPoints Success",
			playerID: 1,
			expectedPlayer: &player.Player{
				ID:      1,
				Balance: 1.5,
			},
			expectedDBError: nil,
		},
		{
			name:            "GetPlayerPoints Error",
			playerID:        2,
			expectedPlayer:  nil,
			expectedDBError: errors.New("wrong id"),
		},
	}
	for _, tt := range tests {
		db.On("PlayerByID", ctx, tt.playerID).Return(tt.expectedPlayer, tt.expectedDBError)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			balance, err := m.GetPlayerPoints(ctx, tt.playerID)
			if tt.expectedDBError != nil {
				assert.Error(t, err, fmt.Sprintf("cannot get player ID: %v", tt.expectedDBError))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedPlayer.Balance, balance)
			}
		})
	}
	db.AssertExpectations(t)
}

func TestManager_TakePointsFromPlayer(t *testing.T) {
	db := &MockDB{}
	m := NewManager(db)
	ctx := context.Background()

	tests := []struct {
		testName                  string
		playerID                  int
		points                    float32
		expectedPlayerByID        *player.Player
		expectedPlayerByIDError   error
		updatePlayer              *player.Player
		expectedUpdatePlayerError error
		expectedBalance           float32
		expectedError             string
	}{
		{
			testName: "Success",
			playerID: 1,
			points:   1.5,
			expectedPlayerByID: &player.Player{
				ID:      1,
				Balance: 4.0,
			},
			updatePlayer: &player.Player{
				ID:      1,
				Balance: 2.5,
			},
			expectedBalance: 2.5,
		},
		{
			testName:                "PlayerByIDError",
			playerID:                2,
			points:                  1.5,
			expectedPlayerByID:      nil,
			expectedPlayerByIDError: errors.New("wrong id"),
			expectedError:           fmt.Sprint("cannot get player ID: wrong id"),
		},
		{
			testName: "BalanceError",
			playerID: 3,
			points:   10,
			expectedPlayerByID: &player.Player{
				ID:      3,
				Balance: 4.0,
			},
			expectedError: ErrNotEnoughBalance.Error(),
		},
		{
			testName: "UpdatePlayerError",
			playerID: 4,
			points:   1.5,
			expectedPlayerByID: &player.Player{
				ID:      4,
				Balance: 4.0,
			},
			updatePlayer: &player.Player{
				ID:      4,
				Balance: 2.5,
			},
			expectedUpdatePlayerError: errors.New("update error"),
			expectedError:             fmt.Sprintf("update error"),
		},
	}
	for _, tt := range tests {
		db.On("PlayerByID", ctx, tt.playerID).Return(tt.expectedPlayerByID, tt.expectedPlayerByIDError)
		if tt.updatePlayer != nil {
			db.On("UpdatePlayer", ctx, tt.playerID, *tt.updatePlayer).Return(tt.expectedUpdatePlayerError)
		}
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			balance, err := m.TakePointsFromPlayer(ctx, tt.playerID, tt.points)
			if tt.expectedError != "" {
				assert.Error(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedPlayerByID.Balance, balance)
			}
		})
	}
	db.AssertExpectations(t)
}

func TestManager_FundPointsToPlayer(t *testing.T) {
	db := &MockDB{}
	m := NewManager(db)
	ctx := context.Background()

	tests := []struct {
		testName                  string
		playerID                  int
		points                    float32
		expectedPlayerByID        *player.Player
		expectedPlayerByIDError   error
		updatePlayer              *player.Player
		expectedUpdatePlayerError error
		expectedBalance           float32
		expectedError             string
	}{
		{
			testName: "Success",
			playerID: 1,
			points:   1.5,
			expectedPlayerByID: &player.Player{
				ID:      1,
				Balance: 4.0,
			},
			updatePlayer: &player.Player{
				ID:      1,
				Balance: 5.5,
			},
			expectedBalance: 5.5,
		},
		{
			testName:                "PlayerByIDError",
			playerID:                2,
			points:                  1.5,
			expectedPlayerByID:      nil,
			expectedPlayerByIDError: errors.New("wrong id"),
			expectedError:           fmt.Sprint("cannot get player ID: wrong id"),
		},
		{
			testName: "UpdatePlayerError",
			playerID: 4,
			points:   1.5,
			expectedPlayerByID: &player.Player{
				ID:      4,
				Balance: 4.0,
			},
			updatePlayer: &player.Player{
				ID:      4,
				Balance: 5.5,
			},
			expectedUpdatePlayerError: errors.New("update error"),
			expectedError:             fmt.Sprintf("update error"),
		},
	}
	for _, tt := range tests {
		db.On("PlayerByID", ctx, tt.playerID).Return(tt.expectedPlayerByID, tt.expectedPlayerByIDError)
		if tt.updatePlayer != nil {
			db.On("UpdatePlayer", ctx, tt.playerID, *tt.updatePlayer).Return(tt.expectedUpdatePlayerError)
		}
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			balance, err := m.FundPointsToPlayer(ctx, tt.playerID, tt.points)
			if tt.expectedError != "" {
				assert.Error(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedPlayerByID.Balance, balance)
			}
		})
	}
	db.AssertExpectations(t)
}

func TestManager_RemovePlayer(t *testing.T) {
	db := &MockDB{}
	m := NewManager(db)
	ctx := context.Background()

	tests := []struct {
		testName      string
		playerID      int
		expectedError error
	}{
		{
			testName:      "Success",
			playerID:      1,
			expectedError: nil,
		},
		{
			testName:      "Error",
			playerID:      -1,
			expectedError: errors.New("wrong id"),
		},
	}
	for _, tt := range tests {
		db.On("DeletePlayer", ctx, tt.playerID).Return(tt.expectedError)
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := m.RemovePlayer(ctx, tt.playerID)
			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
	db.AssertExpectations(t)
}
