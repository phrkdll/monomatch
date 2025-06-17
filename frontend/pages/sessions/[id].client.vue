<script setup lang="ts">
import { AddPlayerRequest } from "~/types/session"
import type { SessionStateResponse } from "~/types/session-state"

const route = useRoute()
const playerId = usePlayerId()
const { apiHost, apiPort } = useRuntimeConfig().public
const sessionId = route.params.id as string

const connecting = ref(true)

const { send, data, status } = useWebSocket<SessionStateResponse>(`ws://${apiHost}:${apiPort}/ws/sessions?sessionId=${sessionId}`, {
	onError: () => navigateTo("/"),
	onConnected: () => connecting.value = false,
})

const playerName = ref<string>("")

async function onSubmit() {
	const request = AddPlayerRequest.safeParse({ id: await playerId.get(), name: playerName.value })

	if (request.success) {
		send(JSON.stringify(request.data))
	}
}
</script>

<template>
	<form
		v-if="!connecting"
		@submit.prevent="onSubmit"
	>
		<article>
			<header>Session: {{ data?.name }}</header>

			<fieldset role="group">
				<input
					v-model="playerName"
					type="text"
					placeholder="Player name"
				>
				<button>Join</button>
			</fieldset>

			<pre>{{ JSON.stringify(data, undefined, 2) }}</pre>
			<footer class="flex justify-between">
				<small>Created {{ $dayjs().to(data?.createdAt) }}</small>
				<small>Connection: {{ status }}</small>
			</footer>
		</article>
	</form>
</template>
