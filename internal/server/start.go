package server

import (
	"fmt"
	"log"

	"github.com/bifidokk/go-api/internal/config"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func Start(conf *config.Config) {
	initSentry(conf.Env.SentryDsn)

	router := gin.Default()
	router.Use(sentrygin.New(sentrygin.Options{}))

	if err := router.SetTrustedProxies(nil); err != nil {
		log.Printf("server: %s", err)
	}

	config.RegisterRepositories(conf)
	config.RegisterServices(conf)
	config.RegisterValidators(conf)
	RegisterRoutes(router, conf)

	router.Run(":8080")
}

func initSentry(sentryDsn string) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDsn,
		EnableTracing: true,
		TracesSampleRate: 5.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
	}
}
