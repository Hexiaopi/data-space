export interface Department {
    id: number,
    name: string,
    create_time: string,
    update_time: string,
}

export interface DepartmentListData {
    list: Department[],
    total: number,
}

export interface DepartmentListResponse {
    code: string,
    desc: string,
    data: DepartmentListData,
}