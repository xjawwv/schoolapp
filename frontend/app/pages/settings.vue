<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8">
        <div>
          <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
            Pengaturan Sistem
          </h2>
          <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
            Setting Konfigurasi
          </h1>
        </div>

        <div v-if="toast.message" :class="[
          'border p-3 text-sm font-medium flex items-center space-x-2',
          toast.type === 'success' ? 'bg-[color:var(--color-surface)] border-[color:var(--color-success)] text-[color:var(--color-success)]' : 'bg-[color:var(--color-surface)] border-[color:var(--color-error)] text-[color:var(--color-error)]'
        ]">
          <CheckCircleIcon v-if="toast.type === 'success'" class="w-4 h-4 shrink-0" />
          <AlertCircleIcon v-else class="w-4 h-4 shrink-0" />
          <span>{{ toast.message }}</span>
        </div>

        <div class="space-y-8">
          <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6 shadow-[--shadow-sm] w-full">
            <div class="flex items-center space-x-3 border-b border-[color:var(--color-border)] pb-4">
              <GlobeIcon class="w-5 h-5 text-[color:var(--color-accent)] shrink-0" />
              <div>
                <h3 class="text-md font-bold text-[color:var(--color-heading)]">Identitas Aplikasi</h3>
                <p class="text-xs text-[color:var(--color-muted)]">Atur penamaan nama instansi atau sekolah</p>
              </div>
            </div>

            <form @submit.prevent="saveSiteSettings" class="space-y-4">
              <div class="flex flex-col space-y-1.5 max-w-2xl">
                <label class="text-xs uppercase tracking-wider text-[color:var(--color-text)] font-semibold">Nama Website</label>
                <input
                  v-model="siteSettings.site_name"
                  type="text"
                  class="input w-full font-semibold"
                  placeholder="SMA N 1 METRO"
                  required
                />
              </div>

              <div class="pt-2">
                <button type="submit" class="button-primary w-full md:w-auto" :disabled="loading.site">
                  {{ loading.site ? 'Menyimpan...' : 'Simpan Pengaturan' }}
                </button>
              </div>
            </form>
          </div>

          <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6 shadow-[--shadow-sm] w-full">
            <div class="flex items-center space-x-3 border-b border-[color:var(--color-border)] pb-4">
              <MapPinIcon class="w-5 h-5 text-[color:var(--color-accent)] shrink-0" />
              <div>
                <h3 class="text-md font-bold text-[color:var(--color-heading)]">Lokasi & Operasional Absensi</h3>
                <p class="text-xs text-[color:var(--color-muted)]">Atur koordinat pusat sekolah, radius jangkauan absensi, dan jam operasional absensi mandiri siswa</p>
              </div>
            </div>

            <form @submit.prevent="saveAttendanceSettings" class="space-y-6">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div class="space-y-4">
                  <div class="grid grid-cols-2 gap-4">
                    <div class="flex flex-col space-y-1.5">
                      <label class="text-xs uppercase tracking-wider text-[color:var(--color-text)] font-semibold">Latitude</label>
                      <input
                        v-model="siteSettings.school_latitude"
                        @input="updateMapFromInputs"
                        type="text"
                        class="input w-full font-mono font-semibold"
                        placeholder="-6.1822"
                        required
                      />
                    </div>
                    <div class="flex flex-col space-y-1.5">
                      <label class="text-xs uppercase tracking-wider text-[color:var(--color-text)] font-semibold">Longitude</label>
                      <input
                        v-model="siteSettings.school_longitude"
                        @input="updateMapFromInputs"
                        type="text"
                        class="input w-full font-mono font-semibold"
                        placeholder="106.2736"
                        required
                      />
                    </div>
                  </div>

                  <div class="flex flex-col space-y-1.5">
                    <label class="text-xs uppercase tracking-wider text-[color:var(--color-text)] font-semibold">Radius Maksimal Jangkauan (Meter)</label>
                    <div class="flex items-center space-x-3">
                      <input
                        v-model.number="siteSettings.school_max_distance"
                        type="number"
                        class="input w-full md:w-48 font-mono font-bold border border-[color:var(--color-border)] px-4 py-2.5 bg-[color:var(--color-bg)]"
                        placeholder="Contoh: 250"
                        required
                      />
                      <span class="text-xs font-semibold text-[color:var(--color-muted)] uppercase tracking-wider">Meter</span>
                    </div>
                    <input
                      v-model.number="siteSettings.school_max_distance"
                      type="range"
                      min="50"
                      max="2000"
                      step="10"
                      class="w-full accent-[color:var(--color-accent)] mt-2"
                    />
                  </div>

                  <div class="grid grid-cols-2 gap-4">
                    <div class="flex flex-col space-y-1.5">
                      <label class="text-xs uppercase tracking-wider text-[color:var(--color-text)] font-semibold">Jam Mulai Absen</label>
                      <input
                        v-model="siteSettings.attendance_start_time"
                        type="time"
                        class="input w-full font-mono font-semibold"
                        required
                      />
                    </div>
                    <div class="flex flex-col space-y-1.5">
                      <label class="text-xs uppercase tracking-wider text-[color:var(--color-text)] font-semibold">Jam Selesai Absen</label>
                      <input
                        v-model="siteSettings.attendance_end_time"
                        type="time"
                        class="input w-full font-mono font-semibold"
                        required
                      />
                    </div>
                  </div>
                </div>

                <div class="flex flex-col space-y-2 relative">
                  <label class="text-xs uppercase tracking-wider text-[color:var(--color-text)] font-semibold">Pilih Pusat Koordinat Peta</label>
                  <div class="relative w-full h-[250px] border border-[color:var(--color-border)] rounded-lg bg-[color:var(--color-bg)] overflow-hidden z-10">
                    <div id="map" class="w-full h-full"></div>
                    <div class="absolute top-3 right-3 flex bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-1 rounded-lg shadow-lg space-x-1" style="z-index: 1000 !important;">
                      <button
                        type="button"
                        @click="setMapType('roadmap')"
                        class="px-2.5 py-1 text-[10px] uppercase tracking-wider font-bold rounded transition duration-150 cursor-pointer"
                        :class="mapType === 'roadmap' ? 'bg-[color:var(--color-accent)] text-[color:var(--color-accent-fg)]' : 'text-[color:var(--color-text)] hover:bg-[color:var(--color-bg)]'"
                      >
                        Peta
                      </button>
                      <button
                        type="button"
                        @click="setMapType('satellite')"
                        class="px-2.5 py-1 text-[10px] uppercase tracking-wider font-bold rounded transition duration-150 cursor-pointer"
                        :class="mapType === 'satellite' ? 'bg-[color:var(--color-accent)] text-[color:var(--color-accent-fg)]' : 'text-[color:var(--color-text)] hover:bg-[color:var(--color-bg)]'"
                      >
                        Satelit
                      </button>
                    </div>
                  </div>
                  <p class="text-[10px] text-[color:var(--color-muted)] italic">Geser penanda merah atau klik lokasi mana saja pada peta untuk mengatur koordinat sekolah.</p>
                </div>
              </div>

              <div class="pt-2 border-t border-[color:var(--color-border)]">
                <button type="submit" class="button-primary w-full md:w-auto flex items-center justify-center py-3" :disabled="loading.attendance">
                  <span v-if="loading.attendance" class="animate-spin rounded-full h-3.5 w-3.5 border-2 border-[color:var(--color-accent-fg)] border-t-transparent mr-2"></span>
                  <span>{{ loading.attendance ? 'Menyimpan...' : 'Simpan Setelan Operasional' }}</span>
                </button>
              </div>
            </form>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from "vue"
import {
  Globe as GlobeIcon,
  CheckCircle as CheckCircleIcon,
  AlertCircle as AlertCircleIcon,
  MapPin as MapPinIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()

const siteName = useState("siteName", () => "SMA N 1 METRO")
const currentUser = useState<any>("currentUser", () => null)

const toast = ref<{ message: string; type: "success" | "error" }>({
  message: "",
  type: "success"
})

const loading = ref({
  site: false,
  attendance: false
})

const siteSettings = ref({
  site_name: "",
  school_latitude: "-6.1822",
  school_longitude: "106.2736",
  school_max_distance: "250",
  attendance_start_time: "07:00",
  attendance_end_time: "17:00"
})

let map: any = null
let marker: any = null
let circle: any = null
const mapType = ref("roadmap")
let roadmapLayer: any = null
let satelliteLayer: any = null

function showToast(message: string, type: "success" | "error") {
  toast.value = { message, type }
  setTimeout(() => {
    toast.value.message = ""
  }, 4000)
}

function loadLeaflet(): Promise<any> {
  return new Promise((resolve) => {
    if ((window as any).L) {
      resolve((window as any).L)
      return
    }
    const link = document.createElement("link")
    link.rel = "stylesheet"
    link.href = "https://unpkg.com/leaflet@1.9.4/dist/leaflet.css"
    document.head.appendChild(link)

    const script = document.createElement("script")
    script.src = "https://unpkg.com/leaflet@1.9.4/dist/leaflet.js"
    script.onload = () => {
      resolve((window as any).L)
    }
    document.body.appendChild(script)
  })
}

function updateCircle(L: any, lat: number, lon: number, radius: number) {
  if (circle) {
    circle.setLatLng([lat, lon])
    circle.setRadius(radius)
  } else if (map) {
    circle = L.circle([lat, lon], {
      color: "#eab308",
      fillColor: "#eab308",
      fillOpacity: 0.15,
      weight: 2,
      radius: radius
    }).addTo(map)
  }
}

function setMapType(type: string) {
  if (!map || !(window as any).L) return
  const L = (window as any).L
  mapType.value = type

  if (type === "satellite") {
    if (roadmapLayer) map.removeLayer(roadmapLayer)
    if (!satelliteLayer) {
      satelliteLayer = L.tileLayer("https://mt1.google.com/vt/lyrs=y&x={x}&y={y}&z={z}", {
        attribution: '&copy; <a href="https://maps.google.com">Google Maps</a>'
      })
    }
    satelliteLayer.addTo(map)
  } else {
    if (satelliteLayer) map.removeLayer(satelliteLayer)
    if (!roadmapLayer) {
      roadmapLayer = L.tileLayer("https://mt1.google.com/vt/lyrs=m&x={x}&y={y}&z={z}", {
        attribution: '&copy; <a href="https://maps.google.com">Google Maps</a>'
      })
    }
    roadmapLayer.addTo(map)
  }
}

async function initMap(L: any) {
  const lat = parseFloat(siteSettings.value.school_latitude) || -6.1822
  const lon = parseFloat(siteSettings.value.school_longitude) || 106.2736
  const radius = parseFloat(siteSettings.value.school_max_distance) || 250

  map = L.map("map", { zoomControl: false }).setView([lat, lon], 15)
  L.control.zoom({ position: 'bottomleft' }).addTo(map)

  roadmapLayer = L.tileLayer("https://mt1.google.com/vt/lyrs=m&x={x}&y={y}&z={z}", {
    attribution: '&copy; <a href="https://maps.google.com">Google Maps</a>'
  })
  roadmapLayer.addTo(map)

  marker = L.marker([lat, lon], { draggable: true }).addTo(map)

  updateCircle(L, lat, lon, radius)

  marker.on("dragend", () => {
    const position = marker.getLatLng()
    siteSettings.value.school_latitude = position.lat.toFixed(6)
    siteSettings.value.school_longitude = position.lng.toFixed(6)
    updateCircle(L, position.lat, position.lng, parseFloat(siteSettings.value.school_max_distance) || 250)
  })

  map.on("click", (e: any) => {
    marker.setLatLng(e.latlng)
    siteSettings.value.school_latitude = e.latlng.lat.toFixed(6)
    siteSettings.value.school_longitude = e.latlng.lng.toFixed(6)
    updateCircle(L, e.latlng.lat, e.latlng.lng, parseFloat(siteSettings.value.school_max_distance) || 250)
  })
}

function updateMapFromInputs() {
  if (map && marker && (window as any).L) {
    const L = (window as any).L
    const lat = parseFloat(siteSettings.value.school_latitude)
    const lon = parseFloat(siteSettings.value.school_longitude)
    const radius = parseFloat(siteSettings.value.school_max_distance) || 250
    if (!isNaN(lat) && !isNaN(lon)) {
      const latlng = [lat, lon]
      marker.setLatLng(latlng)
      map.setView(latlng, map.getZoom())
      updateCircle(L, lat, lon, radius)
    }
  }
}

watch(() => siteSettings.value.school_max_distance, (newVal) => {
  if (map && marker && (window as any).L) {
    const L = (window as any).L
    const lat = parseFloat(siteSettings.value.school_latitude)
    const lon = parseFloat(siteSettings.value.school_longitude)
    const radius = parseFloat(newVal as any) || 250
    if (!isNaN(lat) && !isNaN(lon)) {
      updateCircle(L, lat, lon, radius)
    }
  }
})

onMounted(async () => {
  const userJson = localStorage.getItem("user")
  if (userJson && !currentUser.value) {
    currentUser.value = JSON.parse(userJson)
  }

  if (currentUser.value?.role !== "admin") {
    navigateTo("/dashboard")
    return
  }

  try {
    const res: any = await api.get("/api/settings")
    if (res.success && res.data) {
      siteSettings.value.site_name = res.data.site_name || ""
      siteSettings.value.school_latitude = res.data.school_latitude || "-6.1822"
      siteSettings.value.school_longitude = res.data.school_longitude || "106.2736"
      siteSettings.value.school_max_distance = res.data.school_max_distance || "250"
      siteSettings.value.attendance_start_time = res.data.attendance_start_time || "07:00"
      siteSettings.value.attendance_end_time = res.data.attendance_end_time || "17:00"
      siteName.value = res.data.site_name || "SMA N 1 METRO"
    }

    const L = await loadLeaflet()
    await initMap(L)
  } catch (error) {
    showToast("Gagal memuat konfigurasi sistem", "error")
  }
})

async function saveSiteSettings() {
  loading.value.site = true
  try {
    const res: any = await api.put("/api/settings", {
      site_name: siteSettings.value.site_name
    })
    if (res.success) {
      siteName.value = siteSettings.value.site_name
      localStorage.setItem("siteName", siteSettings.value.site_name)
      showToast("Pengaturan identitas aplikasi berhasil disimpan", "success")
    } else {
      showToast(res.message || "Gagal menyimpan pengaturan", "error")
    }
  } catch (error: any) {
    showToast(error.data?.message || "Terjadi kesalahan server", "error")
  } finally {
    loading.value.site = false
  }
}

async function saveAttendanceSettings() {
  loading.value.attendance = true
  try {
    const payload = {
      school_latitude: siteSettings.value.school_latitude,
      school_longitude: siteSettings.value.school_longitude,
      school_max_distance: String(siteSettings.value.school_max_distance),
      attendance_start_time: siteSettings.value.attendance_start_time,
      attendance_end_time: siteSettings.value.attendance_end_time
    }
    const res: any = await api.put("/api/settings", payload)
    if (res.success) {
      showToast("Pengaturan operasional absensi berhasil disimpan", "success")
    } else {
      showToast(res.message || "Gagal menyimpan pengaturan", "error")
    }
  } catch (error: any) {
    showToast(error.data?.message || "Terjadi kesalahan server", "error")
  } finally {
    loading.value.attendance = false
  }
}
</script>
