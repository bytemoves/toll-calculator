package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	//"github.com/docker/docker/client"
	"github.com/bytemoves/toll-calculator/aggregator/client"
	"github.com/sirupsen/logrus"
)
 
type apiFunc func (w http.ResponseWriter , r *http.Request) error 
	



func main () {
	listenAddr := flag.String("listenAddr",":6000","the listen addr of the http server")
	aggregatorServiceAddr := flag.String("aggServiceAddr","http://localhost:3000","the listen address of the aggragator service")
	flag.Parse()
	var (
		client = client.NewHTTPClient(*aggregatorServiceAddr) //agg service
		invHandler = newInvoiceHandler(client) 
)

	
	http.HandleFunc("/invoice",makeAPIFunc(invHandler.handleGetInvoice))
	logrus.Infof("gateway HTTP server running on port %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr,nil))

}

type InvoiceHandler struct {
	client client.Client

}
func newInvoiceHandler (c client.Client) *InvoiceHandler{
	return &InvoiceHandler{
		client: c,
	}
}

func (h *InvoiceHandler) handleGetInvoice(w http.ResponseWriter , r *http.Request) error{
	fmt.Println("hitting the get invoice inside the gate way")
	//acces to agg client
	inv , err := h.client.GetInvoice(context.Background(),44)
	if err != nil{
		return err

	}

	return WriteJSON(w, http.StatusOK, inv)

}

func WriteJSON (w http.ResponseWriter,code int, v any) error {
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}


func makeAPIFunc (fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func (start time.Time)  {
			logrus.WithFields(logrus.Fields{
				"took": time.Since(start),
				"uri": r.RequestURI,

			}).Info("REQ :: ")
			
		}(time.Now())
		if err := fn(w,r); err != nil{
			WriteJSON(w,http.StatusInternalServerError,map[string]string{"error": err.Error()})
		}
	}
}
