import {getRequest, postRequest} from '../utils/request.js';
import { useUserStore } from '../stores/user.js';
import getUserByName from "./getUserByName.js";

// 发送好友请求
export async function sendFriendRequest(senderId, receiverId, token = useUserStore().getToken) {
    try {
        const header = {
            'token': token,
            'Content-Type': 'application/json'
        };
        const response = await postRequest({
            url: '/send_friend_request',
            headers: header,
            data: { sender_id: senderId, receiver_id: receiverId }
        });
        return response;
    } catch (error) {
        throw error.response;
    }
}

// 批准好友请求
export async function approveFriendRequest(senderId, receiverId, token = useUserStore().getToken) {
    try {
        const header = {
            'token': token,
            'Content-Type': 'application/json'
        };
        const response = await postRequest({
            url: '/approve_friend_request',
            headers: header,
            data: { sender_id: senderId, receiver_id: receiverId }
        });
        return response;
    } catch (error) {
        throw error.response;
    }
}

// 拒绝好友请求
export async function rejectFriendRequest(senderId, receiverId, token = useUserStore().getToken) {
    try {
        const header = {
            'token': token,
            'Content-Type': 'application/json'
        };
        const response = await postRequest({
            url: '/reject_friend_request',
            headers: header,
            data: { sender_id: senderId, receiver_id: receiverId }
        });
        return response;
    } catch (error) {
        throw error.response;
    }
}

// 移除好友
export async function removeFriend(friendId, userId = useUserStore().getUserId, token = useUserStore().getToken) {
    try {
        const header = {
            'token': token,
            'Content-Type': 'application/json'
        };
        const response = await postRequest({
            url: '/remove_friend',
            headers: header,
            data: { user_id: userId, friend_id: friendId }
        });
        return response;
    } catch (error) {
        throw error.response;
    }
}

export const getFriendRequests = async (userId ,token = useUserStore().getToken) => {
    try {
        console.log("Token:", token);
        const header = {
            'token': token
        };

        const response = await getRequest({
            url: `/get_sending_request/${userId}`,
            headers: header
        });
        console.log("Response:", response);
        return response;
    } catch (error) {
        console.error('Error:', error);
    }
};

export const getFriend = async (userId = useUserStore().getUsername ,token = useUserStore().getToken) => {
    try {
        console.log("Token:", token);
        const header = {
            'token': token
        };

        const response = await getUserByName.getUserByName(userId, token);
        console.log("Response:", response);
        return response;
    } catch (error) {
        console.error('Error:', error);
    }
};