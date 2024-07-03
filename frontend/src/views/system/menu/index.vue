<template>
    <div>
        <div>
            <el-input v-model="query.name" clearable style="width: 240px;margin-right: 15px;" />
            <el-button type="primary" :icon="Search" @click="handleFilter">
                搜索
            </el-button>
            <el-button type="primary" v-auth="'menu-add'" :icon="Plus" @click="handleCreate(0)">
                添加
            </el-button>
        </div>

        <el-table :data="data" style="width: 100%; margin-top: 20px" row-key="id" border default-expand-all
            :table-layout="'auto'">
            <el-table-column align="center" prop="id" label="ID" sortable width="100" />
            <el-table-column align="center" prop="name" label="名称" sortable width="200" />
            <el-table-column align="center" prop="icon" label="类型" width="100">
                <template #default="{ row }">
                    <el-tag v-if="row.type === 'Menu'" type="success">菜单</el-tag>
                    <el-tag v-else-if="row.type === 'Button'" type="warning">按钮</el-tag>
                    <el-tag v-else-if="row.type === 'Link'" type="info">外链</el-tag>
                </template>
            </el-table-column>
            <el-table-column align="center" prop="icon" label="图标" width="60">
                <template #default="{ row }">
                    <el-icon v-if="row.icon">
                        <component :is="row.icon"></component>
                    </el-icon>
                </template>
            </el-table-column>
            <el-table-column align="center" prop="desc" label="描述" width="200" />
            <el-table-column align="center" prop="order" label="排序" sortable width="80" />
            <el-table-column align="center" prop="path" label="路径" />
            <el-table-column align="center" prop="create_time" label="创建时间" sortable width="180" />
            <el-table-column align="center" prop="update_time" label="更新时间" sortable width="180" />
            <el-table-column align="center" label="操作" fixed="right" width="220">
                <template #default="scope">
                    <el-button size="small" type="primary" v-if="scope.row.type === 'Menu'" v-auth="'menu-add'"
                        @click="handleCreate(scope.row.id)">
                        添加
                    </el-button>
                    <el-button size="small" v-auth="'menu-update'" @click="handleEdit(scope.row)">
                        编辑
                    </el-button>
                    <el-button size="small" v-auth="'menu-delete'" type="danger"
                        @click="handleDelete(scope.$index, scope.row)">
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>

        <el-dialog :title="textMap[dialogStatus]" v-model="dialogFormVisible" draggable style="max-width: 500px">
            <el-form :model="temp" label-width="auto">
                <el-form-item label="名称">
                    <el-input v-model="temp.name" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="类型">
                    <el-select v-model="temp.type" :disable="dialogStatus === 'update'" clearable placeholder="请选择类型">
                        <el-option label="菜单" value="Menu"></el-option>
                        <el-option label="按钮" value="Button"></el-option>
                        <el-option label="外链" value="Link"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item v-show="temp.type === 'Menu'" label="图标">
                    <SelectIcon v-model:icon-value="temp.icon" />
                </el-form-item>
                <el-form-item label="描述">
                    <el-input v-model="temp.desc" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item v-show="temp.type === 'Menu'" label="排序">
                    <el-input-number v-model="temp.order" :min="0"></el-input-number>
                </el-form-item>
                <el-form-item v-show="temp.type === 'Menu'" label="路径">
                    <el-input v-model="temp.path" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item v-show="temp.type === 'Menu'" label="组件">
                    <el-input v-model="temp.component" autocomplete="off"></el-input>
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
import { ref, onMounted, reactive } from 'vue'
import { treeMenu, createMenu, updateMenu, deleteMenu } from '@/api/menu/index'
import type { MenuTree } from '@/api/acl/type'
import SelectIcon from "@/components/SelectIcon/index.vue"
import { ElMessage, ElMessageBox, ElNotification } from 'element-plus'
import { Search, Plus } from '@element-plus/icons-vue'

const query = reactive({
    name: '',
})
const data = ref<MenuTree[]>([])
const temp = ref({
    id: 0,
    name: '',
    type: '',
    icon: '',
    desc: '',
    order: 0,
    path: '',
    component: '',
    parent_id: 0,
})
const resetTemp = () => {
    temp.value = {
        id: 0,
        name: '',
        type: '',
        icon: '',
        desc: '',
        order: 0,
        path: '',
        component: '',
        parent_id: 0,
    }
}

function getMenuTree() {
    treeMenu(query).then(response => {
        data.value = response.data
    })
}

function handleFilter() {
    getMenuTree()
}

const dialogStatus = ref('')
const dialogFormVisible = ref(false)
const textMap = {
    update: '编辑',
    create: '创建'
}

const handleCreate = (id: number) => {
    resetTemp()
    temp.value.parent_id = id
    dialogStatus.value = 'create'
    dialogFormVisible.value = true
}

function createData() {
    createMenu(temp.value).then(() => {
        getMenuTree()
        dialogFormVisible.value = false
        ElNotification({
            title: 'Success',
            message: '创建成功',
            type: 'success',
        })
    })
}

const handleEdit = (row: MenuTree) => {
    resetTemp()
    temp.value = Object.assign({}, row)
    dialogStatus.value = 'update'
    dialogFormVisible.value = true
}

function updateData() {
    updateMenu(temp.value.id, temp.value).then(() => {
        getMenuTree()
        dialogFormVisible.value = false
        ElNotification({
            title: 'Success',
            message: '修改成功',
            type: 'success',
        })
    })
}

const handleDelete = (index: number, row: MenuTree) => {
    ElMessageBox.confirm(
        '删除菜单及其子项，确定继续执行？',
        '警告',
        {
            confirmButtonText: '确定',
            type: 'warning',
        }
    )
        .then(() => {
            deleteMenu(row.id).then(() => {
                data.value.splice(index, 1)
                ElMessage({
                    type: 'success',
                    message: '删除成功',
                })
            })
        })
}

onMounted(() => {
    getMenuTree()
})

</script>
<style scoped></style>