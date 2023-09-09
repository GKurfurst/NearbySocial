import {postRequest} from '../utils/request.js'
import {useUserStore} from "../stores/user.js";
import {handleToken} from "../utils/token.js";

const login = async (username, password) => {
    try {

        const response = await postRequest({
            url : "/login",
            headers: {'Content-Type': 'application/json'},
            data : {
                username : username,
                password : password
            }
        })
        return response;

    } catch (error) {
        if (error.response.status === 422) {
            alert(error.response.data.message)
        }
    }
}
export default {
    login
}