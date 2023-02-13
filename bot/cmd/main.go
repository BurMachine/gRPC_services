package main

import (
	"burmachineBot/internal/bot"
	"burmachineBot/internal/config"
	"burmachineBot/internal/grpcClient"
	"flag"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	cfgPath := flag.String("config", "./config.yaml", "Path to yaml configuration file")
	flag.Parse()

	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	conf := &config.Conf{}
	err := conf.LoadConfig(*cfgPath)
	if err != nil {
		log.Fatal().Msg("config loading failed: " + err.Error())
	}

	tgbot := &bot.Bot{
		Logger: &log,
	}
	err = tgbot.BotRegistration(conf.Apikey)
	if err != nil {
		log.Fatal().Msg("bot registration failed: " + err.Error())
	}
	grpcSinalChan := make(chan int)
	go grpcClient.RunClient(conf.AddrGrpc, &log, grpcSinalChan)
	tgbot.Run(grpcSinalChan)
}
