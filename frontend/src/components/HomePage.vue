<script setup lang="ts">
import ExpenseRow from '@/components/ExpenseRow.vue'
import router from '@/router'
import { useUserdataStore } from '@/stores/userdata'
import { computed, onMounted, ref } from 'vue'

// Pull username from state
const userdata = useUserdataStore()

const expenseAmount = ref(0.0)
const expenseReason = ref('')
const isValidExpenseAmount = ref(true)

// Modify presentation when getting, update value number when setting
const expenseAmountStr = computed({
  get() {
    return `$${expenseAmount.value}`
  },
  set(newVal) {
    console.log('New val: ' + newVal)
    if (newVal !== undefined) {
      if (newVal.charAt(0) === '$') {
        newVal = newVal.slice(1, newVal.length)
      }

      if (isValidExpenseStr(newVal)) {
        isValidExpenseAmount.value = true
        expenseAmount.value = parseFloat(newVal)
        console.log('Valid number : ' + newVal)
      } else {
        isValidExpenseAmount.value = false
        expenseAmount.value = 0.0
      }
    } else {
      isValidExpenseAmount.value = false
      expenseAmount.value = 0.0
    }
  }
})

// Check to see if the expense string input by the user is valid, input stripped of currency symbol
function isValidExpenseStr(input: any): boolean {
  // Check to make sure we don't exceed two decimal places
  if (input.includes('.')) {
    if (input.split('.')[1].length > 2) {
      return false
    }
  }

  // And lastly just check that it's a number
  return !isNaN(input)
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
        class="expense-form-amount"
      />
      <p v-if="!isValidExpenseAmount" class="error-text">
        {{
          'Invalid expense amount, please ensure it is all numbers and has no more than two decimal places.'
        }}
      </p>
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
