import request from "@/utils/request"
import type { MenuTreeResponse } from "@/api/acl/type"

enum API {
    Menu_URL = '/api/v1/menus',
}

export const treeMenu = (params: any) => {
    return request.get<any, MenuTreeResponse>(`${API.Menu_URL}/tree`, { params })
}

export const createMenu = (data: any) => {
    return request.post<any, any>(API.Menu_URL, data)
}

export const updateMenu = (id: number, data: any) => {
    return request.put<any, any>(`${API.Menu_URL}/${id}`, data)
}

export const deleteMenu = (id: number) => {
    return request.delete<any, any>(`${API.Menu_URL}/${id}`)
}