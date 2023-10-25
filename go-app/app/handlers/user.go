package handlers

import (
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/elastic/go-elasticsearch/v7"
	"go.mongodb.org/mongo-driver/mongo"

	"projector-test-app/internal/models"
	"projector-test-app/internal/repository"
)

type User struct {
	mongoClient   *mongo.Client
	elasticClient *elasticsearch.Client
}

func NewUser(mongoClient *mongo.Client, elasticClient *elasticsearch.Client) *User {
	return &User{
		mongoClient:   mongoClient,
		elasticClient: elasticClient,
	}
}

func (h *User) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	userRepo := repository.NewUser(h.mongoClient)
	if _, err := userRepo.GetAll(r.Context()); err != nil {
		panic(err)
	}

	w.Write([]byte("Success - OK\n"))
}

func (h *User) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	gofakeit.Struct(&newUser)

	userRepo := repository.NewUser(h.mongoClient)
	if _, err := userRepo.Save(r.Context(), &newUser); err != nil {
		panic(err)
	}

	w.Write([]byte("Success - OK\n"))
}
