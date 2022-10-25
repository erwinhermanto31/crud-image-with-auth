package action

import (
	"context"
	"log"
	"time"

	"github.com/erwinhermanto31/crud-image-with-auth/entity"
	"github.com/erwinhermanto31/crud-image-with-auth/repo"
	"github.com/erwinhermanto31/crud-image-with-auth/utils"
	"github.com/erwinhermanto31/crud-image-with-auth/utils/errors"
)

type Login struct {
	repoUser repo.IUsers
}

func NewLogin() *Login {
	return &Login{
		repoUser: repo.NewUsers(),
	}
}

func (h *Login) Handler(ctx context.Context, req entity.Users) (token string, err error) {

	if req.Username == "" {
		return "", errors.ErrNotFound("Please input username")
	}

	if req.Password == "" {
		return "", errors.ErrNotFound("Please input password")
	}

	data, err := h.repoUser.FindUsers(ctx, req)
	if err != nil {
		log.Printf("[Handler] FindUser : %v", err)
		return "", err
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	subject := "login"
	token = utils.GenerateToken(data.Username, int32(data.Id), expirationTime.Unix(), subject)

	// res.Token = token
	// res.ExpiredTimestamp = expirationTime.Unix()
	log.Println(token)

	return token, nil
}
