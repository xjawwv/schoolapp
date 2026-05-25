<template>
  <nav class="h-20 bg-[color:var(--color-surface)] border-b border-[color:var(--color-border)] px-6 lg:px-8 flex items-center justify-between z-20 relative">
    <div class="flex items-center space-x-4">
      <button @click="isSidebarOpen = true" class="lg:hidden text-[color:var(--color-text)] hover:text-[color:var(--color-heading)] cursor-pointer focus:outline-none shrink-0 p-1.5 hover:bg-[color:var(--color-bg)] transition duration-150">
        <MenuIcon class="w-6 h-6" />
      </button>
      <span class="text-xs sm:text-sm tracking-[0.12em] sm:tracking-[0.15em] font-bold text-[color:var(--color-accent)] uppercase truncate">
        Sistem Manajemen Kesiswaan
      </span>
    </div>

    <div class="relative" v-if="user">
      <button
        @click="isMenuOpen = !isMenuOpen"
        class="flex items-center space-x-3 focus:outline-none cursor-pointer group"
      >
        <div class="w-9 h-9 rounded-full bg-[color:var(--color-bg)] border border-[color:var(--color-accent)] flex items-center justify-center text-[color:var(--color-accent)] text-sm font-bold shrink-0 group-hover:opacity-90 transition duration-150 shadow-[--shadow-sm]">
          {{ user.name.charAt(0) }}
        </div>
        <div class="hidden sm:flex flex-col text-left">
          <span class="text-sm font-semibold text-[color:var(--color-heading)] group-hover:text-[color:var(--color-accent)] transition duration-150 leading-tight">{{ user.name }}</span>
          <span class="text-[10px] text-[color:var(--color-muted)] uppercase tracking-widest mt-0.5">{{ user.role }}</span>
        </div>
        <ChevronDownIcon class="w-4 h-4 text-[color:var(--color-muted)] group-hover:text-[color:var(--color-text)] transition duration-150 shrink-0" />
      </button>

      <div
        v-if="isMenuOpen"
        class="absolute right-0 mt-2 w-56 bg-[color:var(--color-surface)] border border-[color:var(--color-border)] shadow-[--shadow-md] py-2 z-30 flex flex-col"
      >
        <div class="px-4 py-2 sm:hidden border-b border-[color:var(--color-border)] mb-1">
          <div class="text-sm font-semibold text-[color:var(--color-heading)]">{{ user.name }}</div>
          <div class="text-[10px] text-[color:var(--color-muted)] uppercase tracking-wider mt-0.5">{{ user.role }}</div>
        </div>

        <button
          @click="toggleTheme"
          class="flex items-center space-x-3 px-4 py-2.5 text-xs uppercase tracking-wider text-[color:var(--color-text)] hover:bg-[color:var(--color-bg)] hover:text-[color:var(--color-heading)] transition duration-150 cursor-pointer text-left w-full"
        >
          <SunIcon v-if="theme === 'dark'" class="w-4 h-4 text-[color:var(--color-accent)] shrink-0" />
          <MoonIcon v-else class="w-4 h-4 text-[color:var(--color-accent)] shrink-0" />
          <span>{{ theme === 'dark' ? 'Mode Terang' : 'Mode Gelap' }}</span>
        </button>

        <div class="h-[1px] bg-[color:var(--color-border)] my-1"></div>

        <button
          @click="logout"
          class="flex items-center space-x-3 px-4 py-2.5 text-xs uppercase tracking-wider text-[color:var(--color-error)] hover:bg-[color:var(--color-bg)] hover:opacity-90 transition duration-150 cursor-pointer text-left w-full"
        >
          <LogOutIcon class="w-4 h-4 shrink-0" />
          <span>Keluar</span>
        </button>
      </div>

      <div
        v-if="isMenuOpen"
        @click="isMenuOpen = false"
        class="fixed inset-0 z-20 cursor-default"
      ></div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import {
  LogOut as LogOutIcon,
  Sun as SunIcon,
  Moon as MoonIcon,
  Menu as MenuIcon,
  ChevronDown as ChevronDownIcon
} from "lucide-vue-next"

const user = ref<{ name: string; email: string; role: string } | null>(null)
const theme = ref("dark")
const isSidebarOpen = useState("sidebarOpen", () => false)
const isMenuOpen = ref(false)

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
  isMenuOpen.value = false
}

function logout() {
  localStorage.removeItem("token")
  localStorage.removeItem("user")
  isMenuOpen.value = false
  navigateTo("/login")
}
</script>
