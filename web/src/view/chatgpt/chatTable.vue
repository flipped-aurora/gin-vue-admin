<template>
  <div  class="gva-table-box">
    <div v-if="!chatToken">
      <warning-bar title="在资源权限中将此角色的资源权限清空 或者不包含创建者的角色 即可屏蔽此客户资源的显示" />
      <el-input class="query-ipt" v-model="ipt" placeholder="请输入token进行保存" clearable />
      <el-button type="primary" @click="save">保存</el-button>
      <div class="secret">
        <p>
          <el-icon><FolderChecked /></el-icon>
          获取SK链接：<el-link type="success">{{ skURL }}</el-link>
        </p>
        <p>
          <el-icon><Connection /></el-icon>
          GVA会严格保护您的隐私，不会将敏感数据透露给任第三方
        </p>
      </div>
    </div>
    <div v-else>
      <el-form :model="form" label-width="120px">
        <el-form-item label="删除当前sk：">
          <el-popover placement="top" width="160">
            <p>确定要删除并返回吗？</p>
            <div style="text-align: right; margin-top: 8px;">
              <el-button type="primary" @click="deleteSK">确定</el-button>
            </div>
            <template #reference>
              <el-button type="primary" link icon="delete">删除</el-button>
            </template>
          </el-popover>
        </el-form-item>
        <el-form-item label="查询db名称：">
          <el-input
              v-model="form.dbname"
              placeholder="请输入要查询的内容"
              class="input-with-select query-ipt"
              clearable
          >
            <template #prepend>
              <el-select v-model="form.select" placeholder="请选择库" style="width: 115px">
                <el-option
                    v-for="(item, index) in dbArr"
                    :key="index"
                    :label="item.label"
                    :value="item.value"
                />
              </el-select>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="查询db描述：">
          <el-input
              v-model="form.textarea"
              :autosize="{ minRows: 2, maxRows: 4 }"
              type="textarea"
              clearable
              placeholder="Please input"
          />
        </el-form-item>
        <el-button type="primary" @click="handleQueryTable">查询</el-button>
      </el-form>
      <div class="tables">
        <el-table
            ref="multipleTable"
            :data="tableData"
            style="width: 100%"
            tooltip-effect="dark"
            row-key="ID"
            v-if="tableData.length"
        >
<!--          <el-table-column-->
<!--              v-for="(item, index) in tableData"-->
<!--              :key="index"-->
<!--              :prop="item.address"-->
<!--              label="Address" />-->
          <el-table-column align="left" label="address" prop="address" width="120" />
        </el-table>
        <p v-else class="text">请在对话框输入你需要AI帮你查询的内容：）</p>
    </div>
  </div>
  </div>
</template>

<script setup>
import { ElMessage } from 'element-plus'
import { getTable } from '@/api/chatgpt'
import { ref } from 'vue'

const chatToken = ref(null)
const ipt = ref('')
const skURL = ref('www.bing.com')
// 页面初始化的时候去后台接口查询拿sk 如果不存在就展示保存页面
const getSK = () => {
  //此处模拟sk不存在的情况 需要重新输入
  chatToken.value = false
}
getSK()
const save = async () => {
  chatToken.value = true
  // const res = await getTable({
  //   dbname: 'gva',
  //   chat: `帮我找到角色id为${ipt.value}的用户`,
  // })
  // if(res.code === 0) {
  //   ElMessage({
  //     message: '保存成功',
  //     type: 'success',
  //   })
  // }

}

const form = ref({
  dbname: '',
  select: '',
  textarea: ''
})
const dbArr = ref([
  {
    label: '测试1',
    value: 1
  },
  {
    label: '测试2',
    value: 2
  },
  {
    label: '测试3',
    value: 3
  }
])
const tableData = ref([])
const data = ref({
  currentPage: 1,
  pageSize: 10,
  totalCount: 0,
})
const deleteSK = () => {
  ipt.value = ''
  Object.keys(form).forEach(key => (form[key] = ''))
  chatToken.value = false
  tableData.value = []
}

const handleQueryTable = () => {
  // 把所选内容提交至后台
  console.log(form.value)
  let obj = [
    {
      date: '2016-05-03',
      name: 'Tom',
      address: 'No. 189, Grove St, Los Angeles',
    }
  ]
  // 根据后台返回值动态渲染表格
  tableData.value = obj
}
</script>

<style scoped lang="scss">
.secret{
  width: 96%;
  padding: 30px;
  margin-top: 20px;
  background: #F5F5F5;
  p {
    line-height: 30px;
  }
}
.query-ipt{
  width: 300px;
  margin-right: 30px;
}
.content{
  p {
    font-size: 16px;
    line-height: 20px;
  }
  padding: 10px;
  width: 100%;
  background: #F5F5F5;
  margin-top: 30px;
}
.tables{
  width: 100%;
  margin-top: 30px;
  .text{
    font-size: 18px;
    color: #6B7687;
    text-align: center;
  }
}
</style>
