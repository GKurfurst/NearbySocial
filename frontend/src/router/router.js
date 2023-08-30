import { createRouter, createWebHistory } from 'vue-router'
import Login from "../views/Login.vue";
import Friends from "../views/FriendList.vue";
import Chat from "../views/Chat.vue";
import Register from "../views/Register.vue";
import App from "../App.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: App },
        { path: '/login', component: Login },
        { path: '/chat', component: Chat },
        {
            path: '/chat/:friendId',
            name: 'Chat',
            component: Chat
        },
        { path: '/friends', component: Friends},
        { path: '/register', component: Register}
    ]
})

export default router