package model

type Articles struct {
	Id        int    `xorm:"not null pk autoincr comment('文章id') INT(10)"`
	UserId    string `xorm:"not null comment('文章作者user_id') unique VARCHAR(16)"`
	Columnist string `xorm:"comment('文章分组') VARCHAR(64)"`
	Title     string `xorm:"comment('文章名') VARCHAR(255)"`
	Tag       string `xorm:"comment('文章标签') VARCHAR(4096)"`
	Content   string `xorm:"comment('文章内容') MEDIUMTEXT"`
}
