package handlers

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)



type RegisterEntryRequest struct{
    //maybe theres more data to be put here 
    ServiceCode     string `json:"serviceCode"`
    Url             string `json:"url"`
    ServiceName     string `json:"serviceName"`
    Ttl             *int64 `json:"ttl,ommitempty"`
}



type TimeWrapper time.Time;

type ListRegistyEntryResponse struct{
    Registries  []RegistryEntryResponse `json:"registries"`
}

type RegistryEntryResponse struct{
    //maybe theres more data to be put here 
    ID              uuid.UUID `json:"id"` //only for response
    ServiceCode     string `json:"serviceCode"`
    Url             string `json:"url"`
    ServiceName     string `json:"serviceName"`
    Expiration      TimeWrapper `json:"expiration"`  
}



func (t TimeWrapper) MarshalJSON() ([]byte, error){    

    stamp := fmt.Sprintf("\"%s\"",time.Time(t))

    return []byte(stamp), nil

}
