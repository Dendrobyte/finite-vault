import './assets/main.css'

import { createPinia } from 'pinia'
import { createApp } from 'vue'
import vue3GoogleLogin from 'vue3-google-login'

import App from './App.vue'
import router from './router'

const pinia = createPinia()
const app = createApp(App)

app.use(pinia)
app.use(router)
app.use(vue3GoogleLogin, {
    clientId: '290525277072-98ucb4hr01c9rd0gutgpoih6lgdfsn31.apps.googleusercontent.com'
})

app.mount('#app')
