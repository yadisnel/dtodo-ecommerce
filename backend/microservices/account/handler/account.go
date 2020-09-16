package handler

import (
	"context"
	log "github.com/micro/micro/v3/service/logger"
	accountpb "github.com/yadisnel/kupiti/backend/microservices/account/proto"
)

type Account struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Account) Call(ctx context.Context, req *accountpb.Request, rsp *accountpb.Response) error {
	log.Info("Received Account.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Account) Stream(ctx context.Context, req *accountpb.StreamingRequest, stream accountpb.Account_StreamStream) error {
	log.Infof("Received Account.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&accountpb.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Account) PingPong(ctx context.Context, stream accountpb.Account_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&accountpb.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
