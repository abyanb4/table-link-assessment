package server

import (
	"context"
	"fmt"

	"github.com/abyanb4/table-link/resources/db"

	pb "github.com/abyanb4/table-link/proto"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	// the logic should moved to modules

	user, err := db.GetUser(req.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	token := "random-token"

	return &pb.LoginResponse{
		Status:  true,
		Message: "Login successful",
		Data: &pb.LoginData{
			AccessToken: token,
		},
	}, nil
}
