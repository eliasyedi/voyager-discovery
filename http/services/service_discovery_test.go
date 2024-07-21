package services

import (
	"testing"
	"voyager-discovery/http/store"
)







func TestServiceRegistration(t *testing.T){
    s := store.NewInMemmoryStore()
    d := NewDiscoveryService(s)
    value := store.RegisterEntry{
        ServiceCode: "api",
        Url: "localhost",
        ServiceName: "test",
    }
    for i:=1 ;i<5 ; i++{
        go func(){
            d.Create(value)
        }()
    }
}
