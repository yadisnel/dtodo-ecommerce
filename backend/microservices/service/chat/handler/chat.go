package handler

import (
	"context"
	chatpb "github.com/yadisnel/kupiti/backend/microservices/service/chat/proto"
)

type Chat struct{}

func (chat Chat) WriteMessage(ctx context.Context, request *chatpb.RequestWriteMessage, response *chatpb.ResponseWriteMessage) error {
	panic("implement me")
}

func (chat Chat) RemoveMessage(ctx context.Context, request *chatpb.RequestRemoveMessage, response *chatpb.ResponseRemoveMessage) error {
	panic("implement me")
}








