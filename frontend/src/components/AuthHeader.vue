<template>
  <el-row class="login-header" style="position: fixed; top: 0; left: 0; right: 0; background-color: #fff; z-index: 1000;">
    <el-col :span="8">
      <el-image @click="goHome" class="logo" src="https://element-plus.org/images/element-plus-logo.svg" fit="contain" />
    </el-col>
  </el-row>
</template>

<script>
export default {
  methods: {
    goHome() {
      this.$router.push('/');
    },
  },
};
</script>

<style scoped>
.login-header {
  height: 80px; /* 调整导航栏高度 */
  box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.1); /* 添加阴影效果 */
  display: flex;
  align-items: center;
  padding: 0 20px;
  background-color: #fff;
}

.logo {
  max-width: 80px;
  cursor: pointer;
}

.title {
  font-size: 24px;
  color: #333;
  margin-left: 20px;
}
</style>
