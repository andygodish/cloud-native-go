package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type eventServiceHandler struct {} 

func (eh *eventServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request)
func (eh *eventServiceHandler) AllEventHandler(w http.ResponseWriter, r * http.Request)
func (eh *eventServiceHandler) NewEventHandler(w http.ResponseWriter, r *http.Request)

func ServeAPI(endpoint string) error {
	handler := &eventServiceHandler{}
	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods("GET").Path("/{SearchCriteria}/{Search}").HandlerFunc(handler.FindEventHandler)
	eventsRouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsRouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)
	return http.ListenAndServe(endpoint, r)
}