import request from "@/utils/request"
import type { LoginParams, LoginResponse } from "./type"

enum API {
    LOGIN_URL = '/api/v1/user/login',
    USER_INFO_URL = '/api/v1/user/info'
}

export const login = (data: LoginParams) => {
    return request.post<any, LoginResponse>(API.LOGIN_URL, data)
}