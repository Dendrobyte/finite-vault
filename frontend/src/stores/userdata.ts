import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { type User, type UserTransaction } from '../types/User';

export const useUserdataStore = defineStore('userdata', () => {
  const localStorageKey = 'infgame_userdata'

  // Establish state variables
  const username = ref('')
  const email = ref('')
  const balance = ref(0.0)
  const dailyNumber = 4
  const expenses = ref([] as {amount: number, description: string, creation_ts: number}[])
  const loggedIn = ref(false)
  const authToken = ref('')

  // Computeds become getters
  // TODO: Make these into functions with = () ?
  const getBalance = computed(() => balance.value)

  const getExpenses = computed(() => expenses.value.reverse())

  const getAuthToken = computed(() => authToken.value);

  const isLoggedIn = computed(() => loggedIn.value);

  // Functions become actions
  function incrementForDay() { // TODO: I don't need this function here
    balance.value += dailyNumber
  }

  function fileNewExpense(amount: number, reason: string) {
    balance.value -= amount
    const timestamp = Date.now()
    expenses.value.push({ amount: amount, description: reason, creation_ts: timestamp })
  }

  function addExpense(tnx: UserTransaction) {
    expenses.value.push({amount: tnx.amount, description: tnx.description, creation_ts: tnx.creation_ts})
  }

  // Responsible for updating the state and logging in the user
  // We know at this point that OAuth login has been validated
  // TODO: (next step) -- hold on to auth token, will send on subsequent requests
  function logInUser(user: User) {
    username.value = user.username
    email.value = user.email
    balance.value = user.balance
    authToken.value = user.auth_token
    loggedIn.value = true
    return true
  }

  // Check to see if there is a JWT in storage, and if so load it in and "log in" the user
  // The isLoggedIn field will be checked later
  function loadUserFromLocalStorage() {
    // NTS: Alternatively: https://github.com/prazdevs/pinia-plugin-persistedstate
    const localUserData: User = JSON.parse(localStorage.getItem(localStorageKey) || '{}')
    if (localUserData.auth_token !== undefined) {
      logInUser(localUserData);
    }
  }

  function logOutUser() {
    username.value = ''
    email.value = ''
    balance.value = 0.0
    authToken.value = ''
    loggedIn.value = false
    localStorage.clear()
  }

  return {
    username,
    email,
    balance,
    dailyNumber,
    getBalance,
    getAuthToken,
    incrementForDay,
    logInUser,
    fileNewExpense,
    getExpenses,
    isLoggedIn,
    loadUserFromLocalStorage,
    addExpense,
    logOutUser
  }
})
