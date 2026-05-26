<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-3xl w-full mx-auto space-y-6 sm:space-y-8">
        <div class="flex items-center space-x-4">
          <NuxtLink
            to="/announcement"
            class="p-2 rounded-md border border-[color:var(--color-border)] bg-[color:var(--color-surface)] text-[color:var(--color-muted)] hover:text-[color:var(--color-heading)] hover:bg-[color:var(--color-bg)] transition duration-150 cursor-pointer"
          >
            <ArrowLeftIcon class="w-5 h-5" />
          </NuxtLink>
          <div>
            <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-1">
              Detail Notifikasi
            </h2>
            <h1 class="text-2xl font-bold text-[color:var(--color-heading)] tracking-tight">
              Riwayat Pengumuman
            </h1>
          </div>
        </div>

        <div v-if="loading" class="flex items-center justify-center py-20">
          <div class="animate-spin rounded-full h-8 w-8 border-2 border-[color:var(--color-accent)] border-t-transparent"></div>
        </div>

        <div v-else-if="error" class="bg-[color:var(--color-surface)] border border-[color:var(--color-error)] p-6 text-center space-y-3">
          <AlertCircleIcon class="w-10 h-10 text-[color:var(--color-error)] mx-auto" />
          <p class="text-sm text-[color:var(--color-error)] font-medium">{{ error }}</p>
          <NuxtLink to="/announcement" class="button-ghost text-xs inline-block mt-2 py-2 px-4">
            Kembali ke Papan Pengumuman
          </NuxtLink>
        </div>

        <div v-else-if="notification" class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] overflow-hidden">
          <div class="px-6 sm:px-8 py-5 border-b border-[color:var(--color-border)] bg-[color:var(--color-bg)]">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 flex items-center justify-center bg-[color:var(--color-accent)]/10 border border-[color:var(--color-accent)]/20 rounded-lg shrink-0">
                <BellIcon class="w-5 h-5 text-[color:var(--color-accent)]" />
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold mb-0.5">Notifikasi Sistem</p>
                <div class="flex items-center gap-2 text-xs text-[color:var(--color-muted)] font-mono">
                  <ClockIcon class="w-3.5 h-3.5 shrink-0" />
                  <span>{{ formatTime(notification.created_at) }}</span>
                </div>
              </div>
            </div>
          </div>

          <div class="px-6 sm:px-8 py-6 sm:py-8 space-y-4">
            <div>
              <p class="text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold mb-2">Judul</p>
              <h3 class="text-xl font-bold text-[color:var(--color-heading)] leading-snug">
                {{ notification.title }}
              </h3>
            </div>

            <div class="border-t border-[color:var(--color-border)] pt-4">
              <p class="text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold mb-2">Deskripsi</p>
              <p class="text-sm text-[color:var(--color-text)] leading-relaxed whitespace-pre-wrap">
                {{ notification.description }}
              </p>
            </div>
          </div>

          <div class="px-6 sm:px-8 py-4 border-t border-[color:var(--color-border)] bg-[color:var(--color-bg)] flex items-center justify-between">
            <p class="text-xs text-[color:var(--color-muted)] font-mono">
              ID: {{ notification.id }}
            </p>
            <NuxtLink
              to="/announcement"
              class="button-ghost text-xs font-semibold py-2 px-4 flex items-center gap-1.5"
            >
              <ArrowLeftIcon class="w-3.5 h-3.5" />
              <span>Kembali</span>
            </NuxtLink>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import {
  ArrowLeft as ArrowLeftIcon,
  AlertCircle as AlertCircleIcon,
  Bell as BellIcon,
  Clock as ClockIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const route = useRoute()
const api = useApi()
const notification = ref<any>(null)
const loading = ref(true)
const error = ref("")

function formatTime(dateStr: string): string {
  if (!dateStr) return "-"
  const d = new Date(dateStr)
  return d.toLocaleString("id-ID", {
    year: "numeric",
    month: "long",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit"
  })
}

async function loadNotification() {
  loading.value = true
  error.value = ""
  try {
    const res: any = await api.get(`/api/notifications/${route.params.id}`)
    if (res.success && res.data) {
      notification.value = res.data
    } else {
      error.value = res.message || "Notifikasi tidak ditemukan"
    }
  } catch (err: any) {
    error.value = err.data?.message || "Gagal memuat detail notifikasi"
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadNotification()
})
</script>
