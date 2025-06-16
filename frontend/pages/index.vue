<script setup lang="ts">
import type { Preset, PresetsResponse } from "~/types/presets"
import { CreateSessionRequest, CreateSessionResponse } from "~/types/session"
import type { TypedResult } from "~/types/typed-result"

const { data: presetResponse } = useLazyFetch<TypedResult<PresetsResponse>>("/api/presets")

const sessionName = ref<string>("")
const selectedPreset = ref<Preset | undefined>(undefined)
const symbols = ref("")

async function onSubmit() {
	const request = await CreateSessionRequest.safeParseAsync({
		sessionName: sessionName.value,
		preset: selectedPreset.value?.id,
		symbols: selectedPreset.value?.id === "custom" && symbols.value ? JSON.parse(symbols.value || "[]") : undefined,
	})

	if (request.error) {
		console.log(request.error.flatten())
		return
	}

	try {
		const response = await $fetch<TypedResult<CreateSessionResponse>>("/api/sessions", {
			method: "POST",
			body: request.data,
			headers: {
				"Content-Type": "application/json",
			},
		})

		const parsedResponse = CreateSessionResponse.safeParse(response.data)
		if (parsedResponse.success) {
			navigateTo(`/sessions/${parsedResponse.data.id}`)
		}
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
					v-for="preset in presetResponse?.data?.presets"
					:key="preset.id"
					:value="preset"
				>
					{{ preset.displayName }}
				</option>
			</select>
			<small v-if="selectedPreset">{{ selectedPreset.description }}</small>
			<textarea
				v-if="selectedPreset?.id === 'custom'"
				v-model="symbols"
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
