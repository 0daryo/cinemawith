package entity

// User ... ユーザーエンティティ
type User struct {
	ID         int64  `db:"id"`
	UID        int64  `db:"uid"`
	Pr         string `db:"pr"`
	Name       string `db:"name"`
	ProfileUrl string `db:"profile_url"`
	Sex        int64  `db:"sex"`
	Age        int64  `db:"age"`
	Nickname   string `db:"nickname"`
	Favorite   string `db:"favorite"`
	BaseEntity
}
