package service

import (
	"context"

	"github.com/0daryo/cinemawith/src/domain/model"
)

// User ... ユーザーサービスのインターフェイス
type User interface {
	Get(ctx context.Context, userID int64) (*model.User, error)
}
