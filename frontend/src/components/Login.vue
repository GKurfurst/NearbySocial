<template>
  <el-col :span="8" class="login-form">
    <el-input v-model="username" placeholder="Username" />
    <el-input v-model="password" placeholder="Password" type="password" />
    <el-button type="primary" @click="login">Login</el-button>
    <el-button type="text" @click="goToRegister">Register</el-button>
    <p>{{ username }}</p>
    <div v-if="loginStatus === 'success'" class="success-message">
      Login successful! If it's not redirecting automatically, You can manually <router-link to="/auth/login">login</router-link>.
    </div>
    <div v-else-if="loginStatus === 'error'" class="error-message">
      Login failed. Please try again.
    </div>
  </el-col>
</template>

<script>
import loginAPI from '../api/login.js';
import {handlePossibleToken, handleToken} from "../utils/token.js";
import {useUserStore} from "../stores/user.js";
import getUserByName from "../api/getUserByName.js";
import {handleMultiJSONResponse} from "../utils/response.js";
import {updatePosition} from "../api/position.js";

export default {
  data() {
    return {
      username: '',
      password: '',
      loginStatus: '',
    };
  },
  methods: {
    async login() {
      try {
        const response = handleMultiJSONResponse(await loginAPI.login(this.username, this.password));

        console.log('ff');
        console.log(response[0].code);
        console.log(response[0]);
        // 提取第一个JSON对象
        const firstJsonObject = response[0];
        // 处理第一个JSON对象
        if (firstJsonObject.code === 200) {
          console.log(firstJsonObject.message); // 输出 "登录成功"
          // 提取第二个JSON对象
          const token = handleToken(response);
          // 提取第三个JSON对象
          const thirdJsonObject = response[2];
          // 处理第三个JSON对象
          if (thirdJsonObject.code === 200) {
            console.log(thirdJsonObject.message); // 输出 "用户状态已更新"
          }
          const rawResponse = await getUserByName.getUserByName(this.username, token);
          console.log('faf')
          console.log(rawResponse);
          const userData = handleMultiJSONResponse(rawResponse);
          console.log(userData);
          useUserStore().setUser({username: userData.Name, userId: userData.UserId,token: token});
          console.log(useUserStore().getUserId);
          console.log('reqa')
          handlePossibleToken(await updatePosition());
          this.loginStatus = 'success';
          this.$router.push('/main');
        } else {
          console.error("登录失败");
          this.loginStatus = 'error';
        }
      } catch (error) {
        console.error(error);
        this.loginStatus = 'error';
      }
    },
    goToRegister() {
      this.$router.push('/auth/register');
    },


  },
};
</script>

<style scoped>
.login-form {
  padding: 20px;
  background-color: #f0f0f0;
  border-radius: 5px;
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
}

.login-form .el-button--text {
  margin-top: 10px;
  color: #007bff;
}
</style>
