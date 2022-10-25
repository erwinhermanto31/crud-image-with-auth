package repo

import (
	"context"

	"github.com/erwinhermanto31/crud-image-with-auth/entity"
	"github.com/erwinhermanto31/crud-image-with-auth/repo/mysql"
	"github.com/erwinhermanto31/crud-image-with-auth/utils"
	"github.com/jmoiron/sqlx"
)

type Images struct {
	iMysql mysql.IMysql
	AppDB  *sqlx.DB
}

func NewImages() *Images {
	return &Images{
		iMysql: mysql.NewClient(),
		AppDB:  mysql.AppDB,
	}
}

func (r *Images) FindImages(ctx context.Context, req entity.Images, url string) (res []entity.Images, err error) {
	query := &utils.Query{
		Filter: map[string]interface{}{
			"user_id$eq!": req.UserId,
		},
	}

	err = r.iMysql.Select(ctx, r.AppDB, &res, query, mysql.QueryFindImage)
	if err != nil {
		return res, err
	}
	for i, v := range res {
		res[i].ImageURL = url + v.ImageURL
	}
	return res, err
}

func (r *Images) UploadImages(ctx context.Context, req entity.Images) (err error) {
	_, err = r.iMysql.CreateOrUpdate(ctx, r.AppDB, &req, mysql.QueryInsertImage)
	if err != nil {
		return err
	}
	return nil
}
