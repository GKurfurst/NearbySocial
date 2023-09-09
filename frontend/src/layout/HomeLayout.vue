<template>
  <div class="home-container">
    <div class="content-bg"></div>
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
  padding-bottom: 40px;
}

.content {
  max-width: 800px;
  margin: 0 auto;
  padding: 60px 20px;
  z-index: 1;
}
.content-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-image: url('https://cn.bing.com/th?id=OHR.ManhattanAerial_EN-CN3036997176_1920x1200.jpg&rf=LaDigue_1920x1200.jpg');
  background-size: cover;
  filter: blur(8px);
  z-index: -1;
}
.title {
  font-size: 50px;
  color: #333;
  margin-bottom: 10px;
  background: linear-gradient(to right, #6be114, #f4f6f3);
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}

.text {
  font-size: 30px;
  color: #4857f8;
}
.logo {
  margin-top: 200px;
}
.description {
  margin-top: 20px;
}
</style>
