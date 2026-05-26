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
              Detail Popup Visual
            </h2>
            <h1 class="text-2xl font-bold text-[color:var(--color-heading)] tracking-tight">
              Pengumuman Gambar
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

        <div v-else-if="popup" class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] overflow-hidden">
          <div class="px-6 sm:px-8 py-5 border-b border-[color:var(--color-border)] bg-[color:var(--color-bg)]">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 flex items-center justify-center bg-[color:var(--color-accent)]/10 border border-[color:var(--color-accent)]/20 rounded-lg shrink-0">
                <ImageIcon class="w-5 h-5 text-[color:var(--color-accent)]" />
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold mb-0.5">Popup Visual Darurat</p>
                <div class="flex items-center gap-2 text-xs text-[color:var(--color-muted)] font-mono">
                  <ClockIcon class="w-3.5 h-3.5 shrink-0" />
                  <span>{{ formatTime(popup.created_at) }}</span>
                </div>
              </div>
            </div>
          </div>

          <div class="px-6 sm:px-8 py-6 sm:py-8 space-y-5">
            <div>
              <p class="text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold mb-2">Judul</p>
              <h3 class="text-xl font-bold text-[color:var(--color-heading)] leading-snug">
                {{ popup.title }}
              </h3>
            </div>

            <div class="border-t border-[color:var(--color-border)] pt-5">
              <p class="text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold mb-3">Gambar Pengumuman</p>
              <div
                class="relative border border-[color:var(--color-border)] rounded-lg overflow-hidden bg-black/5 cursor-pointer group"
                @click="openFullscreen"
              >
                <img
                  :src="getPhotoUrl(popup.photo)"
                  :alt="popup.title"
                  class="w-full max-h-[500px] object-contain mx-auto transition duration-200 group-hover:scale-[1.01]"
                />
                <div class="absolute inset-0 flex items-center justify-center opacity-0 group-hover:opacity-100 transition duration-200 bg-black/20">
                  <div class="flex items-center gap-2 bg-black/70 text-white text-xs font-semibold uppercase tracking-wider py-2 px-4 rounded-full backdrop-blur-sm">
                    <MaximizeIcon class="w-4 h-4" />
                    <span>Lihat Layar Penuh</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="px-6 sm:px-8 py-4 border-t border-[color:var(--color-border)] bg-[color:var(--color-bg)] flex items-center justify-between">
            <p class="text-xs text-[color:var(--color-muted)] font-mono">
              ID: {{ popup.id }}
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

    <Teleport to="body">
      <Transition name="lightbox">
        <div
          v-if="fullscreenOpen"
          class="fixed inset-0 z-[9999] bg-black/90 backdrop-blur-md flex items-center justify-center"
          @click.self="closeFullscreen"
        >
          <button
            @click="closeFullscreen"
            class="absolute top-5 right-5 z-10 p-2.5 rounded-full bg-white/10 hover:bg-white/20 text-white transition duration-200 cursor-pointer"
          >
            <XIcon class="w-6 h-6" />
          </button>

          <img
            v-if="popup"
            :src="getPhotoUrl(popup.photo)"
            :alt="popup.title"
            class="max-w-[92vw] max-h-[90vh] object-contain select-none rounded-lg shadow-2xl"
          />
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue"
import {
  ArrowLeft as ArrowLeftIcon,
  AlertCircle as AlertCircleIcon,
  Image as ImageIcon,
  Clock as ClockIcon,
  Maximize2 as MaximizeIcon,
  X as XIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const route = useRoute()
const api = useApi()
const config = useRuntimeConfig()
const popup = ref<any>(null)
const loading = ref(true)
const error = ref("")
const fullscreenOpen = ref(false)

function getPhotoUrl(photo: string): string {
  if (!photo) return ""
  if (photo.startsWith("http")) return photo
  return `${config.public.apiUrl}${photo}`
}

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

function openFullscreen() {
  fullscreenOpen.value = true
  document.body.style.overflow = "hidden"
}

function closeFullscreen() {
  fullscreenOpen.value = false
  document.body.style.overflow = ""
}

function handleKeydown(e: KeyboardEvent) {
  if (e.key === "Escape" && fullscreenOpen.value) {
    closeFullscreen()
  }
}

async function loadPopup() {
  loading.value = true
  error.value = ""
  try {
    const res: any = await api.get(`/api/popups/${route.params.id}`)
    if (res.success && res.data) {
      popup.value = res.data
    } else {
      error.value = res.message || "Popup tidak ditemukan"
    }
  } catch (err: any) {
    error.value = err.data?.message || "Gagal memuat detail popup"
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadPopup()
  window.addEventListener("keydown", handleKeydown)
})

onUnmounted(() => {
  window.removeEventListener("keydown", handleKeydown)
  document.body.style.overflow = ""
})
</script>

<style scoped>
.lightbox-enter-active {
  transition: opacity 0.25s ease;
}
.lightbox-leave-active {
  transition: opacity 0.2s ease;
}
.lightbox-enter-from,
.lightbox-leave-to {
  opacity: 0;
}
</style>
