import { getRequest } from '../utils/request.js';
import {useUserStore} from "../stores/user.js";

const getUserByName = async (name ,token = useUserStore().getToken) => {
    try {
        console.log("Token:", token);
        const header = {
            'token': token
        };

        const response = await getRequest({
            url: `/get_users_by_name/${name}`,
            headers: header
        });
        console.log("Response:", response);
        return response;
    } catch (error) {
        console.error('Error:', error);
    }
};

export default {
    getUserByName
}
