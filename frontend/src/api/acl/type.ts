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
    name: string,
    path: string,
    icon: string,
    type: string,
    component: string,
    children: MenuTree[],
}

export interface MenuTreeResponse {
    code: string,
    desc: string,
    data: MenuTree[],
}