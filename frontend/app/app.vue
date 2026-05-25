<template>
  <NuxtPage />
</template>

<script setup lang="ts">
import { onMounted } from "vue"
import { useApi } from "~/composables/useApi"

const siteName = useState("siteName", () => "SMA N 1 METRO")
const api = useApi()

onMounted(async () => {
  const cachedTheme = localStorage.getItem("theme") || "dark"
  document.documentElement.setAttribute("data-theme", cachedTheme)

  const cachedSiteName = localStorage.getItem("siteName")
  if (cachedSiteName) {
    siteName.value = cachedSiteName
  }

  try {
    const res: any = await api.get("/api/settings")
    if (res.success && res.data && res.data.site_name) {
      siteName.value = res.data.site_name
      localStorage.setItem("siteName", res.data.site_name)
    }
  } catch (e) {
    console.error(e)
  }
})
</script>
