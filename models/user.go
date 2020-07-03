package models

type User struct {
	ID       uint32 `xorm: "'user_id' pk autoincr INT"`
	Phone    string `xorm: "'phone' index not null unique VARCHAR(13)"`
	Nickname string `xorm: "'nickname' VARCHAR(32)"`
}
