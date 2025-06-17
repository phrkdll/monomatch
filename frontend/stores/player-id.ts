import { v4 } from "uuid"

export const usePlayerId = defineStore("player-id", {
	state: () => ({
		id: "",
	}),
	actions: {
		async get() {
			if (!this.id) {
				this.id = v4()
			}

			return this.id
		},
	},
})
