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