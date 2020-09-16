package handler

import (
	"context"
	log "github.com/micro/micro/v3/service/logger"
	shoppb "github.com/yadisnel/kupiti/backend/microservices/service/shop/proto"
)

type Shop struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Shop) Call(ctx context.Context, req *shoppb.Request, rsp *shoppb.Response) error {
	log.Info("Received Shop.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Shop) Stream(ctx context.Context, req *shoppb.StreamingRequest, stream shoppb.Shop_StreamStream) error {
	log.Infof("Received Shop.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&shoppb.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Shop) PingPong(ctx context.Context, stream shoppb.Shop_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&shoppb.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
