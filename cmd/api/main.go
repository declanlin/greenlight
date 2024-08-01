package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Declare a string to specify the application version number. TO BE REPLACED IN BUILD FILE.
const version = "1.0.0"

// Define a config struct to hold all the configuration settings for our application.
// - port: the network port the server will listen on
// - env: the name of the current operating environment for the application (development, staging, production, etc.)
// These configuration settings will be read in from command-line settings when the application starts.
type config struct {
	port int
	env  string
}

// Define an application struct to hold the dependencies for HTTP handlers, helpers, and middleware.

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// Read the value of the port and env command-line flags into the config struct.
	// Use port 4000 and "development" as the default value for these config fields.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	// Parse the flags from the command line.
	flag.Parse()

	// Initialize a new logger to write messages to the standard output stream, prefixed with the current date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Declare an instance of the application struct.
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Declare an new servemux and add a /v1/healthcheck route which dispatches requests
	// to the healthcheckHandler.
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// Declare a new http.Server with some sensible timeout settings.
	// The server will listen on the specified port and uses the servemux as the handler for request/response handling.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server.
	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
