import request from "@/utils/request"
import type { LoginLogListResponse } from "./type"

enum API {
    LoginLog_URL = '/api/v1/log/logins',
}

export const listLoginLog = (params: any) => {
    return request.get<any, LoginLogListResponse>(API.LoginLog_URL, { params })
}