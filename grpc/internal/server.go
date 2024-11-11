package internal

import (
	"context"
	"fmt"
	pb "grpc-service/protos/grpc-service"
	"math/rand/v2"
	"strconv"
	"sync"
)

type User struct {
	name string
	typ  pb.UserType
}

type UserService struct {
	pb.UnimplementedUserServiceServer

	users map[int]User
	mu    *sync.RWMutex
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[int]User),
		mu:    &sync.RWMutex{},
	}
}

func (s *UserService) GetUser(ctx context.Context, r *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	reqId := r.GetId()
	if reqId == "" {
		return nil, fmt.Errorf("malformed request")
	}

	id, err := strconv.Atoi(reqId)
	if err != nil {
		return nil, err
	}

	s.mu.RLock()
	user, ok := s.users[id]
	s.mu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	fmt.Println("requested users with id: ", id)

	return &pb.GetUserResponse{
		User: &pb.UserRead{Id: r.Id, Name: user.name, Type: user.typ},
	}, nil

}

func (s *UserService) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	grpcUser := r.GetUser()
	if grpcUser == nil {
		return nil, fmt.Errorf("malformed request")
	}

	id := rand.IntN(10000)

	user := User{
		name: grpcUser.Name,
		typ:  grpcUser.Type,
	}

	s.mu.Lock()
	s.users[id] = user
	s.mu.Unlock()

	fmt.Println("users updated:", s.users)

	return &pb.CreateUserResponse{
		Id: strconv.Itoa(id),
	}, nil
}
