package models

type Track struct {
	ID      string `track_id`
	Novel   string `novel_id`
	User    string `user_id`
	Current string `current 当前所在章节序号`
	Favor   bool   `favor 是否收藏`
}
