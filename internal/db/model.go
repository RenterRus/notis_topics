// Code generated by mfd-generator v0.4.4; DO NOT EDIT.

//nolint:all
//lint:file-ignore U1000 ignore unused code, it's generated
package db

var Columns = struct {
	Topics struct {
		ID, UserID, Name, SubID string
	}
}{
	Topics: struct {
		ID, UserID, Name, SubID string
	}{
		ID:     "id",
		UserID: "userid",
		Name:   "name",
		SubID:  "subid",
	},
}

var Tables = struct {
	Topics struct {
		Name, Alias string
	}
}{
	Topics: struct {
		Name, Alias string
	}{
		Name:  "topics",
		Alias: "t",
	},
}

type Topics struct {
	tableName struct{} `pg:"topics,alias:t,discard_unknown_columns"`

	ID     int64  `pg:"id,pk"`
	UserID int64  `pg:"userid,pk"`
	Name   string `pg:"name,pk"`
	SubID  *int64 `pg:"subid"`
}
