export interface LoginLog {
    id: number,
    user_name: string,
    user_agent: string,
    remote_ip: string
    create_time: string,
    update_time: string,
}

export interface LoginLogListData {
    list: LoginLog[],
    total: number,
}

export interface LoginLogListResponse {
    code: string,
    desc: string,
    data: LoginLogListData,
}