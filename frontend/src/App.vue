<script setup lang="ts">
import { RouterView } from 'vue-router'
import { useUserdataStore } from './stores/userdata'

/* FOR HOSTING TO PROD */
// Check the launch params here for environments: https://cli.vuejs.org/guide/mode-and-env.html

useUserdataStore() // Initialize store (maybe not needed since called on welcome screen?)

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
/* TODO: Figure out how to make this max-mobile size within the browser */
.main-container {
  width: 60%;
  margin: auto;
  margin-top: 16vh;
  text-align: center;
}

/* TODO: Oh yea, I should make this at some point */
.footer {
  background-color: green;
}
</style>
