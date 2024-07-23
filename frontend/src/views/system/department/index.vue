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
            <el-button v-auth="'department-add'" type="primary" :icon="Plus" @click="handleCreate">
                添加
            </el-button>
        </div>

        <el-table :data="list" style="width: 100%; margin-top: 20px" row-key="id" border default-expand-all
            :table-layout="'auto'">
            <el-table-column align="center" prop="id" label="ID" sortable width="100" />
            <el-table-column align="center" prop="name" label="名称" sortable width="200" />
            <el-table-column align="center" prop="desc" label="描述" min-width="100px" />
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
                    <el-button v-auth="'department-update'" type="primary" size="small" @click="handleUpdate(row)">
                        编辑
                    </el-button>
                    <el-button v-auth="'department-delete'" type="danger" size="small"
                        @click="handleDelete(row, $index)">
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>

        <el-pagination v-model:current-page="listQuery.page_num" v-model:page-size="listQuery.page_size"
            :page-sizes="[10, 20, 30, 40]" :disabled="false" background layout="total, sizes, prev, pager, next, jumper"
            :total="total" @size-change="handleSizeChange" @current-change="handleCurrentChange"
            style="padding-top: 10px;" />


        <el-dialog :title="textMap[dialogStatus]" v-model="dialogFormVisible" style="width: 400px;">
            <el-form ref="dataForm" :model="temp" label-position="left" label-width="70px">
                <el-form-item label="名称" prop="title">
                    <el-input v-model="temp.name" />
                </el-form-item>
                <el-form-item label="描述">
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
    </div>
</template>

<script setup lang='ts'>
import { ref, reactive, onMounted } from 'vue'
import { listDepartment, createDepartment, updateDepartment, deleteDepartment } from '@/api/department/index'
import type { Department } from '@/api/department/type'
import { Search, Plus } from '@element-plus/icons-vue'
import { ElNotification } from 'element-plus'

const list = ref<Department[]>([])
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
    create: '创建',
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
const stateOptions = [0, 1, 2]
const dialogFormVisible = ref(false)

function getDepartmentList() {
    listDepartment(listQuery).then(response => {
        list.value = response.data.list
        total.value = response.data.total
        listLoading.value = false
    })
}

function handleFilter() {
    getDepartmentList()
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

function handleDelete(row, index) {
    deleteDepartment(row.id).then(() => {
        list.value.splice(index, 1)
    })
}
function createData() {
    createDepartment(temp.value).then(() => {
        getDepartmentList()
        dialogFormVisible.value = false
        ElNotification({
            title: 'Success',
            message: '创建成功',
            type: 'success',
        })
    })
}

function updateData() {
    updateDepartment(temp.value.id, temp.value).then(() => {
        getDepartmentList()
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
})
const handleSizeChange = (val: number) => {
    listQuery.page_size = val
    getDepartmentList()
}
const handleCurrentChange = (val: number) => {
    listQuery.page_num = val
    getDepartmentList()
}
</script>
<style scoped></style>