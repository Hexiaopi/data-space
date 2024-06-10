import request from "@/utils/request"
import type { InfoResponse } from "./type"

enum API {
    Info_URL = "/api/v1/user",
}

export const info = () => {
    return request.get<any, InfoResponse>(API.Info_URL)
}