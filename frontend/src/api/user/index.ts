import request from "@/utils/request"
import type { UserListResponse } from "./type"

enum API {
    User_URL = "/api/v1/users",
}

export const listUser = (params: any) => {
    return request.get<any, UserListResponse>(API.User_URL, { params })
}

export const createUser = (data: any) => {
    return request.post<any, any>(API.User_URL, data)
}

export const updateUser = (id: number, data: any) => {
    return request.put<any, any>(`${API.User_URL}/${id}`, data)
}

export const deleteUser = (id: number) => {
    return request.delete<any, any>(`${API.User_URL}/${id}`)
}

export const getUser = (id: number) => {
    return request.get<any, any>(`${API.User_URL}/${id}`)
}
