package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"projector-test-app/internal/models"
)

const UserCollections = "users"

type User interface {
	Save(user *models.User) (*models.User, error)
	GetAll() ([]models.User, error)
}

type user struct {
	conn *mongo.Client
}

func NewAccount(conn *mongo.Client) User {
	return &user{
		conn: conn,
	}
}

func (u *user) Save(user *models.User) (*models.User, error) {
	coll := u.conn.Database("db").Collection(UserCollections)

	_, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *user) GetAll() ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}
