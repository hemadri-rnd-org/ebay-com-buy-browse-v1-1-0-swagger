package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/browse-api/mcp-server/config"
	"github.com/browse-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetitemsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["item_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("item_ids=%v", val))
		}
		if val, ok := args["item_group_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("item_group_ids=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/item/%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Items
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetitemsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_item",
		mcp.WithDescription("This method retrieves the details of specific items that the buyer needs to make a purchasing decision. Note: This is a (Limited Release) available only to select Partners. For this method, only the following fields are returned: eligibleForInlineCheckout, estimatedAvailabilities, itemAffiliateWebURL, itemID, ItemWebURL, legacyItemID, price, sellerItemRevision, and topRatedBuyingExperience. URLs for this method Production URL: https://api.ebay.com/buy/browse/v1/item/ Sandbox URL: https://api.sandbox.ebay.com/buy/browse/v1/item/ Request headers This method uses the X-EBAY-C-ENDUSERCTX request header to support revenue sharing for eBay Partner Networks and to improve the accuracy of shipping and delivery time estimations. For details see, Request headers in the Buying Integration Guide. Restrictions For a list of supported sites and other restrictions, see API Restrictions. eBay Partner Network: In order to be commissioned for your sales, you must use the URL returned in the itemAffiliateWebUrl field to forward your buyer to the ebay.com site."),
		mcp.WithString("item_ids", mcp.Description("A list of item IDs. Item IDs are the eBay RESTful identifier of items. RESTful Item ID Format: v1|#|# For example: v1|272394640372|0 or v1|162846450672|461882996982 In any given request, either item_ids or item_group_ids can be retrieved. Attempting to retrieve both will result in an error. In a request, multiple item_ids can be passed as comma separated values. Maximum allowed itemIDs: 20 For more information about item IDs for RESTful APIs, see the Legacy API compatibility section of the Buy APIs Overview.")),
		mcp.WithString("item_group_ids", mcp.Description("A list of item group IDs. Item group IDs are the eBay RESTful identifier of item groups. RESTful Group Item ID Format: ############ For example: 330017835749 or 330017835740 In any given request, either item_ids or item_group_ids can be retrieved. Attempting to retrieve both will result in an error. In a request, multiple item_group_ids can be passed as comma separated values. Maximum allowed itemGroupIDs: 10")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetitemsHandler(cfg),
	}
}
