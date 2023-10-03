package metadata

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	RequestID = "request-id"
)

type MD[T any] interface {
	WithContext(parent context.Context) context.Context
	WithRequestID(string) T
	GetRequestID() string
}

type MetaData struct {
	RequestID string
}

func (m MetaData) WithRequestID(requestID string) MetaData {
	m.RequestID = requestID
	return m
}

func (m MetaData) WithContext(parent context.Context) context.Context {
	md := metadata.Pairs(
		RequestID, m.GetRequestID(),
	)
	return metadata.NewOutgoingContext(parent, md)
}

func (m MetaData) GetRequestID() string {
	return m.RequestID
}
