<template>
  <el-aside class="sidebar">
    <el-menu :default-active="activeMenu" class="sidebar-menu" @select="handleMenuSelect" background-color="#333" text-color="white" active-text-color="white">
      <el-menu-item index="/nearby-users">附近用户</el-menu-item>
      <el-menu-item index="/find-users">查找用户</el-menu-item>
      <el-menu-item index="/friend-list">好友列表</el-menu-item>
      <el-menu-item index="/chat/debug">聊天</el-menu-item>
    </el-menu>


  </el-aside>
</template>

<script>
import ChatSidebar from "./ChatSidebar.vue";
export default {
  data() {
    return {
      activeMenu: '', // 当前选中的菜单项
    };
  },
  // watch: {
  //   $route(to) {
  //     const routeMapping = {
  //       '/find-users': 'FindUsers',
  //       '/friend-list': 'FriendList',
  //       '/chat': 'Chat',
  //     };
  //     this.selectedComponent = routeMapping[to.path];
  //   },
  // },
  methods: {
    handleMenuSelect(index) {
      this.activeMenu = index;
      this.$router.push("/main"+index); // 跳转到选中的路由
    },
  },
};
</script>

<style scoped>

</style>
