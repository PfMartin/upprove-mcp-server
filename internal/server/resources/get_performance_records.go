package resources

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
)

func (resourcesHandler *ResourcesHandler) GetPerformanceRecords(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	performanceRecords, err := resourcesHandler.DbStore.GetAllPerformanceRecords(ctx)
	if err != nil {
		return nil, err
	}

	performanceRecordsBytes, err := json.Marshal(performanceRecords)
	if err != nil {
		return nil, err
	}

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      "performanceRecords://all",
			MIMEType: "application/json",
			Text:     string(performanceRecordsBytes),
		},
	}, nil
}
