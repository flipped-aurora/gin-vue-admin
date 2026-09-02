<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
          <el-form-item label="用户ID">
              <el-input v-model.number="searchInfo.userId" placeholder="搜索用户ID" />
          </el-form-item>
        <el-form-item label="状态">
             <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
                 <el-option label="有效" :value="true" />
                 <el-option label="无效" :value="false" />
             </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDrawer">签发</el-button>
      </div>
      <el-table
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="用户" min-width="150">
             <template #default="scope">
                 {{ scope.row.user.nickName }} ({{ scope.row.user.userName }})
             </template>
        </el-table-column>
        <el-table-column align="left" label="角色ID" prop="authorityId" width="100" />
        <el-table-column align="left" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status ? 'success' : 'danger'">
              {{ scope.row.status ? '有效' : '已作废' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="过期时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.expiresAt) }}</template>
        </el-table-column>
         <el-table-column align="left" label="备注" prop="remark" min-width="150" show-overflow-tooltip />
        <el-table-column align="left" label="操作" width="220">
          <template #default="scope">
            <el-button type="primary" link icon="link" @click="openCurl(scope.row)">Curl示例</el-button>
            <el-popover v-if="scope.row.status" v-model:visible="scope.row.visible" placement="top" width="160">
              <p>确定要作废吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button size="small" type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button size="small" type="primary" @click="invalidateToken(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button icon="delete" type="danger" link @click="scope.row.visible = true">作废</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-drawer v-model="drawerVisible" size="400px" title="签发 API Token">
         <el-form ref="formRef" :model="form" label-width="80px">
             <el-form-item label="用户" required>
                 <el-select 
                    v-model="form.userId" 
                    placeholder="请选择用户" 
                    filterable 
                    style="width:100%"
                    @change="handleUserChange"
                 >
                     <el-option
                        v-for="item in userOptions"
                        :key="item.ID"
                        :label="`${item.nickName} (${item.userName})`"
                        :value="item.ID"
                     />
                 </el-select>
             </el-form-item>
             <el-form-item label="角色" required>
                 <el-select v-model="form.authorityId" placeholder="请选择角色" style="width:100%" :disabled="!form.userId">
                     <el-option
                        v-for="item in authorityOptions"
                        :key="item.authorityId"
                        :label="`${item.authorityName} (${item.authorityId})`"
                        :value="item.authorityId"
                     />
                 </el-select>
             </el-form-item>
            <el-form-item label="有效期">
                <el-select v-model="form.days" placeholder="请选择" style="width:100%">
                    <el-option label="1天" :value="1" />
                    <el-option label="7天" :value="7" />
                    <el-option label="30天" :value="30" />
                    <el-option label="90天" :value="90" />
                    <el-option label="永久" :value="-1" />
                </el-select>
            </el-form-item>
            <el-form-item label="备注">
                <el-input v-model="form.remark" type="textarea" />
            </el-form-item>
         </el-form>
         <template #footer>
             <div style="flex: auto">
                 <el-button @click="drawerVisible = false">取消</el-button>
                 <el-button type="primary" @click="submitIssuer">签发JWT</el-button>
             </div>
         </template>
    </el-drawer>

    <el-dialog v-model="tokenDialogVisible" title="签发成功" width="500px">
        <div style="text-align: center; margin-bottom: 20px;">
            <el-alert title="请立即复制保存，关闭后将无法再次查看完整Token" type="warning" :closable="false" show-icon />
        </div>
        <el-input type="textarea" :rows="6" v-model="tokenResult" readonly />
        <template #footer>
            <el-button @click="copyText(tokenResult)">复制</el-button>
            <el-button type="primary" @click="tokenDialogVisible = false">关闭</el-button>
        </template>
    </el-dialog>

    <el-drawer v-model="curlDrawerVisible" size="500px" title="Curl 示例">
        <div style="padding: 10px;">
            <p style="margin-bottom: 10px;">Header 方式:</p>
            <el-input type="textarea" :rows="4" v-model="curlHeader" readonly />
            <el-button style="margin-top: 5px;" size="small" @click="copyText(curlHeader)">复制</el-button>
            
            <el-divider />
            
            <p style="margin-bottom: 10px;">Cookie 方式:</p>
            <el-input type="textarea" :rows="4" v-model="curlCookie" readonly />
            <el-button style="margin-top: 5px;" size="small" @click="copyText(curlCookie)">复制</el-button>
        </div>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  getApiTokenList,
  createApiToken,
  deleteApiToken
} from '@/api/sysApiToken'
import { getUserList } from '@/api/user'
import { ref, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const drawerVisible = ref(false)
const tokenDialogVisible = ref(false)
const tokenResult = ref('')
const curlDrawerVisible = ref(false)
const curlHeader = ref('')
const curlCookie = ref('')

const form = ref({
    userId: '',
    authorityId: '',
    days: 30,
    remark: ''
})

const userOptions = ref([])
const authorityOptions = ref([])

const getTableData = async () => {
  const table = await getApiTokenList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const openDrawer = async () => {
    form.value = { userId: '', authorityId: '', days: 30, remark: '' }
    authorityOptions.value = []
    drawerVisible.value = true
    if (userOptions.value.length === 0) {
        const res = await getUserList({ page: 1, pageSize: 999 })
        if (res.code === 0) {
            userOptions.value = res.data.list
        }
    }
}

const handleUserChange = (val) => {
    form.value.authorityId = ''
    const user = userOptions.value.find(u => u.ID === val)
    if (user) {
        authorityOptions.value = user.authorities || []
        // 默认选中第一个
        if (authorityOptions.value.length > 0) {
            form.value.authorityId = authorityOptions.value[0].authorityId
        }
    } else {
        authorityOptions.value = []
    }
}

const submitIssuer = async () => {
    if (!form.value.userId || !form.value.authorityId) {
        ElMessage.warning("请选择用户和角色")
        return
    }
    const res = await createApiToken(form.value)
    if (res.code === 0) {
        tokenResult.value = res.data.token
        drawerVisible.value = false
        tokenDialogVisible.value = true
        getTableData()
    }
}

const invalidateToken = async (row) => {
    row.visible = false
    const res = await deleteApiToken({ ID: row.ID })
    if (res.code === 0) {
        ElMessage.success("作废成功")
        getTableData()
    }
}

const openCurl = (row) => {
    // 假设 API Host 为当前 origin
    const origin = window.location.origin
    // 构造示例 URL
    const url = `${origin}/api/menu/getMenu`
    
    curlHeader.value = `curl -X POST "${url}" \ 
  -H "x-token: ${row.token}" \ 
  -H "Content-Type: application/json"`
    
    curlCookie.value = `curl -X POST "${url}" \ 
  -b "x-token=${row.token}" \ 
  -H "Content-Type: application/json"`

    curlDrawerVisible.value = true
}

const copyText = (text) => {
    if (!text) return
    const input = document.createElement('textarea')
    input.value = text
    document.body.appendChild(input)
    input.select()
    document.execCommand('copy')
    document.body.removeChild(input)
    ElMessage.success('复制成功')
}

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

getTableData()
</script>

<style scoped>
</style>
