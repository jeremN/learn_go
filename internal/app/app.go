package app

import (
	"fmt"
	"learn_go/internal/api"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)


	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger: logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

func (a *Application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: OK")
}