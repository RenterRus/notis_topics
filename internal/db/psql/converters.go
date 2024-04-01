package psql

import (
	"topics/internal/db"

	"github.com/AlekSi/pointer"
)

func readToResponse(t []db.Topics) []db.Read {
	res := make([]db.Read, 0, len(t))

	for _, v := range t {
		res = append(res, db.Read{
			ID:     v.ID,
			UserID: v.UserID,
			Name:   v.Name,
			SubID:  pointer.Get(v.SubID),
		})
	}

	return res
}
