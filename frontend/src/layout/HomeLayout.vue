<template>
  <div class="home-container">
    <Headerbar />
    <div class="content">
      <div class="logo">
        <el-image src="https://element-plus.org/images/element-plus-logo.svg" fit="contain" style="max-width: 150px;" />
      </div>
      <div class="description">
        <h2 class="title">
          Welcome to Our Project
        </h2>
        <p class="text">
          NearbySocial connects everyone.
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import Headerbar from '../components/Headerbar.vue';
export default {
  // computed: {
  //   location: async function () {
  //     const location = await getUserLocation();
  //     console.log(location);
  //     return location;
  //   }
  // },
  components: {
    Headerbar,
  },
  methods: {
  }
};
</script>

<style scoped>
.home-container {
  text-align: center;
  background-color: #f0f0f0;
  padding-bottom: 40px;
}

.content {
  max-width: 800px;
  margin: 0 auto;
  padding: 60px 20px;
}

.title {
  font-size: 24px;
  color: #333;
  margin-bottom: 10px;
  background: linear-gradient(to right, #48b604, #c1e1b2);
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}

.text {
  font-size: 16px;
  color: #666;
}
</style>
