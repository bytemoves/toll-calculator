package client

import (
	"context"

	"github.com/bytemoves/toll-calculator/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)



type GRPCClient struct{
	Endpoints string
	client types.AggregatorClient

}

func NewGRPCClient(endpoint string) (*GRPCClient, error){
	conn , err := grpc.Dial(endpoint,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		return nil, err

	}
	c := types.NewAggregatorClient(conn)
	
	return &GRPCClient{
		Endpoints: endpoint,
		client: c,
	}, nil

}

func (c *GRPCClient) Aggregate(ctx context.Context, req *types.AggregateRequest) error{
	_,err := c.client.Aggregate(ctx, req)
	return err

	
}