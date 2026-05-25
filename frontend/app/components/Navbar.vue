<template>
  <nav class="h-16 bg-[color:var(--color-surface)] border-b border-[color:var(--color-border)] px-6 flex items-center justify-between z-10">
    <div class="flex items-center space-x-4">
      <span class="text-xs tracking-[0.15em] font-semibold text-[color:var(--color-accent)] uppercase">
        Sistem Manajemen Kesiswaan
      </span>
    </div>
    <div class="flex items-center space-x-6" v-if="user">
      <div class="flex items-center space-x-3">
        <div class="w-8 h-8 rounded-full bg-[color:var(--color-border)] flex items-center justify-center text-[color:var(--color-accent)] border border-[color:var(--color-accent)] text-sm font-semibold">
          {{ user.name.charAt(0) }}
        </div>
        <div class="flex flex-col text-left">
          <span class="text-sm font-medium text-[color:var(--color-heading)]">{{ user.name }}</span>
          <span class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider">{{ user.role }}</span>
        </div>
      </div>
      <div class="flex items-center space-x-4 border-l border-[color:var(--color-border)] pl-4">
        <button @click="toggleTheme" class="flex items-center space-x-2 text-xs uppercase tracking-wider text-[color:var(--color-text)] hover:opacity-80 transition duration-150 cursor-pointer">
          <SunIcon v-if="theme === 'dark'" class="w-4 h-4 text-[color:var(--color-accent)]" />
          <MoonIcon v-else class="w-4 h-4 text-[color:var(--color-accent)]" />
          <span>{{ theme === 'dark' ? 'Terang' : 'Gelap' }}</span>
        </button>
        <button @click="logout" class="flex items-center space-x-2 text-xs uppercase tracking-wider text-[color:var(--color-error)] hover:opacity-80 transition duration-150 cursor-pointer">
          <LogOutIcon class="w-4 h-4" />
          <span>Keluar</span>
        </button>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import {
  LogOut as LogOutIcon,
  Sun as SunIcon,
  Moon as MoonIcon
} from "lucide-vue-next"

const user = ref<{ name: string; email: string; role: string } | null>(null)
const theme = ref("dark")

onMounted(() => {
  const cached = localStorage.getItem("user")
  if (cached) {
    user.value = JSON.parse(cached)
  }

  const cachedTheme = localStorage.getItem("theme") || "dark"
  theme.value = cachedTheme
  document.documentElement.setAttribute("data-theme", cachedTheme)
})

function toggleTheme() {
  const newTheme = theme.value === "dark" ? "light" : "dark"
  theme.value = newTheme
  localStorage.setItem("theme", newTheme)
  document.documentElement.setAttribute("data-theme", newTheme)
}

function logout() {
  localStorage.removeItem("token")
  localStorage.removeItem("user")
  navigateTo("/login")
}
</script>
