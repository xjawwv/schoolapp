<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8">
        <div>
          <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
            Pengaturan Sistem
          </h2>
          <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
            Setting Konfigurasi
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
              <GlobeIcon class="w-5 h-5 text-[color:var(--color-accent)] shrink-0" />
              <div>
                <h3 class="text-md font-bold text-[color:var(--color-heading)]">Identitas Aplikasi</h3>
                <p class="text-xs text-[color:var(--color-muted)]">Atur penamaan nama instansi atau sekolah</p>
              </div>
            </div>

            <form @submit.prevent="saveSiteSettings" class="space-y-4">
              <div class="flex flex-col space-y-1.5 max-w-md">
                <label class="text-xs uppercase tracking-wider text-[color:var(--color-text)] font-semibold">Nama Website</label>
                <input
                  v-model="siteSettings.site_name"
                  type="text"
                  class="input w-full font-semibold"
                  placeholder="SMA N 1 METRO"
                  required
                />
              </div>

              <div class="pt-2">
                <button type="submit" class="button-primary w-full md:w-auto" :disabled="loading.site">
                  {{ loading.site ? 'Menyimpan...' : 'Simpan Pengaturan' }}
                </button>
              </div>
            </form>
          </div>

          <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6 shadow-[--shadow-sm] w-full">
            <div class="flex items-center space-x-3 border-b border-[color:var(--color-border)] pb-4">
              <LockIcon class="w-5 h-5 text-[color:var(--color-accent)] shrink-0" />
              <div>
                <h3 class="text-md font-bold text-[color:var(--color-heading)]">Keamanan Akun</h3>
                <p class="text-xs text-[color:var(--color-muted)]">Perbarui kata sandi akun administrator Anda</p>
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
                  id="show-password-settings"
                  class="h-4 w-4 rounded border-[color:var(--color-border)] bg-[color:var(--color-bg)] text-[color:var(--color-accent)] focus:ring-[color:var(--color-accent)] focus:ring-offset-[color:var(--color-surface)] cursor-pointer"
                />
                <label for="show-password-settings" class="text-xs text-[color:var(--color-text)] font-semibold select-none cursor-pointer">
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
import { ref, onMounted } from "vue"
import {
  Globe as GlobeIcon,
  Lock as LockIcon,
  CheckCircle as CheckCircleIcon,
  AlertCircle as AlertCircleIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()
const siteName = useState("siteName", () => "SMA N 1 METRO")
const showPassword = ref(false)

const toast = ref<{ message: string; type: "success" | "error" }>({
  message: "",
  type: "success"
})

const loading = ref({
  site: false,
  password: false
})

const siteSettings = ref({
  site_name: ""
})

const passwordForm = ref({
  old_password: "",
  new_password: "",
  confirm_password: ""
})

function showToast(message: string, type: "success" | "error") {
  toast.value = { message, type }
  setTimeout(() => {
    toast.value.message = ""
  }, 4000)
}

onMounted(async () => {
  try {
    const res: any = await api.get("/api/settings")
    if (res.success && res.data) {
      siteSettings.value.site_name = res.data.site_name || ""
      siteName.value = res.data.site_name || "SMA N 1 METRO"
    }
  } catch (error) {
    showToast("Gagal memuat konfigurasi sistem", "error")
  }
})

async function saveSiteSettings() {
  loading.value.site = true
  try {
    const res: any = await api.put("/api/settings", {
      site_name: siteSettings.value.site_name
    })
    if (res.success) {
      siteName.value = siteSettings.value.site_name
      localStorage.setItem("siteName", siteSettings.value.site_name)
      showToast("Pengaturan identitas aplikasi berhasil disimpan", "success")
    } else {
      showToast(res.message || "Gagal menyimpan pengaturan", "error")
    }
  } catch (error: any) {
    showToast(error.data?.message || "Terjadi kesalahan server", "error")
  } finally {
    loading.value.site = false
  }
}

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
      showToast("Kata sandi administrator berhasil diperbarui", "success")
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
