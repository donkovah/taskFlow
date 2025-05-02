package service

import (
	"be/src/domain/interfaces"
	"be/src/domain/models"
	"be/src/infrastructure/config"
	"context"
)

type AuthService struct {
	service   interfaces.AuthService
	jwtSecret []byte
}

func NewAuthService(service interfaces.AuthService) *AuthService {
	return &AuthService{
		service:   service,
		jwtSecret: []byte(config.App.JWTSecret),
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*models.LoginResponse, error) {
	return s.service.Login(ctx, email, password)
}

func (s *AuthService) Register(ctx context.Context, user *models.User) (*models.User, error) {
	return s.service.Register(ctx, user)
}

func (s *AuthService) Logout(ctx context.Context, token string) error {
	return s.service.Logout(ctx, token)
}
