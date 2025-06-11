export const useColorTheme = defineStore("colorTheme", {
	state: () => ({
		light: true,
		theme: "light",
	}),
	actions: {
		async toggle() {
			this.light = !this.light
			this.theme = this.light ? "light" : "dark"
		},
	},
})
