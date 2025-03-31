import type { UserInfo } from "@/api/user/type"

export interface LoginParams {
    username: string,
    password: string,
}

interface LoginResponseData {
    access_token: string,
}

export interface LoginResponse {
    code: string,
    desc: string,
    data: LoginResponseData,
}

export interface MenuTree {
    id: number,
    name: string,
    desc: string,
    path: string,
    icon: string,
    type: string,
    order: number,
    component: string,
    parent_id: number,
    children: MenuTree[],
}

export interface MenuTreeResponse {
    code: string,
    desc: string,
    data: MenuTree[],
}


export interface InfoResponse {
    code: string,
    desc: string,
    data: UserInfo,
}