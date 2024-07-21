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
    //returns a all register services
    GetRegistered() store.RegisterEntry

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



//temp
var i int = 1;

func (d *DiscoveryServiceInMemmory) Create(registry store.RegisterEntry){
    log.Println("storing value with id"+ fmt.Sprint(i));
    d.store.Store(registry)
}


func (d *DiscoveryServiceInMemmory) GetRegistered(id uuid.UUID) store.RegisterEntry{
    data, err := d.store.Get(id)

    if err != nil {
       log.Println(err.Error()) 
    }

    return data

     
//todo 
}


