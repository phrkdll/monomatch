export interface TypedResult<T> {
	hasError: boolean
	errorMessage: string | null
	val: T
}
