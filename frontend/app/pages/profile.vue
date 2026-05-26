<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-4xl w-full mx-auto space-y-6 sm:space-y-8">
        <div>
          <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
            Informasi Pribadi
          </h2>
          <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
            Profil Pengguna
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

        <div class="space-y-8">
          <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6 shadow-[--shadow-sm] w-full">
            <div class="flex items-center space-x-3 border-b border-[color:var(--color-border)] pb-4">
              <UserIcon class="w-5 h-5 text-[color:var(--color-accent)] shrink-0" />
              <div>
                <h3 class="text-md font-bold text-[color:var(--color-heading)]">Foto & Detail Akun</h3>
                <p class="text-xs text-[color:var(--color-muted)]">Kelola informasi foto profil dan data akun Anda</p>
              </div>
            </div>

            <div class="flex flex-col sm:flex-row items-center sm:items-start gap-6">
              <div class="relative group cursor-pointer" @click="triggerFileInput">
                <div class="w-24 h-24 rounded-full overflow-hidden border border-[color:var(--color-border)] bg-[color:var(--color-bg)] flex items-center justify-center relative">
                  <img v-if="avatarUrl" :src="avatarUrl" class="w-full h-full object-cover" />
                  <UserIcon v-else class="w-10 h-10 text-[color:var(--color-muted)]" />

                  <div class="absolute inset-0 bg-black/60 opacity-0 group-hover:opacity-100 flex items-center justify-center transition duration-150">
                    <UploadIcon class="w-5 h-5 text-white" />
                  </div>
                </div>
                <input ref="fileInput" type="file" class="hidden" accept="image/*" @change="handleAvatarUpload" />
              </div>

              <div class="space-y-2 text-center sm:text-left flex-1">
                <h4 class="text-lg font-bold text-[color:var(--color-heading)]">{{ currentUser?.name || '-' }}</h4>
                <p class="text-sm font-mono text-[color:var(--color-text)]">{{ currentUser?.email || '-' }}</p>
                <div class="pt-1">
                  <span :class="[
                    'text-xs font-bold uppercase tracking-wider px-2 py-0.5 border',
                    currentUser?.role === 'admin' ? 'border-red-500/30 bg-red-500/10 text-red-500' :
                    currentUser?.role === 'guru' ? 'border-green-500/30 bg-green-500/10 text-green-500' :
                    currentUser?.role === 'siswa' ? 'border-blue-500/30 bg-blue-500/10 text-blue-500' :
                    'border-pink-500/30 bg-pink-500/10 text-pink-500'
                  ]">
                    {{ currentUser?.role }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6 shadow-[--shadow-sm] w-full">
            <div class="flex items-center space-x-3 border-b border-[color:var(--color-border)] pb-4">
              <LockIcon class="w-5 h-5 text-[color:var(--color-accent)] shrink-0" />
              <div>
                <h3 class="text-md font-bold text-[color:var(--color-heading)]">Keamanan Akun</h3>
                <p class="text-xs text-[color:var(--color-muted)]">Perbarui kata sandi keamanan akun Anda</p>
              </div>
            </div>

            <form @submit.prevent="updatePassword" class="space-y-4">
              <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div class="flex flex-col space-y-1.5">
                  <label class="text-xs uppercase tracking-wider text-[color:var(--color-text)] font-semibold">Password Lama</label>
                  <input
                    v-model="passwordForm.old_password"
                    :type="showPassword ? 'text' : 'password'"
                    class="input w-full font-mono"
                    placeholder="••••••••"
                    required
                  />
                </div>

                <div class="flex flex-col space-y-1.5">
                  <label class="text-xs uppercase tracking-wider text-[color:var(--color-text)] font-semibold">Password Baru</label>
                  <input
                    v-model="passwordForm.new_password"
                    :type="showPassword ? 'text' : 'password'"
                    class="input w-full font-mono"
                    placeholder="••••••••"
                    required
                  />
                </div>

                <div class="flex flex-col space-y-1.5">
                  <label class="text-xs uppercase tracking-wider text-[color:var(--color-text)] font-semibold">Konfirmasi Password Baru</label>
                  <input
                    v-model="passwordForm.confirm_password"
                    :type="showPassword ? 'text' : 'password'"
                    class="input w-full font-mono"
                    placeholder="••••••••"
                    required
                  />
                </div>
              </div>

              <div class="flex items-center space-x-2 pt-1">
                <input
                  v-model="showPassword"
                  type="checkbox"
                  id="show-password-profile"
                  class="h-4 w-4 rounded border-[color:var(--color-border)] bg-[color:var(--color-bg)] text-[color:var(--color-accent)] focus:ring-[color:var(--color-accent)] focus:ring-offset-[color:var(--color-surface)] cursor-pointer"
                />
                <label for="show-password-profile" class="text-xs text-[color:var(--color-text)] font-semibold select-none cursor-pointer">
                  Tampilkan Kata Sandi
                </label>
              </div>

              <div class="pt-2">
                <button type="submit" class="button-primary w-full md:w-auto" :disabled="loading.password">
                  {{ loading.password ? 'Mengubah...' : 'Perbarui Password' }}
                </button>
              </div>
            </form>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from "vue"
import {
  Lock as LockIcon,
  CheckCircle as CheckCircleIcon,
  AlertCircle as AlertCircleIcon,
  User as UserIcon,
  Upload as UploadIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"],
  alias: ["/profiles"]
})

const api = useApi()
const config = useRuntimeConfig()

const showPassword = ref(false)
const toast = ref<{ message: string; type: "success" | "error" }>({
  message: "",
  type: "success"
})

const loading = ref({
  password: false
})

const passwordForm = ref({
  old_password: "",
  new_password: "",
  confirm_password: ""
})

const currentUser = useState<any>("currentUser", () => null)
const fileInput = ref<HTMLInputElement | null>(null)

const avatarUrl = computed(() => {
  if (!currentUser.value?.avatar) return ""
  if (currentUser.value.avatar.startsWith("http")) return currentUser.value.avatar
  return `${config.public.apiUrl}${currentUser.value.avatar}`
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

async function handleAvatarUpload(event: Event) {
  const target = event.target as HTMLInputElement
  if (!target.files || target.files.length === 0) return

  const file = target.files[0]
  if (!file) return

  const formData = new FormData()
  formData.append("avatar", file)

  try {
    const res: any = await api.post("/api/profile/avatar", formData)
    if (res.success && res.data?.avatar) {
      currentUser.value.avatar = res.data.avatar
      localStorage.setItem("user", JSON.stringify(currentUser.value))
      showToast("Foto profil berhasil diperbarui", "success")
    } else {
      showToast(res.message || "Gagal memperbarui foto profil", "error")
    }
  } catch (error: any) {
    showToast(error.data?.message || "Koneksi ke server gagal", "error")
  }
}

onMounted(() => {
  const userJson = localStorage.getItem("user")
  if (userJson && !currentUser.value) {
    currentUser.value = JSON.parse(userJson)
  }
})

async function updatePassword() {
  if (passwordForm.value.new_password !== passwordForm.value.confirm_password) {
    showToast("Konfirmasi password baru tidak cocok", "error")
    return
  }

  loading.value.password = true
  try {
    const res: any = await api.put("/api/settings/password", {
      old_password: passwordForm.value.old_password,
      new_password: passwordForm.value.new_password
    })
    if (res.success) {
      showToast("Kata sandi berhasil diperbarui", "success")
      passwordForm.value = {
        old_password: "",
        new_password: "",
        confirm_password: ""
      }
    } else {
      showToast(res.message || "Gagal memperbarui password", "error")
    }
  } catch (error: any) {
    showToast(error.data?.message || "Password lama salah", "error")
  } finally {
    loading.value.password = false
  }
}
</script>
