package handler

import (
	"context"
	categorypb "github.com/yadisnel/kupiti/backend/microservices/service/category/proto"
)

type Category struct{}

func (c Category) Call(ctx context.Context, request *categorypb.Request, response *categorypb.Response) error {
	panic("implement me")
}

func (c Category) Stream(ctx context.Context, request *categorypb.StreamingRequest, stream categorypb.Category_StreamStream) error {
	panic("implement me")
}

func (c Category) PingPong(ctx context.Context, stream categorypb.Category_PingPongStream) error {
	panic("implement me")
}


