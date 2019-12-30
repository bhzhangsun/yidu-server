package models

type Subscribe struct {
	Pick string `xorm:"not null comment('被关注') CHAR(28)"`
	Fans string `xorm:"not null comment('关注者') CHAR(28)"`
}
