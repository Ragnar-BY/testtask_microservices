package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"strconv"

	"github.com/Ragnar-BY/testtask_microservices/apigateway/server/mocks"
	"github.com/Ragnar-BY/testtask_microservices/proto"
	"github.com/gavv/httpexpect"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//go:generate mockery -name=PlayerServiceClient -dir ./proto -output ./webserver/server/mocks

const (
	badPlayerID = "98765432109876543210" // too long for int.
)

func TestAPI_AddPlayerHandler(t *testing.T) {
	players := &mocks.PlayerServiceClient{}
	gate := newGateway(players, mux.NewRouter())
	server := httptest.NewServer(gate)
	defer server.Close()
	e := httpexpect.New(t, server.URL)

	type rpcArguments struct {
		playerName  string
		returnID    int32
		returnError error
	}
	tt := []struct {
		name           string
		rpcArgs        *rpcArguments
		expectedStatus int
		expectedValue  int
	}{
		{
			name:           "Success",
			rpcArgs:        &rpcArguments{playerName: "player1", returnID: 1, returnError: nil},
			expectedStatus: http.StatusCreated,
			expectedValue:  1,
		},
		{
			name:           "WrongName",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "RPCError",
			rpcArgs:        &rpcArguments{playerName: "player2", returnID: 0, returnError: errors.New("cannot add new player")},
			expectedStatus: http.StatusBadRequest,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			playerName := ""
			if tc.rpcArgs != nil {
				players.On("CreateNewPlayer", mock.Anything, &proto.CreatePlayerRequest{Name: tc.rpcArgs.playerName}).Return(&proto.CreatePlayerReply{Id: tc.rpcArgs.returnID}, tc.rpcArgs.returnError)
				playerName = tc.rpcArgs.playerName
			}
			expect := e.Request(http.MethodPost, "/add").WithQuery("name", playerName).Expect()

			expect.Status(tc.expectedStatus)
			if tc.expectedValue != 0 {
				expect.JSON().Number().Equal(tc.expectedValue)
			}
		})
	}
	players.AssertExpectations(t)
}

func TestManagerRouter_balancePlayerHandler(t *testing.T) {
	players := &mocks.PlayerServiceClient{}
	gate := newGateway(players, mux.NewRouter())
	server := httptest.NewServer(gate)
	defer server.Close()
	e := httpexpect.New(t, server.URL)

	type rpcArguments struct {
		playerID      int32
		returnBalance float32
		returnError   error
	}
	tt := []struct {
		name            string
		path            string
		rpcArgs         *rpcArguments
		expectedStatus  int
		expectedBalance float32
	}{
		{
			name:            "Success",
			path:            "1",
			rpcArgs:         &rpcArguments{playerID: 1, returnBalance: 1.5, returnError: nil},
			expectedStatus:  http.StatusOK,
			expectedBalance: 1.5,
		},
		{
			name:           "RPCError",
			path:           "2",
			rpcArgs:        &rpcArguments{playerID: 2, returnBalance: 0, returnError: errors.New("some error")},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "PlayerParseIDError",
			path:           badPlayerID,
			rpcArgs:        nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "WrongID",
			path:           "wrongid",
			rpcArgs:        nil,
			expectedStatus: http.StatusNotFound,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.rpcArgs != nil {
				players.On("GetPlayerPoints", mock.Anything, &proto.PlayerIDRequest{PlayerID: tc.rpcArgs.playerID}).Return(&proto.PlayerBalanceReply{Balance: tc.rpcArgs.returnBalance}, tc.rpcArgs.returnError)
			}
			expect := e.Request(http.MethodGet, "/balance/"+tc.path).Expect()

			expect.Status(tc.expectedStatus)
			if tc.expectedBalance > 0 {
				expect.JSON().Number().Equal(tc.expectedBalance)
			}
		})
	}
	players.AssertExpectations(t)
}

func TestManagerRouter_fundPointsHandler(t *testing.T) {
	players := &mocks.PlayerServiceClient{}
	gate := newGateway(players, mux.NewRouter())
	server := httptest.NewServer(gate)
	defer server.Close()
	e := httpexpect.New(t, server.URL)

	t.Run("Success", func(t *testing.T) {
		players.On("FundPointsToPlayer", mock.Anything, &proto.PlayerIDPointRequest{PlayerID: 1, Points: 2.5}).Return(&proto.PlayerBalanceReply{Balance: 1.5}, nil)
		e.Request(http.MethodPut, "/fund/1").WithQuery("points", 2.5).
			Expect().Status(http.StatusOK).JSON().Number().Equal(1.5)
	})
	t.Run("PlayerParseIDOrPointsError", func(t *testing.T) {
		e.Request(http.MethodPut, "/fund/"+badPlayerID).WithQuery("points", 2.5).
			Expect().Status(http.StatusBadRequest)
	})
	t.Run("RPCError", func(t *testing.T) {
		players.On("FundPointsToPlayer", mock.Anything, &proto.PlayerIDPointRequest{PlayerID: 3, Points: 2.5}).Return(nil, errors.New("some error"))
		e.Request(http.MethodPut, "/fund/3").WithQuery("points", 2.5).
			Expect().Status(http.StatusBadRequest)
	})
}

func TestManagerRouter_takePointsHandler(t *testing.T) {
	players := &mocks.PlayerServiceClient{}
	gate := newGateway(players, mux.NewRouter())
	server := httptest.NewServer(gate)
	defer server.Close()
	e := httpexpect.New(t, server.URL)

	t.Run("Success", func(t *testing.T) {
		players.On("TakePointsFromPlayer", mock.Anything, &proto.PlayerIDPointRequest{PlayerID: 1, Points: 2.5}).Return(&proto.PlayerBalanceReply{Balance: 1.5}, nil)
		e.Request(http.MethodPut, "/take/1").WithQuery("points", 2.5).
			Expect().Status(http.StatusOK).JSON().Number().Equal(1.5)
	})
	t.Run("PlayerParseIDOrFloatError", func(t *testing.T) {
		e.Request(http.MethodPut, "/take/"+badPlayerID).WithQuery("points", 2.5).
			Expect().Status(http.StatusBadRequest)
	})
	t.Run("DBManagerError", func(t *testing.T) {
		players.On("TakePointsFromPlayer", mock.Anything, &proto.PlayerIDPointRequest{PlayerID: 3, Points: 2.5}).Return(nil, errors.New("some error"))
		e.Request(http.MethodPut, "/take/3").WithQuery("points", 2.5).
			Expect().Status(http.StatusBadRequest)
	})
}

func TestManagerRouter_RemovePlayer(t *testing.T) {
	players := &mocks.PlayerServiceClient{}
	gate := newGateway(players, mux.NewRouter())
	server := httptest.NewServer(gate)
	defer server.Close()
	e := httpexpect.New(t, server.URL)

	tt := []struct {
		name           string
		playerID       int
		returnError    error
		expectedStatus int
	}{
		{
			name:           "Success",
			playerID:       1,
			returnError:    nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "RPCError",
			playerID:       2,
			returnError:    errors.New("wrong id"),
			expectedStatus: http.StatusBadRequest,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			players.On("RemovePlayer", mock.Anything, &proto.PlayerIDRequest{PlayerID: int32(tc.playerID)}).Return(&proto.Nothing{}, tc.returnError)
			e.Request(http.MethodDelete, "/remove/"+strconv.Itoa(tc.playerID)).
				Expect().Status(tc.expectedStatus)
		})
	}
	t.Run("ParseIDError", func(t *testing.T) {
		e.Request(http.MethodDelete, "/remove/"+badPlayerID).
			Expect().Status(http.StatusBadRequest)
	})
}

func TestGetPlayerIDAndPoints(t *testing.T) {
	tt := []struct {
		name             string
		playerID         string
		points           string
		expectError      bool
		expectedPlayerID int
		expectedPoints   float32
	}{
		{name: "Success", playerID: "1", points: "1.5", expectError: false, expectedPlayerID: 1, expectedPoints: 1.5},
		{name: "BadPlayerID", playerID: badPlayerID, points: "1.5", expectError: true},
		{name: "WrongPoints", playerID: "2", points: "9876543210987654321098765432109876543210.91", expectError: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("", "", nil)
			req = mux.SetURLVars(req, map[string]string{
				"playerId": tc.playerID,
				"points":   tc.points,
			})
			playerID, points, err := getPlayerIDAndPoints(req)
			if !tc.expectError {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedPlayerID, playerID)
				assert.Equal(t, tc.expectedPoints, points)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
