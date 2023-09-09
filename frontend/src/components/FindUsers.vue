<template>
  <div class="find-user">
    <FindUsersSidebar @user-selected="showUserDetails" />
    <div class="user-details" v-if="selectedUser">
      <el-card class="user-card" v-if="selectedUser.Name">
        <div slot="header" class="user-card-header">
          <img
              :src=this.avatarUrl
              alt="User Avatar"
              class="user-avatar"
          />
          <span class="user-name">{{ selectedUser.Name }}</span>
        </div>
        <div class="user-info">
          <p>电话号码: {{ selectedUser.Telephone }}</p>
        </div>
        <div class="user-actions">
          <el-button
              v-if="!friendRequestSent"
              @click="sendFriendRequest"
              type="primary"
              round
          >
            申请成为好友
          </el-button>
          <el-tag v-else type="success">已发送好友请求</el-tag>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script>
import FindUsersSidebar from "./FindUsersSidebar.vue";
import { sendFriendRequest } from "../api/friend.js";
import { useUserStore } from "../stores/user.js";
import { handlePossibleToken } from "../utils/token.js";
import * as https from "https";

export default {
  computed: {
    https() {
      return https
    }
  },
  data() {
    return {
      avatarUrl: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
      selectedUser: {},
      friendRequestSent: false,
    };
  },
  components: {
    FindUsersSidebar,
  },
  methods: {
    showUserDetails(user) {
      this.selectedUser = user;
      this.friendRequestSent = false;
    },
    async sendFriendRequest() {
      try {
        const response = handlePossibleToken(
            await sendFriendRequest(
                useUserStore().getUserId,
                this.selectedUser.UserId
            )
        );
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
.find-user {
  display: flex;
  flex-direction: row;
}

.user-details {
  flex: 1;
  padding: 20px;
}

/* 添加用户卡片样式 */
.user-card {
  width: 100%;
}

.user-card-header {
  display: flex;
  align-items: center;
}

.user-avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  margin-right: 10px;
}

.user-name {
  font-size: 18px;
  font-weight: bold;
}

.user-info {
  margin-top: 20px;
}

.user-actions {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

/* 样式可以根据需要进行调整 */
</style>
