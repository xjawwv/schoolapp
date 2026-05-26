<template>
  <div class="min-h-screen bg-[color:var(--color-bg)] flex flex-col lg:flex-row">
    <div class="hidden lg:flex lg:w-1/2 flex-col justify-between p-16 bg-black border-r border-[color:var(--color-border)] relative overflow-hidden shrink-0">
      <div class="absolute -right-24 -bottom-24 w-96 h-96 rounded-full bg-[color:var(--color-accent)]/10 blur-[120px] pointer-events-none"></div>

      <div>
        <span class="text-xs uppercase tracking-[0.25em] font-semibold text-[color:var(--color-accent)] mb-4 block">
          Portal Digital Kesiswaan
        </span>
        <h2 class="text-7xl font-display font-black tracking-tight text-white leading-[1.05]">
          Sistem<br />
          Manajemen<br />
          Kesiswaan.
        </h2>
      </div>

      <div class="space-y-2 relative z-10">
        <p class="text-lg font-semibold text-white tracking-wide truncate">
          {{ siteName }}
        </p>
        <p class="text-xs text-[color:var(--color-muted)] max-w-sm uppercase tracking-wider leading-relaxed">
          Platform administrasi terpadu untuk pencatatan absensi, pengolahan nilai akademik, dan direktori data siswa.
        </p>
      </div>
    </div>

    <div class="w-full lg:w-1/2 flex items-center justify-center p-8 sm:p-16">
      <div class="w-full max-w-md space-y-8 bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-8 sm:p-10 shadow-[--shadow-lg] relative overflow-hidden">
        <div class="lg:hidden text-center mb-6">
          <span class="text-[10px] uppercase tracking-[0.2em] font-bold text-[color:var(--color-accent)] mb-2 block">
            Portal Digital Kesiswaan
          </span>
          <h1 class="text-3xl font-display font-black tracking-tight text-[color:var(--color-heading)]">
            {{ siteName }}
          </h1>
        </div>

        <div class="space-y-2">
          <h2 class="text-2xl font-bold tracking-tight text-[color:var(--color-heading)]">
            Selamat Datang
          </h2>
          <p class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider">
            Masukkan kredensial akun untuk mengakses sistem
          </p>
        </div>

        <div class="flex p-1 bg-[color:var(--color-bg)] border border-[color:var(--color-border)] rounded-lg">
          <button
            type="button"
            @click="selectPreset('guru')"
            class="flex-1 py-2 text-xs uppercase tracking-wider font-bold transition duration-150 cursor-pointer text-center rounded-md"
            :class="selectedRole === 'guru' ? 'bg-[color:var(--color-accent)] text-white shadow-[--shadow-sm] font-extrabold' : 'text-[color:var(--color-muted)] hover:text-[color:var(--color-text)]'"
          >
            Guru
          </button>
          <button
            type="button"
            @click="selectPreset('siswa')"
            class="flex-1 py-2 text-xs uppercase tracking-wider font-bold transition duration-150 cursor-pointer text-center rounded-md"
            :class="selectedRole === 'siswa' ? 'bg-[color:var(--color-accent)] text-white shadow-[--shadow-sm] font-extrabold' : 'text-[color:var(--color-muted)] hover:text-[color:var(--color-text)]'"
          >
            Siswa
          </button>
        </div>

        <form @submit.prevent="handleLogin" class="space-y-6">
          <div v-if="errorMessage" class="bg-[color:var(--color-bg)] border border-[color:var(--color-error)] p-3 text-sm text-[color:var(--color-error)] font-medium flex items-center space-x-2">
            <AlertCircleIcon class="w-4 h-4 shrink-0" />
            <span>{{ errorMessage }}</span>
          </div>

          <div class="space-y-1.5">
            <label for="email" class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">
              {{ selectedRole === 'siswa' ? 'NISN' : 'NIP' }}
            </label>
            <input
              id="email"
              v-model="email"
              type="text"
              class="input w-full"
              :placeholder="selectedRole === 'siswa' ? 'Masukkan NISN' : 'Masukkan NIP'"
              required
            />
          </div>

          <div class="space-y-1.5">
            <label for="password" class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">
              Kata Sandi
            </label>
            <div class="relative flex items-center">
              <input
                id="password"
                v-model="password"
                :type="showPassword ? 'text' : 'password'"
                class="input w-full pr-10 font-mono"
                placeholder="••••••••"
                required
                autocomplete="current-password"
              />
              <button
                type="button"
                @click="showPassword = !showPassword"
                class="absolute right-3 text-[color:var(--color-muted)] hover:text-[color:var(--color-text)] transition duration-150 focus:outline-none cursor-pointer"
              >
                <EyeIcon v-if="showPassword" class="w-4 h-4 shrink-0" />
                <EyeOffIcon v-else class="w-4 h-4 shrink-0" />
              </button>
            </div>
          </div>

          <div class="pt-2">
            <button
              type="submit"
              :disabled="isLoading"
              class="button-primary w-full flex items-center justify-center space-x-2 py-3"
            >
              <span v-if="isLoading" class="animate-spin rounded-full h-4 w-4 border-2 border-[color:var(--color-accent-fg)] border-t-transparent"></span>
              <span v-else>Masuk ke Sistem</span>
            </button>
          </div>
        </form>

        <div class="border-t border-[color:var(--color-border)] pt-6 text-center text-xs text-[color:var(--color-muted)] space-y-1 leading-relaxed">
          <div>Admin: <span class="font-mono text-[color:var(--color-text)]">admin@sekolah.com</span> / <span class="font-mono text-[color:var(--color-text)]">admin123</span></div>
          <div>Guru: <span class="font-mono text-[color:var(--color-text)]">12345</span> / <span class="font-mono text-[color:var(--color-text)]">gurusmk</span></div>
          <div>Siswa: <span class="font-mono text-[color:var(--color-text)]">10001</span> / <span class="font-mono text-[color:var(--color-text)]">1</span></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import {
  AlertCircle as AlertCircleIcon,
  Eye as EyeIcon,
  EyeOff as EyeOffIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: []
})

const email = ref("")
const password = ref("")
const isLoading = ref(false)
const errorMessage = ref("")
const showPassword = ref(false)
const siteName = useState("siteName", () => "SMA N 1 METRO")
const selectedRole = ref("siswa")

function selectPreset(role: string) {
  selectedRole.value = role
}

const api = useApi()

onMounted(async () => {
  const cachedSiteName = localStorage.getItem("siteName")
  if (cachedSiteName) {
    siteName.value = cachedSiteName
  }
  try {
    const res: any = await api.get("/api/settings")
    if (res.success && res.data && res.data.site_name) {
      siteName.value = res.data.site_name
      localStorage.setItem("siteName", res.data.site_name)
    }
  } catch (e) {
    console.error(e)
  }
})

async function handleLogin() {
  if (!email.value || !password.value) {
    errorMessage.value = "Semua kolom harus diisi"
    return
  }

  isLoading.value = true
  errorMessage.value = ""

  try {
    const res: any = await api.post("/api/auth/login", {
      email: email.value,
      password: password.value,
      role_group: selectedRole.value
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
