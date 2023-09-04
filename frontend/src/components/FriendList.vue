<template>
  <div class="friend-list">
    <FriendListSidebar
        @acceptRequest="acceptRequest"
        @rejectRequest="rejectRequest"
        @startChat="startChat"
        @removeFriend="removeFriend"
    />
    <div class="friend-details" v-if="selectedFriend">
      <h2>好友详情</h2>
      <p>好友姓名: {{ selectedFriend.name }}</p>
      <!-- 其他好友信息 -->
    </div>
  </div>
</template>

<script>
import FriendListSidebar from "./FriendListSidebar.vue";

export default {
  components: {
    FriendListSidebar,
  },
  data() {
    return {
      selectedFriend: null,
    };
  },
  methods: {

  },
};
</script>

<style scoped>
/* 样式可以根据需要进行调整 */
.friend-list {
  display: flex;
  flex-direction: row;
}

.friend-details {
  flex: 1;
  padding: 20px;
}

.friend-details h2 {
  margin-bottom: 10px;
}
</style>
