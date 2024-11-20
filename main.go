package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/ajswetz/Chirpy/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	// Load .env file
	godotenv.Load()

	// Establish connection to database
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Unable to open connection to database: %v\n", err)
	}
	dbQueries := database.New(db)

	// Get platform from .env
	platform := os.Getenv("PLATFORM")

	// Create server state tracker
	srvState := &apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries,
		platform:       platform,
	}

	// Create new HTTP request multiplexer
	mux := http.NewServeMux()

	////// REGISTER HANDLERS //////

	/// APP ///

	// Register file server handler wrapped by middlewareMetricsInc() to track number of hits
	fileServerHandler := http.StripPrefix("/app/", http.FileServer(http.Dir(".")))
	mux.Handle("/app/", srvState.middlewareMetricsInc(fileServerHandler))

	/// API ///

	// Register server readiness handler
	mux.HandleFunc("GET /api/healthz", readinessHandler)

	// Register create user handler
	mux.HandleFunc("POST /api/users", srvState.createUserHandler)

	// Register create chirp handler
	mux.HandleFunc("POST /api/chirps", srvState.createChirpHandler)

	// Register get all chirps handler
	mux.HandleFunc("GET /api/chirps", srvState.getAllChirpsHandler)

	// Register get single chirp handler
	mux.HandleFunc("GET /api/chirps/{chirpID}", srvState.getSingleChirpHandler)

	/// ADMIN ///

	// Register metrics handler
	mux.HandleFunc("GET /admin/metrics", srvState.metricsHandler)

	// Register reset handler
	mux.HandleFunc("POST /admin/reset", srvState.resetHandler)

	///////////////////////////////

	// Define http.Server listening on port 8080
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start listening on port 8080 and serving responses
	log.Println("Starting Chirpy . . . ")
	err = server.ListenAndServe()
	if err != nil {
		log.Printf("An error occurred: %v\n", err)
	}

}
