package server

import (
	"PfMartin/upprove-mcp-server/internal/db"
	"PfMartin/upprove-mcp-server/internal/server/resources"
	"PfMartin/upprove-mcp-server/internal/server/tools"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/rs/zerolog/log"
)

type Server struct {
	dbStore   db.DbStore
	mcpServer *server.MCPServer
}

func NewServer(dbStore db.DbStore) *Server {
	mcpServer := server.NewMCPServer("upprove-mcp-server", "1.0.0")

	return &Server{
		dbStore,
		mcpServer,
	}
}

func (s *Server) InitTools() {
	toolsHandler := tools.NewToolsHandler(s.dbStore)

	createPerformanceRecordTool := mcp.NewTool(
		"create performance record",
		mcp.WithDescription("Insert a new performance record into the database"),
		mcp.WithString("performanceRecord", mcp.Description("The JSON for the performance record that should be inserted into the database"), mcp.Required()),
	)

	s.mcpServer.AddTool(createPerformanceRecordTool, toolsHandler.CreatePerformanceRecordToolHandler)
}

func (s *Server) InitResources() {
	resourcesHandler := resources.NewResourceshandler(s.dbStore)

	getPerformanceRecordsResource := mcp.NewResource(
		"performanceRecords://all",
		"All performance records in the database",
		mcp.WithResourceDescription("Array of all performance records"),
		mcp.WithMIMEType("application/json"),
	)

	s.mcpServer.AddResource(getPerformanceRecordsResource, resourcesHandler.GetPerformanceRecords)
}

func (s *Server) ServeStdio() {
	httpServer := server.NewStreamableHTTPServer(s.mcpServer)

	if err := httpServer.Start(":8080"); err != nil {
		log.Err(err).Msgf("Server failed to start: %v", err)
	}
	// if err := server.ServeStdio(s.mcpServer); err != nil {
	// 	log.Err(err).Msgf("Server error: %v", err)
	// 	return
	// }
}
