package db

import (
	"context"
)

type CreateNoteRequest struct {
	UserID int64
	Name   string
	SubID  int64
}

type CreateNoteResponse struct {
	ID     int64
	UserID int64
	Name   string
	SubID  int64
}

type ReadRequest struct {
	ID     int64
	UserID int64
}

type Read struct {
	ID     int64
	UserID int64
	Name   string
	SubID  int64
}

type DeleteRequest struct {
	ID     int64
	UserID int64
	Name   string
}

type DeleteResponse struct {
	Message string
}

type UpdateRequest struct {
	ID     int64
	UserID int64
	Name   string
	SubID  int64
}

type UpdateResponse struct {
	ID     int64
	UserID int64
	Name   string
	SubID  int64
}

type TopicStorage interface {
	Create(ctx context.Context, notes CreateNoteRequest) (CreateNoteResponse, error)
	Read(ctx context.Context, req ReadRequest) ([]Read, error)
	Update(ctx context.Context, req UpdateRequest) (UpdateResponse, error)
	Delete(ctx context.Context, req []DeleteRequest) []DeleteResponse
}
