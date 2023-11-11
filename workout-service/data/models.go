package data

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type Models struct {
	WorkoutEntry WorkoutEntry
}

func New(mongo *mongo.Client) Models {
	client = mongo

	return Models{
		WorkoutEntry: WorkoutEntry{},
	}
}
