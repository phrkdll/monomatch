<script setup lang="ts">
import type { Preset, PresetsResponse } from "~/types/presets"
import type { CreateSessionResponse } from "~/types/session"
import { CreateSessionRequestSchema } from "~/types/session"

const { data } = useLazyFetch<PresetsResponse>("/api/presets")

const sessionName = ref<string>("")
const selectedPreset = ref<Preset | undefined>(undefined)
const symbols = ref("")

async function onSubmit() {
	const request = await CreateSessionRequestSchema.safeParseAsync({
		sessionName: sessionName.value,
		preset: selectedPreset.value?.id,
		symbols: selectedPreset.value?.id === "custom" && symbols.value ? JSON.parse(symbols.value || "[]") : undefined,
	})

	console.log("Session created:", request)
	if (request.error) {
		console.log(request.error.flatten())
		return
	}

	try {
		const response = await $fetch<CreateSessionResponse>("/api/sessions/create", {
			method: "POST",
			body: request.data,
			headers: {
				"Content-Type": "application/json",
			},
		})

		console.log("Session created:", response)

		navigateTo(`/sessions/${response.id}`)
	}
	catch (error) {
		console.error("Failed to create session:", error)
	}
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
