import { createRouter, createWebHistory } from 'vue-router'
import Login from "../components/Login.vue";
import App from "../App.vue";
import Chat from "../components/Chat.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: Login },
        { path: '/chat', component: Chat }
    ]
})

export default router