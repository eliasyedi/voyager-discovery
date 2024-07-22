package services

import (
	"encoding/json"
	"sync"
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



func TestGetAllRegistretionApis(t *testing.T){

    s := store.NewInMemmoryStore()
    d := NewDiscoveryService(s)
    value := store.RegisterEntry{
        ServiceCode: "api",
        Url: "localhost",
        ServiceName: "test",
    }
    var wg sync.WaitGroup
    wg.Add(4)
    for i:=1 ;i<5 ; i++{
        go func(){
            d.Create(value)
            wg.Done()
        }()
    }
    wg.Wait()
    data := d.GetAllRegistered()

    t.Log(len(data))
    t.Log(data)
    json,_ := json.Marshal(data)
    t.Log(string(json))

    

}
