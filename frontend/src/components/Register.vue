<template>
  <el-col :span="8" class="register-form">
    <el-input v-model="username" placeholder="Username" />
    <el-input v-model="telephone" placeholder="Telephone" />
    <el-input v-model="password" placeholder="Password" type="password" />
    <el-button type="primary" @click="register">Register</el-button>
    <el-button type="text" @click="goToLogin">Login</el-button>
    <div v-if="registerStatus === 'success'" class="success-message">
      Registration successful! If it's not redirecting automatically, You can manually <router-link to="/auth/login">login</router-link>.
    </div>
    <div v-else-if="registerStatus === 'error'" class="error-message">
      Registration failed. Please try again.
    </div>
  </el-col>
</template>

<script>
import { validateUsername, validatePassword, validateTelephone } from "../utils/validation.js";
import registerAPI from "../api/register.js";
export default {
  data() {
    return {
      username: '',
      password: '',
      telephone: '',
      registerStatus: '',
    };
  },
  methods: {
    async register() {
      if (!validateUsername(this.username)) {
        this.$message.error("Username should be at least 5 characters.");
        return;
      }

      if (!validateTelephone(this.telephone)) {
        //console.log(this.telephone)
        this.$message.error("Telephone should be at least 11 characters.");
        return;
      }

      if (!validatePassword(this.password)) {
        this.$message.error("Password should be at least 6 characters.");
        return;
      }

      try {
        const response = await registerAPI.register(this.username, this.telephone, this.password);
        console.log(response.data)
        console.log('2')
        if (response.data.code === 200) {
          this.registerStatus = 'success';
          this.$router.push('/auth/login');
        } else {
          this.registerStatus = 'error';
        }
      } catch (error) {
        console.log('1')
        console.error(error);
        this.registerStatus = 'error';
      }
    },
    goToLogin() {
      this.$router.push('/auth/login');
    },
  },
};
</script>

<style scoped>
.register-form {
  padding: 20px;
  background-color: #f0f0f0;
  border-radius: 5px;
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
}

.register-form .el-button--text {
  margin-top: 10px;
  color: #007bff;
}
</style>
