package main

import (
	"PfMartin/upprove-mcp-server/config"
	"PfMartin/upprove-mcp-server/internal/db"
	"PfMartin/upprove-mcp-server/logging"
	"fmt"

	"github.com/rs/zerolog/log"
)

func main() {
	logging.NewLogger()

	log.Info().Msg("Loading configuration...")
	conf, err := config.NewConfig("./", ".env")
	if err != nil {
		log.Err(err).Msg("failed to read config")
		return
	}

	upproveStore := db.NewMongoDbStore(conf.DBName, conf.DBUser, conf.DBPassword, conf.DBURI)

	fmt.Println(upproveStore)

	log.Info().Msg("Starting Upprove MCP Server...")
}
