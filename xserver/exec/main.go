package main

import (
	"os"
	"os/signal"

	"github.com/ezeev/go-examples/xserver"
)

func main() {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	server := xserver.NewServer(8081)
	go server.Start()

	<-stop
	server.Shutdown()

}
