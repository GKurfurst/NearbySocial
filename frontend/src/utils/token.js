import {useUserStore} from "../stores/user.js";
import {handleMultiJSONResponse} from "./response.js";
import getUserByName from "../api/getUserByName.js";

export function handleToken(response) {
    if (response.length >= 2) {
        const secondJsonObject = response[1];
        if (secondJsonObject.token && useUserStore().getLoginStatus) {
            console.log("Second JSON Token:", secondJsonObject.token);
            useUserStore().refreshToken(secondJsonObject.token);
        }
        return secondJsonObject.token;
    }
    //还有处理要做
    return null;
}
export function handlePossibleToken(tokenResponse) {
    const response = handleMultiJSONResponse(tokenResponse);
    handleToken(response);
    return response;
}
