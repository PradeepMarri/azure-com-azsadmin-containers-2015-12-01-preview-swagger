package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/storagemanagementclient/mcp-server/config"
	"github.com/storagemanagementclient/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Containers_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		subscriptionIdVal, ok := args["subscriptionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: subscriptionId"), nil
		}
		subscriptionId, ok := subscriptionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: subscriptionId"), nil
		}
		resourceGroupNameVal, ok := args["resourceGroupName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceGroupName"), nil
		}
		resourceGroupName, ok := resourceGroupNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceGroupName"), nil
		}
		farmIdVal, ok := args["farmId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: farmId"), nil
		}
		farmId, ok := farmIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: farmId"), nil
		}
		shareNameVal, ok := args["shareName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: shareName"), nil
		}
		shareName, ok := shareNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: shareName"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api-version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api-version=%v", val))
		}
		if val, ok := args["Intent"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("Intent=%v", val))
		}
		if val, ok := args["MaxCount"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("MaxCount=%v", val))
		}
		if val, ok := args["StartIndex"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("StartIndex=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/subscriptions/%s/resourcegroups/%s/providers/Microsoft.Storage.Admin/farms/%s/shares/%s/containers%s", cfg.BaseURL, subscriptionId, resourceGroupName, farmId, shareName, queryString)
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
		var result []Container
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

func CreateContainers_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_subscriptions_subscriptionId_resourcegroups_resourceGroupName_providers_Microsoft_Storage_Admin_farms_farmId_shares_shareName_containers",
		mcp.WithDescription("Returns the list of containers which can be migrated in the specified share."),
		mcp.WithString("subscriptionId", mcp.Required(), mcp.Description("Subscription Id.")),
		mcp.WithString("resourceGroupName", mcp.Required(), mcp.Description("Resource group name.")),
		mcp.WithString("farmId", mcp.Required(), mcp.Description("Farm Id.")),
		mcp.WithString("shareName", mcp.Required(), mcp.Description("Share name.")),
		mcp.WithString("api-version", mcp.Required(), mcp.Description("REST Api Version.")),
		mcp.WithString("Intent", mcp.Required(), mcp.Description("The container migration intent.")),
		mcp.WithNumber("MaxCount", mcp.Description("The maximum number of containers.")),
		mcp.WithNumber("StartIndex", mcp.Description("The starting index the resource provider uses.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Containers_listHandler(cfg),
	}
}
