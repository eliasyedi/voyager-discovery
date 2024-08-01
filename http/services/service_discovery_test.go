package services

import (
	"encoding/json"
	"sync"
	"testing"
	"voyager-discovery/http/store"
)


var (

    s *store.MapStore = store.NewInMemmoryStore(nil)
    d *DiscoveryServiceInMemmory = NewDiscoveryService(s)
)




func TestServiceRegistration(t *testing.T){
    value := store.RegisterEntry{
        ServiceCode: "api",
        Url: "localhost",
        ServiceName: "test",
        Ttl: 5, 
    }
    for i:=1 ;i<5 ; i++{
        go func(){
            d.Create(value)
        }()
    }
}


func TestConcurrencyWR(t *testing.T){
    t.Run("register",func(t *testing.T) {
        t.Parallel()    
        value := store.RegisterEntry{
            ServiceCode: "api",
            Url: "localhost",
            ServiceName: "test",
            Ttl: 5, 
        }
        for i:=1 ;i<5 ; i++{
            go func(){
                d.Create(value)
            }()
        }
    })
    t.Run("getAll",func(t *testing.T) {
        for i :=0 ; i<5 ; i++{
            data := d.GetAllRegistered()
            t.Log(data)
        }
    })
}

func TestGetAllRegistretionApis(t *testing.T){
    value := store.RegisterEntry{
        ServiceCode: "api",
        Url: "localhost",
        ServiceName: "test",
        Ttl: 5,
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
