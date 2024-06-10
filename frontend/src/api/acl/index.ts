import request from "@/utils/request"
import type { LoginParams, LoginResponse, MenuTreeResponse } from "./type"

enum API {
    LOGIN_URL = '/api/v1/acl/login',
    MENU_URL = '/api/v1/acl/tree'
}

export const login = (data: LoginParams) => {
    return request.post<any, LoginResponse>(API.LOGIN_URL, data)
}

export const tree = () => {
    return request.get<any, MenuTreeResponse>(API.MENU_URL)
}