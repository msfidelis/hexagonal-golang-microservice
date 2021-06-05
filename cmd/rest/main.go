package rest

import (
	"github.com/gin-contrib/logger"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"

	"os"
)

func Run() {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// logger
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)
	subLog := zerolog.New(os.Stdout).With().Logger()	

	// New Router
	router := gin.New()

	//Middlewares
	router.Use(gin.Recovery())

	router.Use(logger.SetLogger(logger.Config{
		Logger:   &subLog,
	}))

	router.Run()

}