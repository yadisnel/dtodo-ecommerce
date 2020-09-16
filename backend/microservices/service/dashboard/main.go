package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/yadisnel/kupiti/backend/microservices/service/dashboard/handler"
	dashboardpb "github.com/yadisnel/kupiti/backend/microservices/service/dashboard/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("dashboard"),
		service.Version("latest"),
	)

	// Register handler
	dashboardpb.RegisterDashboardHandler(srv.Server(), new(handler.Dashboard))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
