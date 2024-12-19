package api

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	weather "github.com/DroppedHard/weather-forecast-api/cmd/service/weather"
	"github.com/DroppedHard/weather-forecast-api/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr	string
	db 		*sql.DB
}

func NewAPIServer (addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr:addr,
		db: db, 
	}
}

func (s *APIServer) Run() error{
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins(strings.Split(config.Envs.ORIGINS_ALLOWED, ","))
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	weatherHandler := weather.NewHandler()
	weatherHandler.RegisterRoutes(subrouter) 
	
	log.Println("Listening on", s.addr)
	
	return http.ListenAndServe(s.addr, handlers.CORS(originsOk, headersOk, methodsOk)(router))
}