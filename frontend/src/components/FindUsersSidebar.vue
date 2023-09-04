<template>
  <div class="find-user-sidebar">
    <el-input v-model="searchText" placeholder="Search user" @input="searchUsers" />
    <el-menu class="user-list">
      <el-menu-item
          @click="selectUser(this.searchResults)"
      >
        {{ this.searchResults.Name }}
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script>
import getUserByName from "../api/getUserByName.js";
import {handleMultiJSONResponse} from "../utils/response.js";
import {handlePossibleToken, handleToken} from "../utils/token.js";
export default {
  data() {
    return {
      searchText: '',
      searchResults: {},
    };
  },
  methods: {
    async searchUsers() {
      // 调用 getUserByName 方法来获取搜索结果
      if (this.searchText.trim() !== '') {
        try {
          const response = handlePossibleToken(await getUserByName.getUserByName(this.searchText));
          console.log('after token check' + response);
          if (response) {
            this.searchResults = response;
          } else {
            this.searchResults = {}; // 没有结果时清空列表
          }
        } catch (error) {
          console.error('Error:', error);
          this.searchResults = {}; // 出错时清空列表
        }
      } else {
        this.searchResults = {}; // 没有输入时清空列表
      }
    },
    selectUser(user) {
      console.log('current user' + user.Name);
      // 当用户点击搜索结果时，将选定的用户传递给父组件
      this.$emit('user-selected', user);
    },
  },
};
</script>

<style scoped>
/* 样式可以根据需要进行调整 */
.find-user-sidebar {
  width: 300px;
  padding: 20px;
}

.user-list {
  list-style-type: none;
  padding: 0;
  cursor: pointer;
}

.user-list li {
  margin: 5px 0;
}
</style>
