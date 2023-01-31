package main

import (
	"nothing-behind.com/sample_gin/db"
	"nothing-behind.com/sample_gin/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
