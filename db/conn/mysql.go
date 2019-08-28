package conn

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(192.168.16.2:3306)/fileserver?charset=utf8")
	if err != nil {
		panic("Failed to connect to mysql, err:" + err.Error())
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		panic("Failed to connect to mysql, err:" + err.Error())
	}
}

// DBConn : 返回数据库连接池对象
func DBConn() *sql.DB {
	return db
}
