<script setup lang="ts">
import type { Preset, PresetsResponse } from "~/types/presets"

const { data } = useLazyFetch<PresetsResponse>("/api/presets")

const sessionName = ref<string>("")
const selectedPreset = ref<Preset | undefined>(undefined)

async function onSubmit() {
	// Handle form submission logic here
	// For example, you might want to send the sessionName and selectedPreset to the server
	console.log("Session Name:", sessionName.value)
	console.log("Selected Preset:", selectedPreset.value)
}
</script>

<template>
	<form @submit.prevent="onSubmit">
		<article>
			<header class="font-semibold">
				Start a session
			</header>
			<input
				id="sessionName"
				v-model="sessionName"
				type="text"
				name="sessionName"
			>
			<select
				v-model="selectedPreset"
				name="selectedPreset"
				aria-label="Select a preset"
				required
			>
				<option
					:value="undefined"
					disabled
					selected
				>
					Select a preset
				</option>
				<option
					v-for="preset in data?.presets"
					:key="preset.id"
					:value="preset"
				>
					{{ preset.displayName }}
				</option>
			</select>
			<small v-if="selectedPreset">{{ selectedPreset.description }}</small>
			<textarea
				v-if="selectedPreset?.id === 'custom'"
				class="font-mono"

				rows="5"
				cols="50"
			/>
			<footer class="flex justify-end gap-2">
				<button
					type="submit"
					class="w-24"
					:disabled="!sessionName || !selectedPreset"
				>
					Create
				</button>
				<button
					type="button"
					class="w-24 secondary"
				>
					Join
				</button>
			</footer>
		</article>
	</form>
</template>
