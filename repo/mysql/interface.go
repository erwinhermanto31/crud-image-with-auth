package mysql

import (
	"context"

	"github.com/erwinhermanto31/crud-image-with-auth/utils"
	"github.com/jmoiron/sqlx"
)

type IMysql interface {
	Get(context.Context, *sqlx.DB, interface{}, *utils.Query, string) error
	Select(context.Context, *sqlx.DB, interface{}, *utils.Query, string) error
	CreateOrUpdate(ctx context.Context, db *sqlx.DB, data interface{}, query string) (lastId int64, err error)
}
