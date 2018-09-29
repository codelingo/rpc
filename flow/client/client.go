package client

import (
	"context"
	"io"

	"github.com/codelingo/rpc/flow"
	"github.com/juju/errors"
	"google.golang.org/grpc"
)

// This package is responsible for all communication to the flow server.

type FlowClient struct {
	pbClient flow.FlowClient
}

func NewFlowClient(conn *grpc.ClientConn) *FlowClient {
	return &FlowClient{
		pbClient: flow.NewFlowClient(conn),
	}
}

func (f *FlowClient) Run(ctx context.Context, req *flow.Request) (chan *flow.Reply, chan error, error) {
	stream, err := f.pbClient.Run(ctx, req)
	if err != nil {
		return nil, nil, errors.Trace(err)
	}

	errorc := make(chan error)
	replyc := make(chan *flow.Reply)

	go func() {
		defer close(errorc)
		defer close(replyc)
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				errorc <- errors.Trace(err)
				return
			}

			if in != nil {
				replyc <- in
			}
		}
	}()

	return replyc, errorc, nil
}
