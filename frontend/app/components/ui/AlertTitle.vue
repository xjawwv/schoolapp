<script setup lang="ts">
const props = defineProps<{
  class?: string
}>()
</script>

<template>
  <h5 :class="['font-semibold leading-none tracking-tight mb-1 text-sm text-[color:var(--color-heading)]', props.class]">
    <slot />
  </h5>
</template>
