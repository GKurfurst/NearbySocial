<template>
  <el-aside class="main-sidebar">
    <el-menu
        :default-active="activeMenu"
        class="sidebar-menu"
        @select="handleMenuSelect"
        background-color="#333"
        text-color="white"
        active-text-color="#ffd04b"
    >
      <el-menu-item index="/nearby-users">
        <el-icon><Compass /></el-icon>
        附近用户
      </el-menu-item>
      <el-menu-item index="/find-users">
        <el-icon><Search /></el-icon>
        查找用户
      </el-menu-item>
      <el-menu-item index="/friend-list">
        <el-icon><User /></el-icon>
        好友列表
      </el-menu-item>
      <el-menu-item index="/chat" v-if=chattedFriends()[0]>
        <el-icon><ChatDotRound /></el-icon>
        聊天
      </el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script>
import ChatSidebar from "./ChatSidebar.vue";
import {ChatDotRound, Compass, Search, User} from "@element-plus/icons-vue";
import {useChatStore} from "../stores/chat.js";
export default {
  components: {ChatDotRound, User, Search, Compass},
  data() {
    return {
      activeMenu: "", // 当前选中的菜单项
    };
  },
  computed: {
    chattedFriends() {
      return useChatStore().getChattedFriends;
    },
  },
  watch: {
    $route(to, from) {
      // 提取第二级目录
      const pathArray = to.path.split("/");
      if (pathArray.length >= 3) {
        this.activeMenu = `/${pathArray[2]}`;
      } else {
        this.activeMenu = "";
      }
    },
  },
  methods: {
    handleMenuSelect(index) {
      this.activeMenu = index;
      if (index === "/chat") {
        this.$router.push("/main/chat/" + useChatStore().getChattedFriends()[0]);
      }
      else {
        this.$router.push("/main" + index);
      }
    },
  },
};
</script>

<style scoped>
.main-sidebar {
  width: 250px;
  background-color: #333; /* 使用你的选择的颜色 */
  color: white;
  position: fixed;
  top: 60px; /* 跟随顶部导航栏下方，根据你的实际需求调整 */
  bottom: 0;
  overflow-y: auto; /* 如果内容溢出，显示滚动条 */
}

.sidebar-menu {
  padding-top: 20px;
}

.el-menu-item {
  font-size: 16px; /* 调整菜单项的字体大小 */
}

.el-menu-item:hover {
  background-color: #2f54eb; /* 调整菜单项的鼠标悬停背景色 */
}

.el-menu-item i {
  margin-right: 10px;
}

.el-menu-item.is-active {
  background-color: #2f54eb; /* 选中状态的背景色 */
}

</style>
