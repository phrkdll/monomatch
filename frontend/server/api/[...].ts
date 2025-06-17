export default defineEventHandler(async (event) => {
	const { apiHost, apiPort } = useRuntimeConfig().public

	const base = `http://${apiHost}:${apiPort}`
	const target = new URL(event.path, base)

	return proxyRequest(event, target.toString())
})
