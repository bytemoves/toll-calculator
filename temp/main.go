package main

import (
	"context"
	//"fmt"
	"log"
	"time"

	//"github.com/bytemoves/toll-calculator/aggregator/client"
	"github.com/bytemoves/toll-calculator/aggregator/client"
	"github.com/bytemoves/toll-calculator/types"
	//"google.golang.org/grpc"
)

func main () {
	c , err := client.NewGRPCClient(":3001")
	if err != nil{
		log.Fatal(err)

	}
	
	if err := c.Aggregate(context.Background(), &types.AggregateRequest{
		ObuID: 1,
		Value: 58.55,
		Unix:  time.Now().UnixNano(),
		}); err != nil {
		log.Fatal(err)
	}
}





// if _, err := c.Aggregate(context.Background(), &types.AggregateRequest{
//     ObuID: 1,
//     Value: 58.55,
//     Unix:  time.Now().UnixNano(),
// 	}); err != nil {
//     log.Fatal(err)
// }