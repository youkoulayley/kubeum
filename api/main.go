package main

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"github.com/youkoulayley/kubeum/api/bootstrap"
	"github.com/youkoulayley/kubeum/api/models"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	// Get the version number
	v := "0.1"
	logrus.Info("Application is loading in v", v)

	// Get the configuration
	cfg := models.Config{}
	err := env.Parse(&cfg)
	if err != nil {
		logrus.Error(err)
	}

	// Setup log level
	var level logrus.Level
	switch strings.ToLower(cfg.LogLevel) {
	case "debug":
		level = logrus.DebugLevel
	case "info":
		level = logrus.InfoLevel
	case "warning":
		level = logrus.WarnLevel
	case "error":
		level = logrus.ErrorLevel
	default:
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)

	logrus.Debugf("Env loaded: %+v", cfg)

	portStr := strconv.Itoa(cfg.Port)

	bootstrap.SetupClient()

	// Initialize router
	r := initializeRouter() // Init router
	logrus.Infof("Start to listen on %s ...", portStr)
	logrus.Fatal(http.ListenAndServe(":"+portStr, r))
}
