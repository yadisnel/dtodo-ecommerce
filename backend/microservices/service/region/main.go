package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/yadisnel/kupiti/backend/microservices/service/region/handler"
	regionpb "github.com/yadisnel/kupiti/backend/microservices/service/region/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("region"),
		service.Version("latest"),
	)

	// Register handler
	regionpb.RegisterRegionHandler(srv.Server(), new(handler.Region))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
