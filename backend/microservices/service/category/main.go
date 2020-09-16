package main

import (
	"github.com/yadisnel/kupiti/backend/microservices/service/category/handler"
	categorypb "github.com/yadisnel/kupiti/backend/microservices/service/category/proto"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("category"),
		service.Version("latest"),
	)

	// Register handler
	categorypb.RegisterCategoryHandler(srv.Server(), new(handler.Category))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
