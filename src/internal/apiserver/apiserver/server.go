package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	Logger *logrus.Logger
}

func NewServer() *server {
	s := &server{
		router: mux.NewRouter(),
		Logger: logrus.New(),
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/tg/{key}", s.GetGoods()).Methods("POST")
	s.router.HandleFunc("/ready/tg/{key}", s.ReadyTG()).Methods("POST")
	s.router.HandleFunc("/ready/{key}", s.Ready()).Methods("GET")
}

func (s *server) configureLogger(config *Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	s.Logger.SetLevel(level)

	return nil
}
