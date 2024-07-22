package store

//theres two roads to take sync.Map o use a sync.RWMutex

import (
	"errors"
	"log"
	"sync"

	"github.com/google/uuid"
)

//interface for storing implementation
type Store interface{
    //store registerEntry
    Store(value RegisterEntry)
    //gets registerEntry
    Get(key any) (RegisterEntry, error)
    //gets all registered apis
    GetAll() []RegisterEntry 
}


type RegisterEntry struct{
    //maybe theres more data to be put here 
    ID              uuid.UUID `json:"id"` //only for response
    ServiceCode     string `json:"serviceCode"`
    Url             string `json:"url"`
    ServiceName     string `json:"serviceName"`
}
//trying out 
//type RegisterEntryResponse struct{
    //maybe theres more data to be put here 
    //ID              uuid.UUID `json:"id"`
    //ServiceCode     string `json:"serviceCode"`
    //Url             string `json:"url"`
    //ServiceName     string `json:"serviceName"`
//}

//holds register map 
type RegisterStore map[any]RegisterEntry

//struct for storage synchronization
type MapStore struct{
    l sync.RWMutex
    store RegisterStore
}


func NewInMemmoryStore()*MapStore{
    return &MapStore{
        store: make(map[any]RegisterEntry),
    }
}


//for development
func (s *MapStore) populate(){
    value := RegisterEntry{
        ServiceCode: "api",
        Url: "localhost",
        ServiceName: "test",
    }
    for i:=1 ;i<5 ; i++{
        s.l.Lock()
        defer s.l.Unlock()
        s.store[i]= value
    }

}



func (s *MapStore) Store(value RegisterEntry){
    s.l.Lock()
    defer s.l.Unlock()
    key, err := uuid.NewUUID()
    if err != nil{
        log.Printf(err.Error())
        return
    }
    value.ID = key
    s.store[key]= value
    log.Printf("storing key value pair [%s], [%s]\n", key, value)
}


func (s *MapStore) Get(key any) (RegisterEntry, error){
    s.l.RLock()
    defer s.l.RUnlock()
    if v, has := s.store[key]; has{
       log.Printf("got key value pair [%s], [%s] \n", key, v)
        return v, nil
    }
    return RegisterEntry{}, errors.New("no data entry found");
}


func (s *MapStore) GetAll() []RegisterEntry{
    s.l.Lock()
    //fix size at the moment for slice no need for append maybe
    registryList := make([]RegisterEntry, 0,len(s.store))
    for _ , value := range s.store{
        registryList=append(registryList, value)
    }
    s.l.Unlock()
    return registryList
}
