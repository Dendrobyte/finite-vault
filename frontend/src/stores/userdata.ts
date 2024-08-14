import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useUserdataStore = defineStore('userdata', () => {
  // Establish state variables
  const username = ref('')
  const balance = ref(0.0)
  const dailyNumber = 4
  const expenses = ref([{ amount: 4, description: 'Nothing' }])
  const loggedIn = ref(false)
  const authToken = ref('')

  // Computeds become getters
  const getBalance = computed(() => balance.value)

  const getExpenses = computed(() => expenses.value)

  const isLoggedIn = computed(() => loggedIn.value);

  const getAuthToken = computed(() => authToken.value);

  // Functions become actions
  function incrementForDay() { // TODO: I don't need this function here
    balance.value += dailyNumber
  }

  function fileNewExpense(amount: number, reason: string) {
    balance.value -= amount
    expenses.value.push({ amount: amount, description: reason })
  }

  // Responsible for updating the state and logging in the user
  // We know at this point that OAuth login has been validated
  // TODO: (next step) -- hold on to auth token, will send on subsequent requests
  function logInUser(inputUsername: string, givenBalance: number) {
    username.value = inputUsername
    balance.value = givenBalance
    loggedIn.value = true
    return true
  }

  return {
    username,
    balance,
    dailyNumber,
    getBalance,
    getAuthToken,
    incrementForDay,
    logInUser,
    fileNewExpense,
    getExpenses,
    isLoggedIn,
  }
})
