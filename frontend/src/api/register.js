import axios from 'axios'
import {postRequest} from '../utils/request.js'
const register = async (username, telephone, password) => {
    try {
        const response = await postRequest({
            url: "/register",
            headers: {'Content-Type': 'application/json'}, // 设置 Content-Type 为 application/json
            data: {
                username: username,
                telephone: telephone,
                password: password
            }
        });
        return response;


    } catch (error) {
        console.error(error);
    }
};
export default {
    register
}