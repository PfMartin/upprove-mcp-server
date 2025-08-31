package main

import (
	"PfMartin/upprove-mcp-server/config"
	"PfMartin/upprove-mcp-server/internal/db"
	"PfMartin/upprove-mcp-server/internal/models"
	"PfMartin/upprove-mcp-server/logging"
	"context"
	"encoding/json"
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
		Description: "Another description",
		Value:       "1",
		Unit:        "hours",
	}

	insertedID, err := upproveStore.CreatePerformanceRecord(ctx, performanceRecordCreate)
	if err != nil {
		log.Err(err).Msg("failed to insert performance record`")
	}

	fmt.Println(insertedID)

	performanceRecords, err := upproveStore.GetAllPerformanceRecords(ctx)
	if err != nil {
		log.Err(err).Msg("failed to get performance records`")
		cancel()
		return
	}

	jsonRecords, err := json.Marshal(performanceRecords)
	if err != nil {
		log.Err(err).Msg("failed to marshal array of performance records to json")
		return
	}

	fmt.Println(string(jsonRecords))

	defer cancel()
	log.Info().Msg("Starting Upprove MCP Server...")
}
