package models

type Chapter struct {
	ID      string   `xorm:"'chapt_id' pk index not null comment('章节id')" VARCHAR(16)`
	Novel   string   `xorm:"'novel_id' not null comment('小说id')" VARCHAR(16)`
	Name    string   `xorm:"not null comment('小说名')"`
	Sources []string `xorm:"not null comment('来源')"`
	Content string   `xorm:" comment('内容') TEXT"`
}
