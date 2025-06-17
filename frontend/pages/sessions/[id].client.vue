<script setup lang="ts">
import { AddPlayerRequest, PlayerReadyRequest } from "~/types/session"
import type { SessionStateResponse } from "~/types/session-state"

const route = useRoute()
const playerId = await usePlayerId().get()
const { apiHost, apiPort } = useRuntimeConfig().public
const sessionId = route.params.id as string
const data = ref<SessionStateResponse>()
const localPlayer = computed(() => data.value?.players?.find(p => p.id === playerId))

const { send, close } = useWebSocket(`ws://${apiHost}:${apiPort}/ws/sessions?sessionId=${sessionId}`, {
	autoReconnect: true,
	autoClose: false,
	onError: () => navigateTo("/"),
	onMessage: (_, event: MessageEvent<string>) => data.value = JSON.parse(event.data),
})

const playerName = ref<string>("")

onUnmounted(() => close())

async function onSubmit() {
	const request = AddPlayerRequest.safeParse({ id: playerId, name: playerName.value })

	if (request.success) {
		send(JSON.stringify(request.data))
	}
}

async function onPlayerReady() {
	const request = PlayerReadyRequest.safeParse({ id: playerId, ready: !localPlayer.value?.ready })

	if (request.success) {
		send(JSON.stringify(request.data))
	}
}
</script>

<template>
	<div>
		<article>
			<header>Session: {{ data?.name }}</header>

			<form
				v-if="!localPlayer"
				@submit.prevent="onSubmit"
			>
				<fieldset role="group">
					<input
						v-model="playerName"
						type="text"
						placeholder="Player name"
					>
					<button>Join</button>
				</fieldset>
			</form>

			<table v-if="localPlayer">
				<thead>
					<tr>
						<td>Name</td>
						<td>Ready?</td>
					</tr>
				</thead>
				<tbody>
					<tr
						v-for="p in data?.players"
						:key="p.id"
					>
						<td>{{ p.playerName }} <small v-if="playerId === p.id">(you)</small></td>
						<td class="!items-center">
							<Icon
								v-if="p.ready"
								name="mynaui:check-solid"
							/>
						</td>
					</tr>
				</tbody>
			</table>

			<footer class="flex justify-between items-center">
				<button
					v-if="localPlayer"
					:class="localPlayer.ready ? 'secondary' : ''"
					class="w-full"
					@click="onPlayerReady"
				>
					{{ localPlayer.ready ? "Cancel" : "Ready" }}
				</button>
			</footer>
		</article>
		<!-- <pre>{{ JSON.stringify(data, undefined, 2) }}</pre> -->

		<AppSymbolCard
			v-if="localPlayer?.ready"
			:symbols="localPlayer.topMostCard.symbols"
			:clickable="true"
			class="mx-auto"
			@symbol-click="console.log"
		/>
	</div>
</template>
