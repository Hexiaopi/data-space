import type { Role } from "@/api/role/type"

interface User {
    id: number,
    name: string,
    avatar: string,
    create_time: string,
    update_time: string,
    roles: [],
    current_role: Role,
}

export interface InfoResponse {
    code: string,
    desc: string,
    data: User,
}