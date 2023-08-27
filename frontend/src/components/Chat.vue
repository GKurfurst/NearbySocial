<template>
  <div class="chat">
    <h1>Welcome, {{ username }}</h1>
    <div class="chat-box">
      <!-- 在这里显示聊天内容 -->
    </div>
    <input v-model="message" @keyup.enter="sendMessage" placeholder="Type your message...">
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: '', // 在登录时获取
      message: ''
    };
  },
  methods: {
    sendMessage() {
      // 在这里处理发送消息逻辑，例如将消息添加到聊天列表
      this.message = '';
    }
  }
};
</script>

<style>
.chat {
  text-align: center;
  margin-top: 100px;
}
.chat-box {
  border: 1px solid #ccc;
  height: 300px;
  overflow: auto;
  margin-bottom: 10px;
  padding: 10px;
}
</style>
