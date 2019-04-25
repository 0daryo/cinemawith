package model

import "github.com/0daryo/cinemawith/src/infrastructure/entity"

// User ... ユーザーモデル
type User struct {
	ID         int64  `json:"id"`
	UID        int64  `json:"uid"`
	Pr         string `json:"pr"`
	Name       string `json:"name"`
	ProfileUrl string `json:"profile_url"`
	Sex        int64  `json:"sex"`
	Age        int64  `json:"age"`
	Nickname   string `json:"nickname"`
	Favorite   string `json:"favorite"`
}

// NewUserFromEntity ... entityからdomain modelへの変換をかねる
func NewUserFromEntity(e *entity.User) *User {
	return &User{
		ID:         e.ID,
		UID:        e.UID,
		Pr:         e.Pr,
		Name:       e.Name,
		ProfileUrl: e.ProfileUrl,
		Sex:        e.Sex,
		Age:        e.Age,
		Nickname:   e.Nickname,
		Favorite:   e.Favorite,
	}
}
