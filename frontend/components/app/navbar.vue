<script lang="ts" setup>
const colorTheme = useColorTheme()
const icon = computed(() => colorTheme.light ? "mynaui:sun" : "mynaui:moon")
</script>

<template>
	<nav class="flex items-center">
		<ul>
			<li class="!flex !items-center">
				<NuxtLink
					to="/"
					class="!flex items-center gap-1"
				>
					<Icon
						name="mynaui:home"
					/>monomatch
				</NuxtLink>
			</li>
		</ul>
		<ul>
			<li class="!flex items-center">
				<NuxtLink
					class="!flex items-center cursor-pointer"
					@click="colorTheme.toggle"
				>
					<Icon :name="icon" />
				</NuxtLink>
			</li>
		</ul>
	</nav>
</template>

<style>
a {
	text-decoration: none;
}
</style>
