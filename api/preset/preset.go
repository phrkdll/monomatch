package preset

import (
	"encoding/json"
	"net/http"

	"github.com/phrkdll/monomatch/internal/utils"
	"github.com/phrkdll/monomatch/pkg/preset"
	"github.com/phrkdll/must/pkg/must"
)

type PresetListResponse struct {
	Presets []preset.Preset `json:"presets"`
}

func listPresets(w http.ResponseWriter, r *http.Request) {
	defer must.Recover()

	response := PresetListResponse{
		Presets: []preset.Preset{
			{ID: preset.PresetId{Inner: "ipv4"}, DisplayName: "IPv4 Addresses", Description: "Randomly generated IPv4 addresses (like 127.0.0.1)."},
			{ID: preset.PresetId{Inner: "binary"}, DisplayName: "Binary Strings", Description: "Randomly generated binary strings (like 0100101)."},
			{ID: preset.PresetId{Inner: "mac"}, DisplayName: "MAC Addresses", Description: "Randomly generated MAC addresses (like de:ad:be:ef)."},
			{ID: preset.PresetId{Inner: "uuid"}, DisplayName: "UUIDs", Description: "Randomly generated UUIDs (like 123e4567-e89b-12d3-a456-426614174000)."},
			{ID: preset.PresetId{Inner: "custom"}, DisplayName: "Custom", Description: "Provide a custom list of at least 57 strings in a JSON array."},
		},
	}

	json := must.Return(json.Marshal(&response)).ElseRespond(w, http.StatusBadRequest)

	utils.SendJSON(w, http.StatusOK, json)
}
