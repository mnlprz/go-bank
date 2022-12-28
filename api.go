package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s *APIServer) Run() {

	r := chi.NewRouter()
	r.Route("/account", func(r chi.Router) {
		r.Get("/", s.handleGetAccounts)
		r.Post("/", s.handleCreateAccount)
	})
	r.Route("/account/{id}", func(r chi.Router) {
		r.Get("/", s.handleGetAccountsByID)
		r.Delete("/", s.handleDeleteAccountsByID)
	})
	r.Route("/transfer", func(r chi.Router) {
		r.Get("/", s.handleTransfer)
	})
	http.ListenAndServe(s.listenAddr, r)
}

func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) {

	accounts, err := s.store.GetAccounts()
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
	}
	WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountsByID(w http.ResponseWriter, r *http.Request) {

	idSrt := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idSrt)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
	}
	account, err := s.store.GetAccountByID(id)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
	}
	WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccountsByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
	}
	err = s.store.DeleteAccountByID(id)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
	}
	WriteJSON(w, http.StatusOK, nil)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) {

	createAccountRequest := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountRequest); err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
	}
	defer r.Body.Close()

	account := NewAccount(createAccountRequest.FirstName, createAccountRequest.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
	}

	WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) {

	transferRequest := new(TransferRequest)

	if err := json.NewDecoder(r.Body).Decode(transferRequest); err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
	}
	defer r.Body.Close()

	WriteJSON(w, http.StatusOK, transferRequest)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {

	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}
