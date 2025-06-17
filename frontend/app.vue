<script setup lang="ts">
const colorTheme = useColorTheme()

const theme = computed(() => colorTheme.theme)
useHead({
	htmlAttrs: {
		"data-theme": theme,
	},
})
</script>

<template>
	<NuxtLayout>
		<NuxtPage />
	</NuxtLayout>
</template>
