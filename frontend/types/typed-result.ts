export interface TypedResult<T> {
	success: boolean
	error: boolean
	message: string | null
	data: T
}
