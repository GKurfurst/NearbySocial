<template>
  <el-aside class="find-user-sidebar">
    <el-input v-model="searchText" placeholder="搜索用户" @input="searchUsers" />
    <el-menu class="user-list">
      <el-menu-item
          v-if="searchResults.Name"
          @click="selectUser(searchResults)"
      >
        {{ searchResults.Name }}
      </el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script>
import getUserByName from "../api/getUserByName.js";
import { handleMultiJSONResponse } from "../utils/response.js";
import { handlePossibleToken, handleToken } from "../utils/token.js";
import {updatePosition} from "../api/position.js";
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
/* 调整搜索框样式 */
.find-user-sidebar {
  width: 300px;
  padding: 20px;
}

/* 调整输入框样式 */
.el-input {
  width: 100%;
}

/* 调整用户列表样式 */
.user-list {
  list-style-type: none;
  padding: 0;
  cursor: pointer;
}

.user-list .el-menu-item {
  padding: 5px 10px;
  margin: 5px 0;
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  background-color: #fff;
  transition: background-color 0.3s ease-in-out;
}

.user-list .el-menu-item:hover {
  background-color: #f0f5ff; /* 鼠标悬停时的背景色 */
}
</style>
