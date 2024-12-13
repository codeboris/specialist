package api

import (
	"net/http"

	"github.com/codeboris/specialist/internal/app/middleware"
	"github.com/codeboris/specialist/storage"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
)

func (a *API) configLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *API) configRouterField() {
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	a.router.Handle(prefix+"/articles/{ID}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.GetArticleByID),
	)).Methods("GET")
	// a.router.HandleFunc(prefix+"/articles/{ID}", a.GetArticleByID).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{ID}", a.DeleteArticleByID).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")
	a.router.HandleFunc(prefix+"/user/auth", a.PostToAuth).Methods("POST")
	a.router.HandleFunc(prefix+"/users", a.GetAllUsers).Methods("GET")
}

func (a *API) configStorageField() error {
	storage := storage.New(a.config.Storage)
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
