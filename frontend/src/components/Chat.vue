<template>
  <div class="chat-container">
    <ChatSidebar @friend-selected="handleFriendSelected" ></ChatSidebar>
    <div class="chat-area">
      <div id = "chat-messages" class="chat-messages">
        <div v-if="showHistory">
          <div v-for="message in historyMessages" class="message" :class="getMessageClass(message)">
            <div class="message-time">{{ message.timestamp }}</div>
            {{ message.content }}
          </div>
        </div>
        <div v-else>
          <div v-for="message in messages" class="message" :class="getMessageClass(message)">
            <div class="message-time">{{ message.timestamp }}</div>
            {{ message.content }}
          </div>
        </div>
      </div>
      <button @click="toggleHistory" class="show-history-button">
        {{ showHistory ? 'Hide History' : 'Show History' }}
      </button>
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
      </div>
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

  created() {
    console.log(useUserStore().isLoggedIn);
    console.log(useUserStore().getUserId);
    console.log(this.$route.params.friendId);
    // if (this.$route.params.friendId === undefined){
    //   this.$router.push(`/main/chat/${useChatStore().getChattedFriends()[0]}`);
    //   //this.$forceUpdate();
    // }
  },
  components: {
    ChatSidebar
  },
  computed: {
    historyMessages() {
      return useChatStore().historyMessages[this.$route.params.friendId].slice().sort((a, b) => Date.parse(a.timestamp) - Date.parse(b.timestamp)) || [];
    },
    messages() {
      return useChatStore().userMessages[this.$route.params.friendId].slice().sort((a, b) => Date.parse(a.timestamp) - Date.parse(b.timestamp)) || [];
    },
  },
  methods: {
    scrollToBottom() {
      const container = document.getElementById('chat-messages');
      container.scrollIntoView(false);
    },
    getMessageClass(message) {
      return {
        'my-message': message.sender_id === useUserStore().getUserId,
        'other-message': message.sender_id !== useUserStore().getUserId,
      };
    },

    toggleHistory() {
      this.$forceUpdate();
      this.showHistory = !this.showHistory;
      this.$forceUpdate();
    },
    startChat() {
      console.log("start chat");
      this.$router.push(`/main/chat/${this.targetFriendId}`);
      this.$forceUpdate();
    },
    sendMessage() {
      if (this.newMessage.trim() === '') return;
      const message = {
        sender_id: useUserStore().getUserId, // 你的 SenderID
        receiver_id: this.$route.params.friendId, // 你的 ReceiverID
        content: this.newMessage,
        timestamp: new Date().toISOString(), // 生成时间戳
      };
      useChatStore().chatServices[this.$route.params.friendId].sendMessage(message);
      this.newMessage = '';
    },
    handleFriendSelected(friendId) {
      this.targetFriendId = friendId;
      this.startChat();
      this.$forceUpdate();
    },
  }
};
</script>

<style scoped>
.chat-container {
  display: flex;
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.chat-messages {
  border: 1px solid #ccc;
  padding: 10px;
  height: 300px;
  overflow-y: auto;
}

.chat-area {
  flex: 1;
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 10px;
  overflow-y: auto;
}
.my-message {
  margin-left: 400px;
  text-align: right;
  background-color: #007bff;
  color: #fff;
  border-radius: 5px;
  padding: 5px 10px;
  max-width: 50%;
}

.other-message {
  text-align: left;
  background-color: #ccc;
  border-radius: 5px;
  padding: 5px 10px;
  margin: 5px 0;
  max-width: 50%;
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