<script setup lang="ts">
import router from '@/router'
import { useSimpleLoginStore } from '@/stores/simplelogin'
import { useUserdataStore } from '@/stores/userdata'
import axios from 'axios'
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'

const userdata = useUserdataStore()
const simpleLoginState = useSimpleLoginStore() // Not sure if I need this anymore. Security thing?
const route = useRoute()

const username = ref('')
const isError = ref(false)
const errorMessage = ref('Something went wrong')

// Move this :hmmge:
const clientId: string = 'finitevault-ipopqvkgvm' // SimpleLogin client ID
const redirect_uri: string = 'http://localhost:5173'

// TODO: Move this also
type User = {
  name: string
  balance: number
}

// TODO: Changing this to use a token stored in local storage and updating the state if the page refreshes or something
function triggerLoginFlow(user: User) {
  userdata.logInUser(user.name, user.balance)
  console.log('Just logged in a user with ' + user.name + ' and a balance of ' + user.balance)
  router.push('home')
}

const oauthCallbackGoogle = async (response: any) => {
  if (response.credential) {
    await axios
      .post('http://localhost:8080/login_google', {
        token: response.credential
      })
      .then((res) => {
        console.log(JSON.stringify(res))
        let userInfo: User = res.data.user_info
        triggerLoginFlow(userInfo)
        isError.value = false
      })
      .catch((err) => {
        isError.value = true
        errorMessage.value = 'Google login failed: ' + err
      })
  } else {
    isError.value = true
    errorMessage.value = 'OAuth login failed.'
  }
}

function startSimpleLoginSignin() {
  console.log('Attempting login start')
  // TODO: Work out a .env here so that on cloudflare you can use the proper https://finite-vault.pages.dev/ domain
  //       Worth coming back to when the basic routes are working and spitting back some information
  let authUrl = `https://app.simplelogin.io/oauth2/authorize?response_type=code&client_id=${clientId}&redirect_uri=${redirect_uri}&scope=profile&state=${{ state: 'noidea' }}`
  location.href = authUrl
}

async function endSimpleLoginSignin(code: string) {
  console.log('Attempting simple login end')
  await axios
    .post('http://localhost:8080/login_proton', {
      token: code, // NTS: If you wanted to type these, they are inferred based on what's given. It's out of our control anyway I think?
      redirect_uri: redirect_uri
    })
    .then((res) => {
      console.log(JSON.stringify(res))
      let userInfo: User = res.data.user_info // TODO: Of course, refactor and modularize
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
})
</script>

<template>
  <div class="greetings">
    <h1 class="welcome-text">Welcome to Finite Vault!</h1>
    <img src="../assets/img/chest_kenney.png" class="logo" />
    <label for="username" class="username-box">Enter your username: </label>
    <input v-model="username" name="username" /><br />
    <p v-if="isError" class="error-text">{{ errorMessage }}</p>
    <GoogleLogin :callback="oauthCallbackGoogle" class="google-button" />
    <button @click="startSimpleLoginSignin" class=".login-button">
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
