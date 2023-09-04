<template>
  <el-container class="layout">
    <!-- 头部区域 -->
    <Header />

    <!-- 主内容区域 -->
    <el-main class="main">
      <el-container>
        <!-- 侧边栏区域 -->
        <Sidebar />

        <!-- 功能部件区域 -->
        <el-main class="widget">
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-main>
  </el-container>
</template>

<script>
import Header from '../components/Headerbar.vue';
import Sidebar from '../components/Sidebar.vue';

export default {
  components: {
    Header,
    Sidebar,
  },
};
</script>

<style scoped>

</style>
