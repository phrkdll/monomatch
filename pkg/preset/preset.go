package preset

import "github.com/phrkdll/strongoid/pkg/strongoid"

//go:generate go run github.com/phrkdll/strongoid/cmd/gen --modules=json,

type PresetId strongoid.Id[string]

type Preset struct {
	ID          PresetId `json:"id"`
	DisplayName string   `json:"displayName"`
	Description string   `json:"description"`
}
