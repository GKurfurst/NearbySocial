<template>
  <el-container class="home-container">
    <Header class="app-header"></Header>
    <el-container class="main-container">
      <Sidebar class="app-sidebar"></Sidebar>
      <el-main class="app-main">
        <router-view></router-view>
      </el-main>
    </el-container>
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
.home-container {
  height: 100%;
}

.app-header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000; /* 确保导航栏在最上层 */
}

.main-container {
  position: fixed;
  top: 60px; /* 跟随顶部导航栏下方 */
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
}

.app-sidebar {
  width: 250px;
  height: 100%;
  background-color: #333; /* 侧边栏背景色 */
  color: white;
  overflow-y: auto;
}

.app-main {
  flex: 1;
  height: 100%;
  padding: 20px; /* 调整主内容区的内边距 */
  margin-left: 250px; /* 侧边栏宽度 */
}
</style>
