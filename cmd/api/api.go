package api

import (
	"database/sql"
	"log"
	"net/http"

	weather "github.com/DroppedHard/weather-forecast-api/cmd/service/weather"
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

	weatherHandler := weather.NewHandler()
	weatherHandler.RegisterRoutes(subrouter) 
	
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}