package dbex_test

import "dbex"

func GetConf() (*dbex.MySqlConnection, error) {
	return dbex.NewMySqlConnection("root:root@tcp(127.0.0.1:3306)/ex?charset=utf8mb4&parseTime=True&loc=Local")
}