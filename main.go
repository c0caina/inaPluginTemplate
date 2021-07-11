package main

import (
	"embed"
	"github.com/df-mc/dragonfly/server"
	"github.com/sirupsen/logrus"
)

//go:embed plugins
var assets embed.FS

func Start(logger *logrus.Logger, server *server.Server) embed.FS {
	logger.Info("I running")

	return assets
}
