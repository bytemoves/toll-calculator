package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"net/http"

	"github.com/bytemoves/toll-calculator/types"
)


func main () {
	listenAddr := flag.String("listenaddr",":3000","the listening address of the HTTP transport ")
	flag.Parse()
	store  := NewMemoryStore()
	var (
		svc = NewInvoiceAggreagator(store)
	)
	makeHTTPTransport(*listenAddr,svc)
 
}

func makeHTTPTransport (listenAddr string , svc Aggregator) {
	fmt.Println("HTTP transport running on port",listenAddr)
	http.HandleFunc("/aggregate",handleAggregate(svc))
	http.ListenAndServe(listenAddr,nil)

}

func  handleAggregate(_ Aggregator) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil{
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	}
}


