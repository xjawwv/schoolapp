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
            Dashboard Utama
          </h1>
        </div>

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
              Siswa Baru Terdaftar
            </h3>
            <NuxtLink to="/students" class="text-xs uppercase tracking-wider text-[color:var(--color-accent)] hover:opacity-85 font-semibold">
              Lihat Semua Siswa
            </NuxtLink>
          </div>

          <div class="overflow-x-auto">
            <table class="w-full text-left border-collapse min-w-[600px]">
              <thead>
                <tr class="border-b border-[color:var(--color-border)]">
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-[18%]">NISN</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-[25%]">Name</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-[15%]">Class</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-[20%]">Gender</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-[22%]">Phone</th>
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
                    v-for="student in stats.recent_students"
                    :key="'1-' + student.id"
                    class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                  >
                    <td class="py-3.5 px-4 text-sm font-mono text-[color:var(--color-accent)] w-[18%]">{{ student.nis }}</td>
                    <td class="py-3.5 px-4 text-sm font-semibold text-[color:var(--color-heading)] w-[25%]">{{ student.name }}</td>
                    <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] w-[15%]">{{ student.class }}</td>
                    <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] w-[20%]">{{ student.gender }}</td>
                    <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] font-mono w-[22%]">{{ student.phone || '-' }}</td>
                  </tr>

                  <template v-if="stats.recent_students.length > 3">
                    <tr
                      v-for="student in stats.recent_students"
                      :key="'2-' + student.id"
                      class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                    >
                      <td class="py-3.5 px-4 text-sm font-mono text-[color:var(--color-accent)] w-[18%]">{{ student.nis }}</td>
                      <td class="py-3.5 px-4 text-sm font-semibold text-[color:var(--color-heading)] w-[25%]">{{ student.name }}</td>
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] w-[15%]">{{ student.class }}</td>
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] w-[20%]">{{ student.gender }}</td>
                      <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] font-mono w-[22%]">{{ student.phone || '-' }}</td>
                    </tr>
                  </template>

                  <tr v-if="stats.recent_students.length === 0">
                    <td colspan="5" class="py-8 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                      No new student data available
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, nextTick } from "vue"
import {
  Users as UsersIcon,
  CalendarCheck as CalendarCheckIcon,
  Award as AwardIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()

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

const scrollContainer = ref<HTMLElement | null>(null)
const isHovered = ref(false)
let currentScroll = 0
let animationFrameId: number | null = null

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
    } else if (stats.value.recent_students.length > 3) {
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

onMounted(async () => {
  try {
    const res: any = await api.get("/api/dashboard/stats")
    if (res.success && res.data) {
      stats.value = res.data
      if (stats.value.recent_students.length > 3) {
        nextTick(() => {
          initAutoScroll()
        })
      }
    }
  } catch (error) {
    console.error("Gagal memuat statistik dashboard", error)
  }
})

onUnmounted(() => {
  if (animationFrameId) {
    cancelAnimationFrame(animationFrameId)
  }
})
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
