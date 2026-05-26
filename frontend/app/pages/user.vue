<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8">
        <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
          <div>
            <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
              Konfigurasi Sistem
            </h2>
            <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
              Manajemen Pengguna
            </h1>
          </div>
          <button @click="openCreateModal" class="button-primary flex items-center space-x-2 self-start md:self-auto py-3">
            <UserPlusIcon class="w-4 h-4" />
            <span>Tambah Pengguna</span>
          </button>
        </div>

        <div v-if="toastMessage" class="bg-[color:var(--color-surface)] border border-[color:var(--color-success)] p-3 text-sm text-[color:var(--color-success)] font-medium flex items-center space-x-2">
          <CheckCircleIcon class="w-4 h-4 shrink-0" />
          <span>{{ toastMessage }}</span>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 sm:gap-6">
          <div class="card shadow-[--shadow-sm] border-l-4 border-l-[color:var(--color-accent)] flex items-center space-x-4">
            <div class="w-12 h-12 rounded-full bg-[color:var(--color-bg)] flex items-center justify-center text-[color:var(--color-accent)]">
              <UserIcon class="w-6 h-6" />
            </div>
            <div>
              <span class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider block font-bold">Total Akun</span>
              <span class="text-2xl font-bold text-[color:var(--color-heading)] mt-0.5 block font-mono">{{ users.length }}</span>
            </div>
          </div>
          <div class="card shadow-[--shadow-sm] border-l-4 border-l-purple-500 flex items-center space-x-4">
            <div class="w-12 h-12 rounded-full bg-purple-500/10 flex items-center justify-center text-purple-500">
              <ShieldIcon class="w-6 h-6" />
            </div>
            <div>
              <span class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider block font-bold">Admin</span>
              <span class="text-2xl font-bold text-[color:var(--color-heading)] mt-0.5 block font-mono">{{ countRole('admin') }}</span>
            </div>
          </div>
          <div class="card shadow-[--shadow-sm] border-l-4 border-l-emerald-500 flex items-center space-x-4">
            <div class="w-12 h-12 rounded-full bg-emerald-500/10 flex items-center justify-center text-emerald-500">
              <UserCheckIcon class="w-6 h-6" />
            </div>
            <div>
              <span class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider block font-bold">Guru</span>
              <span class="text-2xl font-bold text-[color:var(--color-heading)] mt-0.5 block font-mono">{{ countRole('guru') }}</span>
            </div>
          </div>
          <div class="card shadow-[--shadow-sm] border-l-4 border-l-blue-500 flex items-center space-x-4">
            <div class="w-12 h-12 rounded-full bg-blue-500/10 flex items-center justify-center text-blue-500">
              <UsersIcon class="w-6 h-6" />
            </div>
            <div>
              <span class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider block font-bold">Siswa & Siswi</span>
              <span class="text-2xl font-bold text-[color:var(--color-heading)] mt-0.5 block font-mono">{{ countRole('siswa') + countRole('siswi') }}</span>
            </div>
          </div>
        </div>

        <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6">
          <div class="flex flex-col md:flex-row gap-4 items-center justify-between">
            <div class="flex flex-col md:flex-row gap-3 w-full md:max-w-xl">
              <div class="relative flex-1">
                <SearchIcon class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-[color:var(--color-muted)] pointer-events-none" />
                <input
                  v-model="searchQuery"
                  type="text"
                  class="input w-full"
                  style="padding-left: 2.5rem !important;"
                  placeholder="Cari berdasarkan nama atau email..."
                />
              </div>
              <select v-model="filterRole" class="input bg-[color:var(--color-bg)] select-arrow w-full md:w-44">
                <option value="">Semua Peran</option>
                <option value="admin">Admin</option>
                <option value="guru">Guru</option>
                <option value="siswa">Siswa (Laki-laki)</option>
                <option value="siswi">Siswi (Perempuan)</option>
              </select>
            </div>
            <div class="text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-semibold shrink-0">
              Menampilkan: <span class="text-[color:var(--color-accent)]">{{ filteredUsers.length }}</span> Pengguna
            </div>
          </div>

          <div class="overflow-x-auto">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="border-b border-[color:var(--color-border)]">
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-12 text-center">No</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-10">Info</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Nama Lengkap</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Email</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Peran</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Siswa Terkait</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold text-right">Aksi</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="(u, index) in filteredUsers"
                  :key="u.id"
                  class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                >
                  <td class="py-3.5 px-4 text-sm text-[color:var(--color-muted)] font-semibold text-center font-mono w-12">{{ index + 1 }}</td>
                  <td class="py-3.5 px-4">
                    <div
                      class="w-8 h-8 rounded-full border flex items-center justify-center text-xs font-bold font-mono"
                      :class="getAvatarBgClass(u.name)"
                    >
                      {{ u.name.charAt(0).toUpperCase() }}
                    </div>
                  </td>
                  <td class="py-3.5 px-4 text-sm font-semibold text-[color:var(--color-heading)]">
                    {{ u.name }}
                    <span v-if="currentUser?.id === u.id" class="ml-1.5 text-[9px] font-bold uppercase tracking-wider bg-[color:var(--color-accent)] text-[color:var(--color-accent-fg)] py-0.5 px-1.5 rounded-sm">Saya</span>
                  </td>
                  <td class="py-3.5 px-4 text-sm font-mono text-[color:var(--color-text)]">{{ u.email }}</td>
                  <td class="py-3.5 px-4 text-sm font-bold uppercase tracking-wider">
                    <span :class="getRoleBadgeClass(u.role)">
                      {{ u.role }}
                    </span>
                  </td>
                  <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">
                    <span v-if="u.student" class="flex items-center space-x-1.5">
                      <NuxtLink :to="`/students/${u.student.id}`" class="text-[color:var(--color-accent)] font-semibold hover:underline">
                        {{ u.student.name }} ({{ u.student.class }})
                      </NuxtLink>
                    </span>
                    <span v-else class="text-[color:var(--color-muted)] text-xs font-mono">-</span>
                  </td>
                  <td class="py-3.5 px-4 text-sm text-right">
                    <div class="flex items-center justify-end space-x-2">
                      <button
                        @click="openEditModal(u)"
                        class="p-1.5 rounded-md border border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] hover:text-[color:var(--color-accent)] text-[color:var(--color-muted)] transition duration-150 cursor-pointer"
                        title="Ubah data pengguna"
                      >
                        <PencilIcon class="w-3.5 h-3.5" />
                      </button>
                      <button
                        @click="confirmDelete(u)"
                        :disabled="currentUser?.id === u.id"
                        class="p-1.5 rounded-md border border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] hover:text-[color:var(--color-error)] text-[color:var(--color-muted)] transition duration-150 cursor-pointer disabled:opacity-30 disabled:pointer-events-none"
                        title="Hapus pengguna"
                      >
                        <Trash2Icon class="w-3.5 h-3.5" />
                      </button>
                    </div>
                  </td>
                </tr>
                <tr v-if="filteredUsers.length === 0">
                  <td colspan="7" class="py-12 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                    Tidak ada pengguna ditemukan
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </main>

      <div v-if="showFormModal" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm">
        <div class="w-full max-w-lg bg-[color:var(--color-surface)] border border-[color:var(--color-border)] border-l-4 border-l-[color:var(--color-accent)] p-8 shadow-[--shadow-lg] relative">
          <div class="mb-6">
            <h3 class="text-2xl font-bold text-[color:var(--color-heading)] tracking-tight">
              {{ isEdit ? 'Ubah Data Pengguna' : 'Tambah Pengguna Baru' }}
            </h3>
            <p class="text-xs uppercase tracking-wider text-[color:var(--color-muted)] mt-1">
              Lengkapi kredensial otentikasi dan peran akses sistem
            </p>
          </div>

          <div v-if="modalError" class="mb-4 bg-[color:var(--color-bg)] border border-[color:var(--color-error)] p-3 text-sm text-[color:var(--color-error)] font-medium flex items-center space-x-2">
            <AlertCircleIcon class="w-4 h-4 shrink-0" />
            <span>{{ modalError }}</span>
          </div>

          <form @submit.prevent="handleSubmit" class="space-y-4">
            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Nama Lengkap</label>
              <input v-model="formData.name" type="text" class="input w-full" required placeholder="Contoh: Administrator Utama" />
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Alamat Email</label>
              <input v-model="formData.email" type="email" class="input w-full font-mono text-[color:var(--color-accent)]" required placeholder="Contoh: admin@sekolah.com" />
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">
                {{ isEdit ? 'Kata Sandi Baru (Kosongkan jika tidak diubah)' : 'Kata Sandi' }}
              </label>
              <input v-model="formData.password" type="password" class="input w-full font-mono" :required="!isEdit" placeholder="Minimal 6 karakter" />
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="space-y-1.5">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Peran Akses (Role)</label>
                <select v-model="formData.role" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                  <option value="" disabled>Pilih Peran</option>
                  <option value="admin">Admin</option>
                  <option value="guru">Guru</option>
                  <option value="siswa">Siswa (Laki-laki)</option>
                  <option value="siswi">Siswi (Perempuan)</option>
                </select>
              </div>

              <div class="space-y-1.5" v-if="formData.role === 'siswa' || formData.role === 'siswi'">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Biodata Siswa Terhubung</label>
                <select v-model="formData.student_id" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                  <option value="" disabled>Pilih Biodata Siswa</option>
                  <option v-for="std in studentsList" :key="std.id" :value="std.id">
                    {{ std.nisn }} - {{ std.name }} ({{ std.class }})
                  </option>
                </select>
              </div>
            </div>

            <div class="flex items-center justify-end space-x-3 pt-4 border-t border-[color:var(--color-border)]">
              <button type="button" @click="closeFormModal" class="button-ghost text-xs font-semibold uppercase tracking-wider">
                Batal
              </button>
              <button type="submit" :disabled="modalLoading" class="button-primary text-xs font-semibold uppercase tracking-wider flex items-center space-x-1">
                <span v-if="modalLoading" class="animate-spin rounded-full h-3 w-3 border-2 border-[color:var(--color-accent-fg)] border-t-transparent mr-1"></span>
                <span>Simpan</span>
              </button>
            </div>
          </form>
        </div>
      </div>

      <div v-if="showDeleteModal" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm">
        <div class="w-full max-w-md bg-[color:var(--color-surface)] border border-[color:var(--color-border)] border-l-4 border-l-[color:var(--color-error)] p-8 shadow-[--shadow-lg]">
          <h3 class="text-xl font-bold text-[color:var(--color-heading)] tracking-tight mb-2">
            Hapus Pengguna
          </h3>
          <p class="text-sm text-[color:var(--color-text)] mb-6">
            Apakah Anda yakin ingin menghapus kredensial pengguna <span class="font-bold text-[color:var(--color-heading)]">{{ userToDelete?.name }}</span> secara permanen? Akun ini tidak akan dapat login lagi ke dalam sistem.
          </p>
          <div class="flex items-center justify-end space-x-3">
            <button @click="closeDeleteModal" class="button-ghost text-xs font-semibold uppercase tracking-wider">
              Batal
            </button>
            <button @click="handleDelete" class="button-primary !bg-[color:var(--color-error)] !text-white text-xs font-semibold uppercase tracking-wider">
              Hapus Kredensial
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from "vue"
import { useRouter } from "vue-router"
import {
  Users as UsersIcon,
  Plus as PlusIcon,
  Search as SearchIcon,
  Pencil as PencilIcon,
  Trash2 as Trash2Icon,
  UserCheck as UserCheckIcon,
  Shield as ShieldIcon,
  UserPlus as UserPlusIcon,
  CheckCircle as CheckCircleIcon,
  AlertCircle as AlertCircleIcon,
  X as XIcon,
  User as UserIcon
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

const users = ref<any[]>([])
const studentsList = ref<any[]>([])
const searchQuery = ref("")
const filterRole = ref("")
const toastMessage = ref("")

const showFormModal = ref(false)
const isEdit = ref(false)
const selectedUser = ref<any | null>(null)
const modalError = ref("")
const modalLoading = ref(false)

const showDeleteModal = ref(false)
const userToDelete = ref<any | null>(null)

const formData = ref({
  name: "",
  email: "",
  password: "",
  role: "",
  student_id: ""
})

watch(() => formData.value.role, (newRole) => {
  if (newRole !== 'siswa' && newRole !== 'siswi') {
    formData.value.student_id = ""
  }
})

const filteredUsers = computed(() => {
  return users.value.filter(u => {
    const matchSearch =
      u.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      u.email.toLowerCase().includes(searchQuery.value.toLowerCase())

    const matchRole = !filterRole.value || u.role === filterRole.value
    return matchSearch && matchRole
  })
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

  loadUsers()
  loadStudents()
})

async function loadUsers() {
  try {
    const res: any = await api.get("/api/users")
    if (res.success && res.data) {
      users.value = res.data || []
    }
  } catch (error) {
    console.error("Gagal memuat pengguna", error)
  }
}

async function loadStudents() {
  try {
    const res: any = await api.get("/api/students?limit=200")
    if (res.success && res.data) {
      studentsList.value = res.data.students || []
    }
  } catch (error) {
    console.error("Gagal memuat referensi siswa", error)
  }
}

function countRole(role: string): number {
  return users.value.filter(u => u.role === role).length
}

function getAvatarBgClass(name: string): string {
  const colors = [
    "bg-red-500/10 text-red-500 border-red-500/20",
    "bg-blue-500/10 text-blue-500 border-blue-500/20",
    "bg-green-500/10 text-green-500 border-green-500/20",
    "bg-purple-500/10 text-purple-500 border-purple-500/20",
    "bg-orange-500/10 text-orange-500 border-orange-500/20",
    "bg-pink-500/10 text-pink-500 border-pink-500/20",
    "bg-teal-500/10 text-teal-500 border-teal-500/20"
  ]
  let sum = 0
  for (let i = 0; i < name.length; i++) {
    sum += name.charCodeAt(i)
  }
  return colors[sum % colors.length]
}

function getRoleBadgeClass(role: string): string {
  switch (role.toLowerCase()) {
    case "admin":
      return "text-purple-500 bg-purple-500/10 text-xs px-2.5 py-0.5 rounded-sm inline-block"
    case "guru":
      return "text-emerald-500 bg-emerald-500/10 text-xs px-2.5 py-0.5 rounded-sm inline-block"
    case "siswa":
    case "siswi":
      return "text-blue-500 bg-blue-500/10 text-xs px-2.5 py-0.5 rounded-sm inline-block"
    default:
      return "text-[color:var(--color-text)] bg-[color:var(--color-bg)] text-xs px-2.5 py-0.5 rounded-sm inline-block"
  }
}

function openCreateModal() {
  isEdit.value = false
  selectedUser.value = null
  formData.value = {
    name: "",
    email: "",
    password: "",
    role: "",
    student_id: ""
  }
  modalError.value = ""
  showFormModal.value = true
}

function openEditModal(u: any) {
  isEdit.value = true
  selectedUser.value = u
  formData.value = {
    name: u.name,
    email: u.email,
    password: "",
    role: u.role,
    student_id: u.student_id || ""
  }
  modalError.value = ""
  showFormModal.value = true
}

function closeFormModal() {
  showFormModal.value = false
  selectedUser.value = null
}

async function handleSubmit() {
  if (!formData.value.name || !formData.value.email || !formData.value.role) {
    modalError.value = "Kolom Nama, Email, dan Peran wajib diisi"
    return
  }

  if (!isEdit.value && (!formData.value.password || formData.value.password.length < 6)) {
    modalError.value = "Kata sandi wajib diisi minimal 6 karakter"
    return
  }

  if (isEdit.value && formData.value.password && formData.value.password.length < 6) {
    modalError.value = "Kata sandi baru minimal 6 karakter"
    return
  }

  modalLoading.value = true
  modalError.value = ""

  try {
    let payload: any = {
      name: formData.value.name,
      email: formData.value.email,
      role: formData.value.role,
      student_id: formData.value.student_id ? formData.value.student_id : null
    }

    if (formData.value.password) {
      payload.password = formData.value.password
    }

    let res: any
    if (isEdit.value && selectedUser.value) {
      res = await api.put(`/api/users/${selectedUser.value.id}`, payload)
    } else {
      res = await api.post("/api/users", payload)
    }

    if (res.success) {
      showToast(isEdit.value ? "Data pengguna berhasil diubah" : "Pengguna baru berhasil ditambahkan")
      closeFormModal()
      loadUsers()
    } else {
      modalError.value = res.message || "Gagal memproses data"
    }
  } catch (error: any) {
    if (error.response && error.response._data) {
      modalError.value = error.response._data.message || "Gagal memproses data"
    } else {
      modalError.value = "Koneksi ke server terputus"
    }
  } finally {
    modalLoading.value = false
  }
}

function confirmDelete(u: any) {
  if (currentUser.value?.id === u.id) return
  userToDelete.value = u
  showDeleteModal.value = true
}

function closeDeleteModal() {
  showDeleteModal.value = false
  userToDelete.value = null
}

async function handleDelete() {
  if (!userToDelete.value) return

  try {
    const res: any = await api.delete(`/api/users/${userToDelete.value.id}`)
    if (res.success) {
      showToast("Kredensial pengguna berhasil dihapus")
      loadUsers()
    }
  } catch (error) {
    console.error("Gagal menghapus pengguna", error)
  } finally {
    closeDeleteModal()
  }
}

function showToast(msg: string) {
  toastMessage.value = msg
  setTimeout(() => {
    toastMessage.value = ""
  }, 3000)
}
</script>
