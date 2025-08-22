package main

import (
	"github.com/storagemanagementclient/mcp-server/config"
	"github.com/storagemanagementclient/mcp-server/models"
	tools_containers "github.com/storagemanagementclient/mcp-server/tools/containers"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_containers.CreateContainers_migrationstatusTool(cfg),
		tools_containers.CreateContainers_cancelmigrationTool(cfg),
		tools_containers.CreateContainers_listTool(cfg),
		tools_containers.CreateContainers_listdestinationsharesTool(cfg),
		tools_containers.CreateContainers_migrateTool(cfg),
	}
}
