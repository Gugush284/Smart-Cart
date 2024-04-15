package apiserver

import (
	"net/http"
)

func Start(config *Config) error {
	srv := NewServer()

	if err := srv.configureLogger(config); err != nil {
		return err
	}

	srv.Logger.Info("starting api server")
	srv.Logger.Debug(config.SessionKey)

	return http.ListenAndServe(config.BindAddr, srv)
}
