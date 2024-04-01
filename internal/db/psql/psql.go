package psql

import (
	"context"
	"fmt"
	"slices"
	"topics/internal/db"

	"github.com/AlekSi/pointer"
	pg "github.com/go-pg/pg/v10"
	_ "github.com/jackc/pgx"
)

type DataStorage struct {
	conn *pg.DB
}

type DBInfo struct {
	Addr     string
	User     string
	Password string
	Database string
}

func NewPQConnect(db *DBInfo) db.TopicStorage {
	return &DataStorage{
		conn: pg.Connect(&pg.Options{
			Addr:     db.Addr,
			User:     db.User,
			Password: db.Password,
			Database: db.Database,
		}),
	}
}

func (d *DataStorage) Create(ctx context.Context, req db.CreateNoteRequest) (db.CreateNoteResponse, error) {
	model := &db.Topics{
		UserID: req.UserID,
		Name:   req.Name,
	}
	if req.SubID != 0 {
		model.SubID = pointer.To(req.SubID)
	}
	_, err := d.conn.Model(model).Returning("*").Insert()
	if err != nil {
		fmt.Println(err)
		return db.CreateNoteResponse{}, fmt.Errorf("psql.Create: %w", err)
	}

	return db.CreateNoteResponse{
		ID:     model.ID,
		UserID: model.UserID,
		Name:   model.Name,
		SubID:  pointer.Get(model.SubID),
	}, nil
}

func (d *DataStorage) Read(ctx context.Context, req db.ReadRequest) ([]db.Read, error) {
	var model []db.Topics

	_, err := d.conn.ModelContext(ctx, &model).Query(&model, "WITH RECURSIVE getids AS ( SELECT id, userid, name, subid FROM topics  WHERE userid = ?0 and id = ?1 UNION ALL SELECT e.id, e.userid, e.name, e.subid FROM topics e JOIN getids ON e.id = getids.subid) SELECT * FROM getids;",
		req.UserID, req.ID)
	if err != nil {
		return []db.Read{}, fmt.Errorf("psql.Read: %w", err)
	}

	return readToResponse(model), nil
}

func (d *DataStorage) Delete(ctx context.Context, req []db.DeleteRequest) []db.DeleteResponse {
	resp := make([]db.DeleteResponse, 0, len(req))
	m := make(map[int]db.DeleteRequest)
	sl := make([]int, 0, len(req))

	for _, v := range req {
		m[int(v.ID)] = db.DeleteRequest{
			UserID: v.UserID,
			Name:   v.Name,
		}
		sl = append(sl, int(v.ID))
	}

	slices.Sort(sl)

	for i := len(sl) - 1; i >= 0; i-- {
		var model db.Topics
		query := d.conn.Model(&model)
		_, err := query.Where("? = ? and ? = ?", pg.Ident(db.Columns.Topics.ID), sl[i], pg.Ident(db.Columns.Topics.UserID), m[sl[i]].UserID).Delete()
		if err != nil {
			resp = append(resp, db.DeleteResponse{
				Message: fmt.Sprintf("failed delete: %s", m[sl[i]].Name),
			})
		} else {
			resp = append(resp, db.DeleteResponse{
				Message: fmt.Sprintf("%s deleted", m[sl[i]].Name),
			})
		}
	}

	return resp
}

func (d *DataStorage) Update(ctx context.Context, req db.UpdateRequest) (db.UpdateResponse, error) {
	model := &db.Topics{
		ID:     req.ID,
		Name:   req.Name,
		UserID: req.UserID,
	}
	if req.SubID != 0 {
		model.SubID = pointer.To(req.SubID)
	}

	_, err := d.conn.Model(model).Column(db.Columns.Topics.Name).Column(db.Columns.Topics.SubID).Where("? = ?", pg.Ident(db.Columns.Topics.ID), req.ID).Update()
	if err != nil {
		return db.UpdateResponse{}, fmt.Errorf("psql.Update: %w", err)
	}
	return db.UpdateResponse{
		ID:     model.ID,
		UserID: model.UserID,
		SubID:  pointer.Get(model.SubID),
		Name:   model.Name,
	}, nil
}
