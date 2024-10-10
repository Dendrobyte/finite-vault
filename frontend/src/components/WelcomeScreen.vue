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
          'Access-Control-Allow-Origin': '*',
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
          'Access-Control-Allow-Origin': '*',
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
    <h1 class="welcome-text">Finite Vault</h1>
    <img src="../assets/img/chest_kenney.png" class="logo" />
    <p v-if="isError" class="error-text">{{ errorMessage }}</p>
    <div class="login-buttons">
      <button @click="startSimpleLoginSignin" class="login-button">Sign in with Proton</button>
      <!-- <span class="login-button"></span> -->
      <GoogleLogin :callback="oauthCallbackGoogle" class="google-button-wrapper">
        <button class="login-button">Sign in with Google</button>
      </GoogleLogin>
    </div>
  </div>
</template>

<style>
@import '../assets/base.css';

.welcome-text {
  color: var(--core-cream);
  font-weight: 600;
  padding: 0.5em;
  font-size: 6em;
}

.logo {
  width: 80%;
  margin: 0px auto;
  display: block;
}

.login-buttons {
  display: block;
}

.login-button {
  background-color: var(--core-hunter-green);
  color: var(--core-cream);
  width: 60%;
  border: none;
  padding: 1em;
  margin: 0.5em auto;
  font-weight: 800;
  text-align: center;
  text-decoration: none;
  font-size: 2em;
  display: block;
}

.login-button:hover {
  cursor: pointer;
}

/* Use relative position to avoid influencing document flow */
.login-button:active {
  position: relative;
  top: 0.2em;
}

.google-button-wrapper {
  width: 100%;
}

.error-text {
  color: red;
  font-size: 20px;
}
</style>
