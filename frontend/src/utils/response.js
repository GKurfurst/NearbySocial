import loginAPI from "../api/login.js";

export function handleMultiJSONResponse(rawResponse) {
    console.log(rawResponse);
    console.log(rawResponse.data);
    //if rawResponse.data is not object
    if (typeof rawResponse.data !== 'object') {
    console.log(rawResponse.data);
    console.log(rawResponse.data.split('}'));
    const response = rawResponse.data.split('}');
    response.pop();
    for (let i = 0; i < response.length; i++) {
        response[i] += '}';
        console.log(response[i]);
        response[i] = JSON.parse(response[i]);
    }
        console.log(response);
        return response;
    }
    else {
        return rawResponse.data;
    }

}
