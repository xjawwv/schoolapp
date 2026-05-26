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
              Buat Siaran Popup Visual Baru
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
          <form @submit.prevent="submitPopup" class="space-y-6">
            <div class="space-y-2">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">
                Judul Pengumuman
              </label>
              <input
                type="text"
                v-model="title"
                class="input w-full"
                placeholder="Masukkan judul pengumuman popup..."
                required
              />
            </div>

            <div class="space-y-2">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">
                Foto / Gambar Siaran
              </label>

              <div
                @dragover.prevent="dragOver = true"
                @dragleave.prevent="dragOver = false"
                @drop.prevent="handleDrop"
                :class="[
                  'border-2 border-dashed rounded-lg p-8 text-center transition duration-150 flex flex-col items-center justify-center space-y-3 cursor-pointer',
                  dragOver
                    ? 'border-[color:var(--color-accent)] bg-[color:var(--color-accent)]/5'
                    : 'border-[color:var(--color-border)] hover:border-[color:var(--color-accent)]/50 bg-[color:var(--color-bg)]'
                ]"
                @click="triggerFileInput"
              >
                <input
                  type="file"
                  ref="fileInput"
                  @change="handleFileChange"
                  accept="image/*"
                  class="hidden"
                />

                <div v-if="!previewUrl" class="flex flex-col items-center space-y-2">
                  <UploadIcon class="w-8 h-8 text-[color:var(--color-muted)]" />
                  <p class="text-sm font-medium text-[color:var(--color-text)]">
                    Seret & letakkan file gambar di sini, atau klik untuk memilih
                  </p>
                  <p class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider">
                    Format gambar (JPG, PNG, WEBP, GIF) • Maks 3MB
                  </p>
                </div>

                <div v-else class="relative w-full max-h-64 rounded-md overflow-hidden border border-[color:var(--color-border)] flex items-center justify-center bg-black/40">
                  <img :src="previewUrl" class="max-w-full max-h-60 object-contain" />
                  <button
                    type="button"
                    @click.stop="clearImage"
                    class="absolute top-2 right-2 p-1.5 rounded-full bg-black/60 text-white hover:bg-black/80 transition cursor-pointer"
                  >
                    <XIcon class="w-4 h-4" />
                  </button>
                </div>
              </div>
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
                :disabled="submitting || !selectedFile"
              >
                <span v-if="submitting" class="animate-spin rounded-full h-3.5 w-3.5 border-2 border-[color:var(--color-accent-fg)] border-t-transparent mr-2"></span>
                <SendIcon class="w-4 h-4 mr-2" />
                <span>Siarkan Popup</span>
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
  Send as SendIcon,
  Upload as UploadIcon,
  X as XIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()
const config = useRuntimeConfig()
const currentUser = useState<any>("currentUser", () => null)
const submitting = ref(false)

const title = ref("")
const selectedFile = ref<File | null>(null)
const previewUrl = ref("")
const dragOver = ref(false)
const fileInput = ref<HTMLInputElement | null>(null)

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

function triggerFileInput() {
  fileInput.value?.click()
}

function handleFileChange(e: Event) {
  const target = e.target as HTMLInputElement
  if (target.files && target.files[0]) {
    validateAndSetFile(target.files[0])
  }
}

function handleDrop(e: DragEvent) {
  dragOver.value = false
  if (e.dataTransfer?.files && e.dataTransfer.files[0]) {
    validateAndSetFile(e.dataTransfer.files[0])
  }
}

function validateAndSetFile(file: File) {
  if (!file.type.startsWith("image/")) {
    showToast("Format file harus berupa gambar", "error")
    return
  }

  if (file.size > 3 * 1024 * 1024) {
    showToast("Ukuran foto maksimal adalah 3MB", "error")
    return
  }

  selectedFile.value = file
  previewUrl.value = URL.createObjectURL(file)
}

function clearImage() {
  selectedFile.value = null
  previewUrl.value = ""
  if (fileInput.value) fileInput.value.value = ""
}

async function submitPopup() {
  if (!title.value.trim() || !selectedFile.value) return
  submitting.value = true

  const formData = new FormData()
  formData.append("title", title.value)
  formData.append("photo", selectedFile.value)

  try {
    const token = localStorage.getItem("token")
    const headers: any = {}
    if (token) {
      headers["Authorization"] = `Bearer ${token}`
    }

    const res: any = await $fetch(`${config.public.apiUrl}/api/popups`, {
      method: "POST",
      body: formData,
      headers: headers
    })

    if (res.success) {
      showToast("Popup visual darurat berhasil disiarkan", "success")
      setTimeout(() => {
        navigateTo("/announcement")
      }, 1500)
    } else {
      showToast(res.message || "Gagal menyiarkan popup", "error")
    }
  } catch (error: any) {
    showToast(error.data?.message || "Terjadi kesalahan server saat mengunggah gambar", "error")
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
