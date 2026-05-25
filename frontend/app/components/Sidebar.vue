<template>
  <aside class="w-[240px] bg-[color:var(--color-surface)] border-r border-[color:var(--color-border)] flex flex-col h-screen shrink-0">
    <div class="h-16 flex items-center px-6 border-b border-[color:var(--color-border)]">
      <span class="text-lg font-bold text-[color:var(--color-heading)] tracking-wider">
        SMA N 1 METRO
      </span>
    </div>
    <nav class="flex-1 py-6 px-4 space-y-1">
      <NuxtLink
        v-for="item in menuItems"
        :key="item.path"
        :to="item.path"
        class="flex items-center space-x-3 px-4 py-3 text-sm font-medium transition duration-150 group"
        :class="isActive(item.path)
          ? 'bg-[color:var(--color-bg)] text-[color:var(--color-accent)] border-l-2 border-[color:var(--color-accent)]'
          : 'text-[color:var(--color-text)] hover:bg-[color:var(--color-bg)] hover:text-[color:var(--color-heading)] border-l-2 border-transparent'"
      >
        <component
          :is="item.icon"
          class="w-5 h-5 shrink-0 transition duration-150"
          :class="isActive(item.path) ? 'text-[color:var(--color-accent)]' : 'text-[color:var(--color-muted)] group-hover:text-[color:var(--color-text)]'"
        />
        <span>{{ item.label }}</span>
      </NuxtLink>
    </nav>
    <div class="p-4 border-t border-[color:var(--color-border)] text-xs text-[color:var(--color-muted)] text-center tracking-wider uppercase">
      V1.0.0
    </div>
  </aside>
</template>

<script setup lang="ts">
import { computed } from "vue"
import { useRoute } from "vue-router"
import {
  LayoutDashboard as LayoutDashboardIcon,
  Users as UsersIcon,
  CalendarCheck as CalendarCheckIcon,
  Award as AwardIcon
} from "lucide-vue-next"

const route = useRoute()

const menuItems = [
  { label: "Dashboard", path: "/dashboard", icon: LayoutDashboardIcon },
  { label: "Data Siswa", path: "/students", icon: UsersIcon },
  { label: "Absensi", path: "/attendance", icon: CalendarCheckIcon },
  { label: "Nilai", path: "/grades", icon: AwardIcon }
]

function isActive(path: string): boolean {
  if (path === "/dashboard") {
    return route.path === "/dashboard"
  }
  return route.path.startsWith(path)
}
</script>
