import { createRouter, createWebHistory } from 'vue-router';
import { useUserStore} from "../stores/user.js";

import Login from "../components/Login.vue";
import Register from "../components/Register.vue";
import Home from "../views/Home.vue";
import Main from "../views/Main.vue";
import Auth from "../views/Auth.vue";
import Chat from "../components/Chat.vue";
import FindUsers from "../components/FindUsers.vue";
import FriendList from "../components/FriendList.vue";
import NearbyUsers from "../components/NearbyUsers.vue";
import {useChatStore} from "../stores/chat.js";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: Home },
        {
            path: '/main',
            component: Main,
            meta: { requiresAuth: true }, // 添加需要登录的标志
            children: [
                {
                    path: 'nearby-users',
                    component: NearbyUsers,
                },
                {
                    path: 'find-users',
                    component: FindUsers,
                },
                {
                    path: 'friend-list',
                    component: FriendList,
                },
                {
                    path: 'chat',
                    component: Chat,
                },
                {
                    path: 'chat/:friendId',
                    component: Chat,
                },
            ]
        },
        {
            path: '/auth',
            component: Auth,
            redirect: '/auth/login',
            children: [
                {
                    path: 'login',
                    component: Login,
                },
                {
                    path: 'register',
                    component: Register,
                },
            ]
        },
    ]
});

//添加全局前置守卫来检查登录状态并进行导航
router.beforeEach((to, from, next) => {
    if (to.meta.requiresAuth && !useUserStore().isLoggedIn) {
        // 需要登录但未登录，跳转到登录页面
        next('/auth/login');
    } else {
        next();
    }
});

export default router;
