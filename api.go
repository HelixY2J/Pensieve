package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
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

	fromEmail := EnvString("FROM_EMAIL", "")
	if fromEmail == "" {
		log.Fatal("FROM_EMAIL must be set")
	}

	mailer := NewSmtp2goMailer(fromEmail)
	service := NewService(store, mailer)
	service.RegisterRoutes(subrouter)

	log.Println("starting the API server at ", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
