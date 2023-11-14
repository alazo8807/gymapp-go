package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"workout/data"
)

type WorkoutJSONPayload struct {
	Description string `json:"description"`
}

type ExcerciseJSONPayload struct {
	WorkoutID string `json:"workout_id"`
	MachineID string `json:"machine_id"`
	Name      string `json:"name"`
}

type SetJSONPayload struct {
	WorkoutID   string `json:"workout_id"`
	ExcerciseID string `json:"excercise_id"`
	Weight      int    `bson:"weight" json:"weight"`
	Reps        int    `bson:"reps" json:"reps"`
}

// AddExcerciseToWorkout is the route handler used adding an excercise to a workout.
func (app *Config) GetWorkouts(w http.ResponseWriter, r *http.Request) {
	res, err := app.Models.WorkoutEntry.GetAll()
	if err != nil {
		_ = app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error: false,
		Data:  res,
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}

// AddWorkout is the route handler used for POST /workout requests.
func (app *Config) AddWorkout(w http.ResponseWriter, r *http.Request) {
	var requestPayload WorkoutJSONPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Print("Incorrect payload or empty")
		app.errorJSON(w, errors.New("Incorrect payload"), http.StatusBadRequest)
		return
	}

	entry := data.WorkoutEntry{
		Description: fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-01"), requestPayload.Description),
	}

	insertedId, err := app.Models.WorkoutEntry.Insert(entry)
	if err != nil {
		_ = app.errorJSON(w, err)
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Workout Created",
		Data:    insertedId,
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}

// AddExcerciseToWorkout is the route handler used adding an excercise to a workout.
func (app *Config) AddExcerciseToWorkout(w http.ResponseWriter, r *http.Request) {
	var requestPayload ExcerciseJSONPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Print("Incorrect payload or empty")
		app.errorJSON(w, errors.New("Incorrect payload"), http.StatusBadRequest)
		return
	}

	entry := data.Exercise{
		MachineID: requestPayload.MachineID,
		Name:      requestPayload.Name,
		Sets:      []data.Set{},
	}

	err = app.Models.WorkoutEntry.AddExerciseToWorkout(requestPayload.WorkoutID, entry)
	if err != nil {
		_ = app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Excercise added",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}

// AddSet is the route handler used adding a set to a workout's excercise.
func (app *Config) AddSet(w http.ResponseWriter, r *http.Request) {
	var requestPayload SetJSONPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Print("Incorrect payload or empty")
		app.errorJSON(w, errors.New("Incorrect payload"), http.StatusBadRequest)
		return
	}

	entry := data.Set{
		Weight: requestPayload.Weight,
		Reps:   requestPayload.Reps,
	}

	err = app.Models.WorkoutEntry.AddSet(requestPayload.WorkoutID, requestPayload.ExcerciseID, entry)
	if err != nil {
		_ = app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Set added",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
