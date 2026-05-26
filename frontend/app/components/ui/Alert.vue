<script setup lang="ts">
import { computed } from "vue"

const props = defineProps<{
  variant?: "default" | "destructive" | "success"
  class?: string
}>()

const classes = computed(() => {
  const base = "relative w-full rounded-lg border p-4 [&>svg~*]:pl-7 [&>svg]:absolute [&>svg]:left-4 [&>svg]:top-4"
  const variants = {
    default: "bg-[color:var(--color-surface)] text-[color:var(--color-text)] border-[color:var(--color-border)] border-l-4 border-l-[color:var(--color-accent)] shadow-[--shadow-md] [&>svg]:text-[color:var(--color-accent)]",
    destructive: "bg-[color:var(--color-surface)] text-[color:var(--color-error)] border-[color:var(--color-border)] border-l-4 border-l-[color:var(--color-error)] shadow-[--shadow-md] [&>svg]:text-[color:var(--color-error)]",
    success: "bg-[color:var(--color-surface)] text-[color:var(--color-success)] border-[color:var(--color-border)] border-l-4 border-l-[color:var(--color-success)] shadow-[--shadow-md] [&>svg]:text-[color:var(--color-success)]"
  }
  return [base, variants[props.variant || "default"], props.class].join(" ")
})
</script>

<template>
  <div :class="classes" role="alert">
    <slot />
  </div>
</template>
