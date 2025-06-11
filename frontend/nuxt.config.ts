// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	modules: [
		"@nuxt/eslint",
		"@nuxt/icon",
		"@pinia/nuxt",
		"nuxt-security",
		"@nuxtjs/tailwindcss",
	],
	devtools: { enabled: true },
	css: ["@picocss/pico/css/pico.sand.min.css"],
	runtimeConfig: {
		apiUrl: "http://localhost:1982/",
	},
	compatibilityDate: "2025-05-15",
	eslint: {
		config: {
			stylistic: {
				semi: false,
				quotes: "double",
				commaDangle: "always-multiline",
				indent: "tab",
			},
		},
	},
})
