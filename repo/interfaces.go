package repo

import (
	"context"

	"github.com/erwinhermanto31/crud-image-with-auth/entity"
)

type IUsers interface {
	FindUsers(ctx context.Context, req entity.Users) (res entity.Users, err error)
}

type IImages interface {
	FindImages(ctx context.Context, req entity.Images, url string) (res []entity.Images, err error)
	UploadImages(ctx context.Context, req entity.Images) (err error)
}
