package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"

	"net/http"

	"github.com/bytemoves/toll-calculator/types"
)


func main () {
	listenAddr := flag.String("listenaddr",":3000","the listening address of the HTTP transport ")
	flag.Parse()
	
	var (
		store  = NewMemoryStore()
		svc = NewInvoiceAggreagator(store)
	)
	svc = NewLogMiddleware(svc )
	makeHTTPTransport(*listenAddr,svc)
 
}

func makeHTTPTransport (listenAddr string , svc Aggregator) {
	fmt.Println("HTTP transport running on port",listenAddr)
	http.HandleFunc("/aggregate",handleAggregate(svc))
	http.HandleFunc("/invoice",handleGetInvoice(svc))

	http.ListenAndServe(listenAddr,nil)

}

func handleGetInvoice(svc Aggregator) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
	values,ok := r.URL.Query()["obu"]
	if !ok {
		WriteJSON(w,http.StatusBadRequest, map[string]string{"error": "missing OBU ID"})
		return
		
	}
	
	obuID, err:= strconv.Atoi(values[0])
	if err != nil{
		WriteJSON(w,http.StatusBadRequest,map[string]string{"error": "invalid obu id"})
		return
	}
	 _ = obuID
}
}


func  handleAggregate(svc Aggregator) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := svc.AggregateDistance(distance); err !=nil{
			WriteJSON(w,http.StatusInternalServerError,map[string]string{"error": err.Error()})
		}

	}
}

func WriteJSON(w http.ResponseWriter,status int ,v any) error{
	w.WriteHeader(status)
	w.Header().Add("content-type","application/json")  
	return json.NewEncoder(w).Encode(v)

}
