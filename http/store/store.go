package store

//theres two roads to take sync.Map o use a sync.RWMutex

import (
	"errors"
	"log"
	"sync"
)

//interface for storing implementation
type Store interface{
    //store registerEntry
    Store(key any , value RegisterEntry)
    //gets registerEntry
    Get(key any) (RegisterEntry, error)
}


type RegisterEntry struct{
    //maybe theres more data to be put here 
    ServiceCode     string
    Url             string
    ServiceName     string
}

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



func (s *MapStore) Store(key any , value RegisterEntry){
    s.l.Lock()
    defer s.l.Unlock()
    s.store[key]= value
    log.Printf("storing key value pair [%s], [%s]", key, value)
}


func (s *MapStore) Get(key any) (RegisterEntry, error){
    s.l.RLock()
    defer s.l.RUnlock()
    if v, has := s.store[key]; has{
       log.Printf("got key value pair [%s], [%s]", key, v)
        return v, nil
    }
    return RegisterEntry{}, errors.New("no data entry found");
}
