package repositories

import "github.com/quocbang/grpc-gateway/server/utils/roles"

type LoginRequest struct {
	Username string
}

type LoginReply struct {
	Username     string
	HashPassword []byte
	Roles        roles.Roles
}
