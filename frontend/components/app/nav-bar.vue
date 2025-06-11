<script lang="ts" setup>
const colorTheme = useColorTheme()
const icon = computed(() => colorTheme.light ? "mdi:white-balance-sunny" : "mdi:moon-and-stars")
</script>

<template>
	<div class="navbar bg-teal-700 shadow-sm">
		<div class="flex-1">
			<NuxtLink
				class="btn btn-ghost text-xl"
				to="/"
			>monomatch</NuxtLink>
		</div>
		<div class="flex-none">
			<ul class="menu menu-horizontal px-1">
				<li>
					<NuxtLink
						class="btn btn-ghost text-xl"
						@click="colorTheme.toggle"
					>
						<Icon :name="icon" />
					</NuxtLink>
				</li>
			</ul>
		</div>
	</div>
</template>
