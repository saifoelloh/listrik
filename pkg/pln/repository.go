package pln

import (
	"context"
	"fmt"

	"github.com/saifoelloh/listrik/api/presenter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	ReadUser() (*[]presenter.UserPLN, error)
}

type repository struct {
	Collection *mongo.Collection
}

// NewRepo is the single instance repo that is being created.
func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

// ReadBook is a mongo repository that helps to fetch books
func (r *repository) ReadUser() (*[]presenter.UserPLN, error) {
	var users []presenter.UserPLN
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user presenter.UserPLN
		_ = cursor.Decode(&user)
		fmt.Println(user)
		users = append(users, user)
	}
	return &users, nil
}
