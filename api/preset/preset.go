package preset

import (
	"encoding/json"
	"net/http"

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
		},
	}

	json := must.Return(json.Marshal(&response)).ElseRespond(w, http.StatusBadRequest)

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
