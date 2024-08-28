<script setup lang="ts">
import ExpenseRow from '@/components/ExpenseRow.vue'
import router from '@/router'
import { useUserdataStore } from '@/stores/userdata'
import { computed, onMounted, ref } from 'vue'

// Pull username from state
const userdata = useUserdataStore()

const expenseAmount = ref(0.01)
const expenseReason = ref('')

// Modify presentation when getting and parse numeric value when setting
const expenseAmountStr = computed({
  get() {
    return `$${expenseAmount.value}`
  },
  set(newVal) {
    const numericValue = parseExpenseAmount(newVal)
    if (numericValue !== null) {
      expenseAmount.value = numericValue
    }
  }
})

// Is called when the field updates, modifying the html target directly
function validateExpenseAmount(event: any) {
  const value = event.target.value

  const numericValue = parseExpenseAmount(value)
  if (numericValue !== null) {
    event.target.value = `$${numericValue}`
  } else {
    event.target.value = expenseAmountStr.value
  }
}

function parseExpenseAmount(amountStr: string): number | null {
  // Remove non-numeric characters except the decimal point
  const cleanStr = amountStr.replace(/[^0-9.]/g, '')

  // Convert to a float
  const numericValue = parseFloat(cleanStr)

  // Validate the number
  if (isNaN(numericValue)) {
    return null
  }

  return numericValue
}

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

  <div class="new-expense-section">
    <h2 class="vault-title-text">New Expense</h2>
    <form v-on:submit.prevent="fileExpense()" action="" method="post" class="new-expense-form">
      <input
        name="expenseAmount"
        placeholder="$0.00"
        v-model="expenseAmountStr"
        @input="validateExpenseAmount"
        class="expense-form-amount"
      />
      <input
        name="expenseReason"
        placeholder="Your expense reason here..."
        v-model.trim="expenseReason"
        class="expense-form-reason"
      /><br />
      <button @click="fileExpense()" class="new-expense-submit">Submit</button>
    </form>
  </div>

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

.new-expense-section {
  display: inline;
}

.new-expense-form {
  background-color: red;
  display: inline;
}

.new-expense-form input {
  margin: 0.5em;
  padding: 0.5em;
  border: none;
  font-size: 1.2em;
  font-weight: 600;
}

.expense-form-amount {
  width: 16%;
  color: var(--core-ecru);
  background-color: var(--core-field-drab);
}

.expense-form-reason {
  width: 62%;
}
</style>
