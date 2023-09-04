<template>
  <div class="login">
    <AuthHeader :title="pageTitle" />
    <el-row class="login-content">
      <el-col :span="16" class="login-aside">
        <router-view></router-view>
      </el-col>
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
      pageTitle: 'Auth',
    };
  },
};
</script>

<style scoped>
.login {
  padding: 40px;
}

.login-content {
  margin-top: 40px;
}

.login-aside {
  background-color: #f5f7fa;
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
}
</style>
