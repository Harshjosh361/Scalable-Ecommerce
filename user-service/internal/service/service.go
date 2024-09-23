package service

import (
	"context"
	"errors"
	"harsh/internal/data"
	"harsh/internal/models"
	"harsh/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	userStore *data.UserStore
}

func NewUserService(userStore *data.UserStore) *UserService {
	return &UserService{
		userStore: userStore,
	}
}

func (s *UserService) GetUser(ctx context.Context, id string) (*models.User, error) {
	return s.userStore.GetUser(ctx, id)
}
func (s *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	// check if the user is registered
	exist, err := s.userStore.GetUserByEmail(ctx, user.Email)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	} else if exist != nil {
		return nil, errors.New("user already registered")
	}

	// hash the password
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashed
	return s.userStore.CreateUser(ctx, user)
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userStore.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	// check if password is correct
	err = utils.CheckPasswordHash(password, user.Password)
	if err != nil {
		return "", errors.New("invalid password")
	}

	// generate jwt token
	token, err := utils.GenerateJWT(user.Id.Hex())
	if err != nil {
		return "", errors.New("unable to generate jwt")
	}
	return token, nil

}
