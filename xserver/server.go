package xserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Router     *mux.Router
	HttpServer *http.Server
	Port       int
}

// NewServer returns an initialized instance of a new server
func NewServer(port int) *Server {
	s := &Server{}
	s.Port = port

	s.Router = mux.NewRouter()

	// setup routes
	s.Router.HandleFunc("/ping", s.handlePing())

	s.HttpServer = &http.Server{
		Handler:      s.Router,
		Addr:         fmt.Sprintf("127.0.0.1:%d", s.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return s
}

// Start http server. This method is blocking.
func (s *Server) Start() {
	s.HttpServer.ListenAndServe()
}

// Shutdown stops the HTTP Server
func (s *Server) Shutdown() {

	// clean up for shutdown here!

	//shutdown http server
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	s.HttpServer.Shutdown(ctx)
	log.Println("Server shutdown gracefully")

}

// ServerReady Blocks until the server returns a 200 status code from the /ping endpoint
func (s *Server) ServerReady() bool {

	var client = &http.Client{
		Timeout: time.Second * 5,
	}

	tries := 0
	for {
		tries++
		url := fmt.Sprintf("http://localhost:%d/ping", s.Port)
		resp, _ := client.Get(url)
		if resp.StatusCode == http.StatusOK {
			return true
		} else {
			time.Sleep(time.Second * 1)
		}

		if tries >= 5 {
			panic("Unable to connect to test HTTP server at " + url)
		}

	}

}

func (s *Server) handlePing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	}
}
