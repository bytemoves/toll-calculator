package client

import (
	"github.com/bytemoves/toll-calculator/types"
	"google.golang.org/grpc"
)



type GRPCClient struct{
	Endpoints string
	types.AggregatorClient

}

func NewGRPCClient(endpoint string) (*GRPCClient, error){
	conn , err := grpc.Dial(endpoint,grpc.WithInsecure())
	if err != nil{
		return nil, err

	}
	c := types.NewAggregatorClient(conn)
	
	return &GRPCClient{
		Endpoints: endpoint,
		AggregatorClient: c,
	}, nil

}