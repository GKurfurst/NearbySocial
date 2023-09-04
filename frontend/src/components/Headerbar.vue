<template>
  <el-header class="app-header">
    <span class="logo" @click="goHome">
      <img src="../assets/vue.svg" alt="Logo" />
    </span>
    <span>
      <el-avatar class="user-avatar" :src="userAvatarUrl" size="medium"></el-avatar>
    </span>
    <span class="popover-content">
      <span class="username">用户名</span>
      <el-button @click="gologout" class="popover-button" v-if="isLoggedIn">登出</el-button>
      <el-button @click="openLogin" class="popover-button" v-else>登录</el-button>
      <el-button @click="openRegister" class="popover-button" v-if="!isLoggedIn">注册</el-button>
    </span>
<!--    <el-popover placement="bottom-end" width="150" trigger="hover">-->
<!--      <template #reference>-->
<!--        <el-avatar class="user-avatar" :src="userAvatarUrl" size="medium"></el-avatar>-->
<!--      </template>-->
<!--      <div class="popover-content">-->
<!--        <div class="username">用户名</div>-->
<!--        <el-button @click="openLoginDialog" class="popover-button">登录</el-button>-->
<!--        <el-button @click="openRegisterDialog" class="popover-button">注册</el-button>-->
<!--        <el-button @click="logout" class="popover-button">登出</el-button>-->
<!--      </div>-->
<!--    </el-popover>-->
  </el-header>
</template>

<script>
import {useUserStore} from "../stores/user.js";
import {logout} from "../api/logout.js";
export default {
  data() {
    return {
      userAvatarUrl: 'https://element-plus.org/images/element-plus-logo.svg', // 当前用户头像 URL
    };
  },
  computed: {
    isLoggedIn() {
      // 返回当前用户是否已经登录
      return useUserStore().isLoggedIn;
    },
  },
  methods: {
    goHome() {
      // 执行返回首页逻辑
      this.$router.push('/');
    },
    openLogin() {
      this.$router.push('/auth');
    },
    openRegister() {
      this.$router.push('/auth/register');
    },
    gologout() {
      logout()
      this.$router.push('/');
    },
  },
};
</script>

<style scoped>

</style>
