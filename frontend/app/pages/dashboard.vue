<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8">
        <div>
          <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
            Ikhtisar Sistem
          </h2>
          <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
            {{ user?.role === 'siswa' || user?.role === 'siswi' ? 'Dashboard Siswa' : 'Dashboard Utama' }}
          </h1>
        </div>

        <div v-if="user?.role === 'siswa' || user?.role === 'siswi'" class="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <div class="lg:col-span-3 bg-[color:var(--color-surface)] border border-[color:var(--color-border)] border-l-4 border-l-[color:var(--color-accent)] p-6 space-y-6 shadow-[--shadow-sm]">
            <div class="flex items-center space-x-3 border-b border-[color:var(--color-border)] pb-4">
              <UserIcon class="w-5 h-5 text-[color:var(--color-accent)] shrink-0" />
              <h3 class="text-lg font-bold text-[color:var(--color-heading)]">Profil Siswa</h3>
            </div>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm">
              <div class="space-y-1">
                <span class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider block font-bold">Nama Lengkap</span>
                <span class="text-md font-semibold text-[color:var(--color-heading)] block">{{ studentData?.name || '-' }}</span>
              </div>
              <div class="space-y-1">
                <span class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider block font-bold">NISN / Student ID</span>
                <span class="text-md font-mono text-[color:var(--color-accent)] block">{{ studentData?.nis || '-' }}</span>
              </div>
              <div class="space-y-1">
                <span class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider block font-bold">Kelas</span>
                <span class="text-md font-semibold text-[color:var(--color-heading)] block">{{ studentData?.class || '-' }}</span>
              </div>
              <div class="space-y-1">
                <span class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider block font-bold">Jenis Kelamin</span>
                <span class="text-md text-[color:var(--color-text)] block">{{ studentData?.gender || '-' }}</span>
              </div>
              <div class="space-y-1">
                <span class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider block font-bold">Telepon</span>
                <span class="text-md font-mono text-[color:var(--color-text)] block">{{ studentData?.phone || '-' }}</span>
              </div>
              <div class="space-y-1 md:col-span-2">
                <span class="text-xs text-[color:var(--color-muted)] uppercase tracking-wider block font-bold">Alamat Rumah</span>
                <span class="text-md text-[color:var(--color-text)] block">{{ studentData?.address || '-' }}</span>
              </div>
            </div>
          </div>

          <div class="lg:col-span-3 bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6 shadow-[--shadow-sm]">
            <div class="flex items-center justify-between border-b border-[color:var(--color-border)] pb-4">
              <div class="flex items-center space-x-3">
                <AwardIcon class="w-5 h-5 text-[color:var(--color-accent)] shrink-0" />
                <h3 class="text-lg font-bold text-[color:var(--color-heading)]">Riwayat Nilai Akademik</h3>
              </div>
            </div>

            <div class="overflow-x-auto">
              <table class="w-full text-left border-collapse">
                <thead>
                  <tr class="border-b border-[color:var(--color-border)]">
                    <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Mata Pelajaran</th>
                    <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Semester</th>
                    <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Tahun Ajaran</th>
                    <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Skor & Progress</th>
                  </tr>
                </thead>
                <tbody>
                  <tr
                    v-for="grade in studentGrades"
                    :key="grade.id"
                    class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                  >
                    <td class="py-3.5 px-4 text-sm font-semibold text-[color:var(--color-heading)]">{{ grade.subject }}</td>
                    <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">Semester {{ grade.semester }}</td>
                    <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] font-mono">{{ grade.academic_year }}</td>
                    <td class="py-3.5 px-4 text-sm w-[40%]">
                      <div class="flex items-center space-x-3">
                        <span class="font-bold text-[color:var(--color-accent)] font-mono w-10 text-right">{{ grade.score.toFixed(1) }}</span>
                        <div class="flex-1 h-2 bg-[color:var(--color-bg)] border border-[color:var(--color-border)] rounded-full overflow-hidden">
                          <div :style="{ width: grade.score + '%' }" class="h-full bg-[color:var(--color-accent)] transition duration-500"></div>
                        </div>
                      </div>
                    </td>
                  </tr>
                  <tr v-if="studentGrades.length === 0">
                    <td colspan="4" class="py-12 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                      Belum ada riwayat catatan nilai terkumpul
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <div v-else class="space-y-6 sm:space-y-8">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="card flex flex-col justify-between min-h-[140px] shadow-[--shadow-sm]">
              <div>
                <span class="text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">
                  Total Siswa
                </span>
                <div class="flex items-center space-x-2 mt-2">
                  <UsersIcon class="w-5 h-5 text-[color:var(--color-accent)] shrink-0" />
                  <span class="text-3xl font-bold text-[color:var(--color-heading)]">
                    {{ stats.total_students }}
                  </span>
                </div>
              </div>
              <div class="text-xs text-[color:var(--color-muted)]">
                Siswa terdaftar aktif di sistem
              </div>
            </div>

            <div class="card flex flex-col justify-between min-h-[140px] shadow-[--shadow-sm]">
              <div>
                <span class="text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">
                  Kehadiran Bulan Ini
                </span>
                <div class="flex items-center space-x-2 mt-2">
                  <CalendarCheckIcon class="w-5 h-5 text-[color:var(--color-accent)] shrink-0" />
                  <span class="text-3xl font-bold text-[color:var(--color-heading)]">
                    {{ stats.attendance_rate.toFixed(1) }}%
                  </span>
                </div>
              </div>
              <div class="text-xs text-[color:var(--color-muted)]">
                Rata-rata persentase kehadiran hadir
              </div>
            </div>

            <div class="card flex flex-col justify-between min-h-[140px] shadow-[--shadow-sm]">
              <div>
                <span class="text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">
                  Rata-Rata Nilai
                </span>
                <div class="flex items-center space-x-2 mt-2">
                  <AwardIcon class="w-5 h-5 text-[color:var(--color-accent)] shrink-0" />
                  <span class="text-3xl font-bold text-[color:var(--color-heading)]">
                    {{ stats.average_score.toFixed(1) }}
                  </span>
                </div>
              </div>
              <div class="text-xs text-[color:var(--color-muted)]">
                Rata-rata skor seluruh bidang studi
              </div>
            </div>
          </div>

          <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6">
            <div class="flex items-center justify-between mb-6">
              <h3 class="text-lg font-bold text-[color:var(--color-heading)] tracking-wide">
                Presensi Hari Ini
              </h3>
              <NuxtLink to="/attendance" class="text-xs uppercase tracking-wider text-[color:var(--color-accent)] hover:opacity-85 font-semibold">
                Lihat Semua Presensi
              </NuxtLink>
            </div>

            <div class="overflow-x-auto">
              <table class="w-full text-left border-collapse min-w-[600px]">
                <thead>
                  <tr class="border-b border-[color:var(--color-border)]">
                    <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-[18%]">NISN</th>
                    <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-[25%]">Nama</th>
                    <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-[15%]">Kelas</th>
                    <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-[20%]">Status</th>
                    <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-[22%]">Keterangan</th>
                  </tr>
                </thead>
              </table>

              <div
                ref="scrollContainer"
                class="max-h-[220px] overflow-y-auto w-full relative no-scrollbar min-w-[600px]"
                @mouseenter="isHovered = true"
                @mouseleave="isHovered = false"
                @touchstart="isHovered = true"
                @touchend="isHovered = false"
                @scroll="handleScroll"
              >
                <table class="w-full text-left border-collapse">
                  <tbody>
                    <tr
                      v-for="att in todayAttendancesList"
                      :key="'1-' + att.id"
                      class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                    >
                      <td class="py-3.5 px-4 text-sm font-mono text-[color:var(--color-accent)] w-[18%]">{{ att.student?.nis || '-' }}</td>
                      <td class="py-3.5 px-4 text-sm font-semibold text-[color:var(--color-heading)] w-[25%]">
                        <NuxtLink :to="`/students/${att.student?.id}`" class="hover:text-[color:var(--color-accent)] transition duration-100">
                          {{ att.student?.name || '-' }}
                        </NuxtLink>
                      </td>
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] w-[15%]">{{ att.student?.class || '-' }}</td>
                      <td class="py-3.5 px-4 text-sm font-bold uppercase tracking-wider w-[20%]">
                        <span :class="getStatusClass(att.status)">
                          {{ att.status }}
                        </span>
                      </td>
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] w-[22%] truncate" :title="att.note">{{ att.note || '-' }}</td>
                    </tr>

                    <template v-if="todayAttendancesList.length > 3">
                      <tr
                        v-for="att in todayAttendancesList"
                        :key="'2-' + att.id"
                        class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                      >
                        <td class="py-3.5 px-4 text-sm font-mono text-[color:var(--color-accent)] w-[18%]">{{ att.student?.nis || '-' }}</td>
                        <td class="py-3.5 px-4 text-sm font-semibold text-[color:var(--color-heading)] w-[25%]">
                          <NuxtLink :to="`/students/${att.student?.id}`" class="hover:text-[color:var(--color-accent)] transition duration-100">
                            {{ att.student?.name || '-' }}
                          </NuxtLink>
                        </td>
                        <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] w-[15%]">{{ att.student?.class || '-' }}</td>
                        <td class="py-3.5 px-4 text-sm font-bold uppercase tracking-wider w-[20%]">
                          <span :class="getStatusClass(att.status)">
                            {{ att.status }}
                          </span>
                        </td>
                        <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] w-[22%] truncate" :title="att.note">{{ att.note || '-' }}</td>
                      </tr>
                    </template>

                    <tr v-if="todayAttendancesList.length === 0">
                      <td colspan="5" class="py-8 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                        Belum ada data presensi untuk hari ini
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
import { ref, onMounted, onUnmounted, nextTick, computed } from "vue"
import {
  Users as UsersIcon,
  CalendarCheck as CalendarCheckIcon,
  Award as AwardIcon,
  User as UserIcon,
  CheckCircle as CheckCircleIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()
const user = useState<any>("currentUser", () => null)

const stats = ref<{
  total_students: number
  attendance_rate: number
  average_score: number
  recent_students: Array<{
    id: string
    nis: string
    name: string
    class: string
    gender: string
    phone: string
  }>
}>({
  total_students: 0,
  attendance_rate: 0.0,
  average_score: 0.0,
  recent_students: []
})

const studentData = ref<any>(null)
const todayAttendance = ref<any>(null)
const studentGrades = ref<any[]>([])
const todayAttendancesList = ref<any[]>([])

const checkingIn = ref(false)
const selectedStatus = ref("hadir")
const checkInNote = ref("")

const scrollContainer = ref<HTMLElement | null>(null)
const isHovered = ref(false)
let currentScroll = 0
let animationFrameId: number | null = null

const hasCheckedInToday = computed(() => !!todayAttendance.value)

function handleScroll() {
  const el = scrollContainer.value
  if (el && isHovered.value) {
    currentScroll = el.scrollTop
  }
}

function initAutoScroll() {
  const el = scrollContainer.value
  if (!el) {
    setTimeout(initAutoScroll, 100)
    return
  }

  if (el.scrollHeight <= el.clientHeight) {
    setTimeout(initAutoScroll, 100)
    return
  }

  if (animationFrameId) {
    cancelAnimationFrame(animationFrameId)
  }

  const scrollSpeed = 0.4

  const scroll = () => {
    if (isHovered.value) {
      currentScroll = el.scrollTop
    } else if (todayAttendancesList.value.length > 3) {
      currentScroll += scrollSpeed
      el.scrollTop = currentScroll
      const halfHeight = el.scrollHeight / 2
      if (currentScroll >= halfHeight) {
        currentScroll = 0
        el.scrollTop = 0
      }
    }
    animationFrameId = requestAnimationFrame(scroll)
  }
  animationFrameId = requestAnimationFrame(scroll)
}

async function loadStudentDashboard() {
  if (!user.value?.student_id) return

  try {
    const sRes: any = await api.get(`/api/students/${user.value.student_id}`)
    if (sRes.success) {
      studentData.value = sRes.data
    }
  } catch (e) {
    console.error(e)
  }

  try {
    const gRes: any = await api.get("/api/grades")
    if (gRes.success && gRes.data) {
      studentGrades.value = gRes.data
    }
  } catch (e) {
    console.error(e)
  }

  try {
    const aRes: any = await api.get("/api/attendances")
    if (aRes.success && aRes.data) {
      const todayStr = new Date().toISOString().split("T")[0]
      const found = aRes.data.find((a: any) => a.date.startsWith(todayStr))
      if (found) {
        todayAttendance.value = found
      }
    }
  } catch (e) {
    console.error(e)
  }
}

async function submitCheckIn() {
  if (!user.value?.student_id) return

  checkingIn.value = true
  const todayStr = new Date().toISOString().split("T")[0]

  try {
    const res: any = await api.post("/api/attendances", {
      student_id: user.value.student_id,
      date: todayStr,
      status: selectedStatus.value,
      note: checkInNote.value
    })
    if (res.success && res.data) {
      todayAttendance.value = res.data
      checkInNote.value = ""
    }
  } catch (e) {
    console.error(e)
  } finally {
    checkingIn.value = false
  }
}

onMounted(async () => {
  const cached = localStorage.getItem("user")
  if (cached && !user.value) {
    user.value = JSON.parse(cached)
  }

  if (user.value?.role === "siswa" || user.value?.role === "siswi") {
    loadStudentDashboard()
  } else {
    try {
      const res: any = await api.get("/api/dashboard/stats")
      if (res.success && res.data) {
        stats.value = res.data
      }
    } catch (error) {
      console.error("Gagal memuat statistik dashboard", error)
    }

    try {
      const todayStr = new Date().toISOString().split("T")[0]
      const res: any = await api.get(`/api/attendances?date=${todayStr}`)
      if (res.success && res.data) {
        todayAttendancesList.value = res.data || []
        if (todayAttendancesList.value.length > 3) {
          nextTick(() => {
            initAutoScroll()
          })
        }
      }
    } catch (error) {
      console.error("Gagal memuat absensi hari ini", error)
    }
  }
})

onUnmounted(() => {
  if (animationFrameId) {
    cancelAnimationFrame(animationFrameId)
  }
})

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

<style scoped>
.no-scrollbar::-webkit-scrollbar {
  display: none;
}
.no-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
</style>
