<script setup lang="ts">
defineProps({
  amount: Number,
  description: String
  // TODO: Date
})
</script>

<template>
  <div class="greeting">
    <p>${{ amount }} for {{ description }}</p>
  </div>
</template>

<style></style>
