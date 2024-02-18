package main

import (
	"github.com/Irurnnen/go-score/config"
	"github.com/Irurnnen/go-score/router"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Setup gin-gonic
	gin.SetMode(config.GetRelease())

	// Setup logger
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// Setup gin Router
	r := router.NewRouter()
	log.Info("Score has been successfully started!")

	// Run server
	r.Run(config.GetServer())
}
