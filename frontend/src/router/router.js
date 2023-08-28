import { createRouter, createWebHistory } from 'vue-router'
import Login from "../components/Login.vue";
import Friends from "../components/FriendList.vue";
import Chat from "../components/Chat.vue";
import Register from "../components/Register.vue";
import App from "../App.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: App },
        { path: '/login', component: Login },
        { path: '/chat', component: Chat },
        { path: '/friends', component: Friends},
        { path: '/register', component: Register}
    ]
})

export default router