package data

import (
	"context"
	"errors"
	"harsh/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore struct {
	Collection *mongo.Collection
}

// Initializing the struct
func NewUserStore(db *mongo.Database) *UserStore {
	return &UserStore{
		Collection: db.Collection("users"),
	}
}

func (s *UserStore) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	_, err := s.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserStore) GetUser(ctx context.Context, id string) (*models.User, error) {
	var user models.User

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	filter := bson.M{"_id": objectID}
	// finding user
	if err := s.Collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStore) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	filter := bson.M{"email": email}
	if err := s.Collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
