export interface PresetListResponse {
	presets: Preset[]
}

export interface Preset {
	id: string
	displayName: string
	description: string
}
