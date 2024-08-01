package store

//theres two roads to take sync.Map o use a sync.RWMutex

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

type StoreConfig struct{
    //max size. note: it's only recommended for in memmory, by default should be -1 
    maxSize int
    //max ammuont of time that a service lives if for control and registration 
    maxTtlServices time.Duration
}

func NewDefaultStoreConfig()*StoreConfig{
    return &StoreConfig{
        // -1 will not control size
        maxSize: -1,
        //max time that a services may use for registration
        maxTtlServices: time.Duration(10*time.Second),
    } 
}


func ConfigService(options ...func(*StoreConfig)) *StoreConfig{
    var config *StoreConfig = NewDefaultStoreConfig()
    //config := NewDefaultConfig()
    for _ , opt := range options {
        opt(config)
    }
    return config;
}

//interface for storing implementation
type Store interface{
    //store registerEntry
    Store(value RegisterEntry)
    //gets registerEntry
    Get(key any) (RegisterEntry, error)
    //gets all registered apis
    GetAll() []RegisterEntry 
    
}

//todo there needs to be different structs for response maybe, or wrap it in another struct with extension for response or storage
type RegisterEntry struct{
    //maybe theres more data to be put here 
    Id              uuid.UUID 
    ServiceCode     string 
    Url             string 
    ServiceName     string 
    Ttl             int64 
    //response
    Expiration      time.Time

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
    options *StoreConfig
    l sync.RWMutex
    store RegisterStore
}


func NewInMemmoryStore(options *StoreConfig)*MapStore{
    if options != nil{
        options = NewDefaultStoreConfig()
    }
    return &MapStore{
        store: make(map[any]RegisterEntry),
        options: options,
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
    expiration := time.Now().Add(time.Duration(time.Second * time.Duration(value.Ttl)));
    if err != nil{
        log.Printf(err.Error())
        return
    }
    value.Id = key
    value.Expiration = expiration
    s.store[key]= value
    log.Printf("storing key value pair [%s], [%v]\n", key, value)
}


func (s *MapStore) Get(key any) (RegisterEntry, error){
    s.l.Lock()
    defer s.l.Unlock()
    if v, has := s.store[key]; has{
       log.Printf("got key value pair [%s], [%v] \n", key, v)
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



