package handlers

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
	"go.mongodb.org/mongo-driver/mongo"
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
	fmt.Println("Called HandleGetUsers")
}

func (h *User) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called HandleCreateUser")
}
