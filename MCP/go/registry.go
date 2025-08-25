package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_networkstatus "github.com/apimanagementclient/mcp-server/tools/networkstatus"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_networkstatus.CreateNetworkstatus_listbylocationTool(cfg),
		tools_networkstatus.CreateNetworkstatus_listbyserviceTool(cfg),
	}
}
