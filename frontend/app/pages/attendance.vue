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

                <div v-if="formData.status === 'hadir' && (currentUser?.role === 'siswa' || currentUser?.role === 'siswi')" class="bg-[color:var(--color-bg)] border border-[color:var(--color-border)] p-4 space-y-3 text-xs leading-relaxed">
                  <div class="font-bold text-[color:var(--color-heading)] uppercase tracking-wider text-[10px] flex items-center space-x-1.5">
                    <MapPinIcon class="w-3.5 h-3.5 text-[color:var(--color-accent)]" />
                    <span>Verifikasi Lokasi Anti-Fake GPS</span>
                  </div>

                  <div v-if="gpsStatus === 'detecting'" class="flex items-center space-x-2 text-[color:var(--color-muted)]">
                    <span class="animate-spin rounded-full h-3.5 w-3.5 border-2 border-[color:var(--color-accent)] border-t-transparent"></span>
                    <span>Mendeteksi koordinat satelit GPS...</span>
                  </div>

                  <div v-else-if="gpsStatus === 'error'" class="text-[color:var(--color-error)] flex items-start space-x-1.5 font-medium">
                    <AlertCircleIcon class="w-4 h-4 shrink-0 mt-0.5" />
                    <div class="space-y-1">
                      <p>{{ gpsErrorMsg }}</p>
                      <button type="button" @click="getGPSLocation" class="text-[color:var(--color-accent)] hover:underline font-semibold cursor-pointer">
                        Coba Lagi
                      </button>
                    </div>
                  </div>

                  <div v-else-if="gpsStatus === 'success'" class="space-y-2">
                    <div class="flex items-start space-x-1.5 text-[color:var(--color-text)]">
                      <div class="space-y-0.5 font-mono">
                        <p>Latitude: {{ studentLatitude?.toFixed(6) }}</p>
                        <p>Longitude: {{ studentLongitude?.toFixed(6) }}</p>
                      </div>
                    </div>

                    <div v-if="distanceToSchool !== null && distanceToSchool <= schoolMaxDistance" class="text-[color:var(--color-success)] font-medium flex items-center space-x-1.5 bg-emerald-500/5 border border-emerald-500/20 p-2">
                      <CheckCircleIcon class="w-4 h-4 shrink-0" />
                      <span>Lokasi Terverifikasi: {{ distanceToSchool }}m dari sekolah</span>
                    </div>

                    <div v-else class="text-[color:var(--color-error)] font-medium flex items-start space-x-1.5 bg-red-500/5 border border-red-500/20 p-2">
                      <AlertCircleIcon class="w-4 h-4 shrink-0 mt-0.5" />
                      <div class="space-y-1">
                        <p>Jarak Anda terlalu jauh ({{ distanceToSchool }}m dari sekolah).</p>
                        <p class="text-[10px] text-[color:var(--color-muted)] leading-relaxed">Anda harus berada di dalam radius maksimal {{ schoolMaxDistance }}m dari {{ siteName }} untuk melakukan absensi.</p>
                      </div>
                    </div>
                  </div>
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
                  <button type="submit" :disabled="isLoading || (formData.status === 'hadir' && (currentUser?.role === 'siswa' || currentUser?.role === 'siswi') && (gpsStatus !== 'success' || (distanceToSchool !== null && distanceToSchool > schoolMaxDistance)))" class="button-primary text-xs font-semibold uppercase tracking-wider flex-1 flex items-center justify-center py-3">
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

                  <div class="flex items-center space-x-2 w-full md:w-auto" v-if="currentUser?.role === 'siswa' || currentUser?.role === 'siswi'">
                    <span class="text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-semibold whitespace-nowrap">Filter Periode:</span>
                    <select v-model="selectedPeriod" class="input py-1 text-xs bg-[color:var(--color-bg)] select-arrow w-full md:w-40">
                      <option value="minggu">Minggu Ini</option>
                      <option value="bulan">Bulan Ini</option>
                      <option value="semester">Semester Ini</option>
                      <option value="semua">Semua Riwayat</option>
                    </select>
                  </div>
                  <div class="flex items-center space-x-2 w-full md:w-auto" v-else>
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
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Tanggal</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Waktu</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Status</th>
                      <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Notes</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr
                      v-for="att in filteredAttendances"
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
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] font-mono">
                        <div v-if="att.timestamp" class="font-semibold text-[color:var(--color-heading)]">
                          {{ att.timestamp.split(' ')[1] || att.timestamp }}
                        </div>
                        <div v-else class="text-[color:var(--color-muted)]">-</div>
                        <div v-if="att.latitude" class="text-[9px] text-[color:var(--color-accent)] font-mono scale-95 origin-left mt-0.5">
                          GPS: {{ att.latitude.toFixed(4) }}, {{ att.longitude.toFixed(4) }}
                        </div>
                      </td>
                      <td class="py-3.5 px-4 text-sm font-bold uppercase tracking-wider">
                        <span :class="getStatusClass(att.status)">
                          {{ att.status }}
                        </span>
                      </td>
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">{{ att.note || '-' }}</td>
                    </tr>
                    <tr v-if="filteredAttendances.length === 0">
                      <td colspan="7" class="py-12 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                        Tidak ada catatan absensi dalam periode ini
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
import { ref, onMounted, watch, computed } from "vue"
import {
  AlertCircle as AlertCircleIcon,
  CheckCircle as CheckCircleIcon,
  Search as SearchIcon,
  MapPin as MapPinIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()
const currentUser = useState<any>("currentUser", () => null)
const siteName = useState("siteName", () => "SMA N 1 METRO")

const schoolLatitude = ref(-6.1822)
const schoolLongitude = ref(106.2736)
const schoolMaxDistance = ref(250)
const attendanceStartTime = ref("07:00")
const attendanceEndTime = ref("17:00")
const enableAntiFakeGPS = ref("true")

const studentsList = ref<any[]>([])
const attendancesList = ref<any[]>([])
const filterDate = ref(new Date().toISOString().split("T")[0])
const searchQuery = ref("")
const filterClass = ref("")
const selectedPeriod = ref("bulan")

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

const gpsStatus = ref("")
const gpsErrorMsg = ref("")
const studentLatitude = ref<number | null>(null)
const studentLongitude = ref<number | null>(null)
const distanceToSchool = ref<number | null>(null)

function calculateDistanceClient(lat1: number, lon1: number, lat2: number, lon2: number): number {
  const R = 6371
  const dLat = (lat2 - lat1) * Math.PI / 180
  const dLon = (lon2 - lon1) * Math.PI / 180
  const a = Math.sin(dLat / 2) * Math.sin(dLat / 2) +
            Math.cos(lat1 * Math.PI / 180) * Math.cos(lat2 * Math.PI / 180) *
            Math.sin(dLon / 2) * Math.sin(dLon / 2)
  const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a))
  return R * c * 1000
}

function getGPSLocation() {
  gpsStatus.value = "detecting"
  gpsErrorMsg.value = ""

  if (!navigator.geolocation) {
    gpsStatus.value = "error"
    gpsErrorMsg.value = "Geolocation tidak didukung oleh browser Anda"
    return
  }

  const options = {
    enableHighAccuracy: true,
    timeout: 10000,
    maximumAge: 0
  }

  navigator.geolocation.getCurrentPosition(
    (position) => {
      if (enableAntiFakeGPS.value === "true") {
        const isMocked = (position as any).mocked === true ||
                         (position.coords as any).mocked === true ||
                         (position.coords as any).isFromMockProvider === true ||
                         position.coords.accuracy === 0
        if (isMocked) {
          gpsStatus.value = "error"
          gpsErrorMsg.value = "Absensi ditolak: Terdeteksi penggunaan aplikasi Fake GPS / Pemalsu Lokasi"
          studentLatitude.value = null
          studentLongitude.value = null
          distanceToSchool.value = null
          return
        }
      }

      studentLatitude.value = position.coords.latitude
      studentLongitude.value = position.coords.longitude
      gpsStatus.value = "success"

      const dist = calculateDistanceClient(position.coords.latitude, position.coords.longitude, schoolLatitude.value, schoolLongitude.value)
      distanceToSchool.value = Math.round(dist)
    },
    (error) => {
      gpsStatus.value = "error"
      switch (error.code) {
        case error.PERMISSION_DENIED:
          gpsErrorMsg.value = "Akses lokasi ditolak. Silakan izinkan akses lokasi di browser Anda"
          break
        case error.POSITION_UNAVAILABLE:
          gpsErrorMsg.value = "Informasi lokasi tidak tersedia"
          break
        case error.TIMEOUT:
          gpsErrorMsg.value = "Waktu pengambilan lokasi habis"
          break
        default:
          gpsErrorMsg.value = "Gagal mengambil lokasi GPS"
      }
    },
    options
  )
}

watch(() => formData.value.status, (newStatus) => {
  if (newStatus === "hadir" || newStatus === "alpha") {
    formData.value.note = ""
  }
  if (newStatus === "hadir" && (currentUser.value?.role === 'siswa' || currentUser.value?.role === 'siswi')) {
    getGPSLocation()
  } else {
    gpsStatus.value = ""
    gpsErrorMsg.value = ""
    studentLatitude.value = null
    studentLongitude.value = null
    distanceToSchool.value = null
  }
})

async function loadSchoolSettings() {
  try {
    const res: any = await api.get("/api/settings")
    if (res.success && res.data) {
      schoolLatitude.value = parseFloat(res.data.school_latitude) || -6.1822
      schoolLongitude.value = parseFloat(res.data.school_longitude) || 106.2736
      schoolMaxDistance.value = parseInt(res.data.school_max_distance) || 250
      attendanceStartTime.value = res.data.attendance_start_time || "07:00"
      attendanceEndTime.value = res.data.attendance_end_time || "17:00"
      enableAntiFakeGPS.value = res.data.enable_anti_fake_gps || "true"
    }
  } catch (error) {
    console.error("Gagal memuat konfigurasi absensi", error)
  }
}

const filteredAttendances = computed(() => {
  if (currentUser.value?.role !== 'siswa' && currentUser.value?.role !== 'siswi') {
    return attendancesList.value
  }

  const now = new Date()
  const today = new Date(now.getFullYear(), now.getMonth(), now.getDate())

  return attendancesList.value.filter(att => {
    if (!att.date) return false
    const attDate = new Date(att.date)

    if (selectedPeriod.value === "minggu") {
      const currentDay = today.getDay()
      const distanceToMonday = currentDay === 0 ? 6 : currentDay - 1
      const monday = new Date(today)
      monday.setDate(today.getDate() - distanceToMonday)
      monday.setHours(0, 0, 0, 0)
      
      const sunday = new Date(monday)
      sunday.setDate(monday.getDate() + 6)
      sunday.setHours(23, 59, 59, 999)

      return attDate >= monday && attDate <= sunday
    }

    if (selectedPeriod.value === "bulan") {
      const startOfMonth = new Date(today.getFullYear(), today.getMonth(), 1)
      const endOfMonth = new Date(today.getFullYear(), today.getMonth() + 1, 0, 23, 59, 59, 999)
      return attDate >= startOfMonth && attDate <= endOfMonth
    }

    if (selectedPeriod.value === "semester") {
      const currentMonth = today.getMonth()
      let startOfSemester: Date
      let endOfSemester: Date

      if (currentMonth >= 6) {
        startOfSemester = new Date(today.getFullYear(), 6, 1)
        endOfSemester = new Date(today.getFullYear(), 11, 31, 23, 59, 59, 999)
      } else {
        startOfSemester = new Date(today.getFullYear(), 0, 1)
        endOfSemester = new Date(today.getFullYear(), 5, 30, 23, 59, 59, 999)
      }
      return attDate >= startOfSemester && attDate <= endOfSemester
    }

    return true
  })
})

onMounted(() => {
  const cached = localStorage.getItem("user")
  if (cached && !currentUser.value) {
    currentUser.value = JSON.parse(cached)
  }
  if (currentUser.value?.role === 'siswa' || currentUser.value?.role === 'siswi') {
    formData.value.student_id = currentUser.value.student_id
  }
  loadSchoolSettings()
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
    let url = `/api/attendances`
    const isStudent = currentUser.value?.role === 'siswa' || currentUser.value?.role === 'siswi'
    if (!isStudent) {
      url += `?date=${filterDate.value}`
      if (filterClass.value) {
        url += `&class=${filterClass.value}`
      }
      if (searchQuery.value) {
        url += `&search=${searchQuery.value}`
      }
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

  const isStudent = currentUser.value?.role === 'siswa' || currentUser.value?.role === 'siswi'
  const now = new Date()
  const pad = (n: number) => n.toString().padStart(2, '0')
  const timestampStr = `${now.getFullYear()}-${pad(now.getMonth() + 1)}-${pad(now.getDate())} ${pad(now.getHours())}:${pad(now.getMinutes())}:${pad(now.getSeconds())}`
  const nowTimeStr = `${pad(now.getHours())}:${pad(now.getMinutes())}`

  let isTimeInvalid = false
  if (attendanceStartTime.value <= attendanceEndTime.value) {
    isTimeInvalid = nowTimeStr < attendanceStartTime.value || nowTimeStr > attendanceEndTime.value
  } else {
    isTimeInvalid = nowTimeStr < attendanceStartTime.value && nowTimeStr > attendanceEndTime.value
  }

  if (isStudent && isTimeInvalid) {
    errorMessage.value = `Absensi ditolak: Waktu operasional absensi mandiri hanya diperbolehkan pukul ${attendanceStartTime.value} s.d. ${attendanceEndTime.value}`
    return
  }

  const payload: any = {
    student_id: formData.value.student_id,
    date: formData.value.date,
    status: formData.value.status,
    note: formData.value.note,
    timestamp: timestampStr
  }

  if (isStudent && formData.value.status === 'hadir') {
    if (gpsStatus.value !== 'success' || studentLatitude.value === null || studentLongitude.value === null) {
      errorMessage.value = "Lokasi GPS wajib terverifikasi terlebih dahulu"
      return
    }
    if (distanceToSchool.value !== null && distanceToSchool.value > schoolMaxDistance.value) {
      errorMessage.value = "Absensi ditolak karena Anda berada di luar radius area sekolah"
      return
    }
    payload.latitude = studentLatitude.value
    payload.longitude = studentLongitude.value
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
      res = await api.post("/api/attendances", payload)
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
