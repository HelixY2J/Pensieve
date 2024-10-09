package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	store := NewStore(s.db)
	service := NewService(store)
	service.RegisterRoutes(subrouter)

	log.Println("starting the API server at ", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
