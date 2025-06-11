<script lang="ts" setup>
defineProps<{
	icon?: string
}>()
</script>

<template>
	<h2 class="flex gap-2 items-center">
		<Icon
			v-if="icon"
			class="mt-1"
			:name="icon"
		/>
		<slot />
	</h2>
</template>
