package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"workout/data"
)

type JSONPayload struct {
	Description string `bson:"description" json:"description"`
}

func (app *Config) AddWorkout(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Print("Incorrect payload or empty")
		app.errorJSON(w, errors.New("Incorrect payload"), http.StatusBadRequest)
		return
	}

	entry := data.WorkoutEntry{
		Description: fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-01"), requestPayload.Description),
	}

	err = app.Models.WorkoutEntry.Insert(entry)
	if err != nil {
		_ = app.errorJSON(w, err)
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Workout Created",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
