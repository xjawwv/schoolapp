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
        <HeadlessPopover class="relative mr-2">
          <HeadlessPopoverButton
            class="p-2 text-[color:var(--color-muted)] hover:text-[color:var(--color-heading)] hover:bg-[color:var(--color-bg)] rounded-lg transition duration-150 cursor-pointer focus:outline-none relative"
          >
            <BellIcon class="w-5 h-5" />
            <span v-if="unreadCount > 0" class="absolute top-1.5 right-1.5 block h-2 w-2 rounded-full bg-[color:var(--color-error)] ring-2 ring-[color:var(--color-surface)]"></span>
          </HeadlessPopoverButton>

          <transition
            enter-active-class="transition duration-200 ease-out"
            enter-from-class="translate-y-1 opacity-0"
            enter-to-class="translate-y-0 opacity-100"
            leave-active-class="transition duration-150 ease-in"
            leave-from-class="translate-y-0 opacity-100"
            leave-to-class="translate-y-1 opacity-0"
          >
            <HeadlessPopoverPanel
              class="absolute right-0 mt-2 w-80 bg-[color:var(--color-surface)] border border-[color:var(--color-border)] shadow-[--shadow-lg] z-30 focus:outline-none overflow-hidden"
            >
              <div class="px-4 py-3 border-b border-[color:var(--color-border)] flex items-center justify-between">
                <span class="text-xs uppercase tracking-wider text-[color:var(--color-heading)] font-bold">Notifikasi Masuk</span>
                <button
                  v-if="notifications.length > 0"
                  @click="clearAll"
                  class="text-[10px] text-[color:var(--color-accent)] font-semibold hover:underline cursor-pointer uppercase tracking-wider"
                >
                  Hapus Semua
                </button>
              </div>

              <div class="max-h-64 overflow-y-auto divide-y divide-[color:var(--color-border)]">
                <div v-for="(n, idx) in notifications" :key="idx" class="p-3 text-xs text-[color:var(--color-text)] hover:bg-[color:var(--color-bg)]/50 transition">
                  <p class="font-medium text-[color:var(--color-heading)]">{{ n.message }}</p>
                  <p class="text-[9px] text-[color:var(--color-muted)] mt-1 font-mono">{{ n.time }}</p>
                </div>
                <div v-if="notifications.length === 0" class="py-8 text-center text-xs text-[color:var(--color-muted)] uppercase tracking-wider">
                  Tidak ada notifikasi baru
                </div>
              </div>
            </HeadlessPopoverPanel>
          </transition>
        </HeadlessPopover>

        <HeadlessMenu as="div" class="relative">
          <HeadlessMenuButton
            class="flex items-center group p-2 rounded-lg hover:bg-[color:var(--color-bg)] transition duration-150 cursor-pointer focus:outline-none"
          >
            <div class="h-8 w-8 rounded-full overflow-hidden border border-[color:var(--color-accent)]/20 bg-[color:var(--color-bg)] flex items-center justify-center text-sm font-bold shrink-0">
              <img v-if="avatarUrl" :src="avatarUrl" class="w-full h-full object-cover" />
              <span v-else class="text-[color:var(--color-accent)]">{{ user.name.charAt(0) }}</span>
            </div>
            <div class="ml-3 hidden md:block text-left">
              <p class="text-sm font-medium text-[color:var(--color-heading)] group-hover:text-[color:var(--color-accent)] transition duration-150 leading-tight">
                {{ user.name }}
              </p>
              <p class="text-xs font-medium text-[color:var(--color-muted)] group-hover:text-[color:var(--color-text)] transition duration-150 leading-tight mt-0.5">
                {{ user.email || user.role }}
              </p>
            </div>
          </HeadlessMenuButton>

          <transition
            enter-active-class="transition duration-100 ease-out"
            enter-from-class="transform scale-95 opacity-0"
            enter-to-class="transform scale-100 opacity-100"
            leave-active-class="transition duration-75 ease-in"
            leave-from-class="transform scale-100 opacity-100"
            leave-to-class="transform scale-95 opacity-0"
          >
            <HeadlessMenuItems
              class="absolute right-0 mt-2 w-56 bg-[color:var(--color-surface)] border border-[color:var(--color-border)] shadow-[--shadow-md] py-1 z-30 focus:outline-none flex flex-col"
            >
              <div class="px-4 py-3 md:hidden border-b border-[color:var(--color-border)] mb-1">
                <div class="text-sm font-semibold text-[color:var(--color-heading)]">{{ user.name }}</div>
                <div class="text-[10px] text-[color:var(--color-muted)] uppercase tracking-wider mt-0.5">{{ user.role }}</div>
              </div>

              <HeadlessMenuItem v-slot="{ active }">
                <button
                  @click="toggleTheme"
                  :class="[
                    active ? 'bg-[color:var(--color-bg)] text-[color:var(--color-heading)]' : 'text-[color:var(--color-text)]',
                    'flex items-center gap-3 px-4 py-2.5 text-xs uppercase tracking-wider transition duration-150 cursor-pointer text-left w-full'
                  ]"
                >
                  <SunIcon v-if="theme === 'dark'" class="w-4 h-4 text-[color:var(--color-accent)] shrink-0" />
                  <MoonIcon v-else class="w-4 h-4 text-[color:var(--color-accent)] shrink-0" />
                  <span>{{ theme === 'dark' ? 'Mode Terang' : 'Mode Gelap' }}</span>
                </button>
              </HeadlessMenuItem>

              <div class="h-[1px] bg-[color:var(--color-border)] my-0.5"></div>

              <HeadlessMenuItem v-slot="{ active }">
                <button
                  @click="logout"
                  :class="[
                    active ? 'bg-[color:var(--color-bg)]' : '',
                    'flex items-center gap-3 px-4 py-2.5 text-xs uppercase tracking-wider text-[color:var(--color-error)] transition duration-150 cursor-pointer text-left w-full'
                  ]"
                >
                  <LogOutIcon class="w-4 h-4 shrink-0" />
                  <span>Keluar</span>
                </button>
              </HeadlessMenuItem>
            </HeadlessMenuItems>
          </transition>
        </HeadlessMenu>
      </div>
    </div>
  </header>

  <div class="fixed top-4 right-4 z-[9999] flex flex-col items-end space-y-3 pointer-events-none">
    <transition-group
      enter-active-class="transform ease-out duration-300 transition"
      enter-from-class="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-2"
      enter-to-class="translate-y-0 opacity-100 sm:translate-x-0"
      leave-active-class="transition ease-in duration-200"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <UiAlert
        v-for="toast in activeToasts"
        :key="toast.id"
        class="pointer-events-auto shadow-[--shadow-lg] flex items-start text-left relative backdrop-blur-sm w-[280px] sm:w-[380px]"
      >
        <BellIcon class="w-4 h-4 shrink-0 mt-0.5" />
        <div class="flex-1">
          <UiAlertTitle>{{ toast.title }}</UiAlertTitle>
          <UiAlertDescription class="mt-1">
            {{ toast.message }}
          </UiAlertDescription>
        </div>
        <button
          @click="removeToast(toast.id)"
          class="text-[color:var(--color-muted)] hover:text-[color:var(--color-heading)] transition cursor-pointer absolute right-4 top-4 p-1 hover:bg-[color:var(--color-bg)] rounded-md"
        >
          <XIcon class="w-3.5 h-3.5" />
        </button>
      </UiAlert>
    </transition-group>
  </div>

  <HeadlessDialog
    :open="showPopupAnnouncement"
    @close="showPopupAnnouncement = false"
    class="relative z-[10000]"
  >
    <div class="fixed inset-0 bg-black/75 backdrop-blur-sm" aria-hidden="true" />

    <div class="fixed inset-0 flex items-center justify-center p-4">
      <HeadlessDialogPanel
        class="w-full max-w-lg bg-[color:var(--color-surface)] border border-[color:var(--color-border)] border-l-4 border-l-[color:var(--color-accent)] shadow-2xl rounded-lg overflow-hidden p-6 flex flex-col space-y-4 relative text-left"
      >
        <button
          @click="showPopupAnnouncement = false"
          class="absolute top-4 right-4 text-[color:var(--color-muted)] hover:text-[color:var(--color-heading)] transition cursor-pointer p-1.5 hover:bg-[color:var(--color-bg)] rounded-md focus:outline-none"
        >
          <XIcon class="w-4 h-4" />
        </button>

        <div class="space-y-1">
          <span class="text-[10px] font-bold text-[color:var(--color-accent)] tracking-widest uppercase">
            Pengumuman Penting
          </span>
          <HeadlessDialogTitle class="text-xl font-black text-[color:var(--color-heading)] tracking-tight leading-tight">
            {{ activePopup.title }}
          </HeadlessDialogTitle>
        </div>

        <div class="rounded-md border border-[color:var(--color-border)] overflow-hidden w-full max-h-80 bg-black/20 flex items-center justify-center">
          <img
            v-if="activePopup.photo"
            :src="activePopup.photo"
            class="max-w-full max-h-80 object-contain"
          />
        </div>

        <div class="pt-2 flex items-center justify-end">
          <button
            @click="showPopupAnnouncement = false"
            class="button-primary text-xs py-2.5 px-5 font-semibold"
          >
            Tutup Pengumuman
          </button>
        </div>
      </HeadlessDialogPanel>
    </div>
  </HeadlessDialog>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from "vue"
import {
  LogOut as LogOutIcon,
  Sun as SunIcon,
  Moon as MoonIcon,
  Menu as MenuIcon,
  ChevronDown as ChevronDownIcon,
  Bell as BellIcon,
  X as XIcon
} from "lucide-vue-next"
import io from "socket.io-client"

const api = useApi()
const user = useState<any>("currentUser", () => null)
const config = useRuntimeConfig()
const theme = ref("dark")
const isSidebarOpen = useState("sidebarOpen", () => false)

const unreadCount = ref(0)
const notifications = ref<Array<{ message: string; time: string }>>([])
const activeToasts = ref<Array<{ id: string; title: string; message: string }>>([])
const showPopupAnnouncement = ref(false)
const activePopup = ref({ title: "", photo: "" })

let socket: any = null

const avatarUrl = computed(() => {
  if (!user.value?.avatar) return ""
  if (user.value.avatar.startsWith("http")) return user.value.avatar
  return `${config.public.apiUrl}${user.value.avatar}`
})

function removeToast(id: string) {
  activeToasts.value = activeToasts.value.filter((t) => t.id !== id)
}

function triggerToast(title: string, message: string) {
  const id = Math.random().toString()
  activeToasts.value.push({ id, title, message })
  setTimeout(() => {
    removeToast(id)
  }, 5000)
}

function getPhotoUrl(photo: string): string {
  if (!photo) return ""
  if (photo.startsWith("http")) return photo
  return `${config.public.apiUrl}${photo}`
}

async function fetchNotifications() {
  try {
    const res: any = await api.get("/api/notifications")
    if (res.success && res.data) {
      notifications.value = res.data.map((n: any) => {
        const d = new Date(n.created_at)
        const timeStr = `${d.getHours().toString().padStart(2, '0')}:${d.getMinutes().toString().padStart(2, '0')}:${d.getSeconds().toString().padStart(2, '0')}`
        return { message: `${n.title}: ${n.description}`, time: timeStr }
      })
      unreadCount.value = res.data.filter((n: any) => !n.is_read).length
    }
  } catch (error) {
    console.error("Gagal memuat notifikasi", error)
  }
}

onMounted(() => {
  const cached = localStorage.getItem("user")
  if (cached) {
    user.value = JSON.parse(cached)
  }

  const cachedTheme = localStorage.getItem("theme") || "dark"
  theme.value = cachedTheme
  document.documentElement.setAttribute("data-theme", cachedTheme)

  const token = localStorage.getItem("token")
  if (token) {
    fetchNotifications()

    const socketUrl = config.public.apiUrl || "http://localhost:8081"
    socket = io(socketUrl, {
      transports: ["websocket"]
    })

    socket.on("connect", () => {
      console.log("Websocket connected to:", socketUrl)
    })

    socket.on("connect_error", (err: any) => {
      console.error("Websocket connection error:", err)
    })

    socket.on("disconnect", (reason: string) => {
      console.log("Websocket disconnected:", reason)
    })

    socket.on("notification", (data: any) => {
      const now = new Date()
      const timeStr = `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}:${now.getSeconds().toString().padStart(2, '0')}`
      const title = data.title || "Pemberitahuan Baru"
      const desc = data.description || (typeof data === "string" ? data : "")
      notifications.value.unshift({ message: `${title}: ${desc}`, time: timeStr })
      unreadCount.value++
      triggerToast(title, desc)
    })

    socket.on("popup", (data: any) => {
      activePopup.value = {
        title: data.title || "Pengumuman Visual Baru",
        photo: data.photo ? getPhotoUrl(data.photo) : ""
      }
      showPopupAnnouncement.value = true
    })
  }
})

onUnmounted(() => {
  if (socket) {
    socket.disconnect()
  }
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
  user.value = null
  navigateTo("/login")
}

async function clearAll() {
  notifications.value = []
  unreadCount.value = 0
  try {
    await api.delete("/api/notifications")
  } catch (error) {
    console.error("Gagal membersihkan notifikasi di server", error)
  }
}
</script>
