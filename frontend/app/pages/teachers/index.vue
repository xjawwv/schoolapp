<template>
  <div class="flex h-screen bg-[color:var(--color-bg)] overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col overflow-y-auto">
      <Navbar />
      <main class="p-6 sm:p-8 pt-8 sm:pt-10 max-w-7xl w-full mx-auto space-y-6 sm:space-y-8">
        <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
          <div>
            <h2 class="text-xs uppercase tracking-[0.2em] text-[color:var(--color-accent)] font-semibold mb-2">
              Direktori Akademik
            </h2>
            <h1 class="text-4xl font-bold text-[color:var(--color-heading)] tracking-tight">
              Data Guru Terdaftar
            </h1>
          </div>
          <NuxtLink v-if="currentUser?.role === 'admin'" to="/teachers/create" class="button-primary flex items-center space-x-2 self-start md:self-auto py-3">
            <PlusIcon class="w-4 h-4" />
            <span>Tambah Guru</span>
          </NuxtLink>
        </div>

        <div class="bg-[color:var(--color-surface)] border border-[color:var(--color-border)] p-6 space-y-6">
          <div class="flex flex-col md:flex-row gap-4 items-center justify-between">
            <div class="flex flex-col md:flex-row gap-3 w-full md:max-w-xl">
              <div class="relative flex-1">
                <SearchIcon class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-[color:var(--color-muted)] pointer-events-none" />
                <input
                  v-model="searchQuery"
                  type="text"
                  class="input w-full"
                  style="padding-left: 2.5rem !important;"
                  placeholder="Cari berdasarkan nama, NIP, atau email..."
                />
              </div>
            </div>
            <div class="text-xs uppercase tracking-wider text-[color:var(--color-muted)] font-semibold shrink-0">
              Total Guru: <span class="text-[color:var(--color-accent)]">{{ filteredTeachers.length }}</span>
            </div>
          </div>

          <div class="overflow-x-auto">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="border-b border-[color:var(--color-border)]">
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold w-12 text-center">No</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">NIP</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Nama Lengkap</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold">Email</th>
                  <th class="py-3 px-4 text-xs uppercase tracking-widest text-[color:var(--color-muted)] font-bold text-right">Kontak</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="(teacher, index) in filteredTeachers"
                  :key="teacher.id"
                  class="border-b border-[color:var(--color-border)] hover:bg-[color:var(--color-bg)] transition duration-150"
                >
                  <td class="py-3.5 px-4 text-sm text-[color:var(--color-muted)] font-semibold text-center font-mono w-12">{{ index + 1 }}</td>
                  <td class="py-3.5 px-4 text-sm font-mono text-[color:var(--color-accent)]">{{ teacher.nip || '-' }}</td>
                  <td class="py-3.5 px-4 text-sm font-semibold text-[color:var(--color-heading)]">{{ teacher.name }}</td>
                  <td class="py-3.5 px-4 text-sm text-[color:var(--color-text)] font-mono">{{ teacher.email }}</td>
                  <td class="py-3.5 px-4 text-sm text-right">
                    <a
                      :href="`mailto:${teacher.email}`"
                      class="text-[color:var(--color-muted)] hover:text-[color:var(--color-accent)] transition duration-150 cursor-pointer"
                      title="Kirim Email"
                    >
                      <MailIcon class="w-4 h-4 inline" />
                    </a>
                  </td>
                </tr>
                <tr v-if="filteredTeachers.length === 0">
                  <td colspan="5" class="py-12 text-center text-sm text-[color:var(--color-muted)] uppercase tracking-wider">
                    Tidak ada data guru ditemukan
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
import { ref, onMounted, computed } from "vue"
import {
  Search as SearchIcon,
  Mail as MailIcon,
  Plus as PlusIcon
} from "lucide-vue-next"
import { useApi } from "~/composables/useApi"

definePageMeta({
  middleware: ["auth"]
})

const api = useApi()
const currentUser = useState<any>("currentUser", () => null)
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
  const cached = localStorage.getItem("user")
  if (cached && !currentUser.value) {
    currentUser.value = JSON.parse(cached)
  }
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
</script>
