<template>
  <div class="text-center homepage h-100 mt-5 display-3">
    <div v-for="user in users" :key="user.id">
      {{ user }}
    </div>
  </div>
</template>

<script setup>
const { data: users } = await useFetch('http://localhost:4000/users')
</script>