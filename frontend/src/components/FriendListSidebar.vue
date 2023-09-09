<template>
  <el-aside class="friend-list-sidebar">
    <div class="friend-list-section">
      <h3>好友请求</h3>
      <div v-if="friendRequests.length > 0">
        <!-- 好友请求列表 -->
        <div v-for="(request, index) in friendRequests" :key="index" class="friend-request-item">
          <img :src="this.avatarUrl" alt="User Avatar" class="avatar" />
          <div class="request-details">
            <div class="user-name">{{ request.Sender.Name }}</div>
            <div class="request-time">Request Time: {{ request.CreatedAt }}</div>
          </div>
          <div class="action-buttons">
            <el-button @click.stop="acceptRequest(request)" type="success" size="small"><el-icon><CircleCheck /></el-icon></el-button>
          </div>
          <div class="action-buttons">
            <el-button @click.stop="rejectRequest(request)" type="danger" size="small"><el-icon><CircleClose /></el-icon></el-button>
          </div>
        </div>
      </div>
      <div v-else>暂无好友请求</div>
    </div>

    <div class="friend-list-section">
      <h3>好友列表</h3>
      <div v-if="friends.length > 0">
        <!-- 好友列表 -->
        <div v-for="(friend, index) in friends" :key="index" class="friend-item">
          <img :src="this.avatarUrl" alt="User Avatar" class="avatar" />
          <div class="friend-details">
            <div class="user-name">{{ friend.Name }}</div>
            <div class="telephone">电话: {{ friend.Telephone }}</div>
          </div>
          <div class="action-buttons">
            <el-button @click.stop="startChat(friend)" type="primary" size="small"><el-icon><el-icon><ChatDotRound /></el-icon></el-icon></el-button>
          </div>
          <div class="action-buttons">
            <el-button @click.stop="deleteFriend(friend)" type="danger" size="small"><el-icon><Remove /></el-icon></el-button>
          </div>
        </div>
      </div>
      <div v-else>暂无好友</div>
    </div>
  </el-aside>
</template>

<script>
import { approveFriendRequest, getFriend, getFriendRequests, rejectFriendRequest, removeFriend } from '../api/friend.js'
import { useUserStore } from '../stores/user.js'
import { handlePossibleToken } from '../utils/token.js'
import {ChatDotRound, Check, CircleCheck, CircleClose, Remove} from "@element-plus/icons-vue";
import {useChatStore} from "../stores/chat.js";

export default {
  components: {CircleCheck, CircleClose, ChatDotRound, Remove, Check},
  async created() {
    await this.fetchFriendRequests()
    await this.fetchFriends()
    this.friends.forEach(friend => {
      useChatStore().initChatService(friend.UserId);
    })
  },
  data() {
    return {
      avatarUrl: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
      friendRequests: [],
      activeRequest: [],
      friends: [],
      activeFriend: []
    }
  },
  methods: {
    async fetchFriends() {
      try {
        const response = await getFriend()
        this.friends = response.data.Friends
      } catch (error) {
        console.error('Error fetching friends:', error)
      }
    },
    async fetchFriendRequests() {
      try {
        const rawResponse = await getFriendRequests(useUserStore().getUserId)
        this.friendRequests = rawResponse.data
      } catch (error) {
        console.error('Error fetching friend requests:', error)
      }
    },
    async acceptRequest(request) {
      const rawResponse = await approveFriendRequest(request.Sender.UserId, request.Receiver.UserId)
      handlePossibleToken(rawResponse)
      await this.fetchFriendRequests()
      await this.fetchFriends()
      useChatStore().initChatService(request.Sender.UserId);
    },
    async rejectRequest(request) {
      const rawResponse = await rejectFriendRequest(request.Sender.UserId, request.Receiver.UserId)
      handlePossibleToken(rawResponse)
      await this.fetchFriendRequests()
    },
    startChat(friend) {
      this.$router.push(`/main/chat/${friend.UserId}`)
    },
    async deleteFriend(friend) {
      const rawResponse = await removeFriend(friend.UserId)
      handlePossibleToken(rawResponse)
      await this.fetchFriends()
    }
  }
}
</script>

<style scoped>
.friend-list-sidebar {
  background-color: #ffffff; /* 使用你的选择的背景颜色 */
  padding: 20px;
}

.friend-list-section {
  margin-bottom: 20px;
}

.friend-request-item,
.friend-item {
  display: flex;
  align-items: center;
}

.avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  margin-right: 10px;
}

.action-buttons {
  margin-left: auto;
}

.user-name {
  font-weight: bold;
}

.request-time,
.telephone {
  color: #666;
}

/* 根据需要进行其他样式微调 */
</style>