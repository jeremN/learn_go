package main

import (
	"flag"
	"fmt"
	"learn_go/internal/app"
	"learn_go/internal/routes"
	"net/http"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "The port to listen on")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("Application created")

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler: r,
	}

	app.Logger.Printf("Server started on port %d", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

