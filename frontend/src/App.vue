<script setup lang="ts">
import { RouterView } from 'vue-router'
import { useUserdataStore } from './stores/userdata'

useUserdataStore() // Initialize store (maybe not needed since called on welcome screen?)
// TODO: Check local storage for a token and check that information.
// TODO: Route respectively (e.g. if already logged in / there is a token and it's valid, make sure we redirect to the home page)
</script>

<template>
  <header>
    <!-- TODO: Eventually redirect whether there's an active session/token or not -->
  </header>

  <main class="main-container">
    <RouterView />
  </main>

  <!-- <footer class="footer">Hi :)</footer> -->
</template>

<style>
.main-container {
  width: 50%;
  margin: auto;
  margin-top: 10vh;
  background-color: var(--background-indigo);
  padding: 40% 1vh;
  text-align: center;
  height: 80em;
}

.footer {
  background-color: green;
}
</style>
