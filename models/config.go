package models

type Config struct {
	Db Database `mapstructure:"database"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Name     string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
