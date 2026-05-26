<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8">
        <div>
          <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
            Kehadiran Harian
          </h2>
          <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
            Absensi Kesiswaan
          </h1>
        </div>

        <div v-if="toastMessage" class="bg-[color:var(--color-surface)] border border-[color:var(--color-success)] p-3 text-sm text-[color:var(--color-success)] font-medium flex items-center space-x-2">
          <CheckCircleIcon class="w-4 h-4 shrink-0" />
          <span>{{ toastMessage }}</span>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
          <div class="lg:col-span-4" v-if="currentUser?.role === 'siswa' || currentUser?.role === 'siswi'">
            <div class="card space-y-6 shadow-[--shadow-sm]">
              <div>
                <h3 class="text-lg font-bold text-[color:var(--color-heading)] tracking-wide mb-1">
                  {{ isEdit ? 'Ubah Absensi' : 'Input Absensi Baru' }}
                </h3>
                <p class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider">
                  Catat kehadiran harian siswa secara manual
                </p>
              </div>

              <div v-if="errorMessage" class="bg-[color:var(--color-bg)] border border-[color:var(--color-error)] p-3 text-sm text-[color:var(--color-error)] font-medium flex items-center space-x-2">
                <AlertCircleIcon class="w-4 h-4 shrink-0" />
                <span>{{ errorMessage }}</span>
              </div>

              <form @submit.prevent="handleSubmit" class="space-y-4">
                <div class="space-y-1.5">
                  <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Tanggal</label>
                  <input v-model="formData.date" type="date" class="input w-full font-mono opacity-60 bg-[color:var(--color-surface)]" disabled />
                </div>

                <div class="space-y-1.5" v-if="!isEdit && currentUser?.role !== 'siswa' && currentUser?.role !== 'siswi'">
                  <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Siswa</label>
                  <select v-model="formData.student_id" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                    <option value="" disabled>Pilih Siswa</option>
                    <option v-for="student in studentsList" :key="student.id" :value="student.id">
                      {{ student.nisn }} - {{ student.name }} ({{ student.class }})
                    </option>
                  </select>
                </div>
                <div class="space-y-1.5" v-else>
                  <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Siswa</label>
                  <input type="text" class="input w-full opacity-60 bg-[color:var(--color-surface)]" readonly :value="(currentUser?.role === 'siswa' || currentUser?.role === 'siswi') ? currentUser?.name : formData.student_name" />
                </div>

                <div class="space-y-1.5">
                  <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Status Kehadiran</label>
                  <select v-model="formData.status" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                    <option value="" disabled>Pilih Status</option>
                    <option value="hadir">Hadir</option>
                    <option value="sakit">Sakit</option>
                    <option value="izin">Izin</option>
                    <option value="alpha">Alpha</option>
                  </select>
                </div>

                <div class="space-y-1.5">
                  <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Catatan / Keterangan</label>
                  <textarea
                    v-model="formData.note"
                    class="input w-full h-20 resize-none transition duration-150 disabled:cursor-not-allowed"
                    :class="(formData.status === 'hadir' || formData.status === 'alpha' || !formData.status) ? 'opacity-60 bg-[color:var(--color-surface)] cursor-not-allowed' : ''"
                    :placeholder="(formData.status === 'hadir' || formData.status === 'alpha') ? 'Catatan hanya untuk status Izin or Sakit' : 'Catatan opsional (misal: surat dokter)'"
                    :disabled="formData.status === 'hadir' || formData.status === 'alpha' || !formData.status"
                  ></textarea>
                </div>

                <div class="flex items-center space-x-2 pt-2">
                  <button type="submit" :disabled="isLoading" class="button-primary text-xs font-semibold uppercase tracking-wider flex-1 flex items-center justify-center py-3">
                    <span v-if="isLoading" class="animate-spin rounded-full h-3 w-3 border-2 border-[color:var(--color-accent-fg)] border-t-transparent mr-2"></span>
                    <span>{{ isEdit ? 'Perbarui Absensi' : 'Kirim Absensi' }}</span>
                  </button>
                  <button type="button" v-if="isEdit" @click="resetForm" class="button-ghost text-xs font-semibold uppercase tracking-wider py-3">
                    Batal
                  </button>
                </div>
              </form>
            </div>
          </div>

          <div :class="(currentUser?.role === 'siswa' || currentUser?.role === 'siswi') ? 'lg:col-span-8' : 'lg:col-span-12'">
            <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6">
              <div class="flex flex-col md:flex-row gap-4 items-center justify-between">
                <h3 class="text-lg font-bold text-[color:var(--color-heading)] tracking-wide">
                  Daftar Absensi Harian
                </h3>
                <div class="flex flex-wrap items-center gap-3 shrink-0 w-full md:w-auto">
                  <div class="relative w-full md:w-48" v-if="currentUser?.role !== 'siswa' && currentUser?.role !== 'siswi'">
                    <SearchIcon class="absolute left-3 top-1/2 -translate-y-1/2 w-3.5 h-3.5 text-[color:var(--color-muted)] pointer-events-none" />
                    <input
                      v-model="searchQuery"
                      @input="handleSearch"
                      type="text"
                      class="input py-1 text-xs w-full"
                      style="padding-left: 2.25rem !important;"
                      placeholder="Cari nama/NISN..."
                    />
                  </div>

                  <select v-model="filterClass" class="input py-1 text-xs bg-[color:var(--color-bg)] select-arrow w-full md:w-32" @change="loadAttendanceList" v-if="currentUser?.role !== 'siswa' && currentUser?.role !== 'siswi'">
                    <option value="">Semua Kelas</option>
                    <option value="X-A">X-A</option>
                    <option value="XI-B">XI-B</option>
                    <option value="XII-C">XII-C</option>
                  </select>

                  <div class="flex items-center space-x-2 w-full md:w-auto">
                    <span class="text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-semibold whitespace-nowrap">Filter Tanggal:</span>
                    <input v-model="filterDate" type="date" class="input py-1 text-xs font-mono w-full md:w-auto" @change="loadAttendanceList" />
                  </div>
                </div>
              </div>

              <div class="overflow-x-auto">
                <table class="w-full text-left border-collapse">
                  <thead>
                    <tr class="border-b border-[color:var(--color-border)]">
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">NISN</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Name</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Class</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Date</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Status</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Notes</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr
                      v-for="att in attendancesList"
                      :key="att.id"
                      class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                    >
                      <td class="py-3.5 px-4 text-sm font-mono text-[color:var(--color-accent)]">{{ att.student?.nisn || '-' }}</td>
                      <td class="py-3.5 px-4 text-sm font-semibold text-[color:var(--color-heading)]">
                        <NuxtLink :to="`/students/${att.student?.id}`" class="hover:text-[color:var(--color-accent)] transition duration-100">
                          {{ att.student?.name || '-' }}
                        </NuxtLink>
                      </td>
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">{{ att.student?.class || '-' }}</td>
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] font-mono">{{ formatDate(att.date) }}</td>
                      <td class="py-3.5 px-4 text-sm font-bold uppercase tracking-wider">
                        <span :class="getStatusClass(att.status)">
                          {{ att.status }}
                        </span>
                      </td>
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">{{ att.note || '-' }}</td>
                    </tr>
                    <tr v-if="attendancesList.length === 0">
                      <td colspan="6" class="py-12 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                        No attendance records found for the selected date
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from "vue"
import {
  AlertCircle as AlertCircleIcon,
  CheckCircle as CheckCircleIcon,
  Search as SearchIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()
const currentUser = useState<any>("currentUser", () => null)

const studentsList = ref<any[]>([])
const attendancesList = ref<any[]>([])
const filterDate = ref(new Date().toISOString().split("T")[0])
const searchQuery = ref("")
const filterClass = ref("")

let debounceTimer: any = null

function handleSearch() {
  clearTimeout(debounceTimer)
  debounceTimer = setTimeout(() => {
    loadAttendanceList()
  }, 300)
}
const toastMessage = ref("")
const errorMessage = ref("")
const isLoading = ref(false)
const isEdit = ref(false)
const editingRecordId = ref<number | null>(null)

const formData = ref({
  student_id: "",
  student_name: "",
  date: new Date().toISOString().split("T")[0],
  status: "",
  note: ""
})

watch(() => formData.value.status, (newStatus) => {
  if (newStatus === "hadir" || newStatus === "alpha") {
    formData.value.note = ""
  }
})

onMounted(() => {
  const cached = localStorage.getItem("user")
  if (cached && !currentUser.value) {
    currentUser.value = JSON.parse(cached)
  }
  if (currentUser.value?.role === 'siswa' || currentUser.value?.role === 'siswi') {
    formData.value.student_id = currentUser.value.student_id
  }
  loadStudents()
  loadAttendanceList()
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

async function loadAttendanceList() {
  try {
    let url = `/api/attendances?date=${filterDate.value}`
    if (filterClass.value) {
      url += `&class=${filterClass.value}`
    }
    if (searchQuery.value) {
      url += `&search=${searchQuery.value}`
    }
    const res: any = await api.get(url)
    if (res.success && res.data) {
      attendancesList.value = res.data || []
    }
  } catch (error) {
    console.error("Gagal memuat rekap absensi", error)
  }
}

async function handleSubmit() {
  if (!formData.value.date || !formData.value.status || (!isEdit.value && !formData.value.student_id)) {
    errorMessage.value = "Semua kolom wajib diisi"
    return
  }

  isLoading.value = true
  errorMessage.value = ""

  try {
    let res: any
    if (isEdit.value && editingRecordId.value) {
      res = await api.put(`/api/attendances/${editingRecordId.value}`, {
        status: formData.value.status,
        note: formData.value.note
      })
    } else {
      res = await api.post("/api/attendances", {
        student_id: formData.value.student_id,
        date: formData.value.date,
        status: formData.value.status,
        note: formData.value.note
      })
    }

    if (res.success) {
      showToast(isEdit.value ? "Absensi berhasil diubah" : "Absensi berhasil dicatat")
      resetForm()
      loadAttendanceList()
    } else {
      errorMessage.value = res.message || "Gagal menyimpan data absensi"
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

function startEdit(record: any) {
  isEdit.value = true
  editingRecordId.value = record.id
  formData.value = {
    student_id: record.student_id,
    student_name: `${record.student?.nisn} - ${record.student?.name}`,
    date: record.date.split("T")[0],
    status: record.status,
    note: record.note || ""
  }
  errorMessage.value = ""
}

function resetForm() {
  isEdit.value = false
  editingRecordId.value = null
  formData.value = {
    student_id: "",
    student_name: "",
    date: filterDate.value,
    status: "",
    note: ""
  }
  errorMessage.value = ""
}

function showToast(msg: string) {
  toastMessage.value = msg
  setTimeout(() => {
    toastMessage.value = ""
  }, 3000)
}

function formatDate(dateStr: string): string {
  if (!dateStr) return "-"
  return dateStr.split("T")[0] || "-"
}

function getStatusClass(status: string): string {
  switch (status.toLowerCase()) {
    case "hadir":
      return "text-[color:var(--color-success)] text-xs"
    case "sakit":
      return "text-orange-400 text-xs"
    case "izin":
      return "text-blue-400 text-xs"
    case "alpha":
      return "text-[color:var(--color-error)] text-xs"
    default:
      return "text-[color:var(--color-text)] text-xs"
  }
}
</script>
