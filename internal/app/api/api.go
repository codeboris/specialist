package api

import (
	"net/http"

	"github.com/codeboris/specialist/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
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

	if err := api.configStorageField(); err != nil {
		return err
	}

	return http.ListenAndServe(api.config.APPPort, api.router)
}
