// Code generated by protoc-gen-gogo.
// source: user.proto
// DO NOT EDIT!

/*
Package proto_user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	User
	RegisterUser
	RegisterResponse
	AuthenticateUser
	AuthenticateResult
*/
package proto_user

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// id = 1
type User struct {
	Id               *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	DisplayName      *string `protobuf:"bytes,2,opt,name=display_name" json:"display_name,omitempty"`
	Email            *string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Password         *string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}

func (m *User) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *User) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

// id = 20
type RegisterUser struct {
	User             *User   `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	RedirectUri      *string `protobuf:"bytes,2,opt,name=redirect_uri" json:"redirect_uri,omitempty"`
	Locale           *string `protobuf:"bytes,3,opt,name=locale" json:"locale,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RegisterUser) Reset()         { *m = RegisterUser{} }
func (m *RegisterUser) String() string { return proto.CompactTextString(m) }
func (*RegisterUser) ProtoMessage()    {}

func (m *RegisterUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *RegisterUser) GetRedirectUri() string {
	if m != nil && m.RedirectUri != nil {
		return *m.RedirectUri
	}
	return ""
}

func (m *RegisterUser) GetLocale() string {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return ""
}

// id = 21
type RegisterResponse struct {
	RedirectUri      *string                        `protobuf:"bytes,1,opt,name=redirect_uri" json:"redirect_uri,omitempty"`
	Valid            *bool                          `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Errors           []*RegisterResponse_InputError `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}

func (m *RegisterResponse) GetRedirectUri() string {
	if m != nil && m.RedirectUri != nil {
		return *m.RedirectUri
	}
	return ""
}

func (m *RegisterResponse) GetValid() bool {
	if m != nil && m.Valid != nil {
		return *m.Valid
	}
	return false
}

func (m *RegisterResponse) GetErrors() []*RegisterResponse_InputError {
	if m != nil {
		return m.Errors
	}
	return nil
}

type RegisterResponse_InputError struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	ErrorMessage     *string `protobuf:"bytes,2,req,name=error_message" json:"error_message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RegisterResponse_InputError) Reset()         { *m = RegisterResponse_InputError{} }
func (m *RegisterResponse_InputError) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse_InputError) ProtoMessage()    {}

func (m *RegisterResponse_InputError) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *RegisterResponse_InputError) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

// id = 30
type AuthenticateUser struct {
	Email            *string `protobuf:"bytes,1,req,name=email" json:"email,omitempty"`
	Password         *string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthenticateUser) Reset()         { *m = AuthenticateUser{} }
func (m *AuthenticateUser) String() string { return proto.CompactTextString(m) }
func (*AuthenticateUser) ProtoMessage()    {}

func (m *AuthenticateUser) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *AuthenticateUser) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

// id = 31
type AuthenticateResult struct {
	Sid              *string `protobuf:"bytes,1,opt,name=sid" json:"sid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthenticateResult) Reset()         { *m = AuthenticateResult{} }
func (m *AuthenticateResult) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResult) ProtoMessage()    {}

func (m *AuthenticateResult) GetSid() string {
	if m != nil && m.Sid != nil {
		return *m.Sid
	}
	return ""
}

func init() {
}