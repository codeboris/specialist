package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(cfg *Config) *API {
	return &API{
		config: cfg,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	if err := api.configLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port: ", api.config.APPPort)

	api.configRouterField()

	return http.ListenAndServe(api.config.APPPort, api.router)
}
