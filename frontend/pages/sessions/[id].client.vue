<script setup lang="ts">
import { AddPlayerRequestSchema, type AddPlayerResponse, type GetSessionInfo } from "~/types/session"
import type { TypedResult } from "~/types/typed-result"

const route = useRoute()
const sessionId = route.params.id as string

const { data: sessionResponse } = useLazyFetch<TypedResult<GetSessionInfo>>(`/api/sessions/${sessionId}`)
// const { data, send } = useWebSocket(`/api/sessions/${sessionId}/ws`)

const playerName = ref<string>("")

async function onSubmit() {
	const request = AddPlayerRequestSchema.safeParse({ playerName: playerName.value })

	const joinResponse = await $fetch<TypedResult<AddPlayerResponse>>(`/api/sessions/${sessionId}/join`, {
		method: "POST",
		body: request.data,
		headers: {
			"Accept": "application/json",
			"Content-Type": "application/json",
		},
	})

	console.log(joinResponse)
}
</script>

<template>
	<form @submit.prevent="onSubmit">
		<article>
			<header>Session: {{ sessionResponse?.data?.name }}</header>

			<fieldset role="group">
				<input
					v-model="playerName"
					type="text"
					placeholder="Player name"
				>
				<button>Join</button>
			</fieldset>
			<footer>
				<small>Created {{ $dayjs().to(sessionResponse?.data?.createdAt) }}</small>
			</footer>
		</article>
	</form>
</template>
