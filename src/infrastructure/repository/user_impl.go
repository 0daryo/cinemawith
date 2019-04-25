package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/0daryo/cinemawith/src/domain/repository"
	"github.com/0daryo/cinemawith/src/infrastructure/entity"
	"github.com/0daryo/cinemawith/src/lib/log"
	"github.com/0daryo/cinemawith/src/lib/mysql"
	"github.com/jmoiron/sqlx"
)

type user struct {
	sql *sqlx.DB
}

func (r *user) Get(ctx context.Context, userID int64) (*entity.User, error) {

	user := &entity.User{}

	qb := sq.Select("*").From("users").Where(sq.Eq{"id": userID})

	mysql.DumpSelectQuery(ctx, qb)

	query, attrs, err := qb.ToSql()
	if err != nil {
		log.Errorf(ctx, "sq.Select: %s", err.Error())
		return nil, err
	}

	rows, err := r.sql.QueryxContext(ctx, query, attrs...)
	if err != nil {
		log.Errorf(ctx, "r.sql.Queryx: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Errorf(ctx, "r.sql.Queryx: %s", err.Error())
			return nil, err
		}
		break
	}

	return user, nil
}

// NewUser ... ユーザーレポジトリを取得する
func NewUser(sql *sqlx.DB) repository.User {
	return &user{
		sql: sql,
	}
}
