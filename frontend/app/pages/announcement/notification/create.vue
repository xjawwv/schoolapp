<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-2xl w-full mx-auto space-y-6 sm:space-y-8">
        <div class="flex items-center space-x-4">
          <NuxtLink
            to="/announcement"
            class="p-2 rounded-md border border-[color:var(--color-border)] bg-[color:var(--color-surface)] text-[color:var(--color-muted)] hover:text-[color:var(--color-heading)] hover:bg-[color:var(--color-bg)] transition duration-150 cursor-pointer"
          >
            <ArrowLeftIcon class="w-5 h-5" />
          </NuxtLink>
          <div>
            <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-1">
              Komunikasi Sistem
            </h2>
            <h1 class="text-2xl font-bold text-[color:var(--color-heading)] tracking-tight">
              Buat Notifikasi Baru
            </h1>
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

        <div class="card bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 sm:p-8 shadow-[--shadow-md]">
          <form @submit.prevent="submitNotification" class="space-y-6">
            <div class="space-y-2">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">
                Judul Notifikasi
              </label>
              <input
                type="text"
                v-model="form.title"
                class="input w-full"
                placeholder="Masukkan judul notifikasi pengumuman..."
                required
              />
            </div>

            <div class="space-y-2">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">
                Deskripsi Notifikasi
              </label>
              <textarea
                v-model="form.description"
                class="input w-full h-40 resize-none"
                placeholder="Tuliskan deskripsi notifikasi sistem di sini..."
                required
              ></textarea>
            </div>

            <div class="pt-4 border-t border-[color:var(--color-border)] flex items-center justify-end space-x-3">
              <NuxtLink
                to="/announcement"
                class="button-ghost text-xs font-semibold py-2.5 px-4"
              >
                Batal
              </NuxtLink>
              <button
                type="submit"
                class="button-primary py-2.5 px-6 flex items-center justify-center text-xs"
                :disabled="submitting"
              >
                <span v-if="submitting" class="animate-spin rounded-full h-3.5 w-3.5 border-2 border-[color:var(--color-accent-fg)] border-t-transparent mr-2"></span>
                <SendIcon class="w-4 h-4 mr-2" />
                <span>Siarkan Notifikasi</span>
              </button>
            </div>
          </form>
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
  CheckCircle as CheckCircleIcon,
  Send as SendIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()
const currentUser = useState<any>("currentUser", () => null)
const submitting = ref(false)

const form = ref({
  title: "",
  description: ""
})

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

async function submitNotification() {
  if (!form.value.title.trim() || !form.value.description.trim()) return
  submitting.value = true

  try {
    const res: any = await api.post("/api/notifications", {
      title: form.value.title,
      description: form.value.description
    })
    if (res.success) {
      showToast("Notifikasi berhasil disiarkan ke seluruh sistem", "success")
      setTimeout(() => {
        navigateTo("/announcement")
      }, 1500)
    } else {
      showToast(res.message || "Gagal menyiarkan notifikasi", "error")
    }
  } catch (error: any) {
    showToast(error.data?.message || "Terjadi kesalahan server saat menyimpan data", "error")
  } finally {
    submitting.value = false
  }
}

onMounted(() => {
  const cached = localStorage.getItem("user")
  if (cached) {
    currentUser.value = JSON.parse(cached)
  }

  if (currentUser.value?.role !== "admin") {
    navigateTo("/announcement")
  }
})
</script>
