<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="产品名称:"  prop="title">
          <el-input v-model="searchInfo.title"   placeholder="请输入" />
        </el-form-item>
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
        <el-table-column align="center" label="排序标记" prop="sort" width="80" >
          <template v-slot="scope">
            <el-input v-model="scope.row.sort" :clearable="true" @change="handleSort(scope.row)"  placeholder="请输入" />
          </template>
        </el-table-column>
        <el-table-column align="center" label="ID" prop="ID" width="120" />
        <el-table-column align="center" label="产品名称" prop="title" width="300" />
        <el-table-column align="center" label="缩略图" prop="thumb" width="180" >
          <template #default="scope">
            <CustomPic style="margin-top:8px" :pic-src="scope.row.thumb" pic-type="swiper" />
          </template>
        </el-table-column>
        <el-table-column align="center" label="上传者" prop="sysUser.nickName" width="100" />
        <el-table-column align="center" label="是否显示" prop="enable" width="80">
            <template #default="scope">{{ formatBoolean(scope.row.enable) }}</template>
        </el-table-column>
        <el-table-column align="center" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="center" label="操作">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateCourseFunc(scope.row)">修改</el-button>
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
    <el-dialog v-model="dialogFormVisible" v-if="dialogFormVisible" :before-close="closeDialog" title="新增产品">
      <el-form :model="formData" label-position="right" ref="elFormRef" :inline="true" :rules="rule" label-width="100px">
        <el-form-item label="产品分类:" style="width:45%">
          <el-cascader
            v-model="formData.cateId"
            style="width:100%"
            :options="menuOption"
            :props="{ label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item label="产品名称:"  prop="title" style="width: 45%;">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="缩略图:"  prop="thumb" style="width: 45%;">
          <el-input v-if="formData.thumb" v-model="formData.thumb" :clearable="true"  placeholder="请输入" />
          <el-button v-else  type="primary" @click="openHeaderChange('thumb')">从媒体库选择</el-button>
        </el-form-item>
        <el-form-item label="上传者:"  prop="thumb" style="width: 45%;">
          <el-select v-model="formData.createdBy" class="m-2" placeholder="请选择" size="large">
          <el-option
            v-for="item in userlist"
            :key="item.value"
            :label="item.label"
            :value="item.value"
            :filterable="true"
          />
        </el-select>
        </el-form-item>
        <el-form-item label="产品图片集:"  prop="productThumbs" style="width: 100%;">
          <draggable 
            v-model="formData.thumbs" 
            group="people" 
            @start="drag=true" 
            @end="drag=false" 
            item-key="id">
            <template #item="{element}">
              <div style="margin-right: 10px;float: left;">
                <el-icon class="imgico" @click="handleDelThumbs(element.url)"><CloseBold /></el-icon>
                <el-image
                  style="width: 100px; height: 100px"
                  :src="path + element.url"
                  :zoom-rate="1.2"
                  :preview-src-list="element.thumbs"
                  fit="cover"
                />
              </div>
            </template>
          </draggable>
          <!-- <div v-for="(item,index) in formData.thumbs" :key="index" style="margin-right: 10px;">
              <el-icon class="imgico" @click="handleDelThumbs(item.url)"><CloseBold /></el-icon>
              <el-image
                style="width: 100px; height: 100px"
                :src="path + item.url"
                :zoom-rate="1.2"
                :preview-src-list="formData.thumbs"
                fit="cover"
              />
            </div> -->
          
        <el-button  type="primary" @click="openHeaderChange('thumbs')">从媒体库选择</el-button>
        
        </el-form-item>
        <el-form-item label="详细信息:"  prop="productDetails" style="width: 100%;">
          <!-- <Wangeditor  @update:content="formData.productDetails = $event" :content="formData.productDetails" :height="'300px'"/> -->
        </el-form-item>
        <el-form-item label="条款:"  prop="productClause" style="width: 100%;">
          <!-- <Wangeditor  @update:content="formData.productClause = $event" :content="formData.productClause"/> -->
        </el-form-item>
        <el-form-item label="产品简介:"  prop="desc" style="width: 100%;">
          <!-- <Wangeditor  @update:content="formData.desc = $event" :content="formData.desc"/> -->
        </el-form-item>
        <el-form-item label="是否显示:"  prop="enable" >
          <el-switch v-model="formData.enable" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
      
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="formData" :target-key="targetkey"/>
    <el-drawer v-if="drawer" v-model="drawer" custom-class="auth-drawer" :with-header="false" size="40%" title="产品列表">
      <el-tabs  type="border-card">
        <el-tab-pane label="产品列表">
          <Menus ref="menus" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
  </div>
</template>

<script>
export default {
  name: 'Course'
}
</script>

<script setup>
import {
  createCourse,
  deleteCourse,
  deleteCourseByIds,
  updateCourse,
  findCourse,
  getCourseList
} from '@/api/course'
import {
  getCateMenusList
} from '@/api/cateMenus'
import {
  getUserList
} from '@/api/user'

// 全量引入格式化工具 请按需保留
import Menus from '@/view/webcms/course/components/menus.vue'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
// import Wangeditor from '@/components/wangeditor/index.vue'
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox} from 'element-plus'
import { ref, reactive,watch } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
import draggable from 'vuedraggable'
const userStore = useUserStore()
const path = ref(import.meta.env.VITE_BASE_API + '/')
// 自动化生成的字典（可能为空）以及字段
const formData = ref({
    cateId:0,
    title: '',
    thumb:"",
    productThumbs: '',
    productDetails: '',
    productClause: '',
    desc:'',
    createdBy: "",
    enable: false,
    thumbs: [],
    sort:0,
  sysUser: {},
  drag: false,
})
const handleDelThumbs = (value) => {
  formData.value.thumbs = formData.value.thumbs.filter(item => item.url != value)
}
const userlist = ref([])
userlist.value = []

// 验证规则
const rule = reactive({
  productName : [{
        required: true,
        message: '产品名称不能为空！',
        trigger: ['input','blur'],
    }],
    productThumbs : [{
        required: true,
        message: '产品图片不能为空！',
    }],
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const drawer = ref(false)
const activeRow = ref({})


let targetkey = ref('')
const chooseImg = ref(null)
const openHeaderChange = (key) => {
  targetkey.value = key
  chooseImg.value.open()
}

// 课时列表
const changeRow = (key, value) => {
  activeRow.value[key] = value
}
// 打开右侧弹窗
const opdendrawer = (row) => {
  drawer.value = true
  activeRow.value = row
}

// 排序
const handleSort = async (val) =>{
  val.sort = parseInt(val.sort)
 const res = await updateCourse(val)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '更改成功'
    })
    getTableData()
  }
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}


// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.title === ""){
      searchInfo.value.title=null
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
  const table = await getCourseList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

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
            deleteCourseFunc(row)
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
      const res = await deleteCourseByIds({ ids })
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
const updateCourseFunc = async (row) => {
  getUserinfo()
    const res = await findCourse({ ID: row.ID })
  type.value = 'update'
    formData.value.thumbs = []
    if (res.code === 0) {
      formData.value = res.data.recourse
      setOptions()
      formData.value.thumbs = JSON.parse(formData.value.productThumbs)
      dialogFormVisible.value = true
    }
}


// 删除行
const deleteCourseFunc = async (row) => {
    const res = await deleteCourse({ ID: row.ID })
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
const cateMenusList = ref([])
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
  const table = await getCateMenusList({})
  if (table.code === 0) {
    cateMenusList.value = table.data.list
  }
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

// 获取上传者信息
const getUserinfo = async () => { 
  userlist.value = []
  const userlistdata = await getUserList({ page: 1, pageSize: 1000 })
  userlistdata.data.list.forEach(function (item) {
  userlist.value.push({
    value: item.uuid,
    label: item.nickName,
  })
})
  
  
}

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  setOptions()
  getUserinfo()
  
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
        formData.value = {
            cateId:1,
            title: '',
            thumb:"",
            productThumbs: '',
            productDetails: '',
            productClause: '',
            desc:'',
            createdBy: "",
            enable: false,
            thumbs: [],
            sort:0,
            sysUser:{}
        }
}
// 弹窗确定
const enterDialog = async () => {
     formData.value.productThumbs = JSON.stringify(formData.value.thumbs)
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createCourse(formData.value)
                  break
                case 'update':
                  formData.value.sysUser = {} //清空用户信息 防止 gorm:"foreignKey"
                  res = await updateCourse(formData.value)
                  break
                default:
                  res = await createCourse(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
                formData.value.thumbs = []
              }
      })
}
</script>

<style>

.imgico{
  top: -93px;
  right: -107px;
  z-index: 5;
  color: red;
  cursor: pointer;
}

</style>
