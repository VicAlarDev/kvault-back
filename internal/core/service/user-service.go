package service

import (
	"context"

	"github.com/VicAlarDev/kvault-back/internal/core/domain"
	"github.com/VicAlarDev/kvault-back/internal/core/port"
	"github.com/VicAlarDev/kvault-back/internal/core/util"
)

type UserService struct {
	userRepo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	if user == nil {
		return nil, domain.ErrInternal
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	createdUser, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return createdUser, nil
}

/* func (s *UserService) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]domain.User, error) {
	users, err := s.userRepo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if user == nil {
		return nil, domain.ErrInternal
	}

	existingUser, err := s.userRepo.GetUserByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	emptyData := user.Name == "" && user.Email == "" && user.Password == ""
	sameData := existingUser.Name == user.Name && existingUser.Email == user.Email && existingUser.Password == user.Password

	if emptyData || sameData {
		return nil, domain.ErrNoUpdatedData
	}

	var hashedPassword string

	if user.Password != "" {
		hashedPassword, err = util.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
	}

	user.Password = hashedPassword

	updatedUser, err := s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return updatedUser, nil
} */
