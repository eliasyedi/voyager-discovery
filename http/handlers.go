package handlers

import (
	"io"
	"log"
	"net/http"
	"voyager-discovery/http/services"

	"github.com/gorilla/mux"
)

type handler struct {
	//for inyection of services, etc
    discoveryService services.DiscoveryService
}

const (
	DISCOVERYPREFIX = "/discovery"
)

// pass things for inyection
func NewDiscoveryHandler(discoveryService services.DiscoveryService) *handler {
	return &handler{
        discoveryService: discoveryService,
    }
}

func (h *handler) RegisterHandlers(router *mux.Router) {

	subRouter := router.PathPrefix(DISCOVERYPREFIX).Subrouter()
	subRouter.HandleFunc("/register", h.dummy).Methods("POST")
	subRouter.HandleFunc("/unregister", h.dummy).Methods("POST")
	subRouter.HandleFunc("/registered", h.dummy).Methods("GET")
}

func (h *handler) dummy(w http.ResponseWriter, r *http.Request) {
	//h.AuthService.BeforeSessionCreate()
	log.Println("working!")
}

func (h *handler) HanldlePostUnregister(w http.ResponseWriter, r *http.Request) {
	log.Println("handle post unregister")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("could not deserialize request body")
	}
	log.Println(body)
}

func (h *handler) HandlePostRegister(w http.ResponseWriter, r *http.Request) {
	log.Println("handle post register")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("could not deserialize request body")
	}
	log.Println(body)
}

func (h *handler) HandleGetRegistered(w http.ResponseWriter, r *http.Request) {
	log.Println("handle get registered services")
}
