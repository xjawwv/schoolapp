<template>
  <header class="sticky top-0 z-20 flex-shrink-0 flex h-16 bg-[color:var(--color-surface)] border-b border-[color:var(--color-border)]">
    <div class="flex-1 px-4 flex justify-between">
      <div class="flex-1 flex items-center -ml-4">
        <button
          @click="isSidebarOpen = true"
          class="lg:hidden h-16 w-16 inline-flex items-center justify-center text-[color:var(--color-muted)] hover:text-[color:var(--color-heading)] focus:outline-none cursor-pointer"
        >
          <span class="sr-only">Open sidebar</span>
          <MenuIcon class="h-6 w-6" />
        </button>
      </div>
      <div class="ml-4 flex items-center md:ml-6" v-if="user">
        <div class="relative">
          <button
            @click="isMenuOpen = !isMenuOpen"
            class="flex items-center group p-2 rounded-lg hover:bg-[color:var(--color-bg)] transition duration-150 cursor-pointer focus:outline-none"
          >
            <div class="h-8 w-8 rounded-full bg-[color:var(--color-accent)]/10 text-[color:var(--color-accent)] border border-[color:var(--color-accent)]/20 flex items-center justify-center text-sm font-bold shrink-0">
              {{ user.name.charAt(0) }}
            </div>
            <div class="ml-3 hidden md:block text-left">
              <p class="text-sm font-medium text-[color:var(--color-heading)] group-hover:text-[color:var(--color-accent)] transition duration-150 leading-tight">
                {{ user.name }}
              </p>
              <p class="text-xs font-medium text-[color:var(--color-muted)] group-hover:text-[color:var(--color-text)] transition duration-150 leading-tight mt-0.5">
                {{ user.email || user.role }}
              </p>
            </div>
          </button>

          <div
            v-if="isMenuOpen"
            class="absolute right-0 mt-2 w-56 bg-[color:var(--color-surface)] border border-[color:var(--color-border)] shadow-[--shadow-md] py-1 z-30 flex flex-col"
          >
            <div class="px-4 py-3 md:hidden border-b border-[color:var(--color-border)] mb-1">
              <div class="text-sm font-semibold text-[color:var(--color-heading)]">{{ user.name }}</div>
              <div class="text-[10px] text-[color:var(--color-muted)] uppercase tracking-wider mt-0.5">{{ user.role }}</div>
            </div>

            <button
              @click="toggleTheme"
              class="flex items-center gap-3 px-4 py-2.5 text-xs uppercase tracking-wider text-[color:var(--color-text)] hover:bg-[color:var(--color-bg)] hover:text-[color:var(--color-heading)] transition duration-150 cursor-pointer text-left w-full"
            >
              <SunIcon v-if="theme === 'dark'" class="w-4 h-4 text-[color:var(--color-accent)] shrink-0" />
              <MoonIcon v-else class="w-4 h-4 text-[color:var(--color-accent)] shrink-0" />
              <span>{{ theme === 'dark' ? 'Mode Terang' : 'Mode Gelap' }}</span>
            </button>

            <div class="h-[1px] bg-[color:var(--color-border)] my-0.5"></div>

            <button
              @click="logout"
              class="flex items-center gap-3 px-4 py-2.5 text-xs uppercase tracking-wider text-[color:var(--color-error)] hover:bg-[color:var(--color-bg)] hover:opacity-90 transition duration-150 cursor-pointer text-left w-full"
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
      </div>
    </div>
  </header>
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
