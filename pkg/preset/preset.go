package preset

import (
	"errors"

	"github.com/go-faker/faker/v4"
	"github.com/phrkdll/strongoid/pkg/strongoid"
)

//go:generate go run github.com/phrkdll/strongoid/cmd/gen --modules=json,

type PresetId strongoid.Id[string]

type Preset struct {
	Id          PresetId `json:"id"`
	DisplayName string   `json:"displayName"`
	Description string   `json:"description"`
}

var (
	ErrUnknownPreset = errors.New("unknown preset")
	KnownPresets     = map[PresetId]Preset{
		{Inner: "ipv4"}:   {Id: PresetId{Inner: "ipv4"}, DisplayName: "IPv4 Addresses", Description: "Randomly generated IPv4 addresses (like 127.0.0.1)."},
		{Inner: "mac"}:    {Id: PresetId{Inner: "mac"}, DisplayName: "MAC Addresses", Description: "Randomly generated MAC addresses (like de:ad:be:ef)."},
		{Inner: "uuid"}:   {Id: PresetId{Inner: "uuid"}, DisplayName: "UUIDs", Description: "Randomly generated UUIDs (like 123e4567-e89b-12d3-a456-426614174000)."},
		{Inner: "custom"}: {Id: PresetId{Inner: "custom"}, DisplayName: "Custom", Description: "Provide a custom list of at least 57 strings in a JSON array."},
		// {Id: PresetId{Inner: "binary"}, DisplayName: "Binary Strings", Description: "Randomly generated binary strings (like 0100101)."},
	}
)

func (p Preset) ToSymbols() ([]string, error) {
	switch p.Id.Inner {
	case "ipv4":
		return generateIPv4(), nil
	case "mac":
		return generateMAC(), nil
	case "uuid":
		return generateUUID(), nil
	default:
		return []string{}, ErrUnknownPreset
	}
}

func generateIPv4() []string {
	var symbols []string
	for range 57 {
		symbols = append(symbols, faker.IPv4())
	}

	return symbols
}

func generateMAC() []string {
	var symbols []string
	for range 57 {
		symbols = append(symbols, faker.MacAddress())
	}

	return symbols
}

func generateUUID() []string {
	var symbols []string
	for range 57 {
		symbols = append(symbols, faker.UUIDHyphenated())
	}

	return symbols
}
