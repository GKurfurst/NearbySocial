import { defineStore } from 'pinia';
import {ChatService, getHistoryMessages} from "../api/chat.js";
import {useUserStore} from "./user.js";

export const useChatStore = defineStore('chat', {
    state: () => ({
        chatServices: {},
        userMessages: {},
        historyMessages: {},
        unreadMessages: {},
    }),
    actions: {
        clearChatStore() {
            this.chatServices = {};
            this.userMessages = {};
            this.historyMessages = {};
            this.unreadMessages = {};
        },
        initChatService(friendId) {
            if (!this.chatServices[friendId]) {
                this.chatServices[friendId] = new ChatService(friendId);
                this.fetchHistoryMessages(friendId);
            }

        },
        async fetchHistoryMessages(friendId) {
            const response = await getHistoryMessages(friendId);
            if (response.data.history === null) {
                this.userMessages[friendId] = [];
                this.historyMessages[friendId] = [];
            }
            for (const message of response.data.history) {
                useChatStore().addHistoryMessage(friendId, message);
                if (Date.parse(message.timestamp) > Date.now() - 7 * 24 * 60 * 60 * 1000) {
                    useChatStore().addMessage(friendId, message, false);
                }
            }
            const response2 = await getHistoryMessages(useUserStore().getUserId, friendId);
            for (const message of response2.data.history) {
                useChatStore().addHistoryMessage(friendId, message);
                if (Date.parse(message.timestamp) > Date.now() - 7 * 24 * 60 * 60 * 1000) {
                    useChatStore().addMessage(friendId, message, false);
                }
            }
        },
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
            return Object.keys(this.historyMessages);
        },
        addMessage(friendId, message, notify = true) {
            if (!this.userMessages[friendId]) {
                this.userMessages[friendId] = [];
            }
            this.userMessages[friendId].push(message);
            if (!this.unreadMessages[friendId]) {
                this.unreadMessages[friendId] = 0;
            }
            if (notify){
                this.unreadMessages[friendId]++;
            }

        },

        addHistoryMessage(friendId, message) {
            if (!this.historyMessages[friendId]) {
                this.historyMessages[friendId] = [];
            }
            this.historyMessages[friendId].push(message);
        }
    },
});
