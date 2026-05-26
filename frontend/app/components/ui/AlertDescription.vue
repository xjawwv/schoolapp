<script setup lang="ts">
const props = defineProps<{
  class?: string
}>()
</script>

<template>
  <div :class="['text-xs opacity-90 leading-relaxed font-normal text-[color:var(--color-muted)]', props.class]">
    <slot />
  </div>
</template>
