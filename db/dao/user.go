package dao

import (
	"github.com/JohnHuahuaZhan/fileserver/db/conn"
)

type User struct {
	Username     string
	Email        string
	Phone        string
	SignUpAt     string
	LastActiveAt string
	Status       int
}

func Insert(username, password string) bool {
	stmt, err := conn.DBConn().Prepare(
		"insert ignore into tbl_user (`user_name`,`user_pwd`) values (?,?)")
	if err != nil {
		return false
	}
	defer stmt.Close()
	ret, err := stmt.Exec(username, password)
	if err != nil {
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}
	return false
}
