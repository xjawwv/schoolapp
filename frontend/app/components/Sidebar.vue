<template>
  <div>
    <div
      v-if="isSidebarOpen"
      @click="isSidebarOpen = false"
      class="fixed inset-0 bg-black/60 z-40 lg:hidden"
    ></div>

    <aside
      class="fixed inset-y-0 left-0 z-40 lg:static w-[240px] bg-[color:var(--color-surface)] border-r border-[color:var(--color-border)] flex flex-col h-screen shrink-0 transition-transform duration-300 transform lg:transform-none"
      :class="isSidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'"
    >
      <div class="h-16 flex items-center justify-between px-6 border-b border-[color:var(--color-border)]">
        <span class="text-lg font-bold text-[color:var(--color-heading)] tracking-wider truncate mr-2">
          {{ siteName }}
        </span>
        <button @click="isSidebarOpen = false" class="lg:hidden text-[color:var(--color-text)] hover:text-[color:var(--color-heading)] cursor-pointer">
          <XIcon class="w-5 h-5" />
        </button>
      </div>
      <nav class="flex-1 py-6 px-4 space-y-6 overflow-y-auto no-scrollbar">
        <div v-for="group in menuGroups" :key="group.title" class="space-y-1.5">
          <h4 class="text-[10px] font-bold text-[color:var(--color-muted)] tracking-widest uppercase px-3 mb-2">
            {{ group.title }}
          </h4>
          <NuxtLink
            v-for="item in group.items"
            :key="item.path"
            :to="item.path"
            class="flex items-center space-x-3 px-3 py-2.5 rounded-lg text-sm font-medium transition duration-150 group"
            :class="isActive(item.path)
              ? 'bg-[color:var(--color-bg)] text-[color:var(--color-accent)] font-semibold'
              : 'text-[color:var(--color-text)] hover:bg-[color:var(--color-bg)]/50 hover:text-[color:var(--color-heading)]'"
          >
            <component
              :is="item.icon"
              class="w-5 h-5 shrink-0 transition duration-150"
              :class="isActive(item.path) ? 'text-[color:var(--color-accent)]' : 'text-[color:var(--color-muted)] group-hover:text-[color:var(--color-text)]'"
            />
            <span>{{ item.label }}</span>
          </NuxtLink>
        </div>
      </nav>
      <div class="p-4 border-t border-[color:var(--color-border)] text-xs text-[color:var(--color-muted)] text-center tracking-wider uppercase">
        V1.3.2
      </div>
    </aside>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from "vue"
import { useRoute } from "vue-router"
import {
  LayoutDashboard as LayoutDashboardIcon,
  Users as UsersIcon,
  CalendarCheck as CalendarCheckIcon,
  Settings as SettingsIcon,
  User as UserIcon,
  GraduationCap as GraduationCapIcon,
  X as XIcon,
  Bell as BellIcon
} from "lucide-vue-next"

const route = useRoute()
const isSidebarOpen = useState("sidebarOpen", () => false)
const siteName = useState("siteName", () => "SMA N 1 METRO")
const user = useState<any>("currentUser", () => null)

const menuGroups = computed(() => {
  const role = user.value?.role || ""

  const academicItems = (role === "siswa" || role === "siswi")
    ? [
        { label: "Absensi", path: "/attendance", icon: CalendarCheckIcon },
        { label: "Daftar Guru", path: "/teachers", icon: GraduationCapIcon },
        { label: "Notifikasi", path: "/notification", icon: BellIcon }
      ]
    : [
        { label: "Data Siswa", path: "/students", icon: UsersIcon },
        { label: "Daftar Guru", path: "/teachers", icon: GraduationCapIcon },
        { label: "Absensi", path: "/attendance", icon: CalendarCheckIcon },
        { label: "Notifikasi", path: "/notification", icon: BellIcon }
      ]

  const configItems = []
  configItems.push({ label: "Profil", path: "/profile", icon: UserIcon })
  if (role === "admin") {
    configItems.push({ label: "Pengguna", path: "/user", icon: UsersIcon })
    configItems.push({ label: "Setting", path: "/settings", icon: SettingsIcon })
  }

  return [
    {
      title: "Utama",
      items: [
        { label: "Dashboard", path: "/dashboard", icon: LayoutDashboardIcon }
      ]
    },
    {
      title: "Akademik",
      items: academicItems
    },
    {
      title: "Konfigurasi",
      items: configItems
    }
  ]
})

onMounted(() => {
  const cached = localStorage.getItem("user")
  if (cached && !user.value) {
    user.value = JSON.parse(cached)
  }
})

function isActive(path: string): boolean {
  if (!route || !route.path) return false
  if (path === "/dashboard") {
    return route.path === "/dashboard"
  }
  return route.path.startsWith(path)
}
</script>
