<script setup lang="ts">
import { AddPlayerRequest } from "~/types/session"

const route = useRoute()
const { apiHost, apiPort } = useRuntimeConfig().public
const sessionId = route.params.id as string

const connecting = ref(true)

const { send, data } = useWebSocket(`ws://${apiHost}:${apiPort}/ws/sessions?sessionId=${sessionId}`, {
	onError: () => navigateTo("/"),
	onConnected: () => connecting.value = false,
})

const playerName = ref<string>("")

async function onSubmit() {
	const request = AddPlayerRequest.safeParse({ playerName: playerName.value })

	send(JSON.stringify(request))
}
</script>

<template>
	<form
		v-if="!connecting"
		@submit.prevent="onSubmit"
	>
		<article>
			<header> {{ connecting ? "Connecting..." : `Session: ${data?.name}` }}</header>

			<fieldset role="group">
				<input
					v-model="playerName"
					type="text"
					placeholder="Player name"
				>
				<button>Join</button>
			</fieldset>

			<pre>
				{{ data }}
			</pre>
			<footer class="flex justify-between">
				<small>Created {{ $dayjs().to(data?.createdAt) }}</small>
			</footer>
		</article>
	</form>
</template>
