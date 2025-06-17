import z from "zod"

export const CreateSessionRequest = z.object({
	sessionName: z.string().min(1, "Name is required"),
	preset: z.string(),
	symbols: z.array(z.string()).min(57, "At least 57 symbols are required").optional(),
})
export type CreateSessionRequest = z.infer<typeof CreateSessionRequest>

export const CreateSessionResponse = z.object({
	id: z.string(),
})
export type CreateSessionResponse = z.infer<typeof CreateSessionResponse>

export const GetSessionInfoResponse = z.object({
	id: z.string(),
	name: z.string(),
	createdAt: z.date(),
})
export type GetSessionInfoResponse = z.infer<typeof GetSessionInfoResponse>

export const AddPlayerRequest = z.object({
	messageType: z.string().default("add-player"),
	id: z.string().uuid(),
	name: z.string().min(1, "Player name is required"),
})
export type AddPlayerRequest = z.infer<typeof AddPlayerRequest>

export const AddPlayerResponse = z.object({
	id: z.string(),
	name: z.string(),
})
export type AddPlayerResponse = z.infer<typeof AddPlayerResponse>
