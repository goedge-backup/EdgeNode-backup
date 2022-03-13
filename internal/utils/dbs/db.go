// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dbs

import (
	"context"
	"database/sql"
)

type DB struct {
	rawDB *sql.DB

	enableStat bool
}

func NewDB(rawDB *sql.DB) *DB {
	return &DB{
		rawDB: rawDB,
	}
}

func (this *DB) EnableStat(b bool) {
	this.enableStat = b
}

func (this *DB) Prepare(query string) (*Stmt, error) {
	stmt, err := this.rawDB.Prepare(query)
	if err != nil {
		return nil, err
	}

	var s = NewStmt(stmt, query)
	if this.enableStat {
		s.EnableStat()
	}
	return s, nil
}

func (this *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if this.enableStat {
		defer SharedQueryStatManager.AddQuery(query).End()
	}
	return this.rawDB.ExecContext(ctx, query, args...)
}

func (this *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if this.enableStat {
		defer SharedQueryStatManager.AddQuery(query).End()
	}
	return this.rawDB.Exec(query, args...)
}

func (this *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if this.enableStat {
		defer SharedQueryStatManager.AddQuery(query).End()
	}
	return this.rawDB.Query(query, args...)
}

func (this *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	if this.enableStat {
		defer SharedQueryStatManager.AddQuery(query).End()
	}
	return this.rawDB.QueryRow(query, args...)
}

func (this *DB) Close() error {
	return this.rawDB.Close()
}