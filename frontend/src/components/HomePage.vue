<script setup lang="ts">
import ExpenseRow from '@/components/ExpenseRow.vue'
import router from '@/router'
import { useUserdataStore } from '@/stores/userdata'
import { onMounted, ref } from 'vue'

// Pull username from state
const userdata = useUserdataStore()

const expenseAmount = ref(0.0)
const expenseReason = ref('')
const isValidExpenseAmount = ref(true)
const isFormInvalid = ref(true)

// Modify presentation when getting, update value number when setting
const expenseAmountStr = ref('')

// Every time we update the input, we need to target the value of the specific html element
// to modify the data there. The variable in our "state" is updated separately.
function validateExpenseAmount(event: any) {
  const value = event.target.value

  const isValid = isValidExpenseStr(value)
  if (isValid) {
    expenseAmount.value = getExpenseFromExpenseStr(value)
    expenseAmountStr.value = `$${expenseAmount.value}`
    isValidExpenseAmount.value = true
  } else {
    event.target.value = value.slice(0, value.length - 1)
    isValidExpenseAmount.value = false
  }
}

// Check to see if a string is a valid "expense string"
function isValidExpenseStr(input: any): boolean {
  // Check for empty string
  if (input.length == 0) {
    expenseAmount.value = 0.0
    return false
  }

  // Remove '$' if present
  if (input.charAt(0) === '$') {
    input = input.slice(1, input.length)
  }

  // Check to make sure we don't exceed two decimal places
  if (input.includes('.')) {
    if (input.split('.')[1].length > 2) {
      return false
    }
  }

  // And lastly just check that it's a number
  return !isNaN(input)
}

// Given an already validated string, let's remove the dollar sign and parse it as a float for the expense
function getExpenseFromExpenseStr(expenseStr: string): number {
  expenseStr = expenseStr.slice(1, expenseStr.length)
  if (expenseStr.length === 0) {
    return 0.0
  }
  return parseFloat(expenseStr)
}

function submitExpense() {
  // TODO: Don't forget to validate expense reason (and revalidate final amount, THEN show an error)
  userdata.fileNewExpense(expenseAmount.value, expenseReason.value)
  expenseAmount.value = 0.0
  expenseAmountStr.value = ''
  expenseReason.value = ''
}

// Simple check to make sure we can submit the form
function canSubmitForm(): boolean {
  return isValidExpenseAmount.value && expenseReason.value.length > 0
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
    <h2 class="vault-balance-text">${{ userdata.balance.toFixed(2) }}</h2>
  </div>
  <br />

  <div class="expense-page-section">
    <h2 class="vault-title-text">New Expense</h2>
    <form v-on:submit.prevent="submitExpense()" action="" method="post" class="new-expense-form">
      <input
        name="expenseAmount"
        placeholder="$0.00"
        :value="expenseAmountStr"
        @input="validateExpenseAmount"
        class="expense-form-amount"
      />
      <input
        name="expenseReason"
        placeholder="Your expense reason here..."
        v-model.trim="expenseReason"
        class="expense-form-reason"
      /><br />
      <p v-if="!isFormInvalid" class="error-text">
        {{ 'Expense amount or reason is invalid.' }}
      </p>
      <button
        class="expense-form-button"
        :class="[canSubmitForm() ? 'btn-active' : 'btn-inactive']"
      >
        Submit
      </button>
    </form>
  </div>

  <div class="expense-page-section">
    <h2 class="vault-title-text" style="font-size: 4em">Past Expenses</h2>
    <ExpenseRow
      v-for="(expense, idx) in userdata.getExpenses"
      :key="idx"
      :amount="expense.amount"
      :description="expense.description"
    />
  </div>
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

.expense-page-section {
  display: inline;
}

.new-expense-form {
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
  width: 24%;
  color: var(--core-ecru);
  background-color: var(--core-field-drab);
}

.expense-form-reason {
  width: 60%;
  color: var(--core-field-drab);
  background-color: var(--core-ecru);
}

.expense-form-button {
  font-size: 2em;
  margin: 0.8em;
  padding: 0.5em 1em;
  border-radius: 0;
  border: none;
  font-weight: 600;
  color: var(--white-soft);
  background-color: var(--core-ecru);
}

.expense-form-button:hover {
  cursor: pointer;
}

.expense-form-button:active {
  position: relative;
  top: 0.1em;
}

.btn-inactive {
  opacity: 40%;
  pointer-events: none;
}

.btn-active {
  opacity: 100%;
}
</style>
