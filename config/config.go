package config

const (
	ENV   string = "dev"
	MYSQL string = "yidu:ANnZtwkRsMn3mNJJ@tcp(0.0.0.0:3306)/db_yidu?charset=utf8"
)

type Configurer struct {
	Env  string `json:"env"`
	Port string `json:"port"`

	Mysql string `json:"mysql"`
}
