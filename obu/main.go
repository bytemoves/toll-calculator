package main

import (
	
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/bytemoves/toll-calculator/types"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"

var SendInterval = time.Second * 5


func genLatLong () (float64 , float64){
	return genCord(),genCord()
}

func genCord () float64 {
	n := float64 (rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f

}


func main (){

	obuIDS := generateOBUIDS(20)

	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil{
		log.Fatal(err)
	}
	
	for {
		for i := 0; i < len(obuIDS); i++ {
		lat , long := genLatLong()
			data := types.OBUData{
				OBUID: obuIDS[i],
				Lat: lat,
				Long: long,
			}
			if err := conn.WriteJSON(data); err != nil{
				log.Fatal(err)
			}
			
		}
		
		time.Sleep(SendInterval)
	}

}

func generateOBUIDS(n int) []int{

	ids := make([]int,n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
		
	}
	return ids
}

func init (){
	
	rand.Seed(time.Now().UnixNano())

}