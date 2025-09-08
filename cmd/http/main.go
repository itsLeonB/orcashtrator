package main

import (
	"github.com/itsLeonB/orcashtrator/internal/config"
	"github.com/itsLeonB/orcashtrator/internal/delivery/http"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	srv := http.Setup(config.Load())
	srv.ServeGracefully()
}
