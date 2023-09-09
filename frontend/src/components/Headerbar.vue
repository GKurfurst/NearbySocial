<template>
  <el-header class="app-header">
    <span class="logo" @click="goHome">
      <img src="../assets/vue.svg" alt="Logo" />
    </span>
    <span class="user-section">
      <el-avatar class="user-avatar" :src="userAvatarUrl" size="medium" v-if="isLoggedIn"></el-avatar>
      <span class="username" v-if="isLoggedIn">{{ getUsername() }}</span>
      <el-button @click="handleUserAction" class="popover-button">{{ getActionText() }}</el-button>
    </span>
  </el-header>
</template>

<script>
import { useUserStore } from "../stores/user.js";
import { logout } from "../api/logout.js";

export default {
  data() {
    return {
      userAvatarUrl: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png', // 当前用户头像 URL
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
    handleUserAction() {
      if (this.isLoggedIn) {
        logout();
        this.$router.push('/');
      } else {
        this.openLogin();
      }
    },
    openLogin() {
      this.$router.push('/auth');
    },
    getUsername() {
      return useUserStore().getUsername;
    },
    getActionText() {
      // 根据用户登录状态返回不同的按钮文本
      return this.isLoggedIn ? '登出' : '登录';
    },
  },
};
</script>

<style scoped>
.app-header {
  background-color: #2f54eb; /* 使用你的选择的颜色 */
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
}

.logo {
  cursor: pointer;
}

.user-section {
  display: flex;
  align-items: center;
}

.user-avatar {
  margin-right: 10px;
}

.username {
  color: #fff; /* 设置用户名文本颜色 */
  margin-right: 10px;
}

.popover-button {
  background-color: transparent;
  border: none;
  color: #fff; /* 设置按钮文本颜色 */
  cursor: pointer;
}

.popover-button:hover {
  text-decoration: underline;
}
</style>
