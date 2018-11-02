package datasource

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)




func NewMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/xiuxiubeidanci?charset=utf8mb4")
	if err != nil {
		return nil, err
	}

	return db, nil
}