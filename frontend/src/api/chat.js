import PackedWebSocket from "../utils/ws.js";
import { useChatStore } from "../stores/chat.js";
import {useUserStore} from "../stores/user.js";
import {getRequest} from "../utils/request.js"; // 导入 ChatStore

class ChatService {
    constructor(id) {
        this.friendID = id;
        this.ws = new PackedWebSocket(`ws://localhost:9090/api/chat/${id}?userId=${useUserStore().getUserId}`);
        this.ws.onmessage(this.handleMessage.bind(this));
    }

    sendMessage(message) {
        this.ws.send(JSON.stringify(message));
        useChatStore().addMessage(this.friendID, message);
    }

    handleMessage(event) {
        const message = JSON.parse(event.data);
        useChatStore().addMessage(this.friendID, message);
    }


}

async function getHistoryMessages(friendId, userId = useUserStore().getUserId, token = useUserStore().getToken) {
    try {
        console.log("Token:", token);
        const header = {
            'token': token
        };

        const response = await getRequest({
            url: `/get_chat_history/${userId}/${friendId}`,
            headers: header
        });
        console.log("Response:", response);
        return response;
    } catch (error) {
        console.error('Error fetching history messages:', error);
        throw error;
    }
}

export { ChatService, getHistoryMessages };
