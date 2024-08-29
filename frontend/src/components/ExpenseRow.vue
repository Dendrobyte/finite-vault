<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps({
  amount: Number,
  description: String,
  date: Number // Unix timestamp, to convert to date object
})

// Convert timestamp on transaction to date
// The timestamp should be stored as local timestamp, so no need to convert from a central timezone
const dateString = computed(() => {
  if (props.date === undefined) {
    return 'Date Error'
  }
  const newDate = new Date(props.date * 1000)
  return newDate.toDateString()
})

// TODO: Trash icon?
</script>

<template>
  <div class="past-expense-date-container">
    <p class="past-expense-date">{{ dateString }}</p>
  </div>

  <div class="past-expenses-row">
    <span class="expense-past expense-past-amount"> ${{ amount?.toFixed(2) }} </span>
    <span class="expense-past expense-past-reason">
      <p>{{ description }}</p>
    </span>
  </div>
</template>

<style>
.past-expense-date-container {
  display: inline-block;
  margin: auto;
}

.past-expense-date {
  color: var(--core-cream);
  opacity: 60%;
  text-align: left;
  font-weight: 800;
  font-size: 1.4em;
}

.past-expenses-row {
  margin: auto;
}

.expense-past {
  margin: 0.5em;
  padding: 0.5em;
  font-size: 1.2em;
  font-weight: 800;
  display: inline-block;
  color: var(--core-cream);
  background-color: var(--core-hunter-green);
}

.expense-past-amount {
  text-align: left;
  width: 24%;
  opacity: 60%;
}

.expense-past-reason {
  text-align: left;
  width: 60%;
}

.expense-past-reason p {
  opacity: 60%;
  font-weight: 600;
}

/* TODO: Make a class for "current" expense for one that's tapped on? */
</style>
