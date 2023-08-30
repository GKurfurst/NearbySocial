<template>
  <div class="login">
    <h1>Login</h1>
    <input v-model="username" placeholder="Username">
    <input v-model="password" placeholder="Password">
    <button @click="login">Login</button>
    <p>{{username}}</p>
  </div>
</template>

<script>
import loginAPI from "../api/login.js";
export default {
  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
    login() {
      loginAPI.login(this.username, this.password);
      this.$router.push('/chat');
    }
  }
};
</script>

<style>

</style>
