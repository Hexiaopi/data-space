export interface Role {
    id: number,
    name: string,
    create_time: string,
    update_time: string,
}

export interface RoleList {
    roles: Role[],
    total: number,
}
