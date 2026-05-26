<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8">
        <div class="flex items-center space-x-3">
          <NuxtLink to="/teachers" class="p-2 rounded-md border border-[color:var(--color-border)] bg-[color:var(--color-surface)] text-[color:var(--color-muted)] hover:text-[color:var(--color-heading)] hover:bg-[color:var(--color-bg)] transition duration-150 cursor-pointer">
            <ChevronLeftIcon class="w-4 h-4" />
          </NuxtLink>
          <div>
            <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-1">
              Direktori Akademik
            </h2>
            <h1 class="text-3xl font-bold text-[color:var(--color-heading)] tracking-tight">
              Tambah Guru Baru
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
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="space-y-1.5">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">NIP</label>
                <input v-model="formData.nip" type="text" class="input w-full font-mono text-[color:var(--color-accent)]" required placeholder="Contoh: 198501012010011001" />
              </div>
              <div class="space-y-1.5">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Nama Lengkap</label>
                <input v-model="formData.name" type="text" class="input w-full" required placeholder="Contoh: Budi Santoso" />
              </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="space-y-1.5">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Tempat Lahir</label>
                <input v-model="formData.place_of_birth" type="text" class="input w-full" required placeholder="Contoh: Jakarta" />
              </div>
              <div class="space-y-1.5">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Tanggal Lahir</label>
                <input v-model="formData.date_of_birth" type="date" class="input w-full" required />
              </div>
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Gender</label>
              <select v-model="formData.gender" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                <option value="" disabled>Pilih Gender</option>
                <option value="Laki-laki">Laki-laki</option>
                <option value="Perempuan">Perempuan</option>
              </select>
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Alamat</label>
              <textarea v-model="formData.address" class="input w-full h-20 resize-none" placeholder="Alamat lengkap tempat tinggal guru"></textarea>
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Alamat Email</label>
              <div class="relative">
                <MailIcon class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-[color:var(--color-muted)]" />
                <input v-model="formData.email" type="email" class="input w-full font-mono text-[color:var(--color-accent)]" style="padding-left: 2.5rem !important;" required placeholder="Contoh: budi@sekolah.com" />
              </div>
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Nomor WhatsApp</label>
              <div class="relative">
                <PhoneIcon class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-[color:var(--color-muted)]" />
                <input v-model="formData.whatsapp" type="text" class="input w-full font-mono" style="padding-left: 2.5rem !important;" placeholder="Contoh: 081234567890" />
              </div>
            </div>

            <div class="bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-4 text-xs text-[color:var(--color-text)] space-y-1.5 leading-relaxed">
              <div class="font-bold text-[color:var(--color-heading)] uppercase tracking-wider text-[10px]">Informasi Kata Sandi Default</div>
              <p>Kata sandi akan dibuat otomatis dengan format <span class="font-bold text-[color:var(--color-accent)] font-mono">GURUXXXX</span> dimana XX pertama adalah 2 digit terakhir tahun lahir dan XX terakhir adalah 2 digit terakhir NIP.</p>
              <p v-if="generatedPassword" class="font-mono text-sm font-bold text-[color:var(--color-heading)]">Password: {{ generatedPassword }}</p>
            </div>

            <div class="flex items-center justify-end space-x-3 pt-6 border-t border-[color:var(--color-border)]">
              <NuxtLink to="/teachers" class="button-ghost text-xs font-semibold uppercase tracking-wider">
                Batal
              </NuxtLink>
              <button type="submit" :disabled="isLoading" class="button-primary text-xs font-semibold uppercase tracking-wider flex items-center space-x-1.5 py-3">
                <span v-if="isLoading" class="animate-spin rounded-full h-3.5 w-3.5 border-2 border-[color:var(--color-accent-fg)] border-t-transparent mr-1"></span>
                <UserPlusIcon class="w-4 h-4 shrink-0" v-else />
                <span>Simpan Guru</span>
              </button>
            </div>
          </form>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from "vue"
import { useRouter } from "vue-router"
import {
  ChevronLeft as ChevronLeftIcon,
  AlertCircle as AlertCircleIcon,
  CheckCircle as CheckCircleIcon,
  UserPlus as UserPlusIcon,
  Mail as MailIcon,
  Phone as PhoneIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()
const router = useRouter()
const currentUser = useState<any>("currentUser", () => null)

const isLoading = ref(false)
const errorMessage = ref("")
const toastMessage = ref("")

const formData = ref({
  name: "",
  email: "",
  nip: "",
  place_of_birth: "",
  date_of_birth: "",
  gender: "",
  address: "",
  whatsapp: ""
})

const generatedPassword = computed(() => {
  const dob = formData.value.date_of_birth
  const nip = formData.value.nip
  if (!dob || !nip || nip.length < 2) return ""
  const yearPart = dob.split("-")[0]
  if (!yearPart || yearPart.length < 4) return ""
  const yy = yearPart.slice(-2)
  const nn = nip.slice(-2)
  return `GURU${yy}${nn}`
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
})

async function handleSubmit() {
  if (!formData.value.nip || !formData.value.name || !formData.value.email) {
    errorMessage.value = "Kolom NIP, Nama, dan Email wajib diisi"
    return
  }

  if (!formData.value.place_of_birth || !formData.value.date_of_birth) {
    errorMessage.value = "Tempat dan Tanggal Lahir wajib diisi"
    return
  }

  if (!formData.value.gender) {
    errorMessage.value = "Gender wajib dipilih"
    return
  }

  isLoading.value = true
  errorMessage.value = ""

  try {
    const payload = {
      name: formData.value.name,
      email: formData.value.email,
      password: "",
      role: "guru",
      nip: formData.value.nip,
      place_of_birth: formData.value.place_of_birth,
      date_of_birth: formData.value.date_of_birth,
      gender: formData.value.gender,
      address: formData.value.address,
      whatsapp: formData.value.whatsapp,
      student_id: null
    }

    const res: any = await api.post("/api/users", payload)
    if (res.success) {
      toastMessage.value = "Guru baru berhasil terdaftar di sistem"
      setTimeout(() => {
        router.push("/teachers")
      }, 1500)
    } else {
      errorMessage.value = res.message || "Gagal membuat akun guru baru"
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
