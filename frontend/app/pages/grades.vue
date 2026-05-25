<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-8 max-w-7xl w-full mx-auto space-y-8">
        <div>
          <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
            Hasil Belajar
          </h2>
          <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
            Nilai Akademik Siswa
          </h1>
        </div>

        <div v-if="toastMessage" class="bg-[color:var(--color-surface)] border border-[color:var(--color-success)] p-3 text-sm text-[color:var(--color-success)] font-medium flex items-center space-x-2">
          <CheckCircleIcon class="w-4 h-4 shrink-0" />
          <span>{{ toastMessage }}</span>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
          <div class="lg:col-span-4">
            <div class="card space-y-6 shadow-[--shadow-sm]">
              <div>
                <h3 class="text-lg font-bold text-[color:var(--color-heading)] tracking-wide mb-1">
                  {{ isEdit ? 'Ubah Catatan Nilai' : 'Input Nilai Baru' }}
                </h3>
                <p class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider">
                  Catat nilai evaluasi mata pelajaran siswa
                </p>
              </div>

              <div v-if="errorMessage" class="bg-[color:var(--color-bg)] border border-[color:var(--color-error)] p-3 text-sm text-[color:var(--color-error)] font-medium flex items-center space-x-2">
                <AlertCircleIcon class="w-4 h-4 shrink-0" />
                <span>{{ errorMessage }}</span>
              </div>

              <form @submit.prevent="handleSubmit" class="space-y-4">
                <div class="space-y-1.5" v-if="!isEdit">
                  <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Siswa</label>
                  <select v-model="formData.student_id" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                    <option value="" disabled>Pilih Siswa</option>
                    <option v-for="student in studentsList" :key="student.id" :value="student.id">
                      {{ student.nis }} - {{ student.name }} ({{ student.class }})
                    </option>
                  </select>
                </div>
                <div class="space-y-1.5" v-else>
                  <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Siswa</label>
                  <input type="text" class="input w-full opacity-60 bg-[color:var(--color-surface)]" readonly :value="formData.student_name" />
                </div>

                <div class="space-y-1.5">
                  <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Mata Pelajaran</label>
                  <select v-model="formData.subject" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                    <option value="" disabled>Pilih Bidang Studi</option>
                    <option value="Matematika">Matematika</option>
                    <option value="Fisika">Fisika</option>
                    <option value="Kimia">Kimia</option>
                    <option value="Biologi">Biologi</option>
                    <option value="Bahasa Indonesia">Bahasa Indonesia</option>
                    <option value="Bahasa Inggris">Bahasa Inggris</option>
                    <option value="Sejarah">Sejarah</option>
                    <option value="Geografi">Geografi</option>
                    <option value="Ekonomi">Ekonomi</option>
                    <option value="Sosiologi">Sosiologi</option>
                  </select>
                </div>

                <div class="grid grid-cols-2 gap-4">
                  <div class="space-y-1.5">
                    <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Semester</label>
                    <select v-model.number="formData.semester" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                      <option :value="1">Ganjil (1)</option>
                      <option :value="2">Genap (2)</option>
                    </select>
                  </div>
                  <div class="space-y-1.5">
                    <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Tahun Ajaran</label>
                    <select v-model="formData.academic_year" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                      <option value="2024/2025">2024/2025</option>
                      <option value="2025/2026">2025/2026</option>
                      <option value="2026/2027">2026/2027</option>
                    </select>
                  </div>
                </div>

                <div class="space-y-1.5">
                  <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Nilai Evaluasi (0 - 100)</label>
                  <input
                    v-model.number="formData.score"
                    type="number"
                    step="0.1"
                    min="0"
                    max="100"
                    class="input w-full font-mono text-[color:var(--color-accent)] font-bold"
                    required
                    placeholder="Contoh: 85.5"
                  />
                </div>

                <div class="flex items-center space-x-2 pt-2">
                  <button type="submit" :disabled="isLoading" class="button-primary text-xs font-semibold uppercase tracking-wider flex-1 flex items-center justify-center py-3">
                    <span v-if="isLoading" class="animate-spin rounded-full h-3 w-3 border-2 border-[color:var(--color-accent-fg)] border-t-transparent mr-2"></span>
                    <span>{{ isEdit ? 'Perbarui Nilai' : 'Simpan Nilai' }}</span>
                  </button>
                  <button type="button" v-if="isEdit" @click="resetForm" class="button-ghost text-xs font-semibold uppercase tracking-wider py-3">
                    Batal
                  </button>
                </div>
              </form>
            </div>
          </div>

          <div class="lg:col-span-8">
            <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6">
              <div class="flex flex-col md:flex-row gap-4 items-center justify-between">
                <h3 class="text-lg font-bold text-[color:var(--color-heading)] tracking-wide">
                  Daftar Nilai Siswa
                </h3>
                <div class="flex items-center space-x-2 shrink-0 w-full md:w-auto">
                  <span class="text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-semibold whitespace-nowrap">Filter Siswa:</span>
                  <select v-model="filterStudentID" class="input py-1 text-xs bg-[color:var(--color-bg)] select-arrow w-full md:max-w-xs" @change="loadGradesList">
                    <option value="">Semua Siswa</option>
                    <option v-for="student in studentsList" :key="student.id" :value="student.id">
                      {{ student.name }} ({{ student.class }})
                    </option>
                  </select>
                </div>
              </div>

              <div class="overflow-x-auto">
                <table class="w-full text-left border-collapse">
                  <thead>
                    <tr class="border-b border-[color:var(--color-border)]">
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">NIS</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Nama</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Mata Pelajaran</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Semester</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Tahun Ajaran</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold text-right">Nilai</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold text-right">Aksi</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr
                      v-for="grade in gradesList"
                      :key="grade.id"
                      class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                    >
                      <td class="py-3.5 px-4 text-sm font-mono text-[color:var(--color-accent)]">{{ grade.student?.nis || '-' }}</td>
                      <td class="py-3.5 px-4 text-sm font-semibold text-[color:var(--color-heading)]">
                        <NuxtLink :to="`/students/${grade.student?.id}`" class="hover:text-[color:var(--color-accent)] transition duration-100">
                          {{ grade.student?.name || '-' }}
                        </NuxtLink>
                      </td>
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">{{ grade.subject }}</td>
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">Semester {{ grade.semester }}</td>
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] font-mono">{{ grade.academic_year }}</td>
                      <td class="py-3.5 px-4 text-sm text-right font-bold text-[color:var(--color-accent)] font-mono">{{ grade.score.toFixed(1) }}</td>
                      <td class="py-3.5 px-4 text-sm text-right">
                        <button @click="startEdit(grade)" class="text-xs uppercase tracking-wider text-[color:var(--color-accent)] hover:opacity-80 transition duration-100 cursor-pointer font-semibold">
                          Ubah
                        </button>
                      </td>
                    </tr>
                    <tr v-if="gradesList.length === 0">
                      <td colspan="7" class="py-12 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                        Tidak ada catatan nilai ditemukan
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
import { ref, onMounted } from "vue"
import {
  AlertCircle as AlertCircleIcon,
  CheckCircle as CheckCircleIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()

const studentsList = ref<any[]>([])
const gradesList = ref<any[]>([])
const filterStudentID = ref("")
const toastMessage = ref("")
const errorMessage = ref("")
const isLoading = ref(false)
const isEdit = ref(false)
const editingRecordId = ref<number | null>(null)

const formData = ref({
  student_id: "",
  student_name: "",
  subject: "",
  semester: 1,
  academic_year: "2025/2026",
  score: null as number | null
})

onMounted(() => {
  loadStudents()
  loadGradesList()
})

async function loadStudents() {
  try {
    const res = await api.get("/api/students?limit=200")
    if (res.success && res.data) {
      studentsList.value = res.data.students || []
    }
  } catch (error) {
    console.error("Gagal memuat daftar siswa", error)
  }
}

async function loadGradesList() {
  try {
    let url = "/api/grades"
    if (filterStudentID.value) {
      url += `?student_id=${filterStudentID.value}`
    }
    const res = await api.get(url)
    if (res.success && res.data) {
      gradesList.value = res.data || []
    }
  } catch (error) {
    console.error("Gagal memuat rekap nilai", error)
  }
}

async function handleSubmit() {
  if (!formData.value.subject || formData.value.score === null || !formData.value.academic_year || (!isEdit.value && !formData.value.student_id)) {
    errorMessage.value = "Semua kolom wajib diisi"
    return
  }

  if (formData.value.score < 0 || formData.value.score > 100) {
    errorMessage.value = "Nilai harus berada di antara 0 dan 100"
    return
  }

  isLoading.value = true
  errorMessage.value = ""

  try {
    let res
    if (isEdit.value && editingRecordId.value) {
      res = await api.put(`/api/grades/${editingRecordId.value}`, {
        score: formData.value.score,
        subject: formData.value.subject,
        semester: formData.value.semester,
        academic_year: formData.value.academic_year
      })
    } else {
      res = await api.post("/api/grades", {
        student_id: formData.value.student_id,
        subject: formData.value.subject,
        score: formData.value.score,
        semester: formData.value.semester,
        academic_year: formData.value.academic_year
      })
    }

    if (res.success) {
      showToast(isEdit.value ? "Catatan nilai berhasil diubah" : "Nilai berhasil disimpan")
      resetForm()
      loadGradesList()
    } else {
      errorMessage.value = res.message || "Gagal menyimpan data nilai"
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
    student_name: `${record.student?.nis} - ${record.student?.name}`,
    subject: record.subject,
    semester: record.semester,
    academic_year: record.academic_year,
    score: record.score
  }
  errorMessage.value = ""
}

function resetForm() {
  isEdit.value = false
  editingRecordId.value = null
  formData.value = {
    student_id: "",
    student_name: "",
    subject: "",
    semester: 1,
    academic_year: "2025/2026",
    score: null
  }
  errorMessage.value = ""
}

function showToast(msg: string) {
  toastMessage.value = msg
  setTimeout(() => {
    toastMessage.value = ""
  }, 3000)
}
</script>
