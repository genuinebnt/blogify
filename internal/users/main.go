package main

import (
	"github.com/genuinebnt/blogify/internal/common/logs"
	"github.com/genuinebnt/blogify/internal/common/server"
)

func main() {
	// Initialize the global logger
	logs.Init()

	r := Router()
	server.RunHTTPServer(r)
}
