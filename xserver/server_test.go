package xserver_test

import (
	"testing"

	"github.com/ezeev/go-examples/xserver"
)

func TestServer(t *testing.T) {

	server := xserver.NewServer(8081)
	go server.Start()
	// wait until the server is ready
	server.ServerReady()

	// run tests against the server here!

	server.Shutdown()

}
