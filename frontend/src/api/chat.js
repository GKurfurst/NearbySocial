import { WebSocket } from "../utils/ws.js";
import { useChatStore } from "../stores/chat.js"; // 导入 ChatStore

class ChatService {
    constructor(id) {
        this.friendID = id;
        this.ws = new WebSocket(`ws://localhost:8080/chat/${id}`);
        this.ws.onmessage(this.handleMessage.bind(this));
    }

    sendMessage(message) {
        this.ws.send(JSON.stringify(message));
    }

    handleMessage(event) {
        const message = JSON.parse(event.data);
        useChatStore().addMessage(this.friendID, message);
    }
}

export default ChatService;
