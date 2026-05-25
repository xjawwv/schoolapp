<template>
  <div class="min-h-screen bg-[color:var(--color-bg)] flex items-center justify-center">
    <div class="animate-pulse text-sm uppercase tracking-widest text-[color:var(--color-accent)] font-semibold">
      Loading
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from "vue"

onMounted(() => {
  const token = localStorage.getItem("token")
  if (token) {
    navigateTo("/dashboard")
  } else {
    navigateTo("/login")
  }
})
</script>
