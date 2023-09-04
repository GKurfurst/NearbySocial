<template>
  <el-aside class="friend-list-sidebar">
    <!-- 好友请求 -->
    <h3>好友请求</h3>
    <div v-if="friendRequests.length > 0">
      <el-collapse v-model="activeRequest">
        <el-collapse-item v-for="(request, index) in friendRequests" :key="index">
          <template #title>
            {{ request.Sender.Name }}
            <el-button @click.stop="acceptRequest(request)" type="primary" size="small">接受</el-button>
            <el-button @click.stop="rejectRequest(request)" type="danger" size="small">拒绝</el-button>
          </template>
          Request Time: {{ request.CreatedAt }}
        </el-collapse-item>
      </el-collapse>
    </div>
    <div v-else>暂无好友请求</div>

    <!-- 好友列表 -->
    <h3>好友列表</h3>
    <div v-if="friends.length > 0">
      <el-collapse v-model="activeFriend">
        <el-collapse-item v-for="(friend, index) in friends" :key="index">
          <template #title>
            {{ friend.Name }}
            <el-button @click.stop="startChat(friend)" type="success" size="small">开始聊天</el-button>
            <el-button @click.stop="deleteFriend(friend)" type="danger" size="small">移除好友</el-button>
          </template>
          电话: {{ friend.Telephone }}
        </el-collapse-item>
      </el-collapse>
    </div>
    <div v-else>暂无好友</div>
  </el-aside>
</template>

<script>
import {approveFriendRequest, getFriend, getFriendRequests, rejectFriendRequest, removeFriend} from "../api/friend.js";
import {useUserStore} from "../stores/user.js";
import {handlePossibleToken} from "../utils/token.js";

export default {
  created() {
    // 在组件创建时获取好友请求
    this.fetchFriendRequests();
    this.fetchFriends();
  },
  data() {
    return {
      friendRequests: [],
      activeRequest: [],
      friends: [],
      activeFriend: [],
    };
  },
  methods: {
    async fetchFriends() {
      try {
        // 调用获取好友请求的函数
        const response = await getFriend();
        //还需处理rawResponse可能携带的新token
        this.friends = response.data.Friends;
      } catch (error) {
        console.error('Error fetching friend requests:', error);
      }
    },
    async fetchFriendRequests() {
      try {
        // 调用获取好友请求的函数
        const rawResponse = await getFriendRequests(useUserStore().getUserId);
        //还需处理rawResponse可能携带的新token
        this.friendRequests = rawResponse.data;
      } catch (error) {
        console.error('Error fetching friend requests:', error);
      }
    },
    async acceptRequest(request) {
      const rawResponse = await approveFriendRequest(request.Sender.UserId, request.Receiver.UserId);
      handlePossibleToken(rawResponse);
      await this.fetchFriendRequests();
      await this.fetchFriends();

    },
    async rejectRequest(request) {
      const rawResponse = await rejectFriendRequest(request.Sender.UserId, request.Receiver.UserId);
      handlePossibleToken(rawResponse);
      await this.fetchFriendRequests();
    },
    startChat(friend) {
      this.$router.push(`/main/chat/${friend.UserId}`);
    },
    async deleteFriend(friend) {
      const rawResponse = await removeFriend(friend.UserId);
      handlePossibleToken(rawResponse);
      await this.fetchFriends();
    },
  },
};
</script>
