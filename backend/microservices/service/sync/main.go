package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/yadisnel/kupiti/backend/microservices/service/sync/handler"
	syncpb "github.com/yadisnel/kupiti/backend/microservices/service/sync/proto"

)

func main() {
	// Create service
	srv := service.New(
		service.Name("sync"),
		service.Version("latest"),
	)

	// Register handler
	syncpb.RegisterSyncHandler(srv.Server(), new(handler.Sync))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
