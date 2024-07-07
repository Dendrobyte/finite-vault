import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useSimpleLoginStore = defineStore('simplelogin', () => {
  // Establish state variables
  const accessToken = ref('')
  return {
    accessToken
  }
})
