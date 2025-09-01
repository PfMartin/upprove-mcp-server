package tools

import (
	"PfMartin/upprove-mcp-server/internal/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func (toolsHandler *ToolsHandler) CreatePerformanceRecordToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	performanceRecordString, ok := request.Params.Arguments.(map[string]any)["performanceRecord"].(string)
	if !ok || performanceRecordString == "" {
		return mcp.NewToolResultError("Missing or invalid performanceRecord argument"), nil
	}

	var performanceRecord models.PerformanceRecordCreate

	err := json.Unmarshal([]byte(performanceRecordString), &performanceRecord)
	if err != nil {
		return mcp.NewToolResultError("Invalid performanceRecord argument"), nil
	}

	insertedID, err := toolsHandler.DbStore.CreatePerformanceRecord(ctx, performanceRecord)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to create new performance record: %s", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Created new performance record with id: %s", insertedID)), nil
}
