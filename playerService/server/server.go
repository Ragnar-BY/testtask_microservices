package server

import (
	"context"
	"log"
	"net"

	"github.com/Ragnar-BY/testtask_microservices/playerService/manager"
	pb "github.com/Ragnar-BY/testtask_microservices/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server is server.
type Server struct {
	manager manager.Manager
}

// NewServer creates new Server instance.
func NewServer(mngr manager.Manager) *Server {
	s := &Server{manager: mngr}
	return s
}

// Start starts pkg with addr.
func (s *Server) Start(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterPlayerServiceServer(srv, s)
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// CreateNewPlayer create new player in DB.
func (s *Server) CreateNewPlayer(ctx context.Context, in *pb.CreatePlayerRequest) (*pb.CreatePlayerReply, error) {
	id, err := s.manager.CreateNewPlayer(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePlayerReply{Id: int32(id)}, nil
}

// GetPlayerPoints returns player balance, if player exists.
func (s *Server) GetPlayerPoints(ctx context.Context, in *pb.PlayerIDRequest) (*pb.PlayerBalanceReply, error) {
	balance, err := s.manager.GetPlayerPoints(ctx, int(in.PlayerID))
	if err != nil {
		return nil, err
	}
	return &pb.PlayerBalanceReply{Balance: balance}, nil
}

// TakePointsFromPlayer take points from player and returns new balance.
func (s *Server) TakePointsFromPlayer(ctx context.Context, in *pb.PlayerIDPointRequest) (*pb.PlayerBalanceReply, error) {
	balance, err := s.manager.TakePointsFromPlayer(ctx, int(in.PlayerID), in.Points)
	if err != nil {
		return nil, err
	}
	return &pb.PlayerBalanceReply{Balance: balance}, nil
}

// FundPointsToPlayer funds points to player and returns new balance.
func (s *Server) FundPointsToPlayer(ctx context.Context, in *pb.PlayerIDPointRequest) (*pb.PlayerBalanceReply, error) {
	balance, err := s.manager.FundPointsToPlayer(ctx, int(in.PlayerID), in.Points)
	if err != nil {
		return nil, err
	}
	return &pb.PlayerBalanceReply{Balance: balance}, nil
}

// RemovePlayer removes player, if player exists.
func (s *Server) RemovePlayer(ctx context.Context, in *pb.PlayerIDRequest) (*pb.Nothing, error) {
	err := s.manager.RemovePlayer(ctx, int(in.PlayerID))
	if err != nil {
		return nil, err
	}
	return &pb.Nothing{}, nil
}
