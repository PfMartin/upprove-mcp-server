package resources

import "PfMartin/upprove-mcp-server/internal/db"

type ResourcesHandler struct {
	DbStore db.DbStore
}

func NewResourceshandler(dbStore db.DbStore) *ResourcesHandler {
	return &ResourcesHandler{
		DbStore: dbStore,
	}
}
