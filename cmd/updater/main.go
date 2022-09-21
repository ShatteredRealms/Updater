package main

import (
    "github.com/ShatteredRealms/Updater/internal/log"
    "github.com/ShatteredRealms/Updater/internal/option"
    "github.com/ShatteredRealms/Updater/pkg/srv"
)

func main() {
    config := option.NewConfig()

    logLevel := log.Verbose
    if config.Mode.GetValue() == "release" {
        logLevel = log.Info
    }
    logger := log.NewLogger(logLevel, "")

    logger.Info("Creating server")
    r := srv.NewServer()

    logger.Info("Starting server")
    err := r.Run(config.Address())
    if err != nil {
        logger.Logf(log.Error, "Server shutting down: %v", err)
    }
}
