export function useApi() {
	const config = useRuntimeConfig()
	const baseUrl = (config.public.apiUrl as string) || "http://localhost:8080"

	const request = async (url: string, options: any = {}) => {
		const token = typeof window !== "undefined" ? localStorage.getItem("token") : null

		const headers = { ...options.headers } as any

		if (!(options.body instanceof FormData)) {
			headers["Content-Type"] = "application/json"
		}

		if (token) {
			headers["Authorization"] = `Bearer ${token}`
		}

		try {
			const response = await $fetch.raw(`${baseUrl}${url}`, {
				...options,
				headers,
			})

			return response._data
		} catch (error: any) {
			if (error.response && error.response.status === 401) {
				if (typeof window !== "undefined") {
					localStorage.removeItem("token")
					localStorage.removeItem("user")
					navigateTo("/login")
				}
			}
			throw error
		}
	}

	return {
		get: (url: string, options: any = {}) => request(url, { ...options, method: "GET" }),
		post: (url: string, body: any, options: any = {}) => request(url, { ...options, method: "POST", body }),
		put: (url: string, body: any, options: any = {}) => request(url, { ...options, method: "PUT", body }),
		delete: (url: string, options: any = {}) => request(url, { ...options, method: "DELETE" }),
	}
}
