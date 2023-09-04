<template>
  <Layout>
<!--    <template #main>-->
<!--      <Chat />-->
<!--    </template>-->
  </Layout>
</template>

<script>
import Headerbar from '../components/Headerbar.vue';
import Layout from '../layout/MainLayout.vue';
import Chat from '../components/Chat.vue';

export default {
  components: {
    Headerbar,
    Layout,
    Chat,
  },
};
</script>
