<template>
  <div class="text-center homepage h-100 mt-5 display-3">
    <div class="grid grid-cols-7">
      <div v-for="user in users" :key="user.id" class="border border-gray-200 rounded-lg shadow m-3 p-1 bg-gray-100">
        <span>User_name: {{ user.userName }}</span><br />
        <span>Name : {{ user.name }}</span>
      </div>
    </div>
</div></template>

<script setup>
const users = ref('')
$fetch(`http://localhost:4000/users`, {
        method: `GET`,
        onResponse: ({ response }) => {
            if (response.status === 200) {
                users.value = response._data
            }
        },
        credentials: 'include'
    }).catch(function () {
        console.log("errrrr")
    })
console.log(users.value)
</script>