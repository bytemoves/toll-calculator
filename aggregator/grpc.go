package main

import (
	"context"

	"github.com/bytemoves/toll-calculator/types"
)

type GRPCAggregatorServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NEWAggregatoGRPCServer (svc Aggregator) *GRPCAggregatorServer{
	return &GRPCAggregatorServer{
		svc: svc,
	}
}
//transport layer
 //json  _. 
//business layer _> busines layer type(main type everyone need to convert to)

func (s *GRPCAggregatorServer) Aggregate (ctx context.Context, req *types.AggregateRequest) (*types.None,error){
	distance  := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix: req.Unix,
	}
	return &types.None{},s.svc.AggregateDistance(distance)

}