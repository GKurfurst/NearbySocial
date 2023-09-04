import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
    state: () => ({
        userId: null,
        username: '',
        token: '',
        isLoggedIn: false,
    }),
    getters: {
        getUsername: state => state.username,
        getUserId: state => state.userId,
        getToken: state => state.token,
        getLoginStatus: state => state.isLoggedIn,

    },
    actions: {
        setUser(user) {
            this.userId = user.userId;
            this.username = user.username;
            this.token = user.token;
            this.isLoggedIn = true;
        },
        clearUser() {
            this.userId = null;
            this.username = '';
            this.isLoggedIn = false;
            this.token = '';
        },
        refreshToken(token) {
            this.token = token;
        }
    }
});
