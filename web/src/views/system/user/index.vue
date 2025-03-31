<template>
    <div class="main-box">
        <TreeFilter label="name" title="部门列表(单选)" :data="treeFilterData" :default-value="listQuery.department_id"
            @change="changeTreeFilter" />

        <div class="table-box">
            <div>
                <el-input v-model="listQuery.name" placeholder="名称" style="width: 200px;"
                    @keyup.enter.native="handleFilter" />
                <el-select v-model="listQuery.state" placeholder="状态" style="width: 90px" @change="handleFilter">
                    <el-option v-for="item in stateOptions" :key="item" :label="statusDisplayFilter(item)"
                        :value="item" />
                </el-select>
                <el-button type="primary" :icon="Search" @click="handleFilter">
                    搜索
                </el-button>
                <el-button v-auth="'user-add'" type="primary" :icon="Plus" @click="handleCreate">
                    添加
                </el-button>
            </div>
            <el-table :data="list" style="width: 100%; margin-top: 20px" row-key="id" border default-expand-all
                :table-layout="'auto'">
                <el-table-column align="center" prop="id" label="ID" sortable width="100" />
                <el-table-column align="center" prop="name" label="账户" sortable width="200" />
                <el-table-column align="center" prop="desc" label="姓名" sortable width="200" />
                <el-table-column align="center" prop="avatar" label="头像" width="100px">
                    <template #default="scope">
                        <el-avatar size="large" :src="scope.row.avatar" />
                    </template>
                </el-table-column>
                <el-table-column label="状态" align="center" width="100">
                    <template #default="scope">
                        <el-tag :type="statusTypeFilter(scope.row.state)">
                            {{ statusDisplayFilter(scope.row.state) }}
                        </el-tag>
                    </template>
                </el-table-column>
                <el-table-column align="center" prop="create_time" label="创建时间" sortable width="180" />
                <el-table-column align="center" prop="update_time" label="更新时间" sortable width="180" />
                <el-table-column align="center" label="操作" fixed="right" width="220">
                    <template #default="{ row, $index }">
                        <el-button v-auth="'user-update'" type="primary" size="small" @click="handleUpdate(row)">
                            编辑
                        </el-button>
                        <el-button v-auth="'user-delete'" type="danger" size="small" @click="handleDelete(row, $index)">
                            删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>

            <el-pagination v-model:current-page="listQuery.page_num" v-model:page-size="listQuery.page_size"
                :page-sizes="[10, 20, 30, 40]" :disabled="false" background
                layout="total, sizes, prev, pager, next, jumper" :total="total" @size-change="handleSizeChange"
                @current-change="handleCurrentChange" style="padding-top: 10px;" />
        </div>

        <el-dialog :title="textMap[dialogStatus]" v-model="dialogFormVisible" style="width: 400px;">
            <el-form ref="dataForm" :model="temp" label-position="left" label-width="70px">
                <el-form-item label="账户" prop="title">
                    <el-input v-model="temp.name" :disabled="dialogStatus === 'update'" />
                </el-form-item>
                <el-form-item label="姓名" prop="title">
                    <el-input v-model="temp.desc" />
                </el-form-item>
                <el-form-item v-if="dialogStatus === 'create'" label="密码" prop="title">
                    <el-input type="password" v-model="temp.password" show-password />
                </el-form-item>
                <el-form-item label="头像">
                    <el-input v-model="temp.avatar" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea"
                        placeholder="请输入" />
                </el-form-item>
                <el-form-item label="状态">
                    <el-select v-model="temp.state" class="filter-item" placeholder="请选择">
                        <el-option label="有效" :value="1"></el-option>
                        <el-option label="无效" :value="2"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item lable="角色">
                    <el-select v-model="temp.roleIds" multiple>
                        <el-option v-for="item in roles" :key="item.id" :label="item.name" :value="item.id">
                        </el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <template #footer>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="dialogFormVisible = false">
                        取消
                    </el-button>
                    <el-button type="primary" @click="dialogStatus === 'create' ? createData() : updateData()">
                        确认
                    </el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>

<script setup lang='ts'>
import { ref, reactive, onMounted } from 'vue'
import TreeFilter from "@/components/TreeFilter/index.vue"
import { listDepartment } from '@/api/department/index'
import { listUser, createUser, updateUser, deleteUser,getUser } from '@/api/user/index'
import { User } from '@/api/user/type'
import { Search, Plus } from '@element-plus/icons-vue'
import { ElNotification } from 'element-plus'
import { Role } from '@/api/role/type'
import { listRole } from '@/api/role'

const list = ref<User[]>([])
const total = ref(0)
const listLoading = ref(true)
const listQuery = reactive({
    department_id: 0,
    name: '',
    state: 0,
    page_size: 10,
    page_num: 1,
})
const temp = ref({
    id: 0,
    name: '',
    desc: '',
    avatar: '',
    password: '',
    state: 1,
    department_id: 0,
    roles: [],
    roleIds: []
})
const resetTemp = () => {
    temp.value = {
        id: 0,
        name: '',
        desc: '',
        avatar: '',
        password: '',
        state: 1,
        department_id: listQuery.department_id,
        roles: [],
        roleIds: [],
    }
}

const treeFilterData = ref<any>([]);

const dialogStatus = ref('')
const textMap = {
    update: '编辑',
    create: '创建',
}
const dialogFormVisible = ref(false)
const roles = ref<Role[]>([])
const stateOptions = [0, 1, 2]
const statusTypeFilter = (status: number) => {
    const statusMap = {
        0: 'info',
        1: 'success',
        2: 'danger'
    }
    return statusMap[status]
}
const statusDisplayFilter = (status: number) => {
    const statusMap = {
        0: '全部',
        1: '有效',
        2: '无效'
    }
    return statusMap[status]
}

// 树形筛选切换
const changeTreeFilter = (val: number) => {
    if (val > 0) {
        listQuery.department_id = val
        temp.value.department_id = val
    } else {
        listQuery.department_id = 0
    }
    getUserList()
}

function getDepartmentList() {
    listDepartment(listQuery).then(response => {
        treeFilterData.value = response.data.list
    })
}

function getUserList() {
    listUser(listQuery).then(response => {
        list.value = response.data.list
        total.value = response.data.total
        listLoading.value = false
    })
}

function handleFilter() {
    getUserList()
}

function handleCreate() {
    resetTemp()
    listRole({}).then(response => {
        roles.value = response.data.roles
    })
    dialogStatus.value = 'create'
    dialogFormVisible.value = true
}

function handleUpdate(row) {
    resetTemp()
    listRole({}).then(response => {
        roles.value = response.data.roles
    })
    getUser(row.id).then(response => {
        temp.value = response.data
        temp.value.roleIds = response.data.roles.map(item=>item.id)
    })
    dialogStatus.value = 'update'
    dialogFormVisible.value = true
}

function handleDelete(row, index) {
    deleteUser(row.id).then(() => {
        list.value.splice(index, 1)
    })
}

function createData() {
    createUser(temp.value).then(() => {
        getUserList()
        dialogFormVisible.value = false
        ElNotification({
            title: 'Success',
            message: '创建成功',
            type: 'success',
        })
    })
}

function updateData() {
    updateUser(temp.value.id, temp.value).then(() => {
        getUserList()
        dialogFormVisible.value = false
        ElNotification({
            title: 'Success',
            message: '修改成功',
            type: 'success',
        })
    })
}

onMounted(() => {
    getDepartmentList()
    getUserList()
})

const handleSizeChange = (val: number) => {
    listQuery.page_size = val
    getUserList()
}
const handleCurrentChange = (val: number) => {
    listQuery.page_num = val
    getUserList()
}
</script>
<style scoped></style>