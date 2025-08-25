package main

import (
	"github.com/browse-api/mcp-server/config"
	"github.com/browse-api/mcp-server/models"
	tools_item "github.com/browse-api/mcp-server/tools/item"
	tools_search_by_image "github.com/browse-api/mcp-server/tools/search_by_image"
	tools_shopping_cart "github.com/browse-api/mcp-server/tools/shopping_cart"
	tools_item_summary "github.com/browse-api/mcp-server/tools/item_summary"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_item.CreateCheckcompatibilityTool(cfg),
		tools_search_by_image.CreateSearchbyimageTool(cfg),
		tools_shopping_cart.CreateGetshoppingcartTool(cfg),
		tools_shopping_cart.CreateAdditemTool(cfg),
		tools_shopping_cart.CreateUpdatequantityTool(cfg),
		tools_item.CreateGetitemTool(cfg),
		tools_item.CreateGetitembylegacyidTool(cfg),
		tools_shopping_cart.CreateRemoveitemTool(cfg),
		tools_item.CreateGetitemsTool(cfg),
		tools_item_summary.CreateSearchTool(cfg),
		tools_item.CreateGetitemsbyitemgroupTool(cfg),
	}
}
