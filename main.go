package main

import (
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error to start application: cannot load environment variables: " + err.Error())
	}
}

func main() {
	println("Book lender starting!")
	app := SetupApplication()

	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		port = "8080"
	}

	server := &http.Server{
		Handler:      app.Handler,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	println("Setting up Book Lender server on port:", port)
	if err := server.ListenAndServe(); err != nil {
		panic("Server error: " + err.Error())
	}
}
