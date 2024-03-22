<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
          <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
          <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
              <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
              <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
              <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
          </template>
          </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" />
          <el-table-column align="center" label="ID" prop="ID" width="80" />
          <el-table-column align="center" label="标题" prop="title" width="240" />
          <el-table-column align="center" label="栏目名称" prop="catemenus.name" width="180" />
          <el-table-column align="center" label="缩略图" prop="thumb" width="240" >
            <template #default="scope">
              <CustomPic style="margin-top:8px" :pic-src="scope.row.thumb" pic-type="swiper" />
            </template>
          </el-table-column>
          <el-table-column align="center" label="是否显示" prop="enable" width="120">
              <template #default="scope">{{ formatBoolean(scope.row.enable) }}</template>
          </el-table-column>
          <el-table-column align="center" label="日期" width="180">
              <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
          </el-table-column>
          <el-table-column align="center" label="操作">
              <template #default="scope">
              <el-button type="primary" link icon="edit" class="table-button" @click="updateClassFunc(scope.row)">修改</el-button>
              <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
              </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" v-if="dialogFormVisible" :before-close="closeDialog" title="添加内容">
      <el-form :model="formData" label-position="right" ref="elFormRef" :inline="true" :rules="rule" label-width="90px">
        <el-form-item label="栏目分类:" style="width:45%">
          <el-cascader
            v-model="formData.cateId"
            style="width:100%"
            :options="menuOption"
            :props="{ label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item label="标题:"  prop="title" style="width: 45%;">
          <el-input v-model="formData.title"   placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="缩略图:"  prop="thumb" style="width: 45%;">
            <CustomPic style="margin-top:8px;" v-if="formData.thumb" :pic-src="formData.thumb" pic-type="swiper"/>
            <el-button v-else  type="primary" @click="openHeaderChange">选择</el-button>
            <el-input   v-model="formData.thumb" :clearable="true"  placeholder="请输入" style="width: 65%;" />
        </el-form-item>
        <el-form-item label="是否显示:"  prop="enable" style="width: 45%;">
          <el-switch v-model="formData.enable" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="介绍:"  prop="desc" style="width: 95%;">
          <el-input v-model="formData.desc" type="textarea"   placeholder="简介"/>
        </el-form-item>
        <el-form-item label="内容:"  prop="content">
            <!-- <Wangeditor  @update:content="formData.content = $event" :content="formData.content"/> -->
            <Tinymceditor ref="editor" @update:content="formData.content = $event" :content="formData.content" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="formData" :target-key="`thumb`" />
  </div>
</template>

<script>
export default {
  name: 'Class'
}
</script>

<script setup>
import {
  createClass,
  deleteClass,
  deleteClassByIds,
  updateClass,
  findClass,
  getClassList
} from '@/api/class'
import {
  getCateMenusList
} from '@/api/cateMenus'

// 全量引入格式化工具 请按需保留
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
// import Wangeditor from '@/components/wangeditor/index.vue'
import Tinymceditor from '@/components/tinymce/index.vue'
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive,shallowRef,onBeforeUnmount,watch } from 'vue'


// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  title: '',
  cateId: 1,
  name:"",
  thumb:"",
  videoUrl: '',
  desc:'',
  content: ref('<p>hello</p>'),
  enable: false,
  catemenus: {}
})

const props = defineProps({
  defaultCatid:{
    type:String,
    default:"0"
  }
});


// 验证规则
const rule = reactive({
    title : [{
        required: true,
        message: '',
        trigger: ['input','blur'],
    }],
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const cateMenusList = ref([])
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.enable === ""){
      searchInfo.value.enable=null
  }
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getClassList({ page: page.value, pageSize: pageSize.value,cateId:props.defaultCatid, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 监控catid 变化更新表格数据
watch(props, (newValue, oldValue) => {
  formData.value.cateId = props.defaultCatid
  getTableData()
})


const getCateMenusListData = async() =>{
  const table = await getCateMenusList({})
  if (table.code === 0) {
    cateMenusList.value = table.data.list
  }
}
getCateMenusListData()


// ============== 表格控制部分结束 ===============

const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteClassFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteClassByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateClassFunc = async(row) => {
    const res = await findClass({ ID: row.ID })
    type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reclass
    setOptions()
      dialogFormVisible.value = true
    }
}


// 删除行
const deleteClassFunc = async (row) => {
    const res = await deleteClass({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)
const menuOption = ref([
  {
    ID: '0',
    title: '根菜单'
  }
])
const setOptions = async() => {
  menuOption.value = [
    {
      ID: '0',
      title: '根目录'
    }
  ]
  
  setMenuOptions(cateMenusList.value, menuOption.value, false)
}
const setMenuOptions = (menuData, optionsData, disabled) => {
  menuData &&
    menuData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              title: item.name,
              ID: String(item.ID),
              disabled: disabled || item.ID === formData.value.cateId,
              children: []
            }
            setMenuOptions(
              item.children,
              option.children,
              disabled || item.ID === formData.value.cateId
            )
            optionsData.push(option)
          } else {
            const option = {
              title: item.name,
              ID: String(item.ID),
              disabled: disabled || item.ID === formData.value.cateId
            }
            optionsData.push(option)
          }
        })
}
// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  setOptions()
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
      title: '',
      cateId: props.defaultCatid,
      name:"",
      thumb:"",
      videoUrl: '',
      desc:'',
      content: ref('<p>hello</p>'),
      enable: false,
      catemenus: {}
    }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
       if (!valid) return
      //  转换下
       formData.value.totalHour = parseFloat(formData.value.totalHour) 
        let res
        switch (type.value) {
          case 'create':
            res = await createClass(formData.value)
            break
          case 'update':
            res = await updateClass(formData.value)
            break
          default:
            res = await createClass(formData.value)
            break
        }
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '创建/更改成功'
          })
          closeDialog()
          getTableData()
        }
    
  })
}
</script>

<style scoped>
.el-table .el-table__cell {
  position: unset;
}

</style>
