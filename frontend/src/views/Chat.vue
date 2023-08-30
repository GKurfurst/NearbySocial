<template>
  <div class="chat-container">
    <input v-model="targetFriendId" placeholder="Chat with ...">
    <button @click="startChat">Chat</button>
    <div class="chat-messages">
      <div v-for="message in messages" class="message">
        {{ message }}
      </div>
    </div>
    <div class="chat-input">
      <input v-model="newMessage" @keyup.enter="sendMessage" placeholder="Type your message...">
      <button @click="sendMessage">Send</button>
    </div>
  </div>
</template>

<script>
import {useUserStore} from "../stores/user.js";
import {useChatStore} from "../stores/chat.js";
import ChatService from "../api/chat.js";
export default {
  data() {
    return {
      newMessage: '',
      targetFriendId: '',
      chatService: null
    };
  },

  mounted() {
    console.log(useUserStore().getUser.userId);
    if (this.$route.params.friendId !== undefined) {
      this.initChatService(this.$route.params.friendId);
    }

  },
  computed: {
    messages() {
      const chatStore = useChatStore();
      return useChatStore().userMessages[this.$route.params.friendId] || [];
    },
  },
  methods: {
    startChat() {
      this.$router.push(`/chat/${this.targetFriendId}`);
    },
    initChatService(friendId) {
      this.chatService = new ChatService(friendId);
      },
    sendMessage() {
      if (this.newMessage.trim() === '') return;
      this.chatService.sendMessage(this.newMessage);
      //this.newMessage = '';
    }
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
</style>
