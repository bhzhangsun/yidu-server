package models

type Track struct {
	ID      string `xorm: "'track_id'"`
	Novel   string `xorm: "'novel_id'"`
	User    uint32 `xorm: "'user_id'"`
	Current string `xorm: "'current' comment('当前所在章节序号')"`
	Favor   bool   `xorm: "'favor' comment('是否收藏')"`
}
