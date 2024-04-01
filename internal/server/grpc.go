package server

import (
	"context"
	"fmt"
	"topics/internal/db"
	"topics/pkg/api"
)

type noteServer struct {
	topicDB db.TopicStorage
	api.UnimplementedNotesServer
}

func NewService(topic db.TopicStorage) *noteServer {
	return &noteServer{
		topicDB: topic,
	}
}

// Create implements pkg.NotesServer.
func (n *noteServer) Create(ctx context.Context, pkg *api.CreateRequest) (*api.CreateResponse, error) {
	userID, err := getXUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	res, err := n.topicDB.Create(ctx, db.CreateNoteRequest{
		SubID:  pkg.GetSubId(),
		Name:   pkg.GetName(),
		UserID: userID,
	})
	if err != nil {
		return nil, fmt.Errorf("Create: %w", err)
	}

	return createToResponse(res), nil
}

// Delete implements pkg.NotesServer.
func (n *noteServer) Delete(ctx context.Context, pkg *api.DeleteRequest) (*api.DeleteResponse, error) {
	userID, err := getXUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("delete: %w", err)
	}

	return &api.DeleteResponse{
		Message: deleteToResponse(n.topicDB.Delete(ctx, deleteToRequest(pkg, userID))),
	}, nil
}

// Read implements pkg.NotesServer.
func (n *noteServer) Read(ctx context.Context, pkg *api.ReadRequest) (*api.ReadResponse, error) {
	userID, err := getXUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}

	res, err := n.topicDB.Read(ctx, db.ReadRequest{
		UserID: userID,
		ID:     pkg.GetId(),
	})
	return readToResponse(res), err
}

// Update implements pkg.NotesServer.
func (n *noteServer) Update(ctx context.Context, pkg *api.UpdateRequest) (*api.UpdateResponse, error) {
	userID, err := getXUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}

	resp, err := n.topicDB.Update(ctx, db.UpdateRequest{
		ID:     pkg.GetId(),
		UserID: userID,
		Name:   pkg.GetName(),
		SubID:  pkg.GetSubId(),
	})

	return &api.UpdateResponse{
		Id:     resp.ID,
		Name:   resp.Name,
		SubId:  resp.SubID,
		UserId: resp.UserID,
	}, err
}
