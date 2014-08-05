package service

import "github.com/opentarock/frontend-user-management/server/service/proto_user"

type UserService interface {
	RegisterUser(user *proto_user.User, redirectURI string) (*proto_user.RegisterResponse, error)
}
