export type Preset = {
	id: string
	displayName: string
	description: string
}

export type PresetsResponse = {
	presets: Preset[]
}
