import { defineStore } from 'pinia';

export const useChatStore = defineStore('chat', {
    state: () => ({
        userMessages: {},
    }),
    actions: {
        addMessage(friendId, message) {
            if (!this.userMessages[friendId]) {
                this.userMessages[friendId] = [];
            }
            this.userMessages[friendId].push(message);
        },
    },
});
