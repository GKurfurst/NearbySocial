import { useUserStore } from '../stores/user.js';
import {useChatStore} from "../stores/chat.js"; // 导入 User Store

export function logout() {

    useUserStore().clearUser();
    useChatStore().clearChatStore();
}
