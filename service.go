package main

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Service struct {
	store Storage
}

func NewService(store Storage) *Service {
	return &Service{store: store}
}

func (s *Service) RegisterRoutes(router *mux.Router) {

	// endpoints - parse kindle file
	// send daily insights from cloud

	router.HandleFunc("/user/{userID}/parse-kindle-file", s.handleParseKindleFile).Methods("POST")
	router.HandleFunc("/cloud/send-daily-insights", s.handleParseKindleFile).Methods("POST")
}

func (s *Service) handleParseKindleFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	file, _, err := r.FormFile("file")
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Error in parsing the file: %v", err))
		return
	}
	defer file.Close()

	raw, err := parseKindleFileExtract(file)

	if err != nil {
		WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Error in parsing the file: %v", err))
		return
	}
	userIDint, _ := strconv.Atoi(userID)
	if err := s.createDataFromRawBook(raw, userIDint); err != nil {
		WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error creating data from raw book: %v", err))
		return
	}

	WriteJSON(w, http.StatusOK, "file is parsed successfully")

}

func (s *Service) handleSendDailyInsights(w http.ResponseWriter, r *http.Request) {

}

func parseKindleFileExtract(file multipart.File) (*RawExtractBook, error) {
	decoder := json.NewDecoder(file)

	raw := new(RawExtractBook)
	if err := decoder.Decode(raw); err != nil {
		return nil, err
	}
	return raw, nil
}

func (s *Service) createDataFromRawBook(raw *RawExtractBook, userID string) error {
	return nil
}
