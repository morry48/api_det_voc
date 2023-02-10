package main

import (
	"nothing-behind.com/sample_gin/config"
	"nothing-behind.com/sample_gin/server"
)

func main() {
	config.Init()
	server.Init()
	config.Close()
}
