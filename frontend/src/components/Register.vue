<template>
  <h1>Register</h1>
  <input v-model="username" placeholder="Username">
  <input v-model="telephone" placeholder="Telephone">
  <input v-model="password" placeholder="Password">
  <button @click="Register">Register</button>
  <p>{{username}}</p>
</template>

<script>
import registerService from "../service/registerService.js"
export default {
  data() {
    return {
      username: '',
      telephone: '',
      password: ''
    };
  },
  methods: {
    Register() {
      registerService.register(this.username, this.telephone, this.password);
      this.$router.push('/login');
    }
  }
};
</script>
