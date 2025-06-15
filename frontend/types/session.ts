import z from "zod"

export const CreateSessionRequestSchema = z.object({
	sessionName: z.string().min(1, "Name is required"),
	preset: z.string(),
	symbols: z.array(z.string()).min(57, "At least 57 symbols are required").optional(),
})
export type CreateSessionRequest = z.infer<typeof CreateSessionRequestSchema>

export const CreateSessionResponseSchema = z.object({
	id: z.string(),
})
export type CreateSessionResponse = z.infer<typeof CreateSessionResponseSchema>

export const GetSessionInfoSchema = z.object({
	id: z.string(),
	name: z.string(),
	createdAt: z.date(),
})
export type GetSessionInfo = z.infer<typeof GetSessionInfoSchema>

export const AddPlayerRequestSchema = z.object({
	playerName: z.string().min(1, "Player name is required"),
})
export type AddPlayerRequest = z.infer<typeof AddPlayerRequestSchema>

export const AddPlayerResponseSchema = z.object({
	id: z.string(),
	name: z.string(),
})
export type AddPlayerResponse = z.infer<typeof AddPlayerResponseSchema>
