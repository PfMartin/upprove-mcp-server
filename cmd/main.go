package main

import (
	"PfMartin/upprove-mcp-server/config"
	"PfMartin/upprove-mcp-server/internal/db"
	"PfMartin/upprove-mcp-server/internal/server"
	"PfMartin/upprove-mcp-server/logging"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/rs/zerolog/log"
)

func main() {
	logging.NewLogger()

	ex, err := os.Executable()
	if err != nil {
		log.Err(err).Msg("Failed to get the path of the current executable")
		return
	}

	exPath := path.Dir(ex)
	envPath := fmt.Sprintf("%s/..", exPath)

	if strings.HasPrefix(exPath, "/tmp") {
		envPath = "./"
	}

	log.Info().Msg("Loading configuration...")
	conf, err := config.NewConfig(envPath, ".env")
	if err != nil {
		log.Err(err).Msg("failed to read config")
		return
	}

	dbStore := db.NewMongoDbStore(conf.DBName, conf.DBUser, conf.DBPassword, conf.DBURI)

	mcpServer := server.NewServer(dbStore)
	mcpServer.InitResources()
	mcpServer.InitTools()

	mcpServer.ServeStdio()

}
