package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"voyager-discovery/http/services"
	"voyager-discovery/http/store"

	"github.com/google/uuid"
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
	subRouter.HandleFunc("/register", h.HandlePostRegister).Methods("POST")
	subRouter.HandleFunc("/unregister", h.dummy).Methods("DELETE")
	subRouter.HandleFunc("/registered/{id}", h.HandleGetRegisteredById).Methods("GET")
	subRouter.HandleFunc("/all-registered", h.HandleGetAllRegistered).Methods("GET")
}

func (h *handler) dummy(w http.ResponseWriter, r *http.Request) {
	//h.AuthService.BeforeSessionCreate()
	log.Println("working!")
}

func (h *handler) HanldlePostUnregister(w http.ResponseWriter, r *http.Request) {
	log.Println("handle post unregister")

	body, err := io.ReadAll(r.Body)
	if err != nil {
        //todo logic for errror handling 
		log.Fatalf("could not deserialize request body")
	}
	log.Println(body)
}

func (h *handler) HandlePostRegister(w http.ResponseWriter, r *http.Request) {
	log.Println("handle post register")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("could not deserialize request body")
        //todo logic for errror handling 
	}
    bodyStr := &RegisterEntryRequest{}
    err = json.Unmarshal(body, &bodyStr)
	if err != nil {
		log.Fatalf("could not deserialize request body")
        //todo logic for errror handling 
        
	}

    reqBod := store.RegisterEntry{
        ServiceCode: bodyStr.ServiceCode,
        Url: bodyStr.Url,
        ServiceName: bodyStr.ServiceName,
        Ttl: int64(*bodyStr.Ttl),
    }

    
    h.discoveryService.Create(reqBod)
}

func (h *handler) HandleGetRegisteredById(w http.ResponseWriter, r *http.Request) {
	log.Println("handle get registered services")
    //
    params := mux.Vars(r)
    id :=  params["id"]
    log.Println(id)
    //no need for pathUnescape, its allready unescaped
    //id,err := url.PathUnescape(id)
//    if err !=nil {
 //       log.Println("bad request" )
        //todo logic for errror handling 
  //  }
    idSlice := []byte(id)

    log.Println(idSlice)
    if id == "" {
        log.Println("bad request" )
        //todo logic for errror handling 
    }
    h.discoveryService.GetRegisteredById(uuid.UUID([]byte(id)))
}

func (h *handler) HandleGetAllRegistered(w http.ResponseWriter, r *http.Request) {
	log.Println("handle get registered services")
    entries := h.discoveryService.GetAllRegistered();
    entriesResponse := make([]RegistryEntryResponse, len(entries))
    for _ , data := range entries{
        entry := RegistryEntryResponse{
            ID: data.Id, 
            ServiceCode: data.ServiceCode,
            Url: data.Url,
            ServiceName: data.ServiceName,
            Expiration: TimeWrapper(data.Expiration),

        }
        entriesResponse = append(entriesResponse, entry)
    }
    
    jsonRes, err := json.Marshal(entriesResponse);
    if err != nil {
        //todo logic for errror handling 
        log.Println("theres been an error")
        log.Println(err.Error())
    }
    w.Write(jsonRes);
    return; 
}
