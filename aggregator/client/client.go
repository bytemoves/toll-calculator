package client

import (
	"context"

	"github.com/bytemoves/toll-calculator/types"
)

type Client interface {
	Aggregate (context.Context, *types.AggregateRequest) error 

}

