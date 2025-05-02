package persistence

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

// UserRepositoryImpl implements the UserRepository interface using GORM
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

// GetUsers retrieves all users from the database
func (r *UserRepositoryImpl) GetUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	result := r.db.WithContext(ctx).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetUser retrieves a single user by ID from the database
func (r *UserRepositoryImpl) GetUser(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).First(&user, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// CreateUser creates a new user in the database
func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	result := r.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// UpdateUser updates an existing user in the database
func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	result := r.db.WithContext(ctx).Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// DeleteUser removes a user from the database by ID
func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&models.User{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserByEmail retrieves a single user by email from the database
func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
