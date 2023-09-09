<template>
  <el-aside class="sidebar">
    <el-menu
        :default-active="activeFriend"
        class="sidebar-menu"
        @select="handleFriendSelect"
        background-color="#f0f5ff"
    >
    <el-menu-item v-for="friend in chattedFriends()" :index="friend" :key="friend">
      <div class="friend-item">
        <img :src="this.avatarUrl" alt="User Avatar" class="avatar" />
        <div class="friend-details">
          <div class="friend-name">{{ friend }}</div>

        </div>
        <div v-if="hasUnreadMessages(friend)" class="unread-dot"></div>
      </div>
    </el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script>
import { useChatStore } from "../stores/chat.js";

export default {
  data() {
    return {
      avatarUrl: "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
      activeFriend: "",
    };
  },
  created() {
    // 提取第二级目录
    const pathArray = this.$route.path.split("/");
    if (pathArray.length >= 4 && pathArray[2] === "chat") {
      this.activeFriend = `${pathArray[3]}`;
    }
  },
  watch: {
    $route(to, from) {
      // 提取第二级目录
      const pathArray = to.path.split("/");
      if (pathArray.length >= 4 && pathArray[2] === "chat") {
        this.activeFriend = pathArray[3];
      } else {
        this.activeFriend = "";
      }
    },
  },
  computed: {
    chattedFriends() {
      return useChatStore().getChattedFriends;
    },
  },
  methods: {
    hasUnreadMessages(friendId) {
      if (this.activeFriend === friendId) {
        useChatStore().markMessagesAsRead(friendId);
      }
      return useChatStore().hasUnreadMessages(friendId);
    },
    handleFriendSelect(index) {
      this.activeFriend = index;
      useChatStore().markMessagesAsRead(index); // 将选中的好友的消息标记为已读
      this.$emit("friend-selected", index); // 触发父组件事件，传递选中的好友 ID
    },
  },
};
</script>

<style scoped>
.sidebar {
  background-color: #ffffff; /* 侧边栏背景色 */
  padding: 20px;
}

.sidebar-menu {
  border: none; /* 移除菜单边框 */
}

.friend-item {
  display: flex;
  align-items: center;
  padding: 10px 0;
  cursor: pointer;
}

.avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  margin-right: 10px;
}

.friend-details {
  display: flex;
  flex-direction: column;
}

.friend-name {
  font-weight: bold;
}

.unread-dot {
  background-color: red;
  width: 6px;
  height: 6px;
  border-radius: 50%;
  margin-left: 5px;
}
</style>
