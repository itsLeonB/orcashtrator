package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/ginkgo"
	"github.com/itsLeonB/orcashtrator/internal/config"
	"github.com/itsLeonB/orcashtrator/internal/provider"
)

func Setup(configs config.Config) *ginkgo.HttpServer {
	providers := provider.All(configs)

	gin.SetMode(configs.Env)
	r := gin.New()
	registerRoutes(r, configs, providers.Logger, providers.Services)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", configs.Port),
		Handler:           r,
		ReadTimeout:       configs.Timeout,
		ReadHeaderTimeout: configs.Timeout,
		WriteTimeout:      configs.Timeout,
		IdleTimeout:       configs.Timeout,
	}

	return ginkgo.NewHttpServer(srv, configs.Timeout, providers.Logger, providers.Shutdown)
}
