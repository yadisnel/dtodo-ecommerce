package handler

import (
	"context"
	accountpb "github.com/yadisnel/kupiti/backend/microservices/service/account/proto"
)

type Account struct{}

func (account Account) Token(ctx context.Context, request *accountpb.RequestToken, response *accountpb.ResponseToken) error {
	panic("implement me")
}

func (account Account) Info(ctx context.Context, request *accountpb.RequestInfo, response *accountpb.ResponseInfo) error {
	panic("implement me")
}

func (account Account) UpdateAvatar(ctx context.Context, request *accountpb.RequestUpdateAvatar, response *accountpb.ResponseUpdateAvatar) error {
	panic("implement me")
}

func (account Account) EmqxAuthUser(ctx context.Context, request *accountpb.RequestEmqxAuthUser, response *accountpb.ResponseEmqxAuthUser) error {
	panic("implement me")
}

func (account Account) EmqxAuthAdmin(ctx context.Context, request *accountpb.RequestEmqxAuthAdmin, response *accountpb.ResponseEmqxAuthAdmin) error {
	panic("implement me")
}








