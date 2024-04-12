package main

import "github.com/bytemoves/toll-calculator/types"

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

func (s *GRPCAggregatorServer) AggregateDistance ( req *types.AggregateRequest) error{
	distance  := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix: req.Unix,
	}
	return s.svc.AggregateDistance(distance)

}