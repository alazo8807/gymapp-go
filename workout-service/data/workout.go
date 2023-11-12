package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Database = "gymapp"
const WorkoutCollection = "workouts"

// Insert adds a new workout into the MongoDB collection
func (l *WorkoutEntry) Insert(workout WorkoutEntry) error {
	collection := client.Database(Database).Collection(WorkoutCollection)

	workout.ID = primitive.NewObjectID().Hex()
	workout.Exercises = []Exercise{}
	workout.CreatedAt = time.Now()

	_, err := collection.InsertOne(context.TODO(), workout)
	if err != nil {
		log.Println("Error adding workout:", err)
		return err
	}

	return nil
}

// GetAll retrieves all workouts from the MongoDB collection
func GetAll(client *mongo.Client) ([]WorkoutEntry, error) {
	collection := client.Database(Database).Collection(WorkoutCollection)

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("Error getting workouts:", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var workouts []WorkoutEntry
	err = cursor.All(context.TODO(), &workouts)
	if err != nil {
		log.Println("Error decoding workouts:", err)
		return nil, err
	}

	return workouts, nil
}
