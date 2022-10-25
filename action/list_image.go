package action

import (
	"context"
	"log"

	"github.com/erwinhermanto31/crud-image-with-auth/entity"
	"github.com/erwinhermanto31/crud-image-with-auth/repo"
)

type ListImage struct {
	repoImage repo.IImages
}

func NewListImage() *ListImage {
	return &ListImage{
		repoImage: repo.NewImages(),
	}
}

func (h *ListImage) Handler(ctx context.Context, req entity.Images, url string) (res []entity.Images, err error) {
	res, err = h.repoImage.FindImages(ctx, req, url)
	if err != nil {
		log.Printf("[Handler] ListImages : %v", err)
		return res, err
	}
	return res, nil
}
