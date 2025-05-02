package service

import (
	"be/src/domain/interfaces"
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo  repository.UserRepository
	jwtSecret []byte
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string) interfaces.AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: []byte(jwtSecret),
	}
}

func (s *authService) Login(ctx context.Context, email, password string) (*models.LoginResponse, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		fmt.Println("no user found")
		return nil, errors.New("invalid credentials")
	}

	// Print the hashed password and plain password for debugging
	fmt.Printf("Stored hashed password: %s\n", user.Password)
	fmt.Printf("Provided plain password: %s\n", password)

	// Compare the hashed password from DB with the plain password from request
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Printf("Password comparison failed with error: %v\n", err)
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.JWTClaims{
		UserID:    user.ID.String(), // Convert UUID to string
		Email:     user.Email,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   int(time.Until(expirationTime).Seconds()),
	}, nil
}

func (s *authService) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (s *authService) Register(ctx context.Context, user *models.User) (*models.User, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash the password
	hashedPassword, err := s.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	// Create the user
	createdUser, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// Don't return the password hash
	createdUser.Password = ""
	return createdUser, nil
}

func (s *authService) Logout(ctx context.Context, token string) error {
	// Implement logout logic here
	// This might involve invalidating the token, updating a blacklist, etc.
	return nil
}
