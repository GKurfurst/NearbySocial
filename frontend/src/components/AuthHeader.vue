<template>
  <el-row class="login-header">
    <el-col :span="8">
      <el-image @click="goHome" class="logo" src="https://element-plus.org/images/element-plus-logo.svg" fit="contain" />
    </el-col>
    <el-col :span="16">
      <h1 class="title">{{ title }}</h1>
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
  props: {
    title: {
      type: String,
      required: true,
    },
  },
};
</script>

<style scoped>
.logo {
  max-width: 80px;
}

.title {
  font-size: 24px;
  color: #333;
  margin-left: 20px;
}
</style>
