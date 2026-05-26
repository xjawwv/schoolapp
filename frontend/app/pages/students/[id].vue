<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8" v-if="student">
        <div>
          <NuxtLink to="/students" class="inline-flex items-center space-x-2 text-xs uppercase tracking-wider text-[color:var(--color-accent)] hover:opacity-85 font-semibold mb-4">
            <ChevronLeftIcon class="w-4 h-4" />
            <span>Kembali ke Daftar</span>
          </NuxtLink>
          <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
            <div>
              <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
                Profil Akademik
              </h2>
              <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
                {{ student.name }}
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
                  {{ student.name.charAt(0) }}
                </div>
                <div class="flex flex-col">
                  <span class="text-sm font-semibold text-[color:var(--color-heading)]">NIS: {{ student.nis }}</span>
                  <span class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider">{{ student.class }}</span>
                </div>
              </div>

              <div class="space-y-4">
                <div>
                  <span class="block text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold mb-1">Gender</span>
                  <span class="text-sm text-[color:var(--color-text)] font-semibold">{{ student.gender }}</span>
                </div>
                <div>
                  <span class="block text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold mb-1">Nomor Telepon</span>
                  <span class="text-sm text-[color:var(--color-text)] font-mono font-semibold">{{ student.phone || '-' }}</span>
                </div>
                <div>
                  <span class="block text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold mb-1">Alamat Lengkap</span>
                  <span class="text-sm text-[color:var(--color-text)] leading-relaxed font-semibold">{{ student.address || '-' }}</span>
                </div>
                <div>
                  <span class="block text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold mb-1">Terdaftar Sejak</span>
                  <span class="text-sm text-[color:var(--color-text)] font-mono font-semibold">{{ formatDate(student.created_at) }}</span>
                </div>
              </div>
            </div>

            <div class="card space-y-4 shadow-[--shadow-sm]">
              <h3 class="text-xs uppercase tracking-widest text-[color:var(--color-accent)] font-bold">
                Ringkasan Kehadiran
              </h3>
              <div class="grid grid-cols-2 gap-4">
                <div class="bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-3 text-center">
                  <span class="text-xs text-[color:var(--color-muted)] block uppercase font-semibold">Hadir</span>
                  <span class="text-2xl font-bold text-[color:var(--color-heading)] mt-1 block">{{ attendanceStats.hadir }}</span>
                </div>
                <div class="bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-3 text-center">
                  <span class="text-xs text-[color:var(--color-muted)] block uppercase font-semibold">Sakit</span>
                  <span class="text-2xl font-bold text-[color:var(--color-heading)] mt-1 block">{{ attendanceStats.sakit }}</span>
                </div>
                <div class="bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-3 text-center">
                  <span class="text-xs text-[color:var(--color-muted)] block uppercase font-semibold">Izin</span>
                  <span class="text-2xl font-bold text-[color:var(--color-heading)] mt-1 block">{{ attendanceStats.izin }}</span>
                </div>
                <div class="bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-3 text-center">
                  <span class="text-xs text-[color:var(--color-muted)] block uppercase font-semibold">Alpha</span>
                  <span class="text-2xl font-bold text-[color:var(--color-heading)] mt-1 block text-[color:var(--color-error)]">{{ attendanceStats.alpha }}</span>
                </div>
              </div>
            </div>

            <div class="card space-y-4 shadow-[--shadow-sm]" v-if="currentUser?.role === 'admin'">
              <h3 class="text-xs uppercase tracking-widest text-[color:var(--color-accent)] font-bold">
                Kredensial Akses
              </h3>
              <div class="space-y-3 bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-4 text-xs">
                <div>
                  <span class="text-[color:var(--color-muted)] block uppercase font-bold mb-1">Username / NISN</span>
                  <span class="font-mono text-sm text-[color:var(--color-heading)] font-semibold">{{ student.nis }}</span>
                </div>
                <div>
                  <span class="text-[color:var(--color-muted)] block uppercase font-bold mb-1">Kata Sandi Bawaan</span>
                  <span class="font-mono text-sm text-[color:var(--color-heading)] font-semibold">1</span>
                </div>
                <div class="pt-2 border-t border-[color:var(--color-border)] text-[color:var(--color-muted)] leading-relaxed">
                  Akun login siswa disinkronkan otomatis menggunakan NISN dan sandi default.
                </div>
              </div>
            </div>
          </div>

          <div class="lg:col-span-8 space-y-6">
            <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6">
              <div class="flex items-center space-x-6 border-b border-[color:var(--color-border)] pb-4 mb-6">
                <button
                  @click="activeTab = 'grades'"
                  class="text-xs uppercase tracking-widest font-bold pb-2 relative transition duration-150 cursor-pointer"
                  :class="activeTab === 'grades' ? 'text-[color:var(--color-accent)]' : 'text-[color:var(--color-muted)] hover:text-[color:var(--color-text)]'"
                >
                  Daftar Nilai
                  <span v-if="activeTab === 'grades'" class="absolute bottom-[-17px] left-0 right-0 h-[2px] bg-[color:var(--color-accent)]"></span>
                </button>
                <button
                  @click="activeTab = 'attendance'"
                  class="text-xs uppercase tracking-widest font-bold pb-2 relative transition duration-150 cursor-pointer"
                  :class="activeTab === 'attendance' ? 'text-[color:var(--color-accent)]' : 'text-[color:var(--color-muted)] hover:text-[color:var(--color-text)]'"
                >
                  Riwayat Absensi
                  <span v-if="activeTab === 'attendance'" class="absolute bottom-[-17px] left-0 right-0 h-[2px] bg-[color:var(--color-accent)]"></span>
                </button>
              </div>

              <div v-if="activeTab === 'grades'">
                <div class="overflow-x-auto">
                  <table class="w-full text-left border-collapse">
                    <thead>
                      <tr class="border-b border-[color:var(--color-border)]">
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Subject</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Semester</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Academic Year</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold text-right">Score</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold text-right" v-if="currentUser?.role !== 'siswa' && currentUser?.role !== 'siswi'">Actions</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr
                        v-for="grade in grades"
                        :key="grade.id"
                        class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                      >
                        <td class="py-3.5 px-4 text-sm font-semibold text-[color:var(--color-heading)]">{{ grade.subject }}</td>
                        <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">Semester {{ grade.semester }}</td>
                        <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] font-mono">{{ grade.academic_year }}</td>
                        <td class="py-3.5 px-4 text-sm text-right font-bold text-[color:var(--color-accent)] font-mono">{{ grade.score.toFixed(1) }}</td>
                        <td class="py-3.5 px-4 text-sm text-right" v-if="currentUser?.role !== 'siswa' && currentUser?.role !== 'siswi'">
                          <button @click="openGradeModal(grade)" class="text-xs uppercase tracking-wider text-[color:var(--color-accent)] hover:opacity-80 transition duration-100 cursor-pointer font-semibold">
                            Ubah
                          </button>
                        </td>
                      </tr>
                      <tr v-if="grades.length === 0">
                        <td :colspan="currentUser?.role !== 'siswa' && currentUser?.role !== 'siswi' ? 5 : 4" class="py-12 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                          No grade records found for this student
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>

              <div v-if="activeTab === 'attendance'">
                <div class="overflow-x-auto">
                  <table class="w-full text-left border-collapse">
                    <thead>
                      <tr class="border-b border-[color:var(--color-border)]">
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Date</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Status</th>
                        <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Notes</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr
                        v-for="att in attendances"
                        :key="att.id"
                        class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                      >
                        <td class="py-3.5 px-4 text-sm text-[color:var(--color-heading)] font-mono">{{ formatDateOnly(att.date) }}</td>
                        <td class="py-3.5 px-4 text-sm font-bold uppercase tracking-wider">
                          <span :class="getStatusClass(att.status)">
                            {{ att.status }}
                          </span>
                        </td>
                        <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">{{ att.note || '-' }}</td>
                      </tr>
                      <tr v-if="attendances.length === 0">
                        <td colspan="3" class="py-12 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                          No attendance records found for this student
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>

      <StudentForm
        :show="showForm"
        :student="student"
        @close="closeForm"
        @saved="onStudentSaved"
      />

      <div v-if="showGradeModal" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
        <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 w-full max-w-md space-y-6 shadow-xl relative">
          <div class="flex justify-between items-center border-b border-[color:var(--color-border)] pb-4">
            <div>
              <h3 class="text-lg font-bold text-[color:var(--color-heading)] tracking-wide">Ubah Catatan Nilai</h3>
              <p class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider">Perbarui evaluasi mata pelajaran siswa</p>
            </div>
            <button @click="closeGradeModal" class="text-[color:var(--color-muted)] hover:text-[color:var(--color-text)] cursor-pointer">
              <XIcon class="w-5 h-5" />
            </button>
          </div>

          <div v-if="gradeError" class="bg-[color:var(--color-bg)] border border-[color:var(--color-error)] p-3 text-sm text-[color:var(--color-error)] font-medium flex items-center space-x-2">
            <AlertCircleIcon class="w-4 h-4 shrink-0" />
            <span>{{ gradeError }}</span>
          </div>

          <form @submit.prevent="handleGradeSubmit" class="space-y-4">
            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Mata Pelajaran</label>
              <select v-model="gradeForm.subject" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
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
                <select v-model.number="gradeForm.semester" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                  <option :value="1">Ganjil (1)</option>
                  <option :value="2">Genap (2)</option>
                </select>
              </div>
              <div class="space-y-1.5">
                <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Tahun Ajaran</label>
                <select v-model="gradeForm.academic_year" class="input w-full bg-[color:var(--color-bg)] select-arrow" required>
                  <option value="2024/2025">2024/2025</option>
                  <option value="2025/2026">2025/2026</option>
                  <option value="2026/2027">2026/2027</option>
                </select>
              </div>
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-bold">Nilai Evaluasi (0 - 100)</label>
              <input
                v-model.number="gradeForm.score"
                type="number"
                step="0.1"
                min="0"
                max="100"
                class="input w-full font-mono text-[color:var(--color-accent)] font-bold"
                required
                placeholder="Contoh: 85.5"
              />
            </div>

            <div class="flex items-center space-x-2 pt-2 border-t border-[color:var(--color-border)]">
              <button type="submit" :disabled="gradeLoading" class="button-primary text-xs font-semibold uppercase tracking-wider flex-1 flex items-center justify-center py-3">
                <span v-if="gradeLoading" class="animate-spin rounded-full h-3 w-3 border-2 border-[color:var(--color-accent-fg)] border-t-transparent mr-2"></span>
                <span>Simpan Nilai</span>
              </button>
              <button type="button" @click="closeGradeModal" class="button-ghost text-xs font-semibold uppercase tracking-wider py-3">
                Batal
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
import { useRoute } from "vue-router"
import {
  ChevronLeft as ChevronLeftIcon,
  CheckCircle as CheckCircleIcon,
  X as XIcon,
  AlertCircle as AlertCircleIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"
import StudentForm from "~/components/StudentForm.vue"

definePageMeta({
  middleware: ["auth"]
})

const route = useRoute()
const api = useApi()

const currentUser = useState<any>("currentUser", () => null)
const student = ref<any | null>(null)
const grades = ref<any[]>([])
const attendances = ref<any[]>([])
const activeTab = ref("grades")
const showForm = ref(false)
const toastMessage = ref("")

const showGradeModal = ref(false)
const gradeLoading = ref(false)
const gradeError = ref("")
const selectedGradeId = ref<number | null>(null)

const gradeForm = ref({
  subject: "",
  semester: 1,
  academic_year: "2025/2026",
  score: null as number | null
})

const studentId = computed(() => route.params.id as string)

const attendanceStats = computed(() => {
  const stats = { hadir: 0, sakit: 0, izin: 0, alpha: 0 }
  attendances.value.forEach((att) => {
    const s = att.status as keyof typeof stats
    if (stats[s] !== undefined) {
      stats[s]++
    }
  })
  return stats
})

onMounted(() => {
  const cached = localStorage.getItem("user")
  if (cached && !currentUser.value) {
    currentUser.value = JSON.parse(cached)
  }
  loadAllData()
})

async function loadAllData() {
  await fetchStudentProfile()
  await fetchStudentGrades()
  await fetchStudentAttendance()
}

async function fetchStudentProfile() {
  try {
    const res: any = await api.get(`/api/students/${studentId.value}`)
    if (res.success && res.data) {
      student.value = res.data
    }
  } catch (error) {
    console.error("Gagal memuat profil siswa", error)
  }
}

async function fetchStudentGrades() {
  try {
    const res: any = await api.get(`/api/grades?student_id=${studentId.value}`)
    if (res.success && res.data) {
      grades.value = res.data || []
    }
  } catch (error) {
    console.error("Gagal memuat daftar nilai", error)
  }
}

async function fetchStudentAttendance() {
  try {
    const res: any = await api.get(`/api/attendances?student_id=${studentId.value}`)
    if (res.success && res.data) {
      attendances.value = res.data || []
    }
  } catch (error) {
    console.error("Gagal memuat riwayat kehadiran", error)
  }
}

function openEditModal() {
  showForm.value = true
}

function closeForm() {
  showForm.value = false
}

function onStudentSaved(savedStudent: any) {
  student.value = savedStudent
  toastMessage.value = "Profil siswa berhasil diperbarui"
  setTimeout(() => {
    toastMessage.value = ""
  }, 3000)
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

function formatDateOnly(dateStr: string): string {
  if (!dateStr) return "-"
  const d = new Date(dateStr)
  return d.toISOString().split("T")[0] || "-"
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

function openGradeModal(grade: any) {
  selectedGradeId.value = grade.id
  gradeForm.value = {
    subject: grade.subject,
    semester: grade.semester,
    academic_year: grade.academic_year,
    score: grade.score
  }
  gradeError.value = ""
  showGradeModal.value = true
}

function closeGradeModal() {
  showGradeModal.value = false
  selectedGradeId.value = null
  gradeError.value = ""
}

async function handleGradeSubmit() {
  if (!gradeForm.value.subject || gradeForm.value.score === null || !gradeForm.value.academic_year) {
    gradeError.value = "Semua kolom wajib diisi"
    return
  }

  if (gradeForm.value.score < 0 || gradeForm.value.score > 100) {
    gradeError.value = "Nilai harus berada di antara 0 dan 100"
    return
  }

  gradeLoading.value = true
  gradeError.value = ""

  try {
    const res: any = await api.put(`/api/grades/${selectedGradeId.value}`, {
      score: gradeForm.value.score,
      subject: gradeForm.value.subject,
      semester: gradeForm.value.semester,
      academic_year: gradeForm.value.academic_year
    })

    if (res.success) {
      toastMessage.value = "Catatan nilai berhasil diperbarui"
      closeGradeModal()
      await fetchStudentGrades()
      setTimeout(() => {
        toastMessage.value = ""
      }, 3000)
    } else {
      gradeError.value = res.message || "Gagal memperbarui data nilai"
    }
  } catch (error: any) {
    if (error.response && error.response._data) {
      gradeError.value = error.response._data.message || "Gagal memproses data"
    } else {
      gradeError.value = "Koneksi ke server terputus"
    }
  } finally {
    gradeLoading.value = false
  }
}
</script>
