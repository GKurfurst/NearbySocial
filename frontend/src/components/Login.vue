<template>
  <div class="login">
    <h1>Login</h1>
    <input v-model="username" placeholder="Username">
    <button @click="login">Login</button>
    <p>{{username}}</p>
  </div>
</template>

<script>

export default {
  data() {
    return {
      username: ''
    };
  },
  methods: {
    login() {
      // 在这里处理登录逻辑，例如将用户名存储在状态管理中
      // 然后导航到聊天页面
      this.$router.push('/chat');
    }
  }
};
</script>

<style>
.login {
  text-align: center;
  margin-top: 100px;
}
</style>
