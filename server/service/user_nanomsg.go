package service

import (
	"log"
	"time"

	proto "code.google.com/p/gogoprotobuf/proto"
	nmsg "github.com/op/go-nanomsg"

	"github.com/opentarock/frontend-user-management/server/service/proto_user"
)

// UserServiceNanomsg is an implementation of UserService using nanomsg for message
// transport and protobuf for message serialization.
type UserServiceNanomsg struct {
	userServiceSocket *nmsg.ReqSocket
}

func NewUserServiceNanomsg() (*UserServiceNanomsg, error) {
	socket, err := nmsg.NewReqSocket()
	if err != nil {
		return nil, err
	}
	// Timeout is set because we can't wait for the messages forever to keep
	// the frontend responsive.
	const timeout = 100 * time.Millisecond
	err = socket.SetSendTimeout(timeout)
	err = socket.SetRecvTimeout(timeout)
	if err != nil {
		return nil, err
	}

	// TODO: Make address and port a parameter
	endpoint, err := socket.Connect("tcp://127.0.0.1:6001")
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to endpoint: %s", endpoint.Address)

	return &UserServiceNanomsg{
		userServiceSocket: socket,
	}, nil
}

func (s *UserServiceNanomsg) RegisterUser(
	user *proto_user.User, redirectURI string) (*proto_user.RegisterResponse, error) {

	registerUser := &proto_user.RegisterUser{
		User:        user,
		RedirectUri: proto.String(redirectURI),
	}
	data, err := proto.Marshal(registerUser)
	// If there is a problem with marshalling there is no way to recover and
	// indicates a serious bug.
	if err != nil {
		log.Fatalf("Error marshaling RegisterUser: %s", err)
	}
	_, err = s.userServiceSocket.Send(data, 0)
	if err != nil {
		return nil, err
	}
	responseData, err := s.userServiceSocket.Recv(0)
	if err != nil {
		return nil, err
	}
	response := &proto_user.RegisterResponse{}
	err = proto.Unmarshal(responseData, response)
	// Response that can't be decoded indicates a bug in the code or a serious data
	// corruption.
	if err != nil {
		log.Fatalf("Error unmarshaling RegisterResponse: %s", err)
	}
	return response, nil
}

// Close closes all the sockets and cleans up all the resources associated with
// this service.
// This method might block until all the resources are properly discarded.
func (s *UserServiceNanomsg) Close() {
	s.userServiceSocket.Close()
}
