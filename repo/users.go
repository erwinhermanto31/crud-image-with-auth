package repo

import (
	"context"

	"github.com/erwinhermanto31/crud-image-with-auth/entity"
	"github.com/erwinhermanto31/crud-image-with-auth/repo/mysql"
	"github.com/erwinhermanto31/crud-image-with-auth/utils"
	"github.com/jmoiron/sqlx"
)

type Users struct {
	iMysql mysql.IMysql
	AppDB  *sqlx.DB
}

func NewUsers() *Users {
	return &Users{
		iMysql: mysql.NewClient(),
		AppDB:  mysql.AppDB,
	}
}

func (r *Users) FindUsers(ctx context.Context, req entity.Users) (res entity.Users, err error) {
	query := &utils.Query{
		Filter: map[string]interface{}{
			"username$eq!": req.Username,
			"password$eq!": req.Password,
		},
	}

	err = r.iMysql.Get(ctx, r.AppDB, &res, query, mysql.QueryFindUser)
	if err != nil {
		return res, err
	}
	return res, err
}
