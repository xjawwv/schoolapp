<template>
  <div class="min-h-screen bg-[color:var(--color-bg)] flex items-center justify-center px-4">
    <div class="w-full max-w-md bg-[color:var(--color-surface)] border border-[color:var(--color-border)] border-l-4 border-l-[color:var(--color-accent)] p-8 shadow-[--shadow-lg]">
      <div class="text-center mb-8">
        <h1 class="text-4xl font-bold tracking-tight mb-2 text-[color:var(--color-heading)]">
          SMA N 1 METRO
        </h1>
        <p class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-muted)]">
          Sistem Manajemen Kesiswaan
        </p>
      </div>

      <form @submit.prevent="handleLogin" class="space-y-6">
        <div v-if="errorMessage" class="bg-[color:var(--color-bg)] border border-[color:var(--color-error)] p-3 text-sm text-[color:var(--color-error)] font-medium flex items-center space-x-2">
          <AlertCircleIcon class="w-4 h-4 shrink-0" />
          <span>{{ errorMessage }}</span>
        </div>

        <div class="space-y-2">
          <label for="email" class="block text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">
            Pos Elektronik
          </label>
          <input
            id="email"
            v-model="email"
            type="email"
            class="input w-full"
            placeholder="nama@sekolah.com"
            required
            autocomplete="email"
          />
        </div>

        <div class="space-y-2">
          <label for="password" class="block text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">
            Kata Sandi
          </label>
          <input
            id="password"
            v-model="password"
            type="password"
            class="input w-full"
            placeholder="••••••••"
            required
            autocomplete="current-password"
          />
        </div>

        <button
          type="submit"
          :disabled="isLoading"
          class="button-primary w-full flex items-center justify-center space-x-2 transition duration-150 py-3"
        >
          <span v-if="isLoading" class="animate-spin rounded-full h-4 w-4 border-2 border-[color:var(--color-accent-fg)] border-t-transparent"></span>
          <span v-else>Masuk</span>
        </button>
      </form>

      <div class="mt-8 text-center text-xs text-[color:var(--color-muted)]">
        Gunakan admin@sekolah.com / admin123 untuk masuk
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { AlertCircle as AlertCircleIcon } from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: []
})

const email = ref("")
const password = ref("")
const isLoading = ref(false)
const errorMessage = ref("")

const api = useApi()

async function handleLogin() {
  if (!email.value || !password.value) {
    errorMessage.value = "Semua kolom harus diisi"
    return
  }

  isLoading.value = true
  errorMessage.value = ""

  try {
    const res = await api.post("/api/auth/login", {
      email: email.value,
      password: password.value
    })

    if (res.success && res.data) {
      localStorage.setItem("token", res.data.token)
      localStorage.setItem("user", JSON.stringify(res.data.user))
      navigateTo("/dashboard")
    } else {
      errorMessage.value = res.message || "Gagal masuk"
    }
  } catch (error: any) {
    if (error.response && error.response._data) {
      errorMessage.value = error.response._data.message || "Kredensial salah"
    } else {
      errorMessage.value = "Koneksi ke server gagal"
    }
  } finally {
    isLoading.value = false
  }
}
</script>
