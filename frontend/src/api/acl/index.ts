import request from "@/utils/request"
import type { LoginParams, LoginResponse, MenuTreeResponse, InfoResponse } from "./type"

enum API {
    LOGIN_URL = '/api/v1/acl/login',
    MENU_URL = '/api/v1/acl/tree',
    Info_URL = "/api/v1/acl/user",
}

export const login = (data: LoginParams) => {
    return request.post<any, LoginResponse>(API.LOGIN_URL, data)
}

export const tree = () => {
    return request.get<any, MenuTreeResponse>(API.MENU_URL)
}

export const info = () => {
    return request.get<any, InfoResponse>(API.Info_URL)
}