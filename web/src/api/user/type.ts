import type { Role } from "@/api/role/type"

export interface UserInfo {
    id: number,
    name: string,
    avatar: string,
    create_time: string,
    update_time: string,
    roles: [],
    current_role: Role,
}

export interface User {
    id: number,
    name: string,
    avatar: string,
    create_time: string,
    update_time: string,
}

export interface UserListData {
    list: User[],
    total: number
}

export interface UserListResponse {
    code: string,
    desc: string,
    data: UserListData,
}