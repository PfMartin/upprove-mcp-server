package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

func CreatePerformanceRecordToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_, ok := request.Params.Arguments.(map[string]any)["category"].(string)
	if !ok {
		return mcp.NewToolResultError("Missing or invalid category argument"), nil
	}

	return mcp.NewToolResultText("HELLO WORLD"), nil
}
