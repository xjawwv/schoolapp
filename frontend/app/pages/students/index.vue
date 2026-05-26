<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8">
        <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
          <div>
            <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
              Manajemen Data
            </h2>
            <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
              Data Siswa Terdaftar
            </h1>
          </div>
          <button v-if="currentUser?.role === 'admin'" @click="openCreateModal" class="button-primary flex items-center space-x-2 self-start md:self-auto py-3">
            <PlusIcon class="w-4 h-4" />
            <span>Tambah Siswa</span>
          </button>
        </div>

        <div v-if="toastMessage" class="bg-[color:var(--color-surface)] border border-[color:var(--color-success)] p-3 text-sm text-[color:var(--color-success)] font-medium flex items-center space-x-2">
          <CheckCircleIcon class="w-4 h-4 shrink-0" />
          <span>{{ toastMessage }}</span>
        </div>

        <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6">
          <div class="flex flex-col md:flex-row gap-4 items-center justify-between">
            <div class="flex flex-col md:flex-row gap-3 w-full md:max-w-xl">
              <div class="relative flex-1">
                <SearchIcon class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-[color:var(--color-muted)] pointer-events-none" />
                <input
                  v-model="searchQuery"
                  @input="handleSearch"
                  type="text"
                  class="input w-full"
                  style="padding-left: 2.5rem !important;"
                  placeholder="Cari berdasarkan nama, kelas, atau NISN..."
                />
              </div>
              <select v-model="filterClass" class="input bg-[color:var(--color-bg)] select-arrow w-full md:w-40" @change="handleSearch">
                <option value="">Semua Kelas</option>
                <option value="X-A">X-A</option>
                <option value="XI-B">XI-B</option>
                <option value="XII-C">XII-C</option>
              </select>
            </div>
            <div class="text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-semibold shrink-0">
              Total Siswa: <span class="text-[color:var(--color-accent)]">{{ totalStudents }}</span>
            </div>
          </div>

          <div class="overflow-x-auto">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="border-b border-[color:var(--color-border)]">
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-12 text-center">No</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">NISN</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Name</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Class</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Gender</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Phone</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold text-right" v-if="currentUser?.role === 'admin'">Actions</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="(student, index) in students"
                  :key="student.id"
                  class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                >
                  <td class="py-3.5 px-4 text-sm text-[color:var(--color-muted)] font-semibold text-center font-mono w-12">{{ (currentPage - 1) * 10 + index + 1 }}</td>
                  <td class="py-3.5 px-4 text-sm font-mono text-[color:var(--color-accent)]">{{ student.nisn }}</td>
                  <td class="py-3.5 px-4 text-sm font-semibold text-[color:var(--color-heading)]">
                    <NuxtLink :to="`/students/${student.id}`" class="hover:text-[color:var(--color-accent)] transition duration-100">
                      {{ student.name }}
                    </NuxtLink>
                  </td>
                  <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">{{ student.class }}</td>
                  <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">{{ student.gender }}</td>
                  <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] font-mono">{{ student.phone || '-' }}</td>
                  <td class="py-3.5 px-4 text-sm text-right relative" v-if="currentUser?.role === 'admin'">
                    <button
                      @click="toggleDropdown(student.id)"
                      class="text-[color:var(--color-muted)] hover:text-[color:var(--color-accent)] transition duration-150 p-1 cursor-pointer focus:outline-none"
                    >
                      <MoreHorizontalIcon class="w-5 h-5 inline" />
                    </button>

                    <div
                      v-if="activeDropdownId === student.id"
                      class="absolute right-4 mt-1 w-32 bg-[color:var(--color-surface)] border border-[color:var(--color-border)] shadow-[--shadow-md] py-1 z-20 flex flex-col text-left"
                    >
                      <button
                        @click="handleEditAction(student)"
                        class="flex items-center space-x-2 px-4 py-2.5 text-xs uppercase tracking-wider text-[color:var(--color-text)] hover:bg-[color:var(--color-bg)] hover:text-[color:var(--color-heading)] transition duration-150 w-full text-left cursor-pointer font-semibold"
                      >
                        <PencilIcon class="w-3.5 h-3.5 shrink-0" />
                        <span>Ubah</span>
                      </button>
                      <button
                        @click="handleDeleteAction(student)"
                        class="flex items-center space-x-2 px-4 py-2.5 text-xs uppercase tracking-wider text-[color:var(--color-error)] hover:bg-[color:var(--color-bg)] transition duration-150 w-full text-left cursor-pointer font-semibold"
                      >
                        <Trash2Icon class="w-3.5 h-3.5 shrink-0" />
                        <span>Hapus</span>
                      </button>
                    </div>

                    <div
                      v-if="activeDropdownId === student.id"
                      @click="activeDropdownId = null"
                      class="fixed inset-0 z-10 cursor-default"
                    ></div>
                  </td>
                </tr>
                <tr v-if="students.length === 0">
                  <td :colspan="currentUser?.role === 'admin' ? 7 : 6" class="py-12 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                    No student data found
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <div class="flex items-center justify-between border-t border-[color:var(--color-border)] pt-6" v-if="totalPages > 1">
            <span class="text-xs text-[color:var(--color-muted)] font-semibold uppercase tracking-wider">
              Halaman {{ currentPage }} dari {{ totalPages }}
            </span>
            <div class="flex space-x-2">
              <button
                @click="changePage(currentPage - 1)"
                :disabled="currentPage === 1"
                class="button-ghost py-2 px-3 text-xs font-semibold uppercase tracking-wider disabled:opacity-30 cursor-pointer"
              >
                Kembali
              </button>
              <button
                @click="changePage(currentPage + 1)"
                :disabled="currentPage === totalPages"
                class="button-ghost py-2 px-3 text-xs font-semibold uppercase tracking-wider disabled:opacity-30 cursor-pointer"
              >
                Lanjut
              </button>
            </div>
          </div>
        </div>
      </main>

      <StudentForm
        :show="showForm"
        :student="selectedStudent"
        @close="closeForm"
        @saved="onStudentSaved"
      />

      <div v-if="showDeleteConfirm" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm">
        <div class="w-full max-w-md bg-[color:var(--color-surface)] border border-[color:var(--color-border)] border-l-4 border-l-[color:var(--color-error)] p-8 shadow-[--shadow-lg]">
          <h3 class="text-xl font-bold text-[color:var(--color-heading)] tracking-tight mb-2">
            Hapus Data Siswa
          </h3>
          <p class="text-sm text-[color:var(--color-text)] mb-6">
            Apakah Anda yakin ingin menghapus data <span class="font-bold text-[color:var(--color-heading)]">{{ studentToDelete?.name }}</span>? Aksi ini akan menghapus semua riwayat kehadiran dan nilai siswa tersebut secara permanen.
          </p>
          <div class="flex items-center justify-end space-x-3">
            <button @click="closeDeleteConfirm" class="button-ghost text-xs font-semibold uppercase tracking-wider">
              Batal
            </button>
            <button @click="handleDelete" class="button-primary !bg-[color:var(--color-error)] !text-white text-xs font-semibold uppercase tracking-wider">
              Hapus Permanen
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import {
  Plus as PlusIcon,
  Search as SearchIcon,
  CheckCircle as CheckCircleIcon,
  ChevronLeft as ChevronLeftIcon,
  ChevronRight as ChevronRightIcon,
  MoreHorizontal as MoreHorizontalIcon,
  Pencil as PencilIcon,
  Trash2 as Trash2Icon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"
import StudentForm from "~/components/StudentForm.vue"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()
const currentUser = useState<any>("currentUser", () => null)

const students = ref<any[]>([])
const totalStudents = ref(0)
const currentPage = ref(1)
const totalPages = ref(1)
const searchQuery = ref("")
const filterClass = ref("")

const showForm = ref(false)
const selectedStudent = ref<any | null>(null)
const toastMessage = ref("")

const showDeleteConfirm = ref(false)
const studentToDelete = ref<any | null>(null)
const activeDropdownId = ref<string | null>(null)

function toggleDropdown(id: string) {
  if (activeDropdownId.value === id) {
    activeDropdownId.value = null
  } else {
    activeDropdownId.value = id
  }
}

function handleEditAction(student: any) {
  activeDropdownId.value = null
  openEditModal(student)
}

function handleDeleteAction(student: any) {
  activeDropdownId.value = null
  confirmDelete(student)
}

let debounceTimer: any = null

onMounted(() => {
  const cached = localStorage.getItem("user")
  if (cached && !currentUser.value) {
    currentUser.value = JSON.parse(cached)
  }
  fetchStudents()
})

async function fetchStudents() {
  try {
    const res: any = await api.get(`/api/students?page=${currentPage.value}&limit=10&search=${searchQuery.value}&class=${filterClass.value}`)
    if (res.success && res.data) {
      students.value = res.data.students || []
      totalStudents.value = res.data.total || 0
      totalPages.value = res.data.total_pages || 1
    }
  } catch (error) {
    console.error("Gagal memuat siswa", error)
  }
}

function handleSearch() {
  clearTimeout(debounceTimer)
  debounceTimer = setTimeout(() => {
    currentPage.value = 1
    fetchStudents()
  }, 300)
}

function changePage(page: number) {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    fetchStudents()
  }
}

function openCreateModal() {
  selectedStudent.value = null
  showForm.value = true
}

function openEditModal(student: any) {
  selectedStudent.value = student
  showForm.value = true
}

function closeForm() {
  showForm.value = false
  selectedStudent.value = null
}

function onStudentSaved(savedStudent: any) {
  showToast(selectedStudent.value ? "Data siswa berhasil diubah" : "Siswa baru berhasil ditambahkan")
  fetchStudents()
}

function confirmDelete(student: any) {
  studentToDelete.value = student
  showDeleteConfirm.value = true
}

function closeDeleteConfirm() {
  showDeleteConfirm.value = false
  studentToDelete.value = null
}

async function handleDelete() {
  if (!studentToDelete.value) return

  try {
    const res: any = await api.delete(`/api/students/${studentToDelete.value.id}`)
    if (res.success) {
      showToast("Data siswa dan riwayat relasinya berhasil dihapus")
      fetchStudents()
    }
  } catch (error) {
    console.error("Gagal menghapus siswa", error)
  } finally {
    closeDeleteConfirm()
  }
}

function showToast(msg: string) {
  toastMessage.value = msg
  setTimeout(() => {
    toastMessage.value = ""
  }, 3000)
}
</script>
