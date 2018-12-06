package client

import (
	"context"
	"io"
	"sync"

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

func (f *FlowClient) Run(ctx context.Context, requestc <-chan *flow.Request, opts ...grpc.CallOption) (chan *flow.Reply, chan error, error) {
	stream, err := f.pbClient.Run(ctx, opts...)
	if err != nil {
		return nil, nil, errors.Trace(err)
	}

	// Get server responses
	errorc := make(chan error)
	errWg := &sync.WaitGroup{}
	errWg.Add(2)
	replyc := make(chan *flow.Reply)
	go func() {
		defer errWg.Done()
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

	// Send requests
	go func() {
		defer errWg.Done()
		for request := range requestc {
			err := stream.Send(request)
			if err == io.EOF {
				return
			}
			if err != nil {
				errorc <- errors.Trace(err)
				return
			}
		}
		err := stream.CloseSend()
		if err != nil {
			errorc <- errors.Trace(err)
		}
	}()

	go func() {
		errWg.Wait()
		close(errorc)
	}()

	return replyc, errorc, nil
}
