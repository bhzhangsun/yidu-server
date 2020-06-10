package models

type User struct {
	ID       string `xorm: 'user_id' pk VARCHAR(16)`
	Phone    string `xorm: 'phone' not null index VARCHAR(13)`
	Nickname string `xorm: 'nickname' VARCHAR(32)`
}
