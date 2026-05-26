<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8">
        <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
          <div>
            <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
              Komunikasi Sistem
            </h2>
            <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
              Papan Pengumuman
            </h1>
          </div>
          <div class="flex items-center gap-3" v-if="currentUser?.role === 'admin'">
            <NuxtLink
              to="/announcement/notification/create"
              class="button-primary py-2.5 px-4 flex items-center text-xs cursor-pointer"
            >
              <PlusIcon class="w-4 h-4 mr-2" />
              <span>Buat Notifikasi</span>
            </NuxtLink>
            <NuxtLink
              to="/announcement/popup/create"
              class="button-ghost py-2.5 px-4 flex items-center text-xs cursor-pointer"
            >
              <ImageIcon class="w-4 h-4 mr-2 text-[color:var(--color-accent)]" />
              <span>Buat Popup</span>
            </NuxtLink>
          </div>
        </div>

        <div v-if="toast.message" :class="[
          'border p-3 text-sm font-medium flex items-center space-x-2',
          toast.type === 'success' ? 'bg-[color:var(--color-surface)] border-[color:var(--color-success)] text-[color:var(--color-success)]' : 'bg-[color:var(--color-surface)] border-[color:var(--color-error)] text-[color:var(--color-error)]'
        ]">
          <CheckCircleIcon v-if="toast.type === 'success'" class="w-4 h-4 shrink-0" />
          <AlertCircleIcon v-else class="w-4 h-4 shrink-0" />
          <span>{{ toast.message }}</span>
        </div>

        <HeadlessTabGroup>
          <div class="flex border-b border-[color:var(--color-border)]">
            <HeadlessTabList class="flex space-x-6">
              <HeadlessTab
                v-slot="{ selected }"
                as="template"
              >
                <button
                  :class="[
                    'pb-4 text-sm font-semibold uppercase tracking-wider transition duration-150 cursor-pointer focus:outline-none border-b-2',
                    selected
                      ? 'border-[color:var(--color-accent)] text-[color:var(--color-accent)] font-bold'
                      : 'border-transparent text-[color:var(--color-muted)] hover:text-[color:var(--color-text)]'
                  ]"
                >
                  Riwayat Notifikasi
                </button>
              </HeadlessTab>
              <HeadlessTab
                v-slot="{ selected }"
                as="template"
              >
                <button
                  :class="[
                    'pb-4 text-sm font-semibold uppercase tracking-wider transition duration-150 cursor-pointer focus:outline-none border-b-2',
                    selected
                      ? 'border-[color:var(--color-accent)] text-[color:var(--color-accent)] font-bold'
                      : 'border-transparent text-[color:var(--color-muted)] hover:text-[color:var(--color-text)]'
                  ]"
                >
                  Riwayat Popup Visual
                </button>
              </HeadlessTab>
            </HeadlessTabList>
          </div>

          <HeadlessTabPanels class="pt-6">
            <HeadlessTabPanel class="space-y-6 focus:outline-none">
              <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6">
                <div class="flex flex-col sm:flex-row gap-4 items-center justify-between">
                  <div>
                    <h3 class="text-lg font-bold text-[color:var(--color-heading)] tracking-wide">
                      Daftar Riwayat Notifikasi
                    </h3>
                    <p class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider mt-0.5">
                      Notifikasi otomatis presensi dan siaran pengumuman sistem
                    </p>
                  </div>
                  <button
                    v-if="currentUser?.role === 'admin' && listNotifications.length > 0"
                    type="button"
                    @click="clearNotifications"
                    class="button-ghost border border-[color:var(--color-error)] text-[color:var(--color-error)] hover:bg-[color:var(--color-error)]/5 text-xs font-semibold uppercase tracking-wider py-2.5 px-4 flex items-center space-x-1.5"
                    :disabled="clearingNotifications"
                  >
                    <TrashIcon class="w-4 h-4" />
                    <span>Hapus Semua Notifikasi</span>
                  </button>
                </div>

                <div class="overflow-x-auto">
                  <table class="w-full text-left border-collapse">
                    <thead>
                      <tr class="border-b border-[color:var(--color-border)]">
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-12">No</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-60">Judul</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Deskripsi</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-48">Waktu</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-28 text-center">Aksi</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr
                        v-for="(n, index) in listNotifications"
                        :key="n.id"
                        class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                      >
                        <td class="py-3.5 px-4 text-sm font-mono text-[color:var(--color-accent)]">{{ index + 1 }}</td>
                        <td class="py-3.5 px-4 text-sm font-bold text-[color:var(--color-heading)]">{{ n.title }}</td>
                        <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">{{ n.description }}</td>
                        <td class="py-3.5 px-4 text-xs text-[color:var(--color-muted)] font-mono">
                          {{ formatTime(n.created_at) }}
                        </td>
                        <td class="py-3.5 px-4 text-center">
                          <NuxtLink
                            :to="`/announcement/notification/${n.id}`"
                            class="inline-flex items-center gap-1.5 text-xs font-semibold uppercase tracking-wider text-[color:var(--color-accent)] hover:text-[color:var(--color-heading)] transition duration-150 py-1.5 px-3 border border-[color:var(--color-border)] hover:border-[color:var(--color-accent)] bg-[color:var(--color-surface)] hover:bg-[color:var(--color-accent)]/5"
                          >
                            <EyeIcon class="w-3.5 h-3.5" />
                            <span>Detail</span>
                          </NuxtLink>
                        </td>
                      </tr>
                      <tr v-if="listNotifications.length === 0">
                        <td colspan="5" class="py-12 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                          Belum ada riwayat notifikasi
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
            </HeadlessTabPanel>

            <HeadlessTabPanel class="space-y-6 focus:outline-none">
              <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6">
                <div class="flex flex-col sm:flex-row gap-4 items-center justify-between">
                  <div>
                    <h3 class="text-lg font-bold text-[color:var(--color-heading)] tracking-wide">
                      Daftar Riwayat Popup Visual
                    </h3>
                    <p class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider mt-0.5">
                      Siaran pengumuman gambar melayang darurat di layar utama
                    </p>
                  </div>
                  <button
                    v-if="currentUser?.role === 'admin' && listPopups.length > 0"
                    type="button"
                    @click="clearPopups"
                    class="button-ghost border border-[color:var(--color-error)] text-[color:var(--color-error)] hover:bg-[color:var(--color-error)]/5 text-xs font-semibold uppercase tracking-wider py-2.5 px-4 flex items-center space-x-1.5"
                    :disabled="clearingPopups"
                  >
                    <TrashIcon class="w-4 h-4" />
                    <span>Hapus Semua Popup</span>
                  </button>
                </div>

                <div class="overflow-x-auto">
                  <table class="w-full text-left border-collapse">
                    <thead>
                      <tr class="border-b border-[color:var(--color-border)]">
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-12">No</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-60">Judul</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-36">Gambar</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-48">Waktu</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-28 text-center">Aksi</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr
                        v-for="(p, index) in listPopups"
                        :key="p.id"
                        class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                      >
                        <td class="py-3.5 px-4 text-sm font-mono text-[color:var(--color-accent)]">{{ index + 1 }}</td>
                        <td class="py-3.5 px-4 text-sm font-bold text-[color:var(--color-heading)]">{{ p.title }}</td>
                        <td class="py-3.5 px-4">
                          <div class="h-10 w-16 bg-[color:var(--color-bg)] rounded border border-[color:var(--color-border)] overflow-hidden">
                            <img :src="getPhotoUrl(p.photo)" class="w-full h-full object-cover" />
                          </div>
                        </td>
                        <td class="py-3.5 px-4 text-xs text-[color:var(--color-muted)] font-mono">
                          {{ formatTime(p.created_at) }}
                        </td>
                        <td class="py-3.5 px-4 text-center">
                          <NuxtLink
                            :to="`/announcement/popup/${p.id}`"
                            class="inline-flex items-center gap-1.5 text-xs font-semibold uppercase tracking-wider text-[color:var(--color-accent)] hover:text-[color:var(--color-heading)] transition duration-150 py-1.5 px-3 border border-[color:var(--color-border)] hover:border-[color:var(--color-accent)] bg-[color:var(--color-surface)] hover:bg-[color:var(--color-accent)]/5"
                          >
                            <EyeIcon class="w-3.5 h-3.5" />
                            <span>Detail</span>
                          </NuxtLink>
                        </td>
                      </tr>
                      <tr v-if="listPopups.length === 0">
                        <td colspan="5" class="py-12 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                          Belum ada riwayat popup visual
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
            </HeadlessTabPanel>
          </HeadlessTabPanels>
        </HeadlessTabGroup>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue"
import {
  AlertCircle as AlertCircleIcon,
  CheckCircle as CheckCircleIcon,
  Plus as PlusIcon,
  Image as ImageIcon,
  Trash2 as TrashIcon,
  Eye as EyeIcon
} from "lucide-vue-next"
import io from "socket.io-client"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()
const config = useRuntimeConfig()
const currentUser = useState<any>("currentUser", () => null)
const listNotifications = ref<any[]>([])
const listPopups = ref<any[]>([])
const clearingNotifications = ref(false)
const clearingPopups = ref(false)

const toast = ref<{ message: string; type: "success" | "error" }>({
  message: "",
  type: "success"
})

function showToast(message: string, type: "success" | "error") {
  toast.value = { message, type }
  setTimeout(() => {
    toast.value.message = ""
  }, 4000)
}

function getPhotoUrl(photo: string): string {
  if (!photo) return ""
  if (photo.startsWith("http")) return photo
  return `${config.public.apiUrl}${photo}`
}

async function loadNotificationsList() {
  try {
    const res: any = await api.get("/api/notifications")
    if (res.success && res.data) {
      listNotifications.value = res.data
    }
  } catch (error) {
    showToast("Gagal memuat riwayat notifikasi", "error")
  }
}

async function loadPopupsList() {
  try {
    const res: any = await api.get("/api/popups")
    if (res.success && res.data) {
      listPopups.value = res.data
    }
  } catch (error) {
    showToast("Gagal memuat riwayat popup", "error")
  }
}

async function clearNotifications() {
  clearingNotifications.value = true
  try {
    const res: any = await api.delete("/api/notifications")
    if (res.success) {
      showToast("Seluruh riwayat notifikasi berhasil dibersihkan", "success")
      listNotifications.value = []
    }
  } catch (error) {
    showToast("Gagal membersihkan notifikasi", "error")
  } finally {
    clearingNotifications.value = false
  }
}

async function clearPopups() {
  clearingPopups.value = true
  try {
    const res: any = await api.delete("/api/popups")
    if (res.success) {
      showToast("Seluruh riwayat popup visual berhasil dibersihkan", "success")
      listPopups.value = []
    }
  } catch (error) {
    showToast("Gagal membersihkan popup", "error")
  } finally {
    clearingPopups.value = false
  }
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

let socket: any = null

onMounted(() => {
  const cached = localStorage.getItem("user")
  if (cached) {
    currentUser.value = JSON.parse(cached)
  }
  loadNotificationsList()
  loadPopupsList()

  const token = localStorage.getItem("token")
  if (token) {
    const socketUrl = config.public.apiUrl || "http://localhost:8081"
    socket = io(socketUrl, {
      transports: ["websocket"]
    })

    socket.on("connect", () => {
      console.log("Websocket connected on announcement page")
    })

    socket.on("connect_error", (err: any) => {
      console.error("Websocket connection error on announcement page:", err)
    })

    socket.on("notification", (data: any) => {
      const now = new Date()
      listNotifications.value.unshift({
        id: data.id || Math.random().toString(),
        title: data.title || "Pemberitahuan Baru",
        description: data.description || "",
        created_at: data.created_at || now.toISOString()
      })
    })

    socket.on("popup", (data: any) => {
      const now = new Date()
      listPopups.value.unshift({
        id: data.id || Math.random().toString(),
        title: data.title || "Pengumuman Visual Baru",
        photo: data.photo || "",
        created_at: data.created_at || now.toISOString()
      })
    })
  }
})

onUnmounted(() => {
  if (socket) {
    socket.disconnect()
  }
})
</script>
