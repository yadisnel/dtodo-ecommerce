package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/yadisnel/kupiti/backend/microservices/service/shop/handler"
	shoppb "github.com/yadisnel/kupiti/backend/microservices/service/shop/proto"

)

func main() {
	// Create service
	srv := service.New(
		service.Name("shop"),
		service.Version("latest"),
	)
	// Register handler
	shoppb.RegisterShopHandler(srv.Server(), new(handler.Shop))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
