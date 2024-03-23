<template>
  <div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
          <el-button type="primary" icon="plus" @click="openDialog('0')">新增根菜单</el-button>
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
        <el-table-column align="center" label="层级" prop="" width="120" >
          <template v-slot="scope">
            <el-tag v-if="scope.row.menuLevel === 0">一级</el-tag>
            <el-tag type="success" v-else-if="scope.row.menuLevel === 1"
              >二级</el-tag>
            <el-tag type="warning" v-else>三级</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="center" label="ID" prop="ID" width="120" />
        <!-- <el-table-column align="center" label="上级栏目ID" prop="parentId" width="120" /> -->
        <el-table-column align="center" label="栏目名称" prop="name" width="360" />
        <el-table-column align="center" label="是否隐藏" prop="hidden" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.hidden) }}</template>
        </el-table-column>
        <el-table-column align="center" label="排序标记" prop="sort" width="120" >
          <template v-slot="scope">
            <el-input v-model="scope.row.sort" :clearable="true" @change="handleSort(scope.row)"  placeholder="请输入" />
          </template>
        </el-table-column>
        <el-table-column align="center" label="访问" prop="Url" width="120">
            <template #default="scope"><a :href="scope.row.url" target="_blank" style="color: #4d70ff;">访问</a></template>
        </el-table-column>
        <el-table-column align="center" label="操作">
            <template #default="scope">
              <el-button
              size="small"
              type="primary"
              link
              icon="plus"
              @click="openDialog(scope.row.ID)"
            >添加子菜单</el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateCateMenusFunc(scope.row)">修改</el-button>
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
    <el-dialog v-model="dialogFormVisible" width="60%" v-if="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle">
      <el-form :model="formData" label-position="right" :inline="true" ref="elFormRef" :rules="rule" label-width="90px">
        <el-form-item label="上级栏目ID" style="width:30%">
          <el-cascader
            v-model="formData.parentId"
            style="width:100%"
            :disabled="!isEdit"
            :options="menuOption"
            :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item label="栏目名称:"  prop="name" style="width:30%">
          <el-input v-model="formData.name" :clearable="true" type="textarea"  placeholder="请输入栏目名称 多个栏目换行输入" />
        </el-form-item>
        <el-form-item label="栏目简称:"  prop="short" style="width:30%">
          <el-input v-model="formData.short" :clearable="true" type="text"  placeholder="请输入栏目简称" />
        </el-form-item>
        <el-form-item label="模型管理:"   prop="modeType" style="width:30%">
          <el-select v-model="formData.modeType" class="m-2" placeholder="Select" size="large">
            <el-option
              v-for="item in modeType"
              :key="item.ID"
              :label="item.structCNName"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="栏目类型:"   prop="cateType" style="width:30%">
          <el-radio-group v-model="formData.cateType" >
            <el-radio v-for="(item,index) in cateType" :key="index" value="item.value">{{ item.label }}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item  style="width:30%"></el-form-item>
        
        <el-form-item label="列表页模板:" v-if="formData.cateType != 3"  prop="listTemplate" style="width:30%">
          <el-select v-model="formData.listTemplate"  placeholder="选择模板">
            <el-option
              v-for="item in templateList"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="内容页模板:" v-if="formData.cateType != 3"  prop="showTemplate" style="width:30%">
          <el-select v-model="formData.showTemplate"  placeholder="选择模板">
            <el-option
              v-for="item in templateList"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="外链地址" v-if="formData.cateType == 3"  prop="islink" style="width:40%">
          <el-input v-model="formData.islink" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <br>
        <el-form-item label="封面缩略图:"  prop="cateThumb" style="display: flex;">
          <CustomPic style="margin-top:8px;" v-show="formData.cateThumb" :pic-src="formData.cateThumb" pic-type="swiper"/>
          <div style="width: 60%;">
          <el-input v-if="formData.cateThumb" v-model="formData.cateThumb" :clearable="true"  placeholder="请输入"/>
          <el-button v-else  type="primary" @click="openHeaderChange('cateThumb')">从媒体库选择</el-button>
        </div>
        </el-form-item>
        <el-form-item label="列表缩略图:"  prop="listThumb" style="display: flex;">
          <CustomPic style="margin-top:8px;" v-show="formData.listThumb" :pic-src="formData.listThumb" pic-type="swiper"/>
          <div style="width: 60%;">
          <el-input v-if="formData.listThumb" v-model="formData.listThumb" :clearable="true"  placeholder="请输入"/>
          <el-button v-else  type="primary" @click="openHeaderChange('listThumb')">从媒体库选择</el-button>
          </div>
        </el-form-item>
        <el-form-item label="是否隐藏:"  prop="hidden" style="width:15%">
          <el-switch v-model="formData.hidden" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="跳转子栏目:"  prop="isJump" style="width:15%">
          <el-switch v-model="formData.isJump" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="列表排序:"  prop="orderType" style="width:23%">
          <el-radio-group v-model="formData.orderType" >
            <el-radio v-for="(item,index) in ordertype" :key="index" :value="item.value">{{ item.label }}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="每页数量:"  prop="pageSize" style="width:12%">
          <el-input v-model="formData.pageSize"  type="text" />
        </el-form-item>
        <el-form-item label="介绍:"  prop="desc" style="width:100%">
          <Tinymceditor ref="editor" @update:content="formData.desc = $event" :content="formData.desc" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="formData" :target-key="targetkey" />
  </div>
</template>

<script>
export default {
  name: 'CateMenus'
}
</script>

<script setup>
import {
  createCateMenus,
  deleteCateMenus,
  deleteCateMenusByIds,
  updateCateMenus,
  findCateMenus,
  getCateMenusList,
  getTemplateList,
  getModelList
} from '@/api/cateMenus'


// 全量引入格式化工具 请按需保留
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import Tinymceditor from '@/components/tinymce/index.vue'
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { getDict } from '@/utils/dictionary'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

//模型管理 1 产品管理 2 文章管理 
const modeType = ref([])
const templateList = ref([])

const formData = ref({
  menuLevel: 0,
  parentId: 0,
  name: '',
  short: '',
  hidden: false,
  isJump: false,
  sort: 0,
  islink: '#',
  modeType: 1,
  cateType: 1,
  listTemplate:'',
  showTemplate:'',
  cateThumb:'',
  listThumb :'',
  desc: '',
  orderType: 1,
  pageSize: 10,
})

// 验证规则
const rule = reactive({
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 排序
const handleSort = async (val) =>{
  val.sort = parseInt(val.sort)
 const res = await updateCateMenus(val)
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
let targetkey = ref('')
const chooseImg = ref(null)
const openHeaderChange = (key) => {
  targetkey.value = key
  chooseImg.value.open()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.hidden === ""){
      searchInfo.value.hidden=null
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
// 1 封面 2 列表 3 链接
const cateType = ref([])
const ordertype = ref([])

// 查询
const getTableData = async() => {
  const table = await getCateMenusList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
  cateType.value = await getDict('catetype')
  console.log(cateType.value)
  ordertype.value = await getDict('ordertype')
  // 获取模型列表
  const modellist = await getModelList()
  modeType.value = modellist.data.list
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
            deleteCateMenusFunc(row)
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
      const res = await deleteCateMenusByIds({ ids })
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
const updateCateMenusFunc = async(row) => {
    const res = await findCateMenus({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.recateMenus
        dialogFormVisible.value = true
        setOptions()
        // 获取模板文件列表
        const data = await getTemplateList()
        templateList.value = data.data.list
        isEdit.value = true
    }
    
}


// 删除行
const deleteCateMenusFunc = async (row) => {
    const res = await deleteCateMenus({ ID: row.ID })
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
    ID: 0,
    title: '根菜单'
  }
])
const setOptions = () => {
  menuOption.value = [
    {
      ID: 0,
      title: '根目录'
    }
  ]
  setMenuOptions(tableData.value, menuOption.value, false)
}
const setMenuOptions = (menuData, optionsData, disabled) => {
  menuData &&
    menuData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              title: item.name,
              ID: item.ID,
              disabled: disabled || item.ID === formData.value.ID,
              children: []
            }
            setMenuOptions(
              item.children,
              option.children,
              disabled || item.ID === formData.value.ID
            )
            optionsData.push(option)
          } else {
            const option = {
              title: item.name,
              ID: item.ID,
              disabled: disabled || item.ID === formData.value.ID
            }
            optionsData.push(option)
          }
        })
}


// 添加菜单方法，id为 0则为添加根菜单
const isEdit = ref(false)
const dialogTitle = ref('新增菜单')

// 打开弹窗
const openDialog = async(id) => {
  dialogTitle.value = '新增菜单'
  type.value = 'create'
  isEdit.value = false
  formData.value.parentId = id
  setOptions()
  // 获取模板文件列表
  const data = await getTemplateList()
  templateList.value = data.data.list
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        menuLevel: 0,
        parentId: 0,
        name: '',
        hidden: false,
        sort: 0,
        islink: '#',
        modeType: 1,
        cateType: 1,
        listTemplate:'',
        showTemplate:'',
        cateThumb:'',
        listThumb :'',
        desc: '',
        orderType: 1,
        pageSize: 10,
    }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createCateMenus(formData.value)
                  break
                case 'update':
                  res = await updateCateMenus(formData.value)
                  break
                default:
                  res = await createCateMenus(formData.value)
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

<style  scoped>

</style>
