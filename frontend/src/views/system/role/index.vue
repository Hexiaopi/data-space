<template>
    <div>
        <div>
            <el-input v-model="listQuery.name" placeholder="名称" style="width: 200px;"
                @keyup.enter.native="handleFilter" />
            <el-select v-model="listQuery.state" placeholder="状态" style="width: 90px" @change="handleFilter">
                <el-option v-for="item in stateOptions" :key="item" :label="statusDisplayFilter(item)" :value="item" />
            </el-select>
            <el-button type="primary" :icon="Search" @click="handleFilter">
                搜索
            </el-button>
            <el-button v-auth="'role-add'" type="primary" :icon="Plus" @click="handleCreate">
                添加
            </el-button>
        </div>
        <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row
            style="width: 100%;">
            <el-table-column label="ID" prop="id" align="center" min-width="50px">
            </el-table-column>
            <el-table-column label="角色名称" prop="name" align="center" width="100px">
            </el-table-column>
            <el-table-column label="角色描述" prop="desc" align="center" min-width="100px">
            </el-table-column>
            <el-table-column label="状态" align="center" width="100">
                <template #default="scope">
                    <el-tag :type="statusTypeFilter(scope.row.state)">
                        {{ statusDisplayFilter(scope.row.state) }}
                    </el-tag>
                </template>
            </el-table-column>
            <el-table-column label="创建时间" prop="create_time" align="center" min-width="100px">
            </el-table-column>
            <el-table-column label="更新时间" prop="update_time" align="center" min-width="100px">
            </el-table-column>
            <el-table-column label="操作" align="center" width="200">
                <template #default="{ row, $index }">
                    <el-button v-auth="'role-update'" type="primary" size="small" @click="handleUpdate(row)">
                        编辑
                    </el-button>
                    <el-button type="success" size="small" @click="handleRoleMenu(row.id)">菜单</el-button>
                    <el-button v-auth="'role-delete'" type="danger" size="small" @click="handleDelete(row, $index)">
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>

        <el-pagination v-model:current-page="listQuery.page_num" v-model:page-size="listQuery.page_size"
            :page-sizes="[10, 20, 30, 40]" :small="true" :disabled="false" :background="false"
            layout="total, sizes, prev, pager, next, jumper" :total="total" @size-change="handleSizeChange"
            @current-change="handleCurrentChange" />

        <el-dialog :title="textMap[dialogStatus]" v-model="dialogFormVisible" style="width: 400px;">
            <el-form ref="dataForm" :model="temp" label-position="left" label-width="70px">
                <el-form-item label="角色名称" prop="title">
                    <el-input v-model="temp.name" />
                </el-form-item>
                <el-form-item label="角色描述">
                    <el-input v-model="temp.desc" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea"
                        placeholder="请输入" />
                </el-form-item>
                <el-form-item label="状态">
                    <el-select v-model="temp.state" class="filter-item" placeholder="请选择">
                        <el-option label="有效" :value="1"></el-option>
                        <el-option label="无效" :value="2"></el-option>
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

        <el-dialog title="菜单权限" v-model="dialogMenuVisible" style="width: 400px;">
            <el-form ref="menuForm" :model="temp" label-position="left" label-width="70px">
                <el-tree :data="menus" ref="menuTree" show-checkbox node-key="id" :default-checked-keys="menuIds"
                    default-expand-all :expand-on-click-node="false" :props="defaultMenuProps">
                    <template #default="{ node, data }">
                        <span class="custom-tree-node">
                            <el-tag v-if="data.type === 'Menu'" type="success">菜单</el-tag>
                            <el-tag v-else-if="data.type === 'Button'" type="warning">按钮</el-tag>
                            <el-tag v-else-if="data.type === 'Link'" type="danger">链接</el-tag>
                            <el-tag v-else type="info">未知</el-tag>
                            <span>{{ node.label }}</span>
                        </span>
                    </template>
                </el-tree>
            </el-form>
            <template #footer>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="dialogMenuVisible = false">
                        取消
                    </el-button>
                    <el-button type="primary" @click="updateRoleMenu()">
                        确认
                    </el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>

<script setup lang='ts'>
import { reactive, ref, onMounted } from 'vue'
import { listRole, createRole, updateRole, deleteRole, getRole } from '@/api/role'
import { treeMenu } from '@/api/menu'
import { Search, Plus } from '@element-plus/icons-vue'
import { ElNotification } from 'element-plus'

const list = ref([])
const menus = ref([])
const total = ref(0)
const listLoading = ref(true)
const listQuery = reactive({
    page_num: 1,
    page_size: 10,
    name: '',
    state: 0,
})
const temp = ref({
    id: 0,
    name: '',
    desc: '',
    state: 1,
})
const resetTemp = () => {
    temp.value = {
        id: 0,
        name: '',
        desc: '',
        state: 1,
    }
}
const dialogStatus = ref('')
const textMap = {
    update: '编辑',
    create: '创建'
}
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
const dialogFormVisible = ref(false)
const stateOptions = [0, 1, 2]

const dialogMenuVisible = ref(false)
const defaultMenuProps = {
    children: 'children',
    label: 'desc',
}
const menuIds = ref([])

function handleRoleMenu(id: number) {
    getRole(id).then(response => {
        let menus = []
        for (let i = 0; i < response.data.menus.length; i++) {
            menus.push(response.data.menus[i].id)
        }
        menuIds.value = menus
        dialogMenuVisible.value = true
    })

}

onMounted(() => {
    getRoles()
    getMenus()
})

function getRoles() {
    listLoading.value = true
    listRole(listQuery).then(response => {
        list.value = response.data.roles
        total.value = response.data.total
        listLoading.value = false
    })
}

function getMenus() {
    treeMenu({}).then(response => {
        menus.value = response.data
    })
}

function handleCreate() {
    resetTemp()
    console.log(temp.value)
    dialogStatus.value = 'create'
    dialogFormVisible.value = true
}

function handleUpdate(row) {
    resetTemp()
    temp.value = Object.assign({}, row)
    dialogStatus.value = 'update'
    dialogFormVisible.value = true
}

function handleFilter() {
    listQuery.page_num = 1
    getRoles()
}

function handleDelete(row, index) {
    deleteRole(row.id).then(() => {
        list.value.splice(index, 1)
    })
}

function createData() {
    createRole(temp.value).then(() => {
        getRoles()
        dialogFormVisible.value = false
        ElNotification({
            title: 'Success',
            message: '创建成功',
            type: 'success',
        })
    })
}

function updateData() {
    updateRole(temp.value.id, temp.value).then(() => {
        getRoles()
        dialogFormVisible.value = false
        ElNotification({
            title: 'Success',
            message: '修改成功',
            type: 'success',
        })
    })
}

const handleSizeChange = (val: number) => {
    listQuery.page_size = val
    getRoles()
}
const handleCurrentChange = (val: number) => {
    listQuery.page_num = val
    getRoles()
}
</script>
<style scoped>
.custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    font-size: 14px;
    padding-right: 8px;
}
</style>