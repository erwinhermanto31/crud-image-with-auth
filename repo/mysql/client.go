package mysql

import (
	"context"
	"log"

	"github.com/erwinhermanto31/crud-image-with-auth/utils"
	"github.com/erwinhermanto31/crud-image-with-auth/utils/errors"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	e errors.Error
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Get(ctx context.Context, db *sqlx.DB, data interface{}, query *utils.Query, queryString string) (err error) {
	where, args := query.Where()
	q := queryString
	q += where

	log.Println(db.Rebind(q))
	if err = db.GetContext(ctx, data, db.Rebind(q), args...); err != nil {
		return
	}

	return
}

func (c *Client) Select(ctx context.Context, db *sqlx.DB, data interface{}, query *utils.Query, queryString string) (err error) {
	where, args := query.Where()
	sort := query.Order()
	q := queryString
	q += where
	q += sort

	if err = db.SelectContext(ctx, data, db.Rebind(q), args...); err != nil {
		return
	}

	return
}

func (c *Client) CreateOrUpdate(ctx context.Context, db *sqlx.DB, data interface{}, query string) (lastId int64, err error) {
	res, err := db.NamedExecContext(ctx, query, data)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	lastId, _ = res.LastInsertId()

	return lastId, err
}
