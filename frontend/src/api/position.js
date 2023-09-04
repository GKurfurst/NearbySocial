import {postRequest} from "../utils/request.js";
import {useUserStore} from "../stores/user.js";
import {handleMultiJSONResponse} from "../utils/response.js";
import loginAPI from "./login.js";
import {getUserPosition} from "../utils/position.js";
export async function updatePosition(userId = useUserStore().getUserId) {
    try {
        const header = {
            'token': useUserStore().getToken,
            'Content-Type': 'application/json'
        };
        const coords = await getUserPosition();
        const response = await postRequest({
            url : "/update_location",
            headers: header,
            data : {
                User_Id: userId,
                Latitude: coords.latitude,
                Longitude: coords.longitude
            }
        })
        return response;

    } catch (error) {
        if (error.response.status === 422) {
            alert(error.response.data.message)
        }
    }
}