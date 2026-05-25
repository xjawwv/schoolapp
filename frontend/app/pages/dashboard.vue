<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-8 max-w-7xl w-full mx-auto space-y-8">
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
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="border-b border-[color:var(--color-border)]">
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">NIS</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Nama</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Kelas</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Gender</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Telepon</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="student in stats.recent_students"
                  :key="student.id"
                  class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                >
                  <td class="py-3.5 px-4 text-sm font-mono text-[color:var(--color-accent)]">{{ student.nis }}</td>
                  <td class="py-3.5 px-4 text-sm font-semibold text-[color:var(--color-heading)]">{{ student.name }}</td>
                  <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">{{ student.class }}</td>
                  <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)]">{{ student.gender }}</td>
                  <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] font-mono">{{ student.phone || '-' }}</td>
                </tr>
                <tr v-if="stats.recent_students.length === 0">
                  <td colspan="5" class="py-8 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                    Tidak ada data siswa terbaru
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
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

onMounted(async () => {
  try {
    const res = await api.get("/api/dashboard/stats")
    if (res.success && res.data) {
      stats.value = res.data
    }
  } catch (error) {
    console.error("Gagal memuat statistik dashboard", error)
  }
})
</script>
