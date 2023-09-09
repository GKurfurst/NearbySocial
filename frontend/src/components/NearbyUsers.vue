<template>
  <div>
    <NearbyUsersSidebar @fetch-nearby-users="fetchNearbyUsers" />
    <div class="user-list">
      <el-card
          v-for="user in nearbyUsers"
          :key="user.user.UserId"
          class="user-card"
          shadow="hover"
      >
        <div slot="header" class="user-card-header">
          <img :src="this.avatarUrl" alt="User Avatar" class="user-avatar" />
          {{ user.user.Name }}
        </div>
        <div class="user-info">
          <p>用户号码: {{ user.user.Telephone }}</p>
          <p>距离: {{ Math.round(user.distance) }} 米</p>
          <!-- 其他用户信息 -->
        </div>
        <div class="user-actions">
          <el-button
              v-if="!user.friendRequestSent"
              type="primary"
              round
              @click="sendFriendRequest(user.user.UserId, user)"
          >
            添加好友
          </el-button>
          <el-tag v-else type="success">已发送好友请求</el-tag>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script>
import NearbyUsersSidebar from "./NearbyUsersSidebar.vue";
import { handlePossibleToken } from "../utils/token.js";
import { getNearbyUser } from "../api/getNearbyUser.js";
import { sendFriendRequest } from "../api/friend.js";
import { useUserStore } from "../stores/user.js";

export default {
  data() {
    return {
      avatarUrl: "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
      nearbyUsers: [], // 存储附近用户的数组
    };
  },
  components: {
    NearbyUsersSidebar,
  },
  methods: {
    async fetchNearbyUsers(range) {
      const response = handlePossibleToken(await getNearbyUser(range));
      this.nearbyUsers = response["nearby users"];
    },
    async sendFriendRequest(friendId, user) {
      try {
        // 调用发送好友请求的函数，传递当前用户和选中用户的ID
        const response = handlePossibleToken(await sendFriendRequest(useUserStore().getUserId, friendId));
        user.friendRequestSent = true; // 将该用户的 friendRequestSent 设置为 true
        this.$message.success("好友请求已发送");
      } catch (error) {
        this.$message.error(`发送好友请求失败`);
        console.log(error);
      }
    },
  },
};
</script>

<style scoped>
.user-list {
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-start;
  gap: 10px;
}

.user-card {
  width: 300px;
  margin-bottom: 10px;
}

.user-card-header {
  display: flex;
  align-items: center;
  font-weight: bold;
  font-size: 18px;
}

.user-avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  margin-right: 10px;
}

.user-info {
  margin-top: 10px;
  font-size: 16px;
}

.user-actions {
  margin-top: 10px;
  display: flex;
  justify-content: flex-end;
}
</style>
