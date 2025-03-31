import request from "@/utils/request"
import type { DepartmentListResponse } from "./type"

enum API {
    Department_URL = '/api/v1/departments',
}

export const listDepartment = (params: any) => {
    return request.get<any, DepartmentListResponse>(API.Department_URL, { params })
}

export const createDepartment = (data: any) => {
    return request.post<any, any>(API.Department_URL, data)
}

export const updateDepartment = (id: number, data: any) => {
    return request.put<any, any>(`${API.Department_URL}/${id}`, data)
}

export const deleteDepartment = (id: number) => {
    return request.delete<any, any>(`${API.Department_URL}/${id}`)
}