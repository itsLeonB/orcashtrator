package main

import (
	"log"

	"github.com/itsLeonB/orcashtrator/internal/config"
	"github.com/itsLeonB/orcashtrator/internal/delivery/http"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rotisserie/eris"
)

func main() {
	srv, err := http.Setup(config.Load())
	if err != nil {
		log.Fatalf("error setting up server: %s", eris.ToString(err, true))
	}
	srv.ServeGracefully()
}
