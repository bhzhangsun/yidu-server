package models

import (
	"time"
)

type Article struct {
	ArticleId string    `xorm:"not null pk comment('用户ID') CHAR(28)"`
	Owner     string    `xorm:"not null comment('所属用户') CHAR(28)"`
	Scan      int       `xorm:"not null comment('浏览量') INT(10)"`
	Tags      string    `xorm:"comment('标记') VARCHAR(1024)"`
	Content   string    `xorm:"not null comment('内容') MEDIUMTEXT"`
	Timestamp time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('时间戳') TIMESTAMP"`
}
