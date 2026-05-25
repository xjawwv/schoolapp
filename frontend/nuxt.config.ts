export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  devtools: { enabled: false },
  css: [
    "~/assets/css/index.css"
  ],
  modules: [
    "@nuxtjs/tailwindcss"
  ],
  runtimeConfig: {
    public: {
      apiUrl: process.env.VITE_API_URL || "http://localhost:8080"
    }
  }
})
