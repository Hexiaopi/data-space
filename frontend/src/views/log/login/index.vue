<template>
    <div>
        <div>
            <el-input v-model="listQuery.user_name" placeholder="用户名" style="width: 200px;"
                @keyup.enter.native="handleFilter" />
            <el-select v-model="listQuery.login_result" placeholder="登录结果" style="width: 90px" @change="handleFilter">
                <el-option v-for="item in login_resultOptions" :key="item" :label="statusDisplayFilter(item)"
                    :value="item" />
            </el-select>
            <el-button type="primary" :icon="Search" @click="handleFilter">
                搜索
            </el-button>
        </div>

        <el-table :data="list" style="width: 100%; margin-top: 20px" row-key="id" border default-expand-all
            :table-layout="'auto'">
            <el-table-column align="center" prop="id" label="ID" sortable width="100" />
            <el-table-column align="center" prop="user_name" label="用户名" sortable width="200" />
            <el-table-column align="center" prop="remote_ip" label="访问IP" min-width="100px" />
            <el-table-column align="center" prop="user_agent" label="用户代理" width="200" />
            <el-table-column label="登录结果" align="center" width="100">
                <template #default="scope">
                    <el-tag :type="statusTypeFilter(scope.row.login_result)">
                        {{ statusDisplayFilter(scope.row.login_result) }}
                    </el-tag>
                </template>
            </el-table-column>
            <el-table-column align="center" prop="result_detail" label="结果详情" min-width="100px" />
            <el-table-column align="center" prop="create_time" label="创建时间" sortable width="180" />
            <el-table-column align="center" prop="update_time" label="更新时间" sortable width="180" />
        </el-table>

        <el-pagination v-model:current-page="listQuery.page_num" v-model:page-size="listQuery.page_size"
            :page-sizes="[10, 20, 30, 40]" :disabled="false" background layout="total, sizes, prev, pager, next, jumper"
            :total="total" @size-change="handleSizeChange" @current-change="handleCurrentChange"
            style="padding-top: 10px;" />

    </div>
</template>

<script setup lang='ts'>
import { ref, reactive, onMounted } from 'vue'
import { listLoginLog } from '@/api/log/index'
import { Search } from '@element-plus/icons-vue'

const list = ref<Department[]>([])
const total = ref(0)
const listLoading = ref(true)
const listQuery = reactive({
    page_num: 1,
    page_size: 10,
    user_name: '',
    login_result: 0,
})

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
        1: '成功',
        2: '失败'
    }
    return statusMap[status]
}
const login_resultOptions = [0, 1, 2]

function getList() {
    listLoginLog(listQuery).then(response => {
        list.value = response.data.list
        total.value = response.data.total
        listLoading.value = false
    })
}

function handleFilter() {
    getList()
}


onMounted(() => {
    getList()
})
const handleSizeChange = (val: number) => {
    listQuery.page_size = val
    getList()
}
const handleCurrentChange = (val: number) => {
    listQuery.page_num = val
    getList()
}
</script>
<style scoped></style>