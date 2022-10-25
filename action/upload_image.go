package action

import (
	"context"
	"log"

	"github.com/erwinhermanto31/crud-image-with-auth/entity"
	"github.com/erwinhermanto31/crud-image-with-auth/repo"
)

type UploadImage struct {
	repoImage repo.IImages
}

func NewUploadImage() *UploadImage {
	return &UploadImage{
		repoImage: repo.NewImages(),
	}
}

func (h *UploadImage) Handler(ctx context.Context, req entity.Images) (err error) {
	err = h.repoImage.UploadImages(ctx, req)
	if err != nil {
		log.Printf("[Handler] UploadImages : %v", err)
		return err
	}
	return nil
}
