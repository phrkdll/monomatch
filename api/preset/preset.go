package preset

import (
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

	var presets []preset.Preset
	for _, p := range preset.KnownPresets {
		presets = append(presets, p)
	}

	response := PresetListResponse{
		Presets: presets,
	}

	utils.SendJSON(w, http.StatusOK, response)
}
