import request from "@/utils/request"
import type { RoleList } from "./type"

enum API {
    Role_URL = "/api/v1/roles",
}

export const listRole = (params: any) => {
    return request.get<any, RoleList>(API.Role_URL, { params })
}

export const createRole = (data: any) => {
    return request.post<any, any>(API.Role_URL, data)
}

export const updateRole = (id: number, data: any) => {
    return request.put<any, any>(`${API.Role_URL}/${id}`, data)
}

export const getRole = (id: number) => {
    return request.get<any, any>(`${API.Role_URL}/${id}`)
}

export const deleteRole = (id: number) => {
    return request.delete<any, any>(`${API.Role_URL}/${id}`)
}