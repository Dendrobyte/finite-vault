<script setup lang="ts">
import router from '@/router'
import { useUserdataStore } from '@/stores/userdata'
import { ref } from 'vue'

const userdata = useUserdataStore()

const username = ref('')
const password = ref('')

function loginSuccess() {
  userdata.logInUser(username.value, password.value)
  router.push('home')
}
</script>

<template>
  <div class="greetings">
    <h1>Welcome to Finite Vault!</h1>
    <label for="username">Enter your username: </label>
    <input v-model="username" name="username" /><br />
    <label for="password">NOT SECURE PASSWORD FIELD: </label>
    <input v-mode="password" name="password" /><br />
    <button @click="loginSuccess">Log In</button>
  </div>
</template>

<style></style>
