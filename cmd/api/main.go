package main

import (
	"flag"
	"log"
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
}
