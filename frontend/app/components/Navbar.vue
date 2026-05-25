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
      <button @click="logout" class="flex items-center space-x-2 text-xs uppercase tracking-wider text-[color:var(--color-error)] hover:opacity-80 transition duration-150 cursor-pointer">
        <LogOutIcon class="w-4 h-4" />
        <span>Keluar</span>
      </button>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import { LogOut as LogOutIcon } from "lucide-vue-next"

const user = ref<{ name: string; email: string; role: string } | null>(null)

onMounted(() => {
  const cached = localStorage.getItem("user")
  if (cached) {
    user.value = JSON.parse(cached)
  }
})

function logout() {
  localStorage.removeItem("token")
  localStorage.removeItem("user")
  navigateTo("/login")
}
</script>
