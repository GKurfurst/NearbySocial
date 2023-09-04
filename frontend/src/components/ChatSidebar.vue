<template>
  <el-aside class="sidebar">
    <el-menu
        :default-active="activeFriend"
        class="sidebar-menu"
        @select="handleFriendSelect"
    >
      <el-menu-item v-for="friend of chattedFriends" :index="friend" :key="friend">
        {{ friend }}
        <span v-if="hasUnreadMessages(friend)" class="unread-dot"></span>
      </el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script>
import {useChatStore} from "../stores/chat.js";

export default {
  data() {
    return {
      activeFriend: "",
    };
  },
  computed: {
    chattedFriends() {
      return useChatStore().getChattedFriends();
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
.unread-dot {
  background-color: red;
  width: 6px;
  height: 6px;
  border-radius: 50%;
  display: inline-block;
  margin-left: 5px;
}
</style>
