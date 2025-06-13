<script setup lang="ts">
import { AddPlayerRequestSchema, type AddPlayerResponse, type GetSessionInfo } from "~/types/session"
import type { TypedResult } from "~/types/typed-result"

const route = useRoute()

const { data: sessionResponse } = useLazyFetch<TypedResult<GetSessionInfo>>(`/api/sessions/${route.params.id}`)

const playerName = ref<string>("")

async function onSubmit() {
	const request = AddPlayerRequestSchema.safeParse({ playerName: playerName.value })

	const { data: joinResponse } = await useFetch<TypedResult<AddPlayerResponse>>(`/api/sessions/${route.params.id}/join`, {
		method: "POST",
		body: request.data,
		headers: {
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
				<small>Created at: {{ sessionResponse?.data?.createdAt?.toString() }}</small>
			</footer>
		</article>
	</form>
</template>
