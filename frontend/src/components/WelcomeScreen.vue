<script setup lang="ts">
import router from '@/router'
import { useUserdataStore } from '@/stores/userdata'
import axios from 'axios'
import { ref } from 'vue'

const userdata = useUserdataStore()

const username = ref('')
const isError = ref(false)
const errorMessage = ref('Something went wrong')

// TODO: Changing this to use a token stored in local storage and updating the state if the page refreshes or something
function loginSuccess() {
  if (username.value === '') {
    errorMessage.value = 'Username cannot be empty (for now)'
    isError.value = true
  } else {
    userdata.logInUser(username.value, 4444)
    router.push('home')
  }
}

// TODO: Move this also
type User = {
  name: string
  balance: number
}

// TODO: Rename when you get rid of the above function
function triggerLoginFlow(user: User) {
  userdata.logInUser(user.name, user.balance)
  console.log('Just logged in a user with ' + user.name + ' and a balance of ' + user.balance)
  router.push('home')
}

const oauthCallback = async (response: any) => {
  if (response.credential) {
    await axios
      .post('http://localhost:8080/login', {
        token: response.credential
      })
      .then((res) => {
        let userInfo: User = res.data.user_info
        triggerLoginFlow(userInfo)
        //loginSuccess(userdata.logInUser(res.username, res.balance))
        isError.value = false
      })
      .catch((err) => {
        isError.value = true
        errorMessage.value = 'Error: ' + err
      })
  } else {
    isError.value = true
    errorMessage.value = 'OAuth login failed.'
  }
}
</script>

<template>
  <div class="greetings">
    <h1 class="welcome-text">Welcome to Finite Vault!</h1>
    <img src="../assets/img/chest_kenney.png" class="logo" />
    <label for="username" class="username-box">Enter your username: </label>
    <input v-model="username" name="username" /><br />
    <p v-if="isError" class="error-text">{{ errorMessage }}</p>
    <GoogleLogin :callback="oauthCallback" class="google-button" />
    <button @click="loginSuccess" class="login-button">FALSE BUTTON (for testing)</button>
  </div>
</template>

<style>
@import '../assets/base.css';

.welcome-text {
  color: var(--header-gold);
  padding: 20px;
}

.logo {
  width: 80%;
  margin: 0px auto;
  display: block;
}

.google-button {
  display: block;
  padding: 1em;
  margin: 1em;
}

.login-button {
  background-color: white;
  border: none;
  color: black;
  padding: 1em;
  margin: 2em;
  text-align: center;
  text-decoration: none;
  display: block;
  font-size: 16px;
}

.error-text {
  color: red;
  font-size: 20px;
}
</style>
