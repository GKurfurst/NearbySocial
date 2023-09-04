<template>
  <div class="chat-container">
    <!-- 左侧侧边栏 -->
    <ChatSidebar @friend-selected="handleFriendSelected" ></ChatSidebar>
    <div class="chat-messages">
      <button @click="toggleHistory" class="show-history-button">
        {{ showHistory ? 'Hide History' : 'Show History' }}
      </button>
      <div v-if="showHistory">
        <div v-for="message in historyMessages" :key="message.Timestamp" class="message" :class="getMessageClass(message)">
          <div class="message-time">{{ message.timestamp }}</div>
          {{ message.content }}
        </div>
      </div>
      <div v-else>
        <div v-for="message in messages" :key="message.Timestamp" class="message" :class="getMessageClass(message)">
          <div class="message-time">{{ message.Timestamp }}</div>
          {{ message.Content }}
        </div>
      </div>
    </div>
    <div class="chat-input">
      <el-input
          v-model="newMessage"
          @keyup.enter="sendMessage"
          placeholder="Type your message..."
          class="chat-input-box"
      />
      <el-button @click="sendMessage" class="send-button">Send</el-button>
    </div>
    <div class="chat-content">
      <!-- 其他聊天内容相关代码 -->
    </div>
  </div>
</template>
<script>
import {useUserStore} from "../stores/user.js";
import {useChatStore} from "../stores/chat.js";
import {ChatService, getHistoryMessages} from "../api/chat.js";
import ChatSidebar from "./ChatSidebar.vue";
export default {
  data() {
    return {
      newMessage: '',
      targetFriendId: '',
      chatServices: {},
      showHistory: false,
    };
  },

  mounted() {
    console.log(useUserStore().isLoggedIn);
    console.log(useUserStore().getUserId);
    console.log(this.$route.params.friendId);
    if (this.$route.params.friendId !== undefined) {
      console.log('detected');
      this.initChatService(this.$route.params.friendId);
    }

  },
  beforeRouteUpdate(to, from, next) {
    console.log("route change" + to.params.friendId);
    if (to.params.friendId !== from.params.friendId) {
      if (!this.chatServices[to.params.friendId]) {
        this.initChatService(to.params.friendId);
      }
    }
    next();
  },
  components: {
    ChatSidebar
  },
  computed: {
    historyMessages() {
      return useChatStore().historyMessages[this.$route.params.friendId] || [];
    },
    messages() {
      return useChatStore().userMessages[this.$route.params.friendId] || [];
    },
  },
  methods: {
    async fetchHistoryMessages(friendId) {
      const response = await getHistoryMessages(friendId);
      for (const message of response.data.history) {
        useChatStore().addHistoryMessage(friendId, message);
      }
      const response2 = await getHistoryMessages(useUserStore().getUserId, friendId);
      for (const message of response2.data.history) {
        useChatStore().addHistoryMessage(friendId, message);
      }
    },
    getMessageClass(message) {
      return {
        'my-message': message.sender_id === useUserStore().getUserId,
        'other-message': message.sender_id !== useUserStore().getUserId,
      };
    },

    toggleHistory() {
      this.showHistory = !this.showHistory;
    },
    startChat() {
      console.log("start chat");
      this.$router.push(`/main/chat/${this.targetFriendId}`);
    },
    initChatService(friendId) {
      this.chatServices[friendId] = new ChatService(friendId);
      this.fetchHistoryMessages(friendId);
      },
    sendMessage() {
      if (this.newMessage.trim() === '') return;
      const message = {
        sender_id: useUserStore().getUserId, // 你的 SenderID
        receiver_id: this.$route.params.friendId, // 你的 ReceiverID
        Content: this.newMessage,
        Timestamp: new Date().toISOString(), // 生成时间戳
      };
      this.chatServices[this.$route.params.friendId].sendMessage(message);
      this.newMessage = '';
    },
    handleFriendSelected(friendId) {
      this.targetFriendId = friendId;
      this.startChat();
    },
  }
};
</script>

<style scoped>
.chat-container {
  max-width: 600px;
  margin: 0 auto;
}

.chat-messages {
  border: 1px solid #ccc;
  padding: 10px;
  height: 300px;
  overflow-y: auto;
}

.my-message {
  text-align: right;
  background-color: #007bff;
  color: #fff;
  border-radius: 5px;
  padding: 5px 10px;
  margin: 5px 0;
}

.other-message {
  text-align: left;
  background-color: #ccc;
  border-radius: 5px;
  padding: 5px 10px;
  margin: 5px 0;
}

.message-time {
  font-size: 12px;
  color: #888;
  text-align: center;
}

.message {
  margin-bottom: 10px;
}

.chat-input {
  display: flex;
  align-items: center;
  margin-top: 10px;
}

input {
  flex: 1;
  padding: 5px;
  border: 1px solid #ccc;
}

button {
  padding: 5px 10px;
  background-color: #007bff;
  color: #fff;
  border: none;
  cursor: pointer;
}
.chat-input-box,
.send-button {
  margin-right: 10px;
}

.chat-button {
  margin-top: 10px;
}
</style>
