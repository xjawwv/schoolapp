<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8">
        <div class="flex items-center space-x-3">
          <NuxtLink to="/user" class="p-2 rounded-md border border-[color:var(--color-border)] bg-[color:var(--color-surface)] text-[color:var(--color-muted)] hover:text-[color:var(--color-heading)] hover:bg-[color:var(--color-bg)] transition duration-150 cursor-pointer">
            <ChevronLeftIcon class="w-4 h-4" />
          </NuxtLink>
          <div>
            <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-1">
              Manajemen Pengguna
            </h2>
            <h1 class="text-3xl font-bold text-[color:var(--color-heading)] tracking-tight">
              Tambah Pengguna Baru
            </h1>
          </div>
        </div>

        <div v-if="toastMessage" class="bg-[color:var(--color-surface)] border border-[color:var(--color-success)] p-3.5 text-sm text-[color:var(--color-success)] font-medium flex items-center space-x-2">
          <CheckCircleIcon class="w-4 h-4 shrink-0" />
          <span>{{ toastMessage }}</span>
        </div>

        <div class="card shadow-[--shadow-md] bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-8">
          <div v-if="errorMessage" class="mb-6 bg-[color:var(--color-bg)] border border-[color:var(--color-error)] p-4 text-sm text-[color:var(--color-error)] font-medium flex items-center space-x-2">
            <AlertCircleIcon class="w-4 h-4 shrink-0" />
            <span>{{ errorMessage }}</span>
          </div>

          <form @submit.prevent="handleSubmit" class="space-y-5">
            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Nama Lengkap</label>
              <div class="relative">
                <UserIcon class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-[color:var(--color-muted)]" />
                <input v-model="formData.name" type="text" class="input w-full" style="padding-left: 2.5rem !important;" required placeholder="Contoh: Ahmad Fauzi" />
              </div>
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Alamat Email</label>
              <div class="relative">
                <MailIcon class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-[color:var(--color-muted)]" />
                <input v-model="formData.email" type="email" class="input w-full font-mono text-[color:var(--color-accent)]" style="padding-left: 2.5rem !important;" required placeholder="Contoh: ahmad@sekolah.com" />
              </div>
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Kata Sandi Akun</label>
              <div class="relative">
                <LockIcon class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-[color:var(--color-muted)]" />
                <input v-model="formData.password" type="password" class="input w-full font-mono" style="padding-left: 2.5rem !important;" required placeholder="Minimal 6 karakter" />
              </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="space-y-1.5">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Peran Hak Akses (Role)</label>
                <select v-model="formData.role" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                  <option value="" disabled>Pilih Peran</option>
                  <option value="admin">Admin</option>
                  <option value="guru">Guru</option>
                  <option value="siswa">Siswa (Laki-laki)</option>
                  <option value="siswi">Siswi (Perempuan)</option>
                </select>
              </div>

              <div class="space-y-1.5" v-if="formData.role === 'guru'">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">NIP Guru</label>
                <input v-model="formData.nip" type="text" class="input w-full font-mono text-[color:var(--color-accent)]" required placeholder="Contoh: 12345" />
              </div>

              <div class="space-y-1.5" v-if="formData.role === 'siswa' || formData.role === 'siswi'">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Hubungkan Biodata Siswa</label>
                <select v-model="formData.student_id" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                  <option value="" disabled>Pilih Biodata Siswa</option>
                  <option v-for="std in studentsList" :key="std.id" :value="std.id">
                    {{ std.nisn }} - {{ std.name }} ({{ std.class }})
                  </option>
                </select>
              </div>
            </div>

            <div class="bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-4 text-xs text-[color:var(--color-text)] space-y-1.5 leading-relaxed">
              <div class="font-bold text-[color:var(--color-heading)] uppercase tracking-wider text-[10px]">Ketentuan Kredensial Pengguna</div>
              <p v-if="formData.role === 'guru'">Pengguna berrole <strong>Guru</strong> wajib dibekali NIP yang unik sebagai pengenal sekunder di portal sekolah.</p>
              <p v-else-if="formData.role === 'siswa' || formData.role === 'siswi'">Pengguna berrole <strong>Siswa/Siswi</strong> wajib terhubung dengan biodata siswa di direktori sekolah untuk singkronisasi data kehadiran dan nilai.</p>
              <p v-else>Pengguna berrole <strong>Admin</strong> memiliki kewenangan konfigurasi sistem penuh.</p>
            </div>

            <div class="flex items-center justify-end space-x-3 pt-6 border-t border-[color:var(--color-border)]">
              <NuxtLink to="/user" class="button-ghost text-xs font-semibold uppercase tracking-wider">
                Batal
              </NuxtLink>
              <button type="submit" :disabled="isLoading" class="button-primary text-xs font-semibold uppercase tracking-wider flex items-center space-x-1.5 py-3">
                <span v-if="isLoading" class="animate-spin rounded-full h-3.5 w-3.5 border-2 border-[color:var(--color-accent-fg)] border-t-transparent mr-1"></span>
                <UserPlusIcon class="w-4 h-4 shrink-0" v-else />
                <span>Simpan Pengguna</span>
              </button>
            </div>
          </form>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from "vue"
import { useRouter } from "vue-router"
import {
  ChevronLeft as ChevronLeftIcon,
  AlertCircle as AlertCircleIcon,
  CheckCircle as CheckCircleIcon,
  Lock as LockIcon,
  Mail as MailIcon,
  User as UserIcon,
  UserPlus as UserPlusIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"
import Sidebar from "~/components/Sidebar.vue"
import Navbar from "~/components/Navbar.vue"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()
const router = useRouter()
const currentUser = useState<any>("currentUser", () => null)

const studentsList = ref<any[]>([])
const isLoading = ref(false)
const errorMessage = ref("")
const toastMessage = ref("")

const formData = ref({
  name: "",
  email: "",
  password: "",
  role: "",
  student_id: "",
  nip: ""
})

watch(() => formData.value.role, (newRole) => {
  if (newRole !== 'siswa' && newRole !== 'siswi') {
    formData.value.student_id = ""
  }
  if (newRole !== 'guru') {
    formData.value.nip = ""
  }
})

onMounted(() => {
  const cached = localStorage.getItem("user")
  if (cached && !currentUser.value) {
    currentUser.value = JSON.parse(cached)
  }

  if (currentUser.value?.role !== "admin") {
    router.push("/dashboard")
    return
  }

  loadStudents()
})

async function loadStudents() {
  try {
    const res: any = await api.get("/api/students?limit=200")
    if (res.success && res.data) {
      studentsList.value = res.data.students || []
    }
  } catch (error) {
    console.error("Gagal memuat daftar siswa", error)
  }
}

async function handleSubmit() {
  if (!formData.value.name || !formData.value.email || !formData.value.password || !formData.value.role) {
    errorMessage.value = "Semua kolom utama wajib diisi"
    return
  }

  if (formData.value.password.length < 6) {
    errorMessage.value = "Kata sandi wajib minimal 6 karakter"
    return
  }

  if (formData.value.role === 'guru' && !formData.value.nip) {
    errorMessage.value = "Kolom NIP wajib diisi untuk peran Guru"
    return
  }

  if ((formData.value.role === 'siswa' || formData.value.role === 'siswi') && !formData.value.student_id) {
    errorMessage.value = "Biodata Siswa (NISN) wajib dihubungkan untuk peran Siswa/Siswi"
    return
  }

  isLoading.value = true
  errorMessage.value = ""

  try {
    const payload = {
      name: formData.value.name,
      email: formData.value.email,
      password: formData.value.password,
      role: formData.value.role,
      student_id: (formData.value.role === 'siswa' || formData.value.role === 'siswi') ? formData.value.student_id : null,
      nip: formData.value.role === 'guru' ? formData.value.nip : ""
    }

    const res: any = await api.post("/api/users", payload)
    if (res.success) {
      toastMessage.value = "Pengguna baru berhasil terdaftar di sistem"
      setTimeout(() => {
        router.push("/user")
      }, 1500)
    } else {
      errorMessage.value = res.message || "Gagal membuat pengguna baru"
    }
  } catch (error: any) {
    if (error.response && error.response._data) {
      errorMessage.value = error.response._data.message || "Gagal memproses data"
    } else {
      errorMessage.value = "Koneksi ke server terputus"
    }
  } finally {
    isLoading.value = false
  }
}
</script>
