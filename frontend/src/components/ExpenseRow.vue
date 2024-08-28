<script setup lang="ts">
defineProps({
  amount: Number,
  description: String
  // TODO: Date
})

// TODO: Lowkey everything. Trash icon?
</script>

<template>
  <div class="past-expenses-row">
    <span class="expense-past expense-past-amount"> ${{ amount?.toFixed(2) }} </span>
    <span class="expense-past expense-past-reason">
      <p>{{ description }}</p>
    </span>
  </div>
</template>

<style>
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
