<script lang="ts" setup>
const colorTheme = useColorTheme()
const icon = computed(() => colorTheme.light ? "mdi:white-balance-sunny" : "mdi:moon-and-stars")
</script>

<template>
	<nav class="container">
		<ul>
			<li>
				<strong>
					<NuxtLink to="/">monomatch</NuxtLink>
				</strong>
			</li>
		</ul>
		<ul class="flex gap-2 items-center">
			<li class="!flex !items-center">
				<NuxtLink
					class="!flex !items-center cursor-pointer"
					@click="colorTheme.toggle"
				>
					<Icon
						:name="icon"
					/>
				</NuxtLink>
			</li>
		</ul>
	</nav>
</template>

<style scoped>
a {
	text-decoration: none;
}
</style>
