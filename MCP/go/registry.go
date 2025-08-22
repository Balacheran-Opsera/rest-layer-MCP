package main

import (
	"github.com/api-clarify-io/mcp-server/config"
	"github.com/api-clarify-io/mcp-server/models"
	tools_bundles "github.com/api-clarify-io/mcp-server/tools/bundles"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_bundles.CreateDelete_v1_bundles_bundle_idTool(cfg),
		tools_bundles.CreateDelete_v1_bundles_bundle_id_tracksTool(cfg),
		tools_bundles.CreateDelete_v1_bundles_bundle_id_tracks_track_idTool(cfg),
		tools_bundles.CreateDelete_v1_bundles_bundle_id_metadataTool(cfg),
	}
}
