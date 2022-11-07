package main

import (
	"flag"
	"fmt"
	"ims/src/client"
	"ims/src/server"
	"os"
	"os/signal"
	"syscall"
)

// ./app -mode client -ip 127.0.0.1 -port 8523
func main() {
	var mode string
	var serverIp string
	var serverPort string

	flag.StringVar(&mode, "mode", "server", "")
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "")
	flag.StringVar(&serverPort, "port", "8523", "")
	flag.Parse()
	if mode == "client" {
		client := client.NewClient("127.0.0.1", "8523")
		go client.Start()
	}
	if mode == "server" {
		server := server.NewServer("127.0.0.1", "8523")
		go server.Start()
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sig
	close(sig)
	fmt.Printf("\nbye bye\n")
	os.Exit(0)
}
