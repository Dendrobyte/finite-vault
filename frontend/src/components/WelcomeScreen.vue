<script setup lang="ts">
import router from '@/router'
import { useUserdataStore } from '@/stores/userdata'
import { ref } from 'vue'

const userdata = useUserdataStore()

const username = ref('')
const password = ref('Dummy Password')
const isError = ref(false)
const errorMessage = ref('Something went wrong')

// TODO: Changing this to use a token stored in local storage and updating the state if the page refreshes or something
function loginSuccess() {
  if (username.value === '') {
    errorMessage.value = 'Username cannot be empty (for now)'
    isError.value = true
  } else {
    userdata.logInUser(username.value, password.value)
    router.push('home')
  }
}
</script>

<template>
  <div class="greetings">
    <h1 class="welcome-text">Welcome to Finite Vault!</h1>
    <label for="username" class="username-box">Enter your username: </label>
    <input v-model="username" name="username" /><br />
    <p v-if="isError" class="error-text">{{ errorMessage }}</p>
    <button @click="loginSuccess" class="login-button">Log In With Google</button>
  </div>
</template>

<style>
@import '../assets/base.css';

.welcome-text {
  color: var(--header-gold);
  padding: 20px;
}

.login-button {
  background-color: white;
  border: none;
  color: black;
  padding: 1em;
  margin: 2em;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
}

.error-text {
  color: red;
  font-size: 20px;
}
</style>
