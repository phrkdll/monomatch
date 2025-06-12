<script lang="ts" setup>
const colorTheme = useColorTheme()
const icon = computed(() => colorTheme.light ? "mdi:white-balance-sunny" : "mdi:moon-and-stars")
</script>

<template>
	<div class="bg-base-100 shadow-sm flex justify-between p-2">
		<NuxtLink
			to="/"
			class="btn btn-link text-xl"
		>monomatch</NuxtLink>

		<div class="flex gap-4">
			<NuxtLink
				to="/about"
				class="btn btn-link"
			>About</NuxtLink>
			<button
				class="btn btn-ghost"
				@click="colorTheme.toggle"
			>
				<Icon
					:name="icon"
					size="20"
				/>
			</button>
		</div>
	</div>
</template>
