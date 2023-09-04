<template>
  <div class="find-user">
    <FindUsersSidebar @user-selected="showUserDetails" />
    <div class="user-details" v-if="selectedUser">
      <h2>User Details</h2>
      <p>Username: {{ selectedUser.Name }}</p>
      <!-- 其他用户信息 -->
      <el-button
          v-if="!friendRequestSent"
      @click="sendFriendRequest"
      >
      申请成为好友
      </el-button>
      <el-tag
          v-else
      type="success"
      >
      已发送好友请求
      </el-tag>
    </div>
  </div>
</template>

<script>
import FindUsersSidebar from "./FindUsersSidebar.vue";
import { sendFriendRequest } from "../api/friend.js";
import {useUserStore} from "../stores/user.js";
import {handlePossibleToken} from "../utils/token.js"; // 导入发送好友请求的函数

export default {
  data() {
    return {
      selectedUser: {},
      friendRequestSent: false,
    };
  },
  components: {
    FindUsersSidebar,
  },
  methods: {
    showUserDetails(user) {
      // 当用户在侧边栏中选择一个用户时，更新选定的用户
      this.selectedUser = user;
      this.friendRequestSent = false;
    },
    async sendFriendRequest() {
      try {
        // 调用发送好友请求的函数，传递当前用户和选中用户的ID
        const response = handlePossibleToken(await sendFriendRequest(useUserStore().getUserId, this.selectedUser.UserId));
        this.friendRequestSent = true;
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
/* 样式可以根据需要进行调整 */
.find-user {
  display: flex;
  flex-direction: row;
}

.user-details {
  flex: 1;
  padding: 20px;
}

.user-details h2 {
  margin-bottom: 10px;
}
</style>
