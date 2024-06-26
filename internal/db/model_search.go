// Code generated by mfd-generator v0.4.4; DO NOT EDIT.

//nolint:all
//lint:file-ignore U1000 ignore unused code, it's generated
package db

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

const condition = "?.? = ?"

// base filters
type applier func(query *orm.Query) (*orm.Query, error)

type search struct {
	appliers []applier
}

func (s *search) apply(query *orm.Query) {
	for _, applier := range s.appliers {
		query.Apply(applier)
	}
}

func (s *search) where(query *orm.Query, table, field string, value interface{}) {
	query.Where(condition, pg.Ident(table), pg.Ident(field), value)
}

func (s *search) WithApply(a applier) {
	if s.appliers == nil {
		s.appliers = []applier{}
	}
	s.appliers = append(s.appliers, a)
}

func (s *search) With(condition string, params ...interface{}) {
	s.WithApply(func(query *orm.Query) (*orm.Query, error) {
		return query.Where(condition, params...), nil
	})
}

// Searcher is interface for every generated filter
type Searcher interface {
	Apply(query *orm.Query) *orm.Query
	Q() applier

	With(condition string, params ...interface{})
	WithApply(a applier)
}

type TopicsSearch struct {
	search

	ID     *int64
	UserID *int64
	Name   *string
	SubID  *int64
}

func (ts *TopicsSearch) Apply(query *orm.Query) *orm.Query {
	if ts == nil {
		return query
	}
	if ts.ID != nil {
		ts.where(query, Tables.Topics.Alias, Columns.Topics.ID, ts.ID)
	}
	if ts.UserID != nil {
		ts.where(query, Tables.Topics.Alias, Columns.Topics.UserID, ts.UserID)
	}
	if ts.Name != nil {
		ts.where(query, Tables.Topics.Alias, Columns.Topics.Name, ts.Name)
	}
	if ts.SubID != nil {
		ts.where(query, Tables.Topics.Alias, Columns.Topics.SubID, ts.SubID)
	}

	ts.apply(query)

	return query
}

func (ts *TopicsSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		if ts == nil {
			return query, nil
		}
		return ts.Apply(query), nil
	}
}
