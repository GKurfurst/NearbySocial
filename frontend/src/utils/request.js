// 封装axios，设置axios的默认配置与拦截器
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