import z from "zod"

const SessionRequest = z.object({
	sessionName: z.string().min(1, "Name is required"),
	preset: z.string(),
	symbols: z.array(z.string()).min(57, "At least 57 symbols are required").optional(),
})

export type SessionRequest = z.infer<typeof SessionRequest>
