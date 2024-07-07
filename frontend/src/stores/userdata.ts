import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useUserdataStore = defineStore('userdata', () => {
  // Establish state variables
  const username = ref('')
  const balance = ref(0.0)
  const dailyNumber = 4
  const expenses = ref([{ amount: 4, description: 'Nothing' }])
  const loggedIn = ref(false)

  // Computeds become getters
  const getBalance = computed(() => balance.value)

  const getExpenses = computed(() => expenses.value)

  const isLoggedIn = computed(() => loggedIn.value);

  // Functions become actions
  function incrementForDay() {
    balance.value += dailyNumber
  }

  function fileNewExpense(amount: number, reason: string) {
    balance.value -= amount
    expenses.value.push({ amount: amount, description: reason })
  }

  function logInUser(inputUsername: string, givenBalance: number) {
    username.value = inputUsername
    balance.value = givenBalance
    console.log("state should be updated for these: " + username.value + " | " + username.value)
    loggedIn.value = true
  }

  return {
    username,
    balance,
    dailyNumber,
    getBalance,
    incrementForDay,
    logInUser,
    fileNewExpense,
    getExpenses,
    isLoggedIn
  }
})
