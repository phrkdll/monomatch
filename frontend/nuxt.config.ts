import tailwindcss from "@tailwindcss/vite";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	modules: [
		"@nuxt/eslint",
		"@nuxt/icon",
		"@pinia/nuxt",
		"nuxt-security",
	],
	devtools: { enabled: true },
	css: ['~/assets/main.css'],
	runtimeConfig: {
		public: {
			apiBaseUrl: process.env.NUXT_PUBLIC_API_BASE_URL,
		},
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
	vite: {
		plugins: [
			tailwindcss(),
		],
	},
});
