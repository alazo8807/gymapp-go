package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Database = "gymapp"
const WorkoutCollection = "workouts"

// Insert adds a new workout into the MongoDB collection
func (w *WorkoutEntry) Insert(workout WorkoutEntry) (string, error) {
	collection := client.Database(Database).Collection(WorkoutCollection)

	workout.ID = primitive.NewObjectID().Hex()
	workout.Exercises = []Exercise{}
	workout.CreatedAt = time.Now()

	resp, err := collection.InsertOne(context.TODO(), workout)
	if err != nil {
		log.Println("Error adding workout:", err)
		return "", err
	}

	insertedId := resp.InsertedID.(string)

	return insertedId, nil
}

// GetAll retrieves all workouts from the MongoDB collection
func (w *WorkoutEntry) GetAll() ([]WorkoutEntry, error) {
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
func (w *WorkoutEntry) GetWorkoutByID(workoutID string) (*WorkoutEntry, error) {
	collection := client.Database(Database).Collection(WorkoutCollection)

	var workout WorkoutEntry
	err := collection.FindOne(context.TODO(), bson.M{"_id": workoutID}).Decode(&workout)
	if err != nil {
		log.Println("Error getting workout by ID:", err)
		return nil, err
	}

	return &workout, nil
}

// UpdateWorkout updates a workout in the MongoDB collection
func (w *WorkoutEntry) UpdateWorkout(workoutID string, updatedWorkout WorkoutEntry) error {
	collection := client.Database(Database).Collection(WorkoutCollection)

	filter := bson.M{"_id": workoutID}
	update := bson.M{"$set": updatedWorkout}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Error updating workout:", err)
		return err
	}

	return nil
}

// DeleteWorkout removes a workout from the MongoDB collection
func (w *WorkoutEntry) DeleteWorkout(workoutID string) error {
	collection := client.Database(Database).Collection(WorkoutCollection)

	filter := bson.M{"_id": workoutID}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Error deleting workout:", err)
		return err
	}

	return nil
}

// AddExerciseToWorkoutHandler inserts a new excercise for workout with id workoutID
func (w *WorkoutEntry) AddExerciseToWorkout(workoutID string, exercise Exercise) error {
	// Retrieve the existing workout from the database
	existingWorkout, err := w.GetWorkoutByID(workoutID)
	if err != nil {
		log.Println("Error retrieving workout for that ID:", err)
		return err
	}

	// Add a unique ID to the exercise
	exercise.ID = primitive.NewObjectID().Hex()

	// Append the new exercise to the existing workout
	existingWorkout.Exercises = append(existingWorkout.Exercises, exercise)

	// Update the workout in the database
	if err := w.UpdateWorkout(workoutID, *existingWorkout); err != nil {
		log.Println("Error updating workout", err)
		return err
	}

	return nil
}

// AddSet inserts a new set for workout excercise
func (w *WorkoutEntry) AddSet(workoutID string, excerciseID string, set Set) error {
	// Retrieve the existing workout from the database
	existingWorkout, err := w.GetWorkoutByID(workoutID)
	if err != nil {
		log.Println("Error retrieving workout for that ID:", err)
		return err
	}

	fmt.Println(existingWorkout)
	// Find the target excercise by excerciseID
	var targetExcercise *Exercise
	for i, excercise := range existingWorkout.Exercises {
		if excercise.ID == excerciseID {
			targetExcercise = &existingWorkout.Exercises[i]
			break
		}
	}

	if targetExcercise == nil {
		return fmt.Errorf("Could not find an excercise corresponding to excerciseID: %s", excerciseID)
	}

	// Add a unique ID to the set
	set.ID = primitive.NewObjectID().Hex()
	targetExcercise.Sets = append(targetExcercise.Sets, set)

	// Update the workout in the database
	if err := w.UpdateWorkout(workoutID, *existingWorkout); err != nil {
		log.Println("Error updating workout", err)
		return err
	}

	return nil
}
