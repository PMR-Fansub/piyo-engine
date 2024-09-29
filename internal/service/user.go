package service

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	v1 "piyo-engine/api/v1"
	"piyo-engine/internal/constant"
	"piyo-engine/internal/model"
	"piyo-engine/internal/repository"
)

type UserService interface {
	Register(ctx context.Context, req *v1.RegisterRequest) error
	Login(ctx context.Context, req *v1.LoginRequest) (string, error)
	GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error
}

func NewUserService(
	service *Service,
	userRepo repository.UserRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest) error {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil && !errors.Is(err, v1.ErrNotFound) {
		return v1.ErrInternalServerError
	}
	if user != nil {
		return v1.ErrUsernameAlreadyUse
	}

	user, err = s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return v1.ErrInternalServerError
	} else if user != nil {
		return v1.ErrEmailAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Generate user ID
	userID, err := s.sid.GenString()
	if err != nil {
		return err
	}
	user = &model.User{
		UserID:   userID,
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Status:   constant.UserStatusNormal,
	}

	err = s.tm.Transaction(
		ctx, func(ctx context.Context) error {
			// Create a user
			if err = s.userRepo.Create(ctx, user); err != nil {
				return err
			}
			return nil
		},
	)
	return err
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest) (string, error) {
	var user *model.User
	var err error
	if req.Username != "" {
		user, err = s.userRepo.GetByUsername(ctx, req.Username)
	} else if req.Email != "" {
		user, err = s.userRepo.GetByEmail(ctx, req.Email)
	}
	if err != nil || user == nil {
		return "", v1.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	token, err := s.jwt.GenToken(user.UserID, time.Now().Add(time.Hour*24*30))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetProfileResponseData{
		UserId:    user.UserID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	if req.Email != "" && req.Email != user.Email {
		otherUser, err := s.userRepo.GetByEmail(ctx, req.Email)
		if err != nil {
			return v1.ErrInternalServerError
		} else if otherUser != nil && otherUser.UserID != user.UserID {
			return v1.ErrEmailAlreadyUse
		}
		user.Email = req.Email
	}
	if req.Nickname != "" && req.Nickname != user.Nickname {
		user.Nickname = req.Nickname
	}

	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
