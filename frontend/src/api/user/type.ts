import type { Role } from "@/api/role/type"

export interface User {
    id: number,
    name: string,
    avatar: string,
    create_time: string,
    update_time: string,
    roles: [],
    current_role: Role,
}
