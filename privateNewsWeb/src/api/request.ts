import axios, { AxiosRequestConfig } from "axios";
import router from "@/router/index";

const http = axios.create({
    baseURL: import.meta.env.VITE_BASE_URL,
    timeout: 10000
})

// 请求拦截器
http.interceptors.request.use((config: any) => {
    if (config?.headers) {
        config.headers.Authorization = localStorage.getItem("accessToken") || "";
    }
    //console.log("config.Header = ", config.headers)
    return config;
})

// 响应拦截器
http.interceptors.response.use(resposne => {
    var respData: BasicResponseModel = resposne.data;
    console.log("resposne.data = ", respData.data)
    if (resposne.data.code != 1000) {
        if (resposne.data.code == 1002 || resposne.data.code == 1003) { // token 无权限
            localStorage.removeItem("accessToken");
            router.push("/login")
        }
        console.log(respData)
        return Promise.reject(respData)
    }
    return resposne.data;
}, (e) => {
    console.log(e.response.data)
    return Promise.reject(e.response.data)
})

export interface BasicResponseModel<T = any> {
    code: number;
    message: string;
    data: T;
}


export function get(url: string, params: any) {
    console.log("baseURL = " + import.meta.env.VITE_BASE_URL)
    console.log("MODE = " + import.meta.env.MODE)
    console.log("url = " + url)
    console.log("params = " + params)
    return http.request<BasicResponseModel>(
        {
            url: url,
            method: 'get',
            params,
        }
    );
}

export function post(url: string, data: any) {
    return http.post<BasicResponseModel>(url, data);
}

export default http;