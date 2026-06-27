package user

import (
	"SpotSync/internal/domain/user/dto"
	"fmt"
)

var ErrInvalidCredentials = fmt.Errorf("invalid email or password")

type service struct {
	repo       Repository
	// jwtService auth.JWTService
}

func NewService(repo Repository, ) *service {
	return &service{repo, }
}

func (s *service) CreateUser(req dto.CreateRequest) (*dto.Response, error) {

	user := User{
		Name:  req.Name,
		Email: req.Email,
	}

	// hash password and set to user.Password
	err := user.hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	response := dto.Response{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.String(),
	}

	return &response, nil

}