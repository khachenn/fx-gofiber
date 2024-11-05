package service

import (
	"sturdy-winner-api/config"
	"sturdy-winner-api/internal/api/dto"
	"sturdy-winner-api/internal/repository"

	"github.com/dollarsignteam/go-logger"
)

type UserService struct {
	logger     *logger.Logger
	repository *repository.Handler
	cfg        *config.APIConfig
}

func NewUserService(
	logger *logger.Logger,
	repository *repository.Handler,
	cfg *config.APIConfig,
) UserService {
	return UserService{
		logger:     logger,
		repository: repository,
		cfg:        cfg,
	}
}

func (svc UserService) Users() ([]dto.UserResponse, error) {
	userList, err := svc.repository.Users()
	if err != nil {
		return []dto.UserResponse{}, err
	}
	resp := []dto.UserResponse{}
	for _, user := range userList {
		resp = append(resp, dto.UserResponse{
			Email: user.Email,
			Name:  user.Name,
		})
	}
	return resp, nil
}
