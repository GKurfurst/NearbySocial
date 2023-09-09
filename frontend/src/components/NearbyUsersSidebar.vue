<template>
  <el-aside class="nearby-users-sidebar">
    <el-form :model="form" label-width="80px" label-position="top">
      <el-form-item label="选择范围">
        <el-radio-group v-model="form.range">
          <el-radio-button label="500">500m</el-radio-button>
          <el-radio-button label="1000">1000m</el-radio-button>
          <el-radio-button label="2000">2000m</el-radio-button>
          <el-radio-button label="5000">5000m</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-button type="primary" @click="fetchNearbyUsers">搜索</el-button>
    </el-form>
  </el-aside>
</template>

<script>
export default {
  data() {
    return {
      form: {
        range: '500', // 默认范围
      },
    };
  },
  methods: {
    fetchNearbyUsers() {
      // 在这里触发请求附近用户的方法，可以将选定的范围传递给父组件
      this.$emit('fetch-nearby-users', this.form.range);
    },
  },
};
</script>
<style scoped>
.nearby-users-sidebar {
width: 350px;
padding: 20px;
}
</style>
