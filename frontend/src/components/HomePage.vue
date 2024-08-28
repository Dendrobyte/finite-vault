<script setup lang="ts">
import ExpenseRow from '@/components/ExpenseRow.vue'
import router from '@/router'
import { useUserdataStore } from '@/stores/userdata'
import { onMounted, ref } from 'vue'

// Pull username from state
const userdata = useUserdataStore()

const expenseAmount = ref(0.0)
const expenseReason = ref('')

function fileExpense() {
  userdata.fileNewExpense(expenseAmount.value, expenseReason.value)
  expenseAmount.value = 0.0
  expenseReason.value = ''
}

onMounted(() => {
  // TODO: Redirection is OK bc we'll always redirect on every page refresh... Do I need this everywhere?
  //       Worth thinking this through in terms of how to handle each page load. Then again, it's a SPA so this should be in the parent
  // TODO: Move this to the parent / move the login update to the parent
  if (!userdata.isLoggedIn) {
    console.log('No login session present!')
    router.push('/')
  }
})

/* DEVELOPER BUTTON CALLBACK FUNCTIONS */
// These are used to "simulate" certain behavior. Should be deleted once everything is confirmed.
</script>

<template>
  <!-- START OF DEVELOPER BUTTONS -->

  <button @click="userdata.incrementForDay()">Increment Bal by {{ userdata.dailyNumber }}</button>

  <!-- END OF DEVELOPER BUTTONS -->
  <p style="color: white">Settings will go here</p>

  <h1 class="vault-title-text">{{ userdata.username }}'s Vault</h1>
  <div class="vault-balance">
    <h2 class="vault-balance-text">${{ userdata.balance }}</h2>
  </div>
  <br />

  <label for="expenseAmount">New Expense Cost: </label>
  <input name="expenseAmount" placeholder="0.00" v-model="expenseAmount" /><br />
  <label for="expenseReason">Expense Reason: </label>
  <input name="expenseAmount" placeholder="Groceries" v-model="expenseReason" /><br />
  <button @click="fileExpense()">File Expense</button><br />

  <ExpenseRow
    v-for="(expense, idx) in userdata.getExpenses"
    :key="idx"
    :amount="expense.amount"
    :description="expense.description"
  />
</template>

<style>
.vault-title-text {
  color: var(--core-cream);
  font-weight: 600;
  font-size: 5em;
}

/* TODO: Lazy load this image, or turn to webp */
.vault-balance {
  padding: 3em;
  background-image: url('../assets/img/balance-bg-stroke.webp');
  background-size: contain;
  background-position: center, top;
  background-repeat: no-repeat;
}

.vault-balance-text {
  color: var(--core-cream);
  font-weight: 800;
  font-size: 6em;
}
</style>
