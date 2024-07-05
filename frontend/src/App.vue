<script setup lang="ts">
import { RouterView } from 'vue-router'
import { useUserdataStore } from './stores/userdata'

useUserdataStore() // Initialize store (maybe not needed since called on welcome screen?)
</script>

<template>
  <header>
    <!-- TODO: Eventually redirect whether there's an active session/token or not -->
    <main>
      <RouterView />
    </main>
  </header>
</template>

<style></style>
