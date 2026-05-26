<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8">
        <div>
          <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
            Komunikasi Sistem
          </h2>
          <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
            Notifikasi Pengumuman
          </h1>
        </div>

        <div v-if="toast.message" :class="[
          'border p-3 text-sm font-medium flex items-center space-x-2',
          toast.type === 'success' ? 'bg-[color:var(--color-surface)] border-[color:var(--color-success)] text-[color:var(--color-success)]' : 'bg-[color:var(--color-surface)] border-[color:var(--color-error)] text-[color:var(--color-error)]'
        ]">
          <CheckCircleIcon v-if="toast.type === 'success'" class="w-4 h-4 shrink-0" />
          <AlertCircleIcon v-else class="w-4 h-4 shrink-0" />
          <span>{{ toast.message }}</span>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
          <div class="lg:col-span-4" v-if="currentUser?.role === 'admin'">
            <div class="card space-y-6 shadow-[--shadow-sm]">
              <div>
                <h3 class="text-lg font-bold text-[color:var(--color-heading)] tracking-wide mb-1">
                  Kirim Notifikasi
                </h3>
                <p class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider">
                  Siarkan pesan pengumuman manual ke seluruh pengguna
                </p>
              </div>

              <form @submit.prevent="sendNotification" class="space-y-4">
                <div class="space-y-1.5">
                  <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Pesan Pengumuman</label>
                  <textarea
                    v-model="messageBody"
                    class="input w-full h-32 resize-none transition duration-150"
                    placeholder="Tulis pesan pengumuman di sini..."
                    required
                  ></textarea>
                </div>

                <div class="pt-2">
                  <button type="submit" class="button-primary w-full flex items-center justify-center py-3" :disabled="sending">
                    <span v-if="sending" class="animate-spin rounded-full h-3.5 w-3.5 border-2 border-[color:var(--color-accent-fg)] border-t-transparent mr-2"></span>
                    <SendIcon class="w-4 h-4 mr-2" />
                    <span>Kirim Pengumuman</span>
                  </button>
                </div>
              </form>
            </div>
          </div>

          <div :class="(currentUser?.role === 'admin') ? 'lg:col-span-8' : 'lg:col-span-12'">
            <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6">
              <div class="flex flex-col sm:flex-row gap-4 items-center justify-between">
                <div>
                  <h3 class="text-lg font-bold text-[color:var(--color-heading)] tracking-wide">
                    Riwayat Pemberitahuan
                  </h3>
                  <p class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider mt-0.5">
                    Daftar pemberitahuan absensi otomatis dan siaran pengumuman manual
                  </p>
                </div>
                <button
                  v-if="currentUser?.role === 'admin'"
                  type="button"
                  @click="clearAllNotifications"
                  class="button-ghost border border-[color:var(--color-error)] text-[color:var(--color-error)] hover:bg-[color:var(--color-error)]/5 text-xs font-semibold uppercase tracking-wider py-2.5 px-4 flex items-center space-x-1.5"
                  :disabled="clearing"
                >
                  <TrashIcon class="w-4 h-4" />
                  <span>Bersihkan Semua</span>
                </button>
              </div>

              <div class="overflow-x-auto">
                <table class="w-full text-left border-collapse">
                  <thead>
                    <tr class="border-b border-[color:var(--color-border)]">
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-12">No</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Pesan Pemberitahuan</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-48">Waktu Dikirim</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr
                      v-for="(n, index) in listNotifications"
                      :key="n.id"
                      class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                    >
                      <td class="py-3.5 px-4 text-sm font-mono text-[color:var(--color-accent)]">{{ index + 1 }}</td>
                      <td class="py-3.5 px-4 text-sm font-semibold text-[color:var(--color-heading)]">
                        {{ n.message }}
                      </td>
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] font-mono">
                        {{ formatTime(n.created_at) }}
                      </td>
                    </tr>
                    <tr v-if="listNotifications.length === 0">
                      <td colspan="3" class="py-12 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                        Belum ada riwayat notifikasi
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue"
import {
  AlertCircle as AlertCircleIcon,
  CheckCircle as CheckCircleIcon,
  Send as SendIcon,
  Trash2 as TrashIcon
} from "lucide-vue-next"
import io from "socket.io-client"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()
const currentUser = useState<any>("currentUser", () => null)
const listNotifications = ref<any[]>([])
const messageBody = ref("")
const sending = ref(false)
const clearing = ref(false)

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

async function sendNotification() {
  if (!messageBody.value.trim()) return
  sending.value = true
  try {
    const res: any = await api.post("/api/notifications", {
      message: messageBody.value
    })
    if (res.success) {
      showToast("Pesan notifikasi berhasil disiarkan", "success")
      messageBody.value = ""
      await loadNotificationsList()
    } else {
      showToast(res.message || "Gagal mengirim notifikasi", "error")
    }
  } catch (error: any) {
    showToast(error.data?.message || "Terjadi kesalahan server", "error")
  } finally {
    sending.value = false
  }
}

async function clearAllNotifications() {
  clearing.value = true
  try {
    const res: any = await api.delete("/api/notifications")
    if (res.success) {
      showToast("Seluruh riwayat notifikasi berhasil dibersihkan", "success")
      listNotifications.value = []
    } else {
      showToast(res.message || "Gagal membersihkan notifikasi", "error")
    }
  } catch (error: any) {
    showToast(error.data?.message || "Terjadi kesalahan server", "error")
  } finally {
    clearing.value = false
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

  const token = localStorage.getItem("token")
  if (token) {
    const config = useRuntimeConfig()
    const socketUrl = config.public.apiUrl || "http://localhost:8081"
    socket = io(socketUrl, {
      transports: ["websocket"]
    })

    socket.on("connect", () => {
      console.log("Websocket connected on notification page")
    })

    socket.on("connect_error", (err: any) => {
      console.error("Websocket connection error on notification page:", err)
    })

    socket.on("disconnect", (reason: string) => {
      console.log("Websocket disconnected on notification page:", reason)
    })

    socket.on("notification", (msg: string) => {
      listNotifications.value.unshift({
        id: Math.random().toString(),
        message: msg,
        created_at: new Date().toISOString()
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
