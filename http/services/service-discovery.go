package services

import "voyager-discovery/http/store"

//interface for implementing diferent types of storing data
type DiscoveryService interface{
    //creates a register for a service
    Create()
    //returns a all register services
    GetRegistered()

}


//dont really know about naming conventions in golang
type DiscoveryServiceInMemmory struct{
    //this should hold whatever it needs to use

    Store store.Store 
}



//inyect whats needed for 
func NewDiscoveryService(store store.Store) *DiscoveryServiceInMemmory {
    return &DiscoveryServiceInMemmory{
        Store: store,
    }
}





func (d *DiscoveryServiceInMemmory) Create(){
//todo 
}


func (d *DiscoveryServiceInMemmory) GetRegistered(){
//todo 
}


