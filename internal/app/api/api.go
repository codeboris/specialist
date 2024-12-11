package api

import "github.com/sirupsen/logrus"

type API struct {
	config *Config
	logger *logrus.Logger
}

func New(cfg *Config) *API {
	return &API{
		config: cfg,
		logger: logrus.New(),
	}
}

func (api *API) Start() error {
	if err := api.configLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port: ", api.config.APPPort)
	return nil
}
