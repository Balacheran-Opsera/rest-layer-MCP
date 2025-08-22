package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/api-clarify-io/mcp-server/config"
	"github.com/api-clarify-io/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Delete_v1_bundles_bundle_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		bundle_idVal, ok := args["bundle_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: bundle_id"), nil
		}
		bundle_id, ok := bundle_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: bundle_id"), nil
		}
		url := fmt.Sprintf("%s/v1/bundles/%s", cfg.BaseURL, bundle_id)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
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
		var result map[string]interface{}
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

func CreateDelete_v1_bundles_bundle_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_v1_bundles_bundle_id",
		mcp.WithDescription("Delete a bundle"),
		mcp.WithString("bundle_id", mcp.Required(), mcp.Description("id of a bundle")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Delete_v1_bundles_bundle_idHandler(cfg),
	}
}
