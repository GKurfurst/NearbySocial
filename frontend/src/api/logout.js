import { useUserStore } from '../stores/user.js'; // 导入 User Store

export function logout() {

    const userStore = useUserStore();
    userStore.clearUser();
}
