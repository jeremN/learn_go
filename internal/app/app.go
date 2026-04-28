package app

import (
	"database/sql"
	"fmt"
	"learn_go/internal/api"
	"learn_go/internal/store"
	"learn_go/migrations"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")

	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	workoutStore := store.NewPostgresWorkoutStore(pgDB)

	workoutHandler := api.NewWorkoutHandler(workoutStore, logger)

	app := &Application{
		Logger: logger,
		WorkoutHandler: workoutHandler,
		DB: pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: OK")
}