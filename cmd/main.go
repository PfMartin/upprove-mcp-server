package main

import (
	"PfMartin/upprove-mcp-server/config"
	"PfMartin/upprove-mcp-server/internal/db"
	"PfMartin/upprove-mcp-server/internal/server/tools"
	"PfMartin/upprove-mcp-server/logging"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
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

	toolsHandler := tools.NewToolsHandler(dbStore)

	mcpServer := server.NewMCPServer("upprove-mcp-server", "1.0.0")
	createPerformanceRecordTool := mcp.NewTool(
		"create performance record",
		mcp.WithDescription("Insert a new performance record into the database"),
		mcp.WithString("performanceRecord", mcp.Description("The JSON for the performance record that should be inserted into the database"), mcp.Required()),
	)
	mcpServer.AddTool(createPerformanceRecordTool, toolsHandler.CreatePerformanceRecordToolHandler)

	// Run as stdio server
	if err := server.ServeStdio(mcpServer); err != nil {
		log.Err(err).Msgf("Server error: %v", err)
		return
	}
}
