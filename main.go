package main

import (
	"os"
	server "system-design/caching"
)

func main() {
	port := os.Args[1]
	server.RunServer(port)
}