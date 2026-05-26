<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8">
        <div>
          <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
            Direktori Akademik
          </h2>
          <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
            Daftar Guru & Staff Pengajar
          </h1>
          <p class="text-sm text-[color:var(--color-muted)] mt-2">
            Hubungi atau cari profil guru pengajar Anda secara langsung melalui direktori resmi di bawah ini
          </p>
        </div>

        <div class="flex flex-col md:flex-row gap-4 items-center justify-between">
          <div class="relative w-full md:max-w-md">
            <SearchIcon class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-[color:var(--color-muted)] pointer-events-none" />
            <input
              v-model="searchQuery"
              type="text"
              class="input w-full"
              style="padding-left: 2.5rem !important;"
              placeholder="Cari nama, NIP, atau email guru..."
            />
          </div>
          <div class="text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-semibold shrink-0">
            Total Pengajar: <span class="text-[color:var(--color-accent)]">{{ filteredTeachers.length }}</span> Guru
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="teacher in filteredTeachers"
            :key="teacher.id"
            class="card hover:shadow-[--shadow-md] hover:-translate-y-1 transition duration-200 border-t-4 border-t-emerald-500 flex flex-col justify-between"
          >
            <div class="space-y-4">
              <div class="flex items-center space-x-4 border-b border-[color:var(--color-border)] pb-4">
                <div v-if="teacher.avatar" class="w-14 h-14 rounded-full border border-[color:var(--color-border)] bg-cover bg-center shrink-0" :style="`background-image: url('${teacher.avatar}')`"></div>
                <div
                  v-else
                  class="w-14 h-14 rounded-full border flex items-center justify-center text-lg font-bold font-mono shrink-0"
                  :class="getAvatarBgClass(teacher.name)"
                >
                  {{ teacher.name.charAt(0).toUpperCase() }}
                </div>
                <div class="truncate">
                  <h3 class="text-base font-bold text-[color:var(--color-heading)] truncate" :title="teacher.name">
                    {{ teacher.name }}
                  </h3>
                  <span class="text-[10px] font-bold uppercase tracking-wider bg-emerald-500/10 text-emerald-500 py-0.5 px-2 rounded-sm inline-block mt-1">
                    Guru Pengajar
                  </span>
                </div>
              </div>

              <div class="space-y-3.5 text-sm">
                <div>
                  <span class="block text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold mb-1">NIP (Nomor Induk Pegawai)</span>
                  <span class="font-mono text-sm text-[color:var(--color-accent)] font-bold">{{ teacher.nip || '-' }}</span>
                </div>
                <div>
                  <span class="block text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold mb-1">Alamat Surel / Email</span>
                  <span class="text-sm text-[color:var(--color-text)] font-mono break-all font-semibold">{{ teacher.email }}</span>
                </div>
              </div>
            </div>

            <div class="pt-5 border-t border-[color:var(--color-border)] mt-5">
              <a
                :href="`mailto:${teacher.email}`"
                class="button-ghost w-full justify-center flex items-center space-x-2 text-xs uppercase font-bold tracking-wider py-2.5"
              >
                <MailIcon class="w-4 h-4 shrink-0" />
                <span>Kirim Pesan Email</span>
              </a>
            </div>
          </div>
        </div>

        <div v-if="filteredTeachers.length === 0" class="card shadow-[--shadow-sm] py-16 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-widest leading-relaxed">
          Tidak ada data guru yang cocok dengan pencarian Anda
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from "vue"
import {
  Search as SearchIcon,
  Mail as MailIcon,
  GraduationCap as GraduationCapIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"
import Sidebar from "~/components/Sidebar.vue"
import Navbar from "~/components/Navbar.vue"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()
const teachersList = ref<any[]>([])
const searchQuery = ref("")

const filteredTeachers = computed(() => {
  return teachersList.value.filter(t => {
    const term = searchQuery.value.toLowerCase()
    const nameMatch = t.name.toLowerCase().includes(term)
    const emailMatch = t.email.toLowerCase().includes(term)
    const nipMatch = t.nip && t.nip.toLowerCase().includes(term)
    return nameMatch || emailMatch || nipMatch
  })
})

onMounted(() => {
  loadTeachers()
})

async function loadTeachers() {
  try {
    const res: any = await api.get("/api/teachers")
    if (res.success && res.data) {
      teachersList.value = res.data || []
    }
  } catch (error) {
    console.error("Gagal memuat direktori guru", error)
  }
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
</script>
