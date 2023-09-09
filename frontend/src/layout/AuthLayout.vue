<template>
  <div class="login">
    <div class="login-bg"></div>
    <AuthHeader />
    <el-row class="login-content">
        <router-view></router-view>
    </el-row>
  </div>
</template>

<script>
import AuthHeader from '../components/AuthHeader.vue';

export default {
  components: {
    AuthHeader,
  },
  data() {
    return {
      pageTitle: '身份验证',
    };
  },
};
</script>

<style scoped>
.login {
  position: relative;
  height: 100vh; /* 让登录页面占满整个视口高度 */
  overflow: hidden; /* 隐藏可能溢出的内容 */
}

.login-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-image: url('https://cn.bing.com/th?id=OHR.TinyHummer_EN-CN2919016235_1920x1200.jpg&rf=LaDigue_1920x1200.jpg'); /* 用你的背景图片路径替换 */
  background-size: cover;
  filter: blur(8px); /* 可选：添加模糊效果 */
  z-index: -1;
}

.login-content {
  display: flex;
  padding: 300px;
  margin-top: 40px;
  z-index: 1; /* 确保内容在背景之上 */
}

.login-aside {
  background-color: rgba(255, 255, 255, 0.9); /* 添加半透明的白色背景 */
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
}
</style>
