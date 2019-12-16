package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "root",
			Name:     "goblog",
		},
	}
}

func GetConnection() (db *sql.DB) {
	conf := GetConfig()
	db, err := sql.Open(conf.DB.Dialect, conf.DB.Username+":"+conf.DB.Password+"@/"+conf.DB.Name)
	if err != nil {
		panic(err.Error())
	}
	return db
}
