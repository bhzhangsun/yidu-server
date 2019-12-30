package models

import (
	"time"
)

type Comment struct {
	ArticleId string    `xorm:"not null comment('文章id') CHAR(28)"`
	FromUser  string    `xorm:"not null comment('评论者') CHAR(28)"`
	ToUser    string    `xorm:"not null comment('被评论者') CHAR(28)"`
	IsReply   int       `xorm:"not null default 0 comment('是否回复') TINYINT(1)"`
	Timestamp time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('时间戳') TIMESTAMP"`
	Content   string    `xorm:"not null comment('内容') VARCHAR(1000)"`
}
