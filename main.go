package main

import (
	"github/jabahum/emr-log-analyser/cmd"
	"github/jabahum/emr-log-analyser/util"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("Cannot load config %v" + err.Error())
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		log.Info().Msg("Running in " + config.Environment + " mode")
	} else {
		log.Logger.Info().Msg("Running in " + config.Environment + " mode")
	}

	cmd.Execute()
}
