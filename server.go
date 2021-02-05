package todo

import (
	"context"
	"net/http"
	"time"
)

//Server structure for create server
type Server struct {
	httpServer *http.Server
}

//Run method start server
func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
		MaxHeaderBytes: 1 << 20, //1 mb,
		ReadTimeout: 10 *time.Second,
		WriteTimeout: 10 *time.Second,
	}

	return s.httpServer.ListenAndServe() //start server in endless cycle for
}

//ShutDown shutdown server
func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
} 