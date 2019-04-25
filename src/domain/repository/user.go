package repository

import (
	"context"

	"github.com/0daryo/cinemawith/src/infrastructure/entity"
)

// User ... ユーザーレポジトリのinterface
type User interface {
	Get(ctx context.Context, userID int64) (*entity.User, error)
}
