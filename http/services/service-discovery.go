package services




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

    
}



//inyect whats needed for 
func NewDiscoveryService() *DiscoveryServiceInMemmory {
    return &DiscoveryServiceInMemmory{

    }
}





func (d *DiscoveryServiceInMemmory) Create(){
//todo 
}


func (d *DiscoveryServiceInMemmory) GetRegistered(){
//todo 
}


