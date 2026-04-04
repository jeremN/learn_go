package routes

import (
	"learn_go/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/health", app.HealthCheckHandler)
	router.Get("/workout/{id}", app.WorkoutHandler.HandleGetWorkoutsByID)
	
	router.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)

	return router
}