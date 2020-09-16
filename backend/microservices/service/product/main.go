package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/yadisnel/kupiti/backend/microservices/service/product/handler"
	productpb "github.com/yadisnel/kupiti/backend/microservices/service/product/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("product"),
		service.Version("latest"),
	)

	// Register handler
	productpb.RegisterProductHandler(srv.Server(), new(handler.Product))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
