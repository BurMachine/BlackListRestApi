package server

import (
	"blacklistApi/internal/config"
	"blacklistApi/internal/database"
	"blacklistApi/internal/handlers"
	gorilla "github.com/gorilla/mux"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	conf     config.Conf
	Mux      *gorilla.Router
	Handlers *handlers.Handlers
}

func New(conf config.Conf, storage *database.Storage) *Server {
	return &Server{conf: conf, Mux: gorilla.NewRouter(), Handlers: handlers.New(storage)}
}

func (s *Server) Run() error {
	wg := &sync.WaitGroup{}
	srv := &http.Server{
		Handler:      s.Mux,
		Addr:         s.conf.AddrHttp,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	s.Handlers.RegisteringHandlers(s.Mux)
	var err error
	go func() {
		wg.Add(1)
		err = srv.ListenAndServe()
		wg.Done()
	}()
	s.Handlers.Logger.Info().Msg("server start")
	wg.Wait()
	if err != nil {
		return err
	}
	return nil
}
