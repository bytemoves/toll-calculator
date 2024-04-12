package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"strconv"

	"net/http"

	"github.com/bytemoves/toll-calculator/types"
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
	svc = NewLogMiddleware(svc )
	go makeGRPCTransport(*grpcListenAddr,svc)
	makeHTTPTransport(*httpListenAddr,svc)
 
}
func makeGRPCTransport (listenAddr string,svc Aggregator)  error{
	fmt.Println("GRPC transport running on port",listenAddr)
	//make tcp listener
	ln , err := net.Listen("TCP",listenAddr)
	if  err != nil {
		return err
	}
	defer ln.Close()
	//grpc native server wit options
	server := grpc.NewServer([]grpc.ServerOption{}...)
	
	types.RegisterAggregatorServer(server,NEWAggregatoGRPCServer(svc))
	return server.Serve(ln)
}

func NEWGRPCServer(svc Aggregator) {
	panic("unimplemented")
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
