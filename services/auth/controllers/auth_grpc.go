package controllers

import (
	"errors"
	pb "github.com/jaax2707/ToDoGorm/services/auth/protobuf"
	"github.com/patrickmn/go-cache"
	"golang.org/x/net/context"
)

// AuthorizationMiddleware represents struct of cache
type Server struct {
	cache *cache.Cache
}

// NewCacheMiddleware return AuthorizationMiddleware Object
func NewServer(Cache *cache.Cache) *Server {
	return &Server{cache: Cache}
}

// GetToken is GRPC method witch take token from request
// validated it and return response with error message
func (s *Server) GetToken(c context.Context, req *pb.Request) (*pb.Response, error) {
	t, found := s.cache.Get(req.Token)

	if !found {
		return &pb.Response{}, errors.New("token not found")
	}

	if t != req.Token {
		return &pb.Response{}, errors.New("token is not valid")
	}
	return &pb.Response{}, nil
}
