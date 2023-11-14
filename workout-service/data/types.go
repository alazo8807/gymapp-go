package data

import "time"

// Workout represents a workout consisting of a list of exercises
type WorkoutEntry struct {
	ID          string     `bson:"_id,omitempty" json:"id,omitempty"`
	Description string     `bson:"description" json:"description"`
	Exercises   []Exercise `bson:"excercies" json:"exercises"`
	CreatedAt   time.Time  `bson:"created_at" json:"created_at"`
}

// Set represents a set in a workout with a weight and number of reps
type Set struct {
	ID     string `bson:"_id,omitempty" json:"id,omitempty"`
	Weight int    `bson:"weight" json:"weight"`
	Reps   int    `bson:"reps" json:"reps"`
}

// Exercise represents an exercise in a workout with an ID, optional machine ID, name, and a list of sets
type Exercise struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	MachineID string `bson:"machine_id,omitempty" json:"machine_id,omitempty"`
	Name      string `bson:"name" json:"name"`
	Sets      []Set  `bson:"sets" json:"sets"`
}
