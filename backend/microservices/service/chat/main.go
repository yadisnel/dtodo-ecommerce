package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/yadisnel/kupiti/backend/microservices/service/chat/handler"
	chatpb "github.com/yadisnel/kupiti/backend/microservices/service/chat/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("chat"),
		service.Version("latest"),
	)

	// Register handler
	chatpb.RegisterChatHandler(srv.Server(), new(handler.Chat))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
