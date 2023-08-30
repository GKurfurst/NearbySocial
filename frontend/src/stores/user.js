import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
    state: () => ({
        userId: null,
        username: ''
    }),
    getters: {
        getUser: state => state
    },
    actions: {
        setUser(user) {
            this.userId = user.userId;
            this.username = user.username;
        },
        clearUser() {
            this.userId = null;
            this.username = '';
        }
    }
});
