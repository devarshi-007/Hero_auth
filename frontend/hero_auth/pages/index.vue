<template>
    <div class="login-page" v-if="!isLogin">
        <div class="login-container">
            <!-- <form @submit="login" action="/" class="login-form"> -->
            <h2>Login</h2>
            <div class="form-group">
                <label for="username">Username</label>
                <input type="text" id="username" v-model="username" required>
            </div>
            <div class="form-group">
                <label for="password">Password</label>
                <input type="password" id="password" v-model="password" required>
            </div>
            <button type="submit" @click="login">Login</button>
            <!-- </form> -->
        </div>
    </div>
</template>
  
  
<style scoped>
.login-page {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #f0f0f0;
    /* Set your background color */
}

.login-container {
    background: #fff;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    border-radius: 4px;
    padding: 20px;
}

.login-form {
    display: flex;
    flex-direction: column;
    align-items: center;
    max-width: 300px;
    margin: 0 auto;
    text-align: center;
}

h2 {
    color: #333;
    /* Set your title color */
    font-size: 24px;
    margin: 0 0 20px;
}

.form-group {
    margin: 10px 0;
}

label {
    font-weight: bold;
    font-size: 16px;
    display: block;
    margin-bottom: 6px;
}

input {
    width: 100%;
    padding: 10px;
    margin: 5px 0;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 16px;
}

button {
    background: #007bff;
    /* Set your button background color */
    color: #fff;
    border: none;
    padding: 10px 20px;
    cursor: pointer;
    font-size: 16px;
    margin-top: 10px;
}

button:hover {
    background: #0056b3;
    /* Set your button hover color */
}
</style>
  
  

<script setup>
definePageMeta({
    layout: false
})
import { ref, reactive } from 'vue';
const username = ref('');
const password = ref('');

async function login() {
    console.log("adfadfasdfdafdafafa")
    //     $fetch(`http://localhost:4000/login`, {
    //     method: `GET`,
    //     body: {
    //             "userName": username.value,
    //             "password": password.value
    //     },
    //     onResponse: ({ response }) => {
    //       console.log(response._data)
    //     },
    //   }).catch(function () {
    //     console.log("errpr")
    //   })
    $fetch(`http://localhost:4000/login`, {
        method: `POST`,
        body: JSON.stringify({
            "userName": username.value,
            "password": password.value
        }),
        onResponse: ({ response }) => {
            if (response.status === 200) {
                console.log(response._data)
            }
        },
    }).catch(function () {
        console.log("errrrr")
    })
}
</script>