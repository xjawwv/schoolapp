<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8" v-if="teacher">
        <div>
          <NuxtLink to="/teachers" class="inline-flex items-center space-x-2 text-xs uppercase tracking-wider text-[color:var(--color-accent)] hover:opacity-85 font-semibold mb-4">
            <ChevronLeftIcon class="w-4 h-4" />
            <span>Kembali ke Daftar</span>
          </NuxtLink>
          <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
            <div>
              <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
                Profil Pendidik
              </h2>
              <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
                {{ teacher.name }}
              </h1>
            </div>
            <button v-if="currentUser?.role === 'admin'" @click="openEditModal" class="button-primary text-xs font-semibold uppercase tracking-wider py-3">
              Ubah Profil
            </button>
          </div>
        </div>

        <div v-if="toastMessage" class="bg-[color:var(--color-surface)] border border-[color:var(--color-success)] p-3 text-sm text-[color:var(--color-success)] font-medium flex items-center space-x-2">
          <CheckCircleIcon class="w-4 h-4 shrink-0" />
          <span>{{ toastMessage }}</span>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
          <div class="lg:col-span-4 space-y-6">
            <div class="card space-y-6 shadow-[--shadow-sm]">
              <div class="flex items-center space-x-4 border-b border-[color:var(--color-border)] pb-4">
                <div class="w-12 h-12 rounded-full bg-[color:var(--color-bg)] border border-[color:var(--color-accent)] flex items-center justify-center text-[color:var(--color-accent)] text-lg font-bold">
                  {{ teacher.name.charAt(0) }}
                </div>
                <div class="flex flex-col">
                  <span class="text-sm font-semibold text-[color:var(--color-heading)]">NIP: {{ teacher.nip }}</span>
                  <span class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider">Tenaga Pendidik</span>
                </div>
              </div>

              <div class="space-y-4">
                <div>
                  <span class="block text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold mb-1">Gender</span>
                  <span class="text-sm text-[color:var(--color-text)] font-semibold">{{ teacher.gender }}</span>
                </div>
                <div>
                  <span class="block text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold mb-1">Nomor WhatsApp</span>
                  <span class="text-sm text-[color:var(--color-text)] font-mono font-semibold">{{ teacher.whatsapp || '-' }}</span>
                </div>
                <div>
                  <span class="block text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold mb-1">Alamat Lengkap</span>
                  <span class="text-sm text-[color:var(--color-text)] leading-relaxed font-semibold">{{ teacher.address || '-' }}</span>
                </div>
                <div>
                  <span class="block text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold mb-1">Terdaftar Sejak</span>
                  <span class="text-sm text-[color:var(--color-text)] font-mono font-semibold">{{ formatDate(teacher.created_at) }}</span>
                </div>
              </div>
            </div>

            <div class="card space-y-4 shadow-[--shadow-sm]" v-if="currentUser?.role === 'admin'">
              <h3 class="text-xs uppercase tracking-widest text-[color:var(--color-accent)] font-bold">
                Kredensial Akses
              </h3>
              <div class="space-y-3 bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-4 text-xs">
                <div>
                  <span class="text-[color:var(--color-muted)] block uppercase font-bold mb-1">Username / NIP</span>
                  <span class="font-mono text-sm text-[color:var(--color-heading)] font-semibold">{{ teacher.nip }}</span>
                </div>
                <div>
                  <span class="text-[color:var(--color-muted)] block uppercase font-bold mb-1">Kata Sandi Bawaan</span>
                  <span class="font-mono text-sm text-[color:var(--color-heading)] font-semibold">{{ expectedDefaultPassword }}</span>
                </div>
                <div class="pt-2 border-t border-[color:var(--color-border)] text-[color:var(--color-muted)] leading-relaxed">
                  Akun login pendidik dapat diakses dengan menggunakan NIP atau Email dengan sandi bawaan sistem.
                </div>
              </div>
            </div>
          </div>

          <div class="lg:col-span-8 space-y-6">
            <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6">
              <h3 class="text-xs uppercase tracking-widest font-bold pb-2 text-[color:var(--color-accent)] border-b border-[color:var(--color-border)] mb-6">
                Detail Informasi Pendidik
              </h3>

              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div class="space-y-4">
                  <div class="bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-4">
                    <span class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold mb-1">Nama Lengkap</span>
                    <span class="text-sm font-semibold text-[color:var(--color-heading)]">{{ teacher.name }}</span>
                  </div>
                  <div class="bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-4">
                    <span class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold mb-1">Nomor Induk Pegawai</span>
                    <span class="text-sm font-semibold font-mono text-[color:var(--color-accent)]">{{ teacher.nip }}</span>
                  </div>
                  <div class="bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-4">
                    <span class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold mb-1">Gender</span>
                    <span class="text-sm font-semibold text-[color:var(--color-text)]">{{ teacher.gender }}</span>
                  </div>
                </div>

                <div class="space-y-4">
                  <div class="bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-4">
                    <span class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold mb-1">Tempat & Tanggal Lahir</span>
                    <span class="text-sm font-semibold text-[color:var(--color-heading)]">{{ teacher.place_of_birth || '-' }}, {{ formatDate(teacher.date_of_birth) }}</span>
                  </div>
                  <div class="bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-4">
                    <span class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold mb-1">Surel Email</span>
                    <a :href="`mailto:${teacher.email}`" class="text-sm font-semibold text-[color:var(--color-accent)] hover:opacity-85 font-mono block">
                      {{ teacher.email }}
                    </a>
                  </div>
                  <div class="bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-4">
                    <span class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold mb-1">Nomor Kontak WhatsApp</span>
                    <span class="text-sm font-semibold text-[color:var(--color-text)] font-mono block">{{ teacher.whatsapp || '-' }}</span>
                  </div>
                </div>
              </div>

              <div class="mt-6 bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-4">
                <span class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold mb-1">Alamat Lengkap</span>
                <span class="text-sm text-[color:var(--color-text)] leading-relaxed block">{{ teacher.address || '-' }}</span>
              </div>
            </div>
          </div>
        </div>
      </main>

      <div v-if="showEditModal" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
        <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 w-full max-w-2xl space-y-6 shadow-xl relative max-h-[90vh] overflow-y-auto">
          <div class="flex justify-between items-center border-b border-[color:var(--color-border)] pb-4">
            <div>
              <h3 class="text-lg font-bold text-[color:var(--color-heading)] tracking-wide">Ubah Profil Guru</h3>
              <p class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider">Perbarui informasi profil tenaga pendidik</p>
            </div>
            <button @click="closeEditModal" class="text-[color:var(--color-muted)] hover:text-[color:var(--color-text)] cursor-pointer">
              <XIcon class="w-5 h-5" />
            </button>
          </div>

          <div v-if="editError" class="bg-[color:var(--color-bg)] border border-[color:var(--color-error)] p-3 text-sm text-[color:var(--color-error)] font-medium flex items-center space-x-2">
            <AlertCircleIcon class="w-4 h-4 shrink-0" />
            <span>{{ editError }}</span>
          </div>

          <form @submit.prevent="handleEditSubmit" class="space-y-4">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="space-y-1.5">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">NIP</label>
                <input v-model="editForm.nip" type="text" class="input w-full font-mono text-[color:var(--color-accent)]" required placeholder="Contoh: 198501012010011001" />
              </div>
              <div class="space-y-1.5">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Nama Lengkap</label>
                <input v-model="editForm.name" type="text" class="input w-full" required placeholder="Contoh: Budi Santoso" />
              </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="space-y-1.5">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Tempat Lahir</label>
                <input v-model="editForm.place_of_birth" type="text" class="input w-full" required placeholder="Contoh: Jakarta" />
              </div>
              <div class="space-y-1.5">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Tanggal Lahir</label>
                <input v-model="editForm.date_of_birth" type="date" class="input w-full" required />
              </div>
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Gender</label>
              <select v-model="editForm.gender" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                <option value="" disabled>Pilih Gender</option>
                <option value="Laki-laki">Laki-laki</option>
                <option value="Perempuan">Perempuan</option>
              </select>
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Alamat</label>
              <textarea v-model="editForm.address" class="input w-full h-20 resize-none" placeholder="Alamat lengkap tempat tinggal guru"></textarea>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="space-y-1.5">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Alamat Email</label>
                <div class="relative">
                  <MailIcon class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-[color:var(--color-muted)]" />
                  <input v-model="editForm.email" type="email" class="input w-full font-mono text-[color:var(--color-accent)]" style="padding-left: 2.5rem !important;" required placeholder="Contoh: budi@sekolah.com" />
                </div>
              </div>
              <div class="space-y-1.5">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Nomor WhatsApp</label>
                <div class="relative">
                  <PhoneIcon class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-[color:var(--color-muted)]" />
                  <input v-model="editForm.whatsapp" type="text" class="input w-full font-mono" style="padding-left: 2.5rem !important;" placeholder="Contoh: 081234567890" />
                </div>
              </div>
            </div>

            <div class="flex items-center justify-end space-x-3 pt-6 border-t border-[color:var(--color-border)]">
              <button type="button" @click="closeEditModal" class="button-ghost text-xs font-semibold uppercase tracking-wider py-3">
                Batal
              </button>
              <button type="submit" :disabled="editLoading" class="button-primary text-xs font-semibold uppercase tracking-wider flex items-center justify-center py-3">
                <span v-if="editLoading" class="animate-spin rounded-full h-3 w-3 border-2 border-[color:var(--color-accent-fg)] border-t-transparent mr-2"></span>
                <span>Simpan Perubahan</span>
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from "vue"
import { useRoute, useRouter } from "vue-router"
import {
  ChevronLeft as ChevronLeftIcon,
  CheckCircle as CheckCircleIcon,
  AlertCircle as AlertCircleIcon,
  X as XIcon,
  Mail as MailIcon,
  Phone as PhoneIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const route = useRoute()
const router = useRouter()
const api = useApi()

const currentUser = useState<any>("currentUser", () => null)
const teacher = ref<any | null>(null)
const toastMessage = ref("")

const showEditModal = ref(false)
const editLoading = ref(false)
const editError = ref("")

const editForm = ref({
  name: "",
  email: "",
  nip: "",
  place_of_birth: "",
  date_of_birth: "",
  gender: "",
  address: "",
  whatsapp: ""
})

const teacherId = computed(() => route.params.id as string)

const expectedDefaultPassword = computed(() => {
  if (!teacher.value) return ""
  const dob = teacher.value.date_of_birth
  const nip = teacher.value.nip
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
  fetchTeacherProfile()
})

async function fetchTeacherProfile() {
  try {
    const res: any = await api.get(`/api/teachers/${teacherId.value}`)
    if (res.success && res.data) {
      teacher.value = res.data
    } else {
      router.push("/teachers")
    }
  } catch (error) {
    console.error("Gagal memuat profil guru", error)
    router.push("/teachers")
  }
}

function openEditModal() {
  if (!teacher.value) return
  editForm.value = {
    name: teacher.value.name || "",
    email: teacher.value.email || "",
    nip: teacher.value.nip || "",
    place_of_birth: teacher.value.place_of_birth || "",
    date_of_birth: teacher.value.date_of_birth || "",
    gender: teacher.value.gender || "",
    address: teacher.value.address || "",
    whatsapp: teacher.value.whatsapp || ""
  }
  editError.value = ""
  showEditModal.value = true
}

function closeEditModal() {
  showEditModal.value = false
  editError.value = ""
}

async function handleEditSubmit() {
  if (!editForm.value.nip || !editForm.value.name || !editForm.value.email) {
    editError.value = "Kolom NIP, Nama, dan Email wajib diisi"
    return
  }

  if (!editForm.value.place_of_birth || !editForm.value.date_of_birth) {
    editError.value = "Tempat dan Tanggal Lahir wajib diisi"
    return
  }

  if (!editForm.value.gender) {
    editError.value = "Gender wajib dipilih"
    return
  }

  editLoading.value = true
  editError.value = ""

  try {
    const payload = {
      name: editForm.value.name,
      email: editForm.value.email,
      role: "guru",
      nip: editForm.value.nip,
      place_of_birth: editForm.value.place_of_birth,
      date_of_birth: editForm.value.date_of_birth,
      gender: editForm.value.gender,
      address: editForm.value.address,
      whatsapp: editForm.value.whatsapp,
      student_id: null
    }

    const res: any = await api.put(`/api/users/${teacherId.value}`, payload)
    if (res.success) {
      toastMessage.value = "Profil guru berhasil diperbarui"
      closeEditModal()
      await fetchTeacherProfile()
      setTimeout(() => {
        toastMessage.value = ""
      }, 3000)
    } else {
      editError.value = res.message || "Gagal memperbarui profil guru"
    }
  } catch (error: any) {
    if (error.response && error.response._data) {
      editError.value = error.response._data.message || "Gagal memproses data"
    } else {
      editError.value = "Koneksi ke server terputus"
    }
  } finally {
    editLoading.value = false
  }
}

function formatDate(dateStr: string): string {
  if (!dateStr) return "-"
  const d = new Date(dateStr)
  return d.toLocaleDateString("id-ID", {
    year: "numeric",
    month: "long",
    day: "numeric"
  })
}
</script>
