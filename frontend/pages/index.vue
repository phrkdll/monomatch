<script setup lang="ts">
import type { Preset, PresetListResponse } from "~/types/presets/presets"
import type { CreateSessionResponse } from "~/types/sessions/create"

const busy = ref(false)

const name = ref("")
const selectedPreset = ref("")

const { data } = await useFetch<PresetListResponse>("/api/presets", {
	method: "GET",
})

const presets = ref<Preset[]>([])
presets.value = data.value?.presets || []
console.log(data.value, presets.value)

async function onSubmit() {
	const syms: string[] = []
	for (let i = 0; i < 57; i++) {
		syms.push((i + 1).toString())
	}

	const createResponse = await $fetch<CreateSessionResponse>("/api/sessions/create", {
		method: "POST",
		body: {
			sessionName: name.value.toString(),
			presetId: selectedPreset.value,
			symbols: selectedPreset.value === "custom" ? null : syms,
		},
	})
	console.log(createResponse)

	if (createResponse) {
		navigateTo(`/sessions/${createResponse.id}`)
	}

	console.log(presets.value)
}
</script>

<template>
	<div>
		<form @submit.prevent="onSubmit">
			<article
				class="w-96 mx-auto"
				:aria-busy="busy"
			>
				<header>Welcome to monomatch!</header>
				<p>
					A simple, yet flexible symbol matching game.
				</p>

				<input
					v-model="name"
					v-focus
					type="text"
					placeholder="Session name"
					name="sessionName"
				>

				<select
					id="presets"
					v-model="selectedPreset"
					name="presets"
				>
					<option
						selected
						disabled
						value=""
					>
						Select preset...
					</option>
					<option
						v-for="preset in presets"
						:key="preset.id"
						:value="preset.id"
					>
						{{ preset.displayName }}
					</option>
					<option value="custom">
						Custom
					</option>
				</select>

				<footer class="flex flex-col gap-2">
					<button class="w-full">
						Create room
					</button>
					<button
						type="button"
						class="w-full"
					>
						Join room
					</button>
				</footer>
			</article>
		</form>
	</div>
</template>

<style scoped>
.v-enter-active,
.v-leave-active {
	transition: opacity 0.5s ease;
	transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
	opacity: 0;
	opacity: 0;
}
</style>
