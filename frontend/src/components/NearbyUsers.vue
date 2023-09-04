<template>
  <div>
    <NearbyUsersSidebar @fetch-nearby-users="fetchNearbyUsers" />
    <div class="user-list">
      <el-card
          v-for="user in nearbyUsers"
          :key="user.user.UserId"
          :header="user.user.Name"
          shadow="hover"
          style="width: 300px; margin: 10px;"
      >
        <!-- 在这里显示用户信息，可以根据你的数据结构进行修改 -->
        <p>用户 ID: {{ user.user.UserId }}</p>
        <p>距离: {{ Math.round(user.distance) }}米</p>
        <!-- 其他用户信息 -->

        <!-- 添加好友按钮 -->
        <el-button
            v-if="!user.friendRequestSent"
            type="primary"
            @click="sendFriendRequest(user.user.UserId, user)"
        >
          添加好友
        </el-button>
        <el-tag
            v-else
            type="success"
        >
          已发送好友请求
        </el-tag>
      </el-card>
    </div>
  </div>
</template>

<script>
import NearbyUsersSidebar from "./NearbyUsersSidebar.vue";
import { handlePossibleToken } from "../utils/token.js";
import { getNearbyUser } from "../api/getNearbyUser.js";
import { sendFriendRequest } from "../api/friend.js";
import {useUserStore} from "../stores/user.js";

export default {
  data() {
    return {
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
