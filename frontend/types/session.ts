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
