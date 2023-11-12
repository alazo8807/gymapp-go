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

// GetWorkoutByID retrieves a workout by ID from the MongoDB collection
func GetWorkoutByID(client *mongo.Client, workoutID string) (*WorkoutEntry, error) {
	collection := client.Database(Database).Collection(WorkoutCollection)

	objectID, err := primitive.ObjectIDFromHex(workoutID)
	if err != nil {
		log.Println("Error converting workout ID:", err)
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	var workout WorkoutEntry
	err = collection.FindOne(context.TODO(), filter).Decode(&workout)
	if err != nil {
		log.Println("Error getting workout by ID:", err)
		return nil, err
	}

	return &workout, nil
}

// UpdateWorkout updates a workout in the MongoDB collection
func UpdateWorkout(client *mongo.Client, workoutID string, updatedWorkout WorkoutEntry) error {
	collection := client.Database(Database).Collection(WorkoutCollection)

	objectID, err := primitive.ObjectIDFromHex(workoutID)
	if err != nil {
		log.Println("Error converting workout ID:", err)
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updatedWorkout}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Error updating workout:", err)
		return err
	}

	return nil
}

// DeleteWorkout deletes a workout from the MongoDB collection
func DeleteWorkout(client *mongo.Client, workoutID string) error {
	collection := client.Database(Database).Collection(WorkoutCollection)

	objectID, err := primitive.ObjectIDFromHex(workoutID)
	if err != nil {
		log.Println("Error converting workout ID:", err)
		return err
	}

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Error deleting workout:", err)
		return err
	}

	return nil
}
