package main

import (
	"log"

	proto "code.google.com/p/gogoprotobuf/proto"
	nmsg "github.com/op/go-nanomsg"

	"github.com/opentarock/frontend-user-management/server/service/proto_user"
)

// Implementation of user service backend for testing purposes.
func main() {
	socket, err := nmsg.NewRepSocket()
	if err != nil {
		log.Fatalf("Error creating socket: %s", err)
	}
	defer socket.Close()
	endpoint, err := socket.Bind("tcp://*:6001")
	if err != nil {
		log.Fatalf("Error binding socket: %s", err)
	}

	log.Printf("Bound to endpoint: %s", endpoint.Address)

	for {
		data, err := socket.Recv(0)
		if err != nil {
			log.Fatalf("Error receiving message")
			continue
		}
		registerUser := &proto_user.RegisterUser{}
		err = proto.Unmarshal(data, registerUser)
		if err != nil {
			log.Fatalf("Error unmarshaling RegisterUser: %s", err)
		}

		var response *proto_user.RegisterResponse
		if registerUser.GetUser().GetEmail() == "error@example.com" {
			response = &proto_user.RegisterResponse{
				Valid: proto.Bool(false),
				Errors: []*proto_user.RegisterResponse_InputError{
					&proto_user.RegisterResponse_InputError{
						Name:         proto.String("display_name"),
						ErrorMessage: proto.String("Display Name should be longer that 2 characters."),
					},
					&proto_user.RegisterResponse_InputError{
						Name:         proto.String("email"),
						ErrorMessage: proto.String("Email already used."),
					},
					&proto_user.RegisterResponse_InputError{
						Name:         proto.String("password"),
						ErrorMessage: proto.String("Password is too short."),
					},
				},
			}
		} else {
			redirectUri := registerUser.GetRedirectUri()
			if redirectUri == "" {
				redirectUri = "/user"
			}
			response = &proto_user.RegisterResponse{
				RedirectUri: proto.String(redirectUri),
				Valid:       proto.Bool(true),
			}
		}

		responseData, err := proto.Marshal(response)
		if err != nil {
			log.Fatalf("Error marshaling RegisterResponse: %s", err)
		}
		_, err = socket.Send(responseData, 0)
		if err != nil {
			log.Fatalf("Error sending message")
		}
	}
}
