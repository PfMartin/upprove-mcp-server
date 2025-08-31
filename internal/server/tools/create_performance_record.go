package tools

import (
	"PfMartin/upprove-mcp-server/internal/models"
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/rs/zerolog/log"
)

func CreatePerformanceRecordToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	performanceRecordString, ok := request.Params.Arguments.(map[string]any)["performanceRecord"].(string)
	if !ok || performanceRecordString == "" {
		return mcp.NewToolResultError("Missing or invalid performanceRecord argument"), nil
	}

	var performanceRecord models.PerformanceRecordCreate

	err := json.Unmarshal([]byte(performanceRecordString), &performanceRecord)
	if err != nil {
		return mcp.NewToolResultError("Invalid performanceRecord argument"), nil
	}

	recordBytes, _ := json.Marshal(performanceRecord)

	log.Info().Msg(string(recordBytes))

	return mcp.NewToolResultText("HELLO WORLD"), nil
}
