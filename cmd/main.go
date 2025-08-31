package main

import (
	"PfMartin/upprove-mcp-server/logging"

	"github.com/rs/zerolog/log"
)

func main() {
	logging.NewLogger()

	log.Info().Msg("Starting Upprove MCP Server")
}
