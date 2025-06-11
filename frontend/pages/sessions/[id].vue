<script setup lang="ts">
const route = useRoute()
</script>

<template>
	<article>
		{{ route.params.id }}
	</article>
</template>
