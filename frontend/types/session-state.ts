export type SessionStateResponse = {
	id: string
	name: string
	createdAt: Date
	topMostCard: Card
	remainingCards: number
	players: Player[]
}

export type Player = {
	id: string
	playerName: string
	topMostCard: Card
	cardCount: number
	ready: boolean
}

export type Card = {
	id: string
	symbols: Symbol[]
}

export type Symbol = {
	id: string
	name: string
}
