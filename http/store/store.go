package store

//theres two roads to take sync.Map o use a sync.RWMutex

import (
	"log"
	"sync"
)

//interface for storing implementation
type Store interface{
    //store registerEntry
    Store(key, value any)
    //gets registerEntry
    Get(key any) any
}


type RegisterEntry struct{
    ServiceCode     string
    Url             string
    ServiceName     string
}

//holds register map 
type RegisterStore map[any]any

//struct for storage synchronization
type MapStore struct{
    l sync.RWMutex
    store RegisterStore
}




func NewInMemmoryStore()*MapStore{
    return &MapStore{
        store: make(map[any]any),
    }
}



func (s *MapStore) Store(key , value any){
    s.l.Lock()
    defer s.l.Unlock()
    s.store[key]= value
    log.Printf("storing key value pair [%s], [%s]", key, value)
}



func (s *MapStore) Get(key any) any{
    s.l.RLock()
    defer s.l.RUnlock()
    if v, has := s.store[key]; has{
       log.Printf("got key value pair [%s], [%s]", key, v)
        return v
    }
    return ""
}
