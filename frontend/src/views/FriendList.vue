<template>
  <div class="friend-list">
    <h1>Friend List</h1>
    <ul>
      <li v-for="(friend, index) in friends" :key="index">{{ friend }}</li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      friends: ['Friend 1', 'Friend 2', 'Friend 3'] // 好友列表数据
    };
  }
};
</script>

<style>

</style>
