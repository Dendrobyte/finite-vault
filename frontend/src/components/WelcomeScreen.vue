<script setup lang="ts">
import router from '@/router'
import { useUserdataStore } from '@/stores/userdata'
import axios from 'axios'
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { type User, type ValidatedUser } from '../types/User'

const userdata = useUserdataStore()
const route = useRoute()

const isError = ref(false)
const errorMessage = ref('Something went wrong')

const CLIENT_ID: string = import.meta.env.VITE_SIMPLE_LOGIN_CLIENT_ID // SimpleLogin client ID
const REDIRECT_URI: string = import.meta.env.VITE_ROOT_URI
const BACKEND_URI: string = import.meta.env.VITE_BACKEND_URI

function triggerLoginFlow(user: User) {
  // Make sure it was a successful login
  // TODO: Send along timestamp in the response of a request here, triggers update function in the store
  // TODO: Start a loading icon here or something until we push them to 'home'
  let loginResult: boolean = userdata.logInUser(user)
  if (loginResult) {
    localStorage.setItem('infgame_userdata', JSON.stringify(user))
    router.push('home')
  }
  return loginResult
}

/* OAuth login functions */
const oauthCallbackGoogle = async (response: any) => {
  if (response.credential) {
    await axios
      .post(`${BACKEND_URI}/login/google`, {
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
    .post(
      `${BACKEND_URI}/login/proton`,
      {
        token: code,
        redirect_uri: REDIRECT_URI
      },
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        }
      }
    )
    .then((res) => {
      let userInfo: User = res.data
      triggerLoginFlow(userInfo)
      isError.value = false
    })
    .catch((err) => {
      isError.value = true
      errorMessage.value = 'Proton login failed: ' + err
    })
}

async function validateToken(token: string): Promise<[valid: boolean, email: string]> {
  let isValid: boolean = false
  let email: string = ''
  await axios
    .post(
      `${BACKEND_URI}/validateToken`,
      {
        auth_token: token
      },
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        }
      }
    )
    .then((res) => {
      if (res.status == 422 || res.status == 401) {
        // TODO: Flesh out the logic here based on flawed backend processing
        console.log('Backend was hit but returned a 400 error')
      }
      if (res.status != 200) {
        return
      }
      let validated: ValidatedUser = res.data
      isValid = validated.valid
      email = validated.email
    })
    .catch((err) => {
      isError.value = true
      errorMessage.value = 'Token validation failed: ' + err
    })
  return [isValid, email]
}

onMounted(async () => {
  // SimpleLogin redirect uri set to the same page, so check for a code query
  // TODO: If there's time, make the thing a popup window and make a "mid" component to handle the redirection
  if (route.query.code) {
    endSimpleLoginSignin(String(route.query.code))
  }

  // On the state, check if we have a token in local storage
  userdata.loadUserFromLocalStorage()

  // If they are logged in, go ahead and get that token, validate, etc.
  if (userdata.isLoggedIn == true) {
    let token: string = userdata.getAuthToken
    // TODO: Hit backend to check validity of token
    let [isValid, userEmail] = await validateToken(token)
    if (isValid == true) {
      if (userdata.email == userEmail) {
        router.push('home')
      } else {
        errorMessage.value = 'Token does not match correct email on server. Please log in again.'
        localStorage.clear()
      }
    } else {
      errorMessage.value = 'Token is not valid, continuing to sign-in page.'
      localStorage.clear() // Clear everything if token invalidated
      // TODO: Redirect if expired, user friendly thing if anything
    }
  }

  // Else, load as normal
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
