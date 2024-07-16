package main

import (
	"log"
	"net/http"
	common "voyager-discovery/commons"
	handlers "voyager-discovery/http"

	"github.com/gorilla/mux"
	env "github.com/joho/godotenv"
)





func main(){
    
    mux:= mux.NewRouter()
    //middlewares are going to be execute in the order that they were set
   
    err := env.Load()
    if err!= nil{
        log.Fatal("could not load env file");
    }
    url:= common.EnvStringOrDef("URL", ":3000");


    handler:= handlers.NewDiscoveryHandler();
    handler.RegisterHandlers(mux);


    mux.HandleFunc("/", HandlerNotFound)
    err = http.ListenAndServe(url, mux);
    if err!=nil{
        log.Fatal("no pudo arrancar ", err); 
    }
    log.Printf("aqui estoy \n");


}


//maybe its not required yet
func HandlerNotFound(w http.ResponseWriter, r *http.Request){
    log.Println("handler not found");
    /*
    w.WriteHeader(http.StatusMethodNotAllowed);
    */
    //prefer this over the top one
    //todo -> should make a Global error handler
    http.Error(w,http.StatusText(http.StatusMethodNotAllowed),http.StatusMethodNotAllowed )
}

