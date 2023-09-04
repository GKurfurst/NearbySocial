import {postRequest} from "../utils/request.js";
import {useUserStore} from "../stores/user.js";
import {handleMultiJSONResponse} from "../utils/response.js";
import loginAPI from "./login.js";
import {getUserPosition} from "../utils/position.js";
export async function getNearbyUser(radius, userId = useUserStore().getUserId) {
    try {
        const header = {
            'token': useUserStore().getToken,
            'Content-Type': 'application/json'
        };
        const response = await postRequest({
            url : "/find_nearby_users",
            headers: header,
            data : {
                User_Id: userId,
                Radius: parseFloat(radius)
            }
        })
        return response;

    } catch (error) {
        if (error.response.status === 422) {
            alert(error.response.data.message)
        }
    }
}