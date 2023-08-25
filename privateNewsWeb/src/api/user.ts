import http from "@/api/request";

export interface LoginAndRegParams {
    username: string;
    password: string;
    keyword: string
}


export interface LoginData {
    accessToken: string;
    userId: number,
    username: string

}

// 登录
export function login(params: LoginAndRegParams) {
    return http<LoginData>({
        url: "/auth/login",
        method: "post",
        data: params,
    });
}

// 获取用户信息
export function getUserInfo() {
    return http({
        url: "/v1/user/getUserInfo",
        method: "get",
    });
}

// 更新用户信息
export function editUserInfo() {
    return http({
        url: "/v1/user/getUserInfo",
        method: "get",
    });
}