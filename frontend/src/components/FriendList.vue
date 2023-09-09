<template>
  <div class="friend-list">
    <FriendListSidebar />
    <div class="friend-details" v-if="selectedFriend">
      <h2>好友详情</h2>
      <p>好友姓名: {{ selectedFriend.name }}</p>
      <!-- 其他好友信息 -->
    </div>
  </div>
</template>

<script>
import FriendListSidebar from './FriendListSidebar.vue';

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
    // 保留原有的方法
  },
};
</script>

<style scoped>

/* 根据需要进行其他样式微调 */
</style>
