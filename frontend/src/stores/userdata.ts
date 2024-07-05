import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useUserdataStore = defineStore('userdata', () => {
  // Establish state variables
  const username = ref('')
  const balance = ref(0.0)
  const dailyNumber = 4
  const expenses = ref([{ amount: 4, description: 'Nothing' }])

  // Computeds become getters
  const getBalance = computed(() => balance.value)

  const getExpenses = computed(() => expenses.value)

  // Functions become actions
  function incrementForDay() {
    balance.value += dailyNumber
  }

  function fileNewExpense(amount: number, reason: string) {
    balance.value -= amount
    expenses.value.push({ amount: amount, description: reason })
  }

  function logInUser(inputUsername: string, inputPassword: string) {
    username.value = inputUsername
    balance.value = 17.41
    console.log(`Logged in with password ${inputPassword}`)
  }

  return {
    username,
    balance,
    dailyNumber,
    getBalance,
    incrementForDay,
    logInUser,
    fileNewExpense,
    getExpenses
  }
})
