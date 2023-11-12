package main

import (
	"net/http"
	"workout/data"
)

type JSONPayload struct {
	Description string `bson:"description" json:"description"`
}

func (app *Config) AddWorkout(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPayload

	_ = app.readJSON(w, r, &requestPayload)

	entry := data.WorkoutEntry{
		Description: requestPayload.Description,
	}

	err := app.Models.WorkoutEntry.Insert(entry)
	if err != nil {
		_ = app.errorJSON(w, err)
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Workout Created",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
