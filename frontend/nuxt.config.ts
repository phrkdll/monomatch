import tailwindcss from "@tailwindcss/vite"

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	modules: [
		"@nuxt/eslint",
		"@nuxt/icon",
		"@pinia/nuxt",
		"dayjs-nuxt",
		"@vueuse/nuxt",
	],
	devtools: { enabled: true },
	css: ["~/assets/css/main.css"],
	runtimeConfig: {
		public:	{
			apiHost: "localhost",
			apiPort: 1982,
		},
	},
	compatibilityDate: "2025-05-15", vite: {
		plugins: [
			tailwindcss(),
		],
	},
	dayjs: {
		plugins: ["utc", "timezone", "relativeTime"],
		locales: ["en"],
		defaultLocale: "en",
	},
	eslint: {
		config: {
			stylistic: {
				semi: false,
				quotes: "double",
				indent: "tab",
				commaDangle: "always-multiline",
			},
		},
	},
})
