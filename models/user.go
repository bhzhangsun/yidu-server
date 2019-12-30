package models

type User struct {
	UserId string `xorm:"not null pk comment('用户ID') CHAR(28)"`
	Level  int    `xorm:"not null default 0 comment('用户等级') TINYINT(3)"`
	Name   string `xorm:"not null comment('昵称') CHAR(64)"`
	Phone  string `xorm:"not null comment('手机') unique CHAR(24)"`
	Shadow string `xorm:"not null comment('秘钥hash') CHAR(64)"`
	Qq     string `xorm:"comment('qq') CHAR(64)"`
	Github string `xorm:"comment('github') CHAR(64)"`
	Wechat string `xorm:"comment('wechat') CHAR(64)"`
	Weibo  string `xorm:"comment('weibo') CHAR(64)"`
}
