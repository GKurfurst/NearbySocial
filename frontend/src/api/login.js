import {postRequest} from '../utils/request.js'
import {useUserStore} from "../stores/user.js";

const login = async (username, password) => {
    try {

        // const response = await postRequest({
        //     url : "/login",
        //     headers: {'Content-Type': 'application/x-www-form-urlencoded'},
        //     data : {
        //         username : username,
        //         password : password
        //     }
        // })
        useUserStore().setUser({username : username, userId : password});
    } catch (error) {
        if (error.response.status === 422) {
            alert(error.response.data.message)
        }
    }
}
export default {
    login
}