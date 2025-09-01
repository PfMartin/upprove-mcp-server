package tools

import "PfMartin/upprove-mcp-server/internal/db"

type ToolsHandler struct {
	DbStore db.DbStore
}

func NewToolsHandler(dbStore db.DbStore) *ToolsHandler {
	return &ToolsHandler{
		DbStore: dbStore,
	}
}
