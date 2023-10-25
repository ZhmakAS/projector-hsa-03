package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"projector-test-app/internal/models"
)

const UserCollections = "users"

type User interface {
	Save(ctx context.Context, user *models.User) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
}

type user struct {
	conn *mongo.Client
}

func NewUser(conn *mongo.Client) User {
	return &user{
		conn: conn,
	}
}

func (u *user) Save(ctx context.Context, user *models.User) (*models.User, error) {
	coll := u.conn.Database("db").Collection(UserCollections)

	_, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *user) GetAll(ctx context.Context) ([]models.User, error) {
	cursor, err := u.conn.Database("db").Collection(UserCollections).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
