package main

import (
	"flag"
	"ims/src/client"
	"ims/src/server"
)

// ./app -mode client -ip 127.0.0.1 -port 8523
func main() {
	var mode string
	var serverIp string
	var serverPort string

	flag.StringVar(&mode, "mode", "client", "")
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "")
	flag.StringVar(&serverPort, "port", "8523", "")
	flag.Parse()
	if mode == "client" {
		client := client.NewClient("127.0.0.1", "8523")
		client.Start()
	}
	if mode == "server" {
		server := server.NewServer("127.0.0.1", "8523")
		server.Start()
	}
}
