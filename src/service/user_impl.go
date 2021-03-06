package service

import (
	"context"

	"github.com/0daryo/cinemawith/src/domain/model"
	"github.com/0daryo/cinemawith/src/domain/repository"
	"github.com/0daryo/cinemawith/src/lib/log"
)

type user struct {
	uRepo repository.User
}

func (s *user) Get(ctx context.Context, userID int64) (*model.User, error) {
	user, err := s.uRepo.Get(ctx, userID)
	if err != nil {
		log.Errorf(ctx, "s.uRepo.Get: %s", err.Error())
		return nil, err
	}
	return model.NewUserFromEntity(user), nil
}

// NewUser ... ユーザーサービスを取得する
func NewUser(uRepo repository.User) User {
	return &user{
		uRepo: uRepo,
	}
}
