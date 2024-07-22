package services

import (
	"fmt"
	"log"
	"voyager-discovery/http/store"

	"github.com/google/uuid"
)

//interface for implementing diferent types of storing data
type DiscoveryService interface{
    //creates a register for a service
    Create(registry store.RegisterEntry)
    //returns registered service with id
    GetRegisteredById(id uuid.UUID) store.RegisterEntry
    //return all registered services
    GetAllRegistered() []store.RegisterEntry
}


//dont really know about naming conventions in golang
type DiscoveryServiceInMemmory struct{
    //this should hold whatever it needs to use

    store store.Store 
}



//inyect whats needed for 
func NewDiscoveryService(store store.Store) *DiscoveryServiceInMemmory {
    return &DiscoveryServiceInMemmory{
        store: store,
    }
}




func (d *DiscoveryServiceInMemmory) Create(registry store.RegisterEntry){
    log.Println("storing value {%v} "+ fmt.Sprint(registry));
    d.store.Store(registry)
}


func (d *DiscoveryServiceInMemmory) GetRegisteredById(id uuid.UUID) store.RegisterEntry{
    data, err := d.store.Get(id)

    if err != nil {
       log.Println(err.Error()) 
    }

    return data

     
}


func (d *DiscoveryServiceInMemmory) GetAllRegistered() []store.RegisterEntry{
    return d.store.GetAll() 
}
