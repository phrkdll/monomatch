export default defineEventHandler(async (event) => {
	const apiUrl = useRuntimeConfig().apiUrl

	const path = event.path.replace(/^\/api\//, "")
	const target = new URL(path, apiUrl)

	return proxyRequest(event, target.toString())
})
