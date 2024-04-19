package main

import (
	//"context"
	//"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	//"time"

	"net/http"

	//"github.com/bytemoves/toll-calculator/aggregator/client"
	//"github.com/bytemoves/toll-calculator/aggregator/client"
	"github.com/bytemoves/toll-calculator/types"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)


func main () {
	httpListenAddr := flag.String("listenaddr",":3000","the listening address of the HTTP transport ")
	grpcListenAddr := flag.String("grpc addr",":3001","the listening address of the grpc transport ")
	flag.Parse()
	
	var (
		store  = NewMemoryStore()
		svc = NewInvoiceAggreagator(store)
	)

	svc = NewMetricsMiddleware(svc)
	svc = NewLogMiddleware(svc)

	svc = NewLogMiddleware(svc )
	go func ()  {
		log.Fatal(makeGRPCTransport(*grpcListenAddr,svc))
		
	} ()
	
	log.Fatal(makeHTTPTransport(*httpListenAddr,svc))
	
}


func makeGRPCTransport (listenAddr string,svc Aggregator)  error{
	fmt.Println("GRPC transport running on port",listenAddr)
	//make tcp listener
	ln , err := net.Listen("tcp",listenAddr)
	if  err != nil {
		return err
	}
	defer func () {
		fmt.Println("stopping grpc transport ")
		ln.Close()
	} ()
	//grpc native server wit options
	server := grpc.NewServer([]grpc.ServerOption{}...)
	
	types.RegisterAggregatorServer(server,NEWAggregatoGRPCServer(svc))
	return server.Serve(ln)
}

func NEWGRPCServer(svc Aggregator) {
	panic("unimplemented")
}

func makeHTTPTransport (listenAddr string , svc Aggregator) error {
	fmt.Println("HTTP transport running on port",listenAddr)
	http.HandleFunc("/aggregate",handleAggregate(svc))
	http.HandleFunc("/invoice",handleGetInvoice(svc))
	http.Handle("/metrics",promhttp.Handler())

	return http.ListenAndServe(listenAddr,nil)

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
	 invoice , err := svc.CalculateInvoice(obuID)
	  if err != nil{
		WriteJSON(w,http.StatusInternalServerError,map[string]string{"error": err.Error()})

		return
	  }

	  WriteJSON(w,http.StatusOK,invoice)
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
