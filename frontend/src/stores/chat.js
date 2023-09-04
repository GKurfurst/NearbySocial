import { defineStore } from 'pinia';

export const useChatStore = defineStore('chat', {
    state: () => ({
        userMessages: {},
        historyMessages: {},
        unreadMessages: {},
    }),
    actions: {
        markMessagesAsRead(friendId) {
            if (this.unreadMessages[friendId]) {
                this.unreadMessages[friendId] = 0;
            }
        },
        // 判断指定好友是否有未读消息
        hasUnreadMessages(friendId) {
            return this.unreadMessages[friendId] > 0;
        },
        getChattedFriends() {
            return Object.keys(this.userMessages);
        },
        addMessage(friendId, message) {
            if (!this.userMessages[friendId]) {
                this.userMessages[friendId] = [];
            }
            this.userMessages[friendId].push(message);
            if (!this.unreadMessages[friendId]) {
                this.unreadMessages[friendId] = 0;
            }
            this.unreadMessages[friendId]++;

        },

        addHistoryMessage(friendId, message) {
            if (!this.historyMessages[friendId]) {
                this.historyMessages[friendId] = [];
            }
            this.historyMessages[friendId].push(message);
        }
    },
});
