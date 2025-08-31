package main

import (
	"PfMartin/upprove-mcp-server/config"
	"PfMartin/upprove-mcp-server/internal/db"
	"PfMartin/upprove-mcp-server/internal/models"
	"PfMartin/upprove-mcp-server/logging"
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())

	upproveStore := db.NewMongoDbStore(conf.DBName, conf.DBUser, conf.DBPassword, conf.DBURI)

	performanceRecordCreate := models.PerformanceRecordCreate{
		Category:    "Sports",
		Description: "Any Description",
		Value:       "1",
		Unit:        "hours",
	}

	insertedID, err := upproveStore.CreatePerformanceRecord(ctx, performanceRecordCreate)
	if err != nil {
		log.Err(err).Msg("failed to insert performance record`")
		cancel()
		return
	}

	fmt.Println(insertedID)

	performanceRecords, err := upproveStore.GetAllPerformanceRecords(ctx)
	if err != nil {
		log.Err(err).Msg("failed to get performance records`")
		cancel()
		return
	}

	fmt.Println(performanceRecords)

	defer cancel()
	log.Info().Msg("Starting Upprove MCP Server...")
}
