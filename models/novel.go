package models

type Novel struct {
	ID          string   `xorm:"'novel_id' pk index not null comment('小说ID') VARCHAR(16)"`
	Name        string   `xorm:"not null unique comment('小说名')"`
	Image       string   `xorm:"not null comment('封面链接')"`
	Author      string   `xorm:"not null comment('作者')"`
	Description string   `xorm:"default('') VARCHAR(512)"`
	State       int8     `xorm:"default(0) comment('小说状态：0-完结，1-连载中，2-其他') TINYINT"`
	Class       []string `xorm:"not null comment('分类：[‘男生/女生’, ‘仙侠’]') VARCHAR(255)"`
	Tags        []string `xorm:"not null comment('标签：[升级文, 穿越]')" VARCHAR(255)`
	Favors      uint64   `xorm:"default(0) comment('赞数')"`
	Sources     []string `xorm:"not null comment('来源')"`
	Current     string   `xorm:"not null comment('当前章节') VARCHAR(255)"`
}
