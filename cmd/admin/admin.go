package main

import (
	"fmt"
	"github.com/gorilla/mux"
	domain "multiapi"
	httputil "multiapi/pkg/http"
	"multiapi/pkg/middleware"
	"multiapi/postgres"
	"net/http"
	"time"
)

type AdminAPI struct {
	conf           *Config
	server         *http.Server
	userRepository domain.UserRepository
}

func NewAdminAPI(conf *Config) *AdminAPI {
	db := postgres.NewDB(conf.Database.URI)

	api := &AdminAPI{
		conf:           conf,
		userRepository: postgres.NewUserRepository(db),
	}

	api.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", conf.Server.Address, conf.Server.Port),
		Handler:      api.Routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return api
}

func (a *AdminAPI) Routes() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)

	a.UsersRouter(router)

	return middleware.Chain([]middleware.Middleware{
		middleware.TrimSuffixMiddleware,
		middleware.LoggingMiddleware,
	}, router)
}

type HomeResponse struct {
	Version      string
	DocsEndpoint string
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	res := HomeResponse{
		Version: "1.0.0",
	}

	httputil.RespondJSON(w, res, http.StatusOK)
}
