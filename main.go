package main

import (
	"github.com/JohnHuahuaZhan/fileserver/db/conn"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/static", "./static")
	conn.DBConn().Close()

	r.Run(":8081") // listen and serve on 0.0.0.0:8080

}
