<script setup lang="ts">
import router from '@/router'
import { useUserdataStore } from '@/stores/userdata'
import axios from 'axios'
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { type User } from '../types/User'

const userdata = useUserdataStore()
const route = useRoute()

const isError = ref(false)
const errorMessage = ref('Something went wrong')

const CLIENT_ID: string = import.meta.env.VITE_SIMPLE_LOGIN_CLIENT_ID // SimpleLogin client ID
const REDIRECT_URI: string = import.meta.env.VITE_ROOT_URI
const BACKEND_URI: string = import.meta.env.VITE_BACKEND_URI

function triggerLoginFlow(user: User) {
  // TODO: Send along timestamp in the response of a request here, triggers update function in the store
  userdata.logInUser(user.username, user.balance)
  router.push('home')
}

/* OAuth login functions */
const oauthCallbackGoogle = async (response: any) => {
  if (response.credential) {
    await axios
      .post(`${BACKEND_URI}/login_google`, {
        token: response.credential
      })
      .then((res) => {
        let userInfo: User = res.data.user_info
        triggerLoginFlow(userInfo)
        isError.value = false
      })
      .catch((err) => {
        isError.value = true
        errorMessage.value = 'Google OAuth login failed: ' + err
      })
  } else {
    isError.value = true
    errorMessage.value = 'Google OAuth login failed due to an undefined response.'
  }
}

function startSimpleLoginSignin() {
  // TODO: Work out a .env here so that on cloudflare you can use the proper https://finite-vault.pages.dev/ domain
  let authUrl = `https://app.simplelogin.io/oauth2/authorize?response_type=code&client_id=${CLIENT_ID}&redirect_uri=${REDIRECT_URI}&scope=profile&state=${{ state: 'noidea' }}`
  location.href = authUrl
}

async function endSimpleLoginSignin(code: string) {
  await axios
    .post(`${BACKEND_URI}/login_proton`, {
      token: code,
      redirect_uri: REDIRECT_URI
    })
    .then((res) => {
      let userInfo: User = res.data.user_info
      triggerLoginFlow(userInfo)
      isError.value = false
    })
    .catch((err) => {
      isError.value = true
      errorMessage.value = 'Proton login failed: ' + err
    })
}

onMounted(() => {
  // SimpleLogin redirect uri set to the same page, so check for a code query
  // TODO: If there's time, make the thing a popup window and make a "mid" component to handle the redirection
  if (route.query.code) {
    endSimpleLoginSignin(String(route.query.code))
  }

  // TODO: Check for a token in the local storage. If present, verify on backend and redirect.
  //       Otherwise, show an error message saying that the found token is expired.
})
</script>

<template>
  <div class="greetings">
    <h1 class="welcome-text">Welcome to Finite Vault!</h1>
    <img src="../assets/img/chest_kenney.png" class="logo" />
    <label for="username" class="username-box">Select a login method below</label>
    <p v-if="isError" class="error-text">{{ errorMessage }}</p>
    <GoogleLogin :callback="oauthCallbackGoogle" class="google-button" />
    <button @click="startSimpleLoginSignin" class=".login-button">
      <!-- TODO: I would like to see a component here, perhaps the same one for redirection, where it informs someone about the email "clause" of this app. -->
      Sign in with Proton / Simple Login
    </button>
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
