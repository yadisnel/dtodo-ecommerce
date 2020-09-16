package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	dashboardpb "github.com/yadisnel/kupiti/backend/microservices/service/dashboard/proto"
)

type Dashboard struct{}

func (dashboard Dashboard) AddCategory(ctx context.Context, request *dashboardpb.RequestAddCategory, response *dashboardpb.ResponseAddCategory) error {
	panic("implement me")
}

func (dashboard Dashboard) RemoveCategory(ctx context.Context, request *dashboardpb.RequestRemoveCategory, response *dashboardpb.ResponseRemoveCategory) error {
	panic("implement me")
}

func (dashboard Dashboard) AddSubCategory(ctx context.Context, request *dashboardpb.RequestAddSubCategory, response *dashboardpb.ResponseAddSubCategory) error {
	panic("implement me")
}

func (dashboard Dashboard) RemoveSubCategory(ctx context.Context, request *dashboardpb.RequestRemoveSubCategory, response *dashboardpb.ResponseRemoveSubCategory) error {
	panic("implement me")
}

func (dashboard Dashboard) RetrieveAllCategories(ctx context.Context, request *empty.Empty, stream dashboardpb.Dashboard_RetrieveAllCategoriesStream) error {
	panic("implement me")
}

func (dashboard Dashboard) AddRegion(ctx context.Context, request *dashboardpb.RequestAddRegion, response *dashboardpb.ResponseAddRegion) error {
	panic("implement me")
}

func (dashboard Dashboard) RemoveRegion(ctx context.Context, request *dashboardpb.RequestRemoveRegion, response *dashboardpb.ResponseRemoveRegion) error {
	panic("implement me")
}

func (dashboard Dashboard) UpdateRegion(ctx context.Context, request *dashboardpb.RequestUpdateRegion, response *dashboardpb.ResponseUpdateRegion) error {
	panic("implement me")
}

func (dashboard Dashboard) RetrieveAllRegions(ctx context.Context, request *empty.Empty, stream dashboardpb.Dashboard_RetrieveAllRegionsStream) error {
	panic("implement me")
}





