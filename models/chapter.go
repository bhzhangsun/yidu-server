package models

type Chapter struct {
	ID       string            `xorm: "'chapt_id' pk comment('章节id') VARCHAR(16)"`
	Novel    string            `xorm: "'novel_id' not null comment('小说id') VARCHAR(16)"`
	Name     string            `xorm: "'name' not null comment('章节名')"`
	Sequence int               `xorm: "'sequence' not null comment('章节顺序')"`
	Sources  map[string]string `xorm: "'source' not null comment('来源：源站->path')"`
	Content  string            `xorm: "'content' comment('内容') TEXT"`
}
