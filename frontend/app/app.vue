<template>
  <NuxtPage />
</template>

<script setup lang="ts">
import { onMounted } from "vue"

onMounted(() => {
  const cachedTheme = localStorage.getItem("theme") || "dark"
  document.documentElement.setAttribute("data-theme", cachedTheme)
})
</script>
