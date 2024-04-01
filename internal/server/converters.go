package server

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"topics/internal/db"
	"topics/pkg/api"

	"google.golang.org/grpc/metadata"
)

func readToResponse(t []db.Read) *api.ReadResponse {
	res := make([]*api.ReadResponse_Topics, 0, len(t))

	for _, v := range t {
		res = append(res, &api.ReadResponse_Topics{
			Id:     v.ID,
			UserId: v.ID,
			Name:   v.Name,
			SubId:  v.SubID,
		})
	}

	return &api.ReadResponse{
		Topics: res,
	}
}

func createToResponse(t db.CreateNoteResponse) *api.CreateResponse {
	return &api.CreateResponse{
		Id:     t.ID,
		SubId:  t.SubID,
		Name:   t.Name,
		UserId: t.UserID,
	}
}

func deleteToRequest(pkg *api.DeleteRequest, id int64) []db.DeleteRequest {
	res := make([]db.DeleteRequest, 0, len(pkg.Md))

	for _, v := range pkg.GetMd() {
		res = append(res, db.DeleteRequest{
			ID:     v.GetId(),
			UserID: id,
			Name:   v.GetName(),
		})
	}

	return res
}

func deleteToResponse(pkg []db.DeleteResponse) []string {
	res := make([]string, 0, len(pkg))

	for _, v := range pkg {
		res = append(res, v.Message)
	}

	return res
}

func getXUserID(ctx context.Context) (int64, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, errors.New("user not delivered")
	}

	if user := md.Get("X-User-ID"); len(user) == 0 {
		return 0, errors.New("user not found")
	}

	id, err := strconv.Atoi(md.Get("X-User-ID")[0])
	if err != nil {
		return 0, fmt.Errorf("parse user: %w", err)
	}

	return int64(id), nil
}
