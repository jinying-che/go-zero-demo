// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	shorturlFieldNames          = builder.RawFieldNames(&Shorturl{})
	shorturlRows                = strings.Join(shorturlFieldNames, ",")
	shorturlRowsExpectAutoSet   = strings.Join(stringx.Remove(shorturlFieldNames, "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	shorturlRowsWithPlaceHolder = strings.Join(stringx.Remove(shorturlFieldNames, "`shorten`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"
)

type (
	shorturlModel interface {
		Insert(ctx context.Context, data *Shorturl) (sql.Result, error)
		FindOne(ctx context.Context, shorten string) (*Shorturl, error)
		Update(ctx context.Context, data *Shorturl) error
		Delete(ctx context.Context, shorten string) error
	}

	defaultShorturlModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Shorturl struct {
		Shorten string `db:"shorten"` // shorten key
		Url     string `db:"url"`     // original url
	}
)

func newShorturlModel(conn sqlx.SqlConn) *defaultShorturlModel {
	return &defaultShorturlModel{
		conn:  conn,
		table: "`shorturl`",
	}
}

func (m *defaultShorturlModel) Delete(ctx context.Context, shorten string) error {
	query := fmt.Sprintf("delete from %s where `shorten` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, shorten)
	return err
}

func (m *defaultShorturlModel) FindOne(ctx context.Context, shorten string) (*Shorturl, error) {
	query := fmt.Sprintf("select %s from %s where `shorten` = ? limit 1", shorturlRows, m.table)
	var resp Shorturl
	err := m.conn.QueryRowCtx(ctx, &resp, query, shorten)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShorturlModel) Insert(ctx context.Context, data *Shorturl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, shorturlRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Shorten, data.Url)
	return ret, err
}

func (m *defaultShorturlModel) Update(ctx context.Context, data *Shorturl) error {
	query := fmt.Sprintf("update %s set %s where `shorten` = ?", m.table, shorturlRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Url, data.Shorten)
	return err
}

func (m *defaultShorturlModel) tableName() string {
	return m.table
}
