package pln

import (
	"context"
	"strings"

	"github.com/saifoelloh/listrik/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	ReadUser(params QueryReadUser) (*[]entities.UserPLN, error)
}

type repository struct {
	Collection *mongo.Collection
}

type QueryReadUser struct {
	Name    string
	SortBy  string
	OrderBy string
	Page    uint16
	Show    uint16
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) ReadUser(params QueryReadUser) (*[]entities.UserPLN, error) {
	var users []entities.UserPLN
	filter := bson.D{}
	opt := options.Find()

	if len(strings.TrimSpace(params.Name)) > 0 {
		filter = append(filter, bson.E{
			Key: "name",
			Value: bson.D{
				{
					Key: "$regex",
					Value: primitive.Regex{
						Pattern: ".*" + params.Name + ".*",
						Options: "i",
					},
				},
			},
		})
	}

	if len(strings.TrimSpace(params.OrderBy)) > 0 {
		sortBy := 1
		if len(strings.TrimSpace(params.SortBy)) > 0 {
			if params.SortBy == "asc" {
				sortBy = 1
			} else {
				sortBy = -1
			}
		}
		opt.SetSort(bson.D{{Key: params.OrderBy, Value: sortBy}})
	}

	show := 10
	if params.Show > 0 {
		opt.SetLimit(int64(params.Show))
		show = int(params.Show)
	}

	if params.Page >= 0 {
		skippedData := params.Page * uint16(show)
		opt.SetSkip(int64(skippedData))
	}

	cursor, err := r.Collection.Find(context.Background(), filter, opt)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user entities.UserPLN
		_ = cursor.Decode(&user)
		users = append(users, user)
	}
	return &users, nil
}
