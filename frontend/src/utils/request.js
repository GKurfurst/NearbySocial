
import axios from "axios";

const getRequest = axios.create({
    baseURL : "http://localhost:9090/api",
    method : "get"
});

const postRequest = axios.create({
    baseURL : "http://localhost:9090/api",
    method : "post"
});

export {
    getRequest,
    postRequest};