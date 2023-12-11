<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon>
                  <QuestionFilled />
                </el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期"
            :disabled-date="time => searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
            :disabled-date="time => searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
        </el-form-item>
        <el-form-item label="APPID" prop="AppId">
          <el-input v-model="searchInfo.AppId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="商户号" prop="MchId">
          <el-input v-model="searchInfo.MchId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="商户订单" prop="OutTradeNo">
          <el-input v-model="searchInfo.OutTradeNo" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="支付订单" prop="TransactionId">
          <el-input v-model="searchInfo.TransactionId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="交易类型" prop="TradeType">
          <el-input v-model="searchInfo.TradeType" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="OPENID" prop="OpenId">
          <el-input v-model="searchInfo.OpenId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="订单金额" prop="Total">

          <el-input v-model.number="searchInfo.Total" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="支付金额" prop="PayerTotal">

          <el-input v-model.number="searchInfo.PayerTotal" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="商品标题" prop="GoodsTitle">
          <el-input v-model="searchInfo.GoodsTitle" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="支付方式" prop="PayMent">
          <el-select v-model="searchInfo.PayMent" clearable placeholder="请选择"
            @clear="() => { searchInfo.PayMent = undefined }">
            <el-option v-for="(item, key) in PayMentOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="设备ID" prop="DeviceId">
          <el-input v-model="searchInfo.DeviceId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog" v-auth="btnAuth.新增">新增</el-button>
        <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length"
              @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID" 
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <!-- <el-table-column align="left" label="APPID" prop="AppId" width="120" /> -->
        <el-table-column align="left" label="商户号" prop="MchId" width="160" />
        <el-table-column align="left" label="商户订单" prop="OutTradeNo" width="220" />
        <el-table-column align="left" label="支付订单" prop="TransactionId" width="250" />
        <el-table-column align="left" label="交易类型" prop="TradeType" width="120" />
        <el-table-column align="left" label="交易状态" prop="TradeState" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.TradeState, TradeStateOptions) }}
          </template>
        </el-table-column>
        <!-- <el-table-column align="left" label="银行类型" prop="BankType" width="120" />
        <el-table-column align="left" label="附加数据" prop="Attach" width="120" /> -->
        <el-table-column align="left" label="支付时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.SuccessTime) }}</template>
        </el-table-column>
        <!-- <el-table-column align="left" label="OPENID" prop="OpenId" width="120" /> -->
        <!-- <el-table-column align="left" label="sub_appid" prop="SubOpenId" width="120" /> -->
        <el-table-column align="left" label="订单金额" prop="Total" width="120" />
        <el-table-column align="left" label="支付金额" prop="PayerTotal" width="120" />
        <el-table-column align="left" label="用户ID" prop="UserId" width="120" />
        <el-table-column align="left" label="地址ID" prop="AddrId" width="120" />
        <el-table-column align="left" label="商品ID" prop="GoodsId" width="120" />
        <el-table-column align="left" label="店铺ID" prop="ShopID" width="120" />
        <el-table-column align="left" label="配送ID" prop="TransportId" width="120" />
        <el-table-column align="left" label="结束时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.Endtime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="商品标题" prop="GoodsTitle" width="220" />
        <el-table-column align="left" label="商品价格" prop="GoodsPrice" width="120" />
        <el-table-column align="left" label="商品邮费" prop="PostFee" width="120" />
        <!-- <el-table-column align="left" label="回调地址" prop="CallUrl" width="120" /> -->
        <el-table-column align="left" label="支付方式" prop="PayMent" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.PayMent, PayMentOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="设备ID" prop="DeviceId" width="120" />
        <el-table-column align="left" label="操作" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>
              查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateShopOrdersFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type === 'create' ? '添加' : '修改'"
      destroy-on-close>
      <el-scrollbar height="500px">
        <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
          <el-form-item label="APPID:" prop="AppId">
            <el-input v-model="formData.AppId" :clearable="true" placeholder="请输入APPID" />
          </el-form-item>
          <el-form-item label="商户号:" prop="MchId">
            <el-input v-model="formData.MchId" :clearable="true" placeholder="请输入商户号" />
          </el-form-item>
          <el-form-item label="商户订单:" prop="OutTradeNo">
            <el-input v-model="formData.OutTradeNo" :clearable="true" placeholder="请输入商户订单" />
          </el-form-item>
          <el-form-item label="支付订单:" prop="TransactionId">
            <el-input v-model="formData.TransactionId" :clearable="true" placeholder="请输入支付订单" />
          </el-form-item>
          <el-form-item label="交易类型:" prop="TradeType">
            <el-input v-model="formData.TradeType" :clearable="true" placeholder="请输入交易类型" />
          </el-form-item>
          <el-form-item label="交易状态:" prop="TradeState">
            <el-select v-model="formData.TradeState" placeholder="请选择交易状态" style="width:100%" :clearable="true">
              <el-option v-for="(item, key) in TradeStateOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
          </el-form-item>
          <el-form-item label="银行类型:" prop="BankType">
            <el-input v-model="formData.BankType" :clearable="true" placeholder="请输入银行类型" />
          </el-form-item>
          <el-form-item label="附加数据:" prop="Attach">
            <el-input v-model="formData.Attach" :clearable="true" placeholder="请输入附加数据" />
          </el-form-item>
          <el-form-item label="支付时间:" prop="SuccessTime">
            <el-date-picker v-model="formData.SuccessTime" type="date" style="width:100%" placeholder="选择日期"
              :clearable="true" />
          </el-form-item>
          <el-form-item label="OPENID:" prop="OpenId">
            <el-input v-model="formData.OpenId" :clearable="true" placeholder="请输入OPENID" />
          </el-form-item>
          <el-form-item label="sub_appid:" prop="SubOpenId">
            <el-input v-model="formData.SubOpenId" :clearable="true" placeholder="请输入sub_appid" />
          </el-form-item>
          <el-form-item label="订单金额:" prop="Total">
            <el-input v-model.number="formData.Total" :clearable="true" placeholder="请输入订单金额" />
          </el-form-item>
          <el-form-item label="支付金额:" prop="PayerTotal">
            <el-input v-model.number="formData.PayerTotal" :clearable="true" placeholder="请输入支付金额" />
          </el-form-item>
          <el-form-item label="用户ID:" prop="UserId">
            <el-input v-model.number="formData.UserId" :clearable="true" placeholder="请输入用户ID" />
          </el-form-item>
          <el-form-item label="地址ID:" prop="AddrId">
            <el-input v-model.number="formData.AddrId" :clearable="true" placeholder="请输入地址ID" />
          </el-form-item>
          <el-form-item label="商品ID:" prop="GoodsId">
            <el-input v-model.number="formData.GoodsId" :clearable="true" placeholder="请输入商品ID" />
          </el-form-item>
          <el-form-item label="店铺ID:" prop="ShopID">
            <el-input v-model="formData.ShopID" :clearable="true" placeholder="请输入店铺ID" />
          </el-form-item>
          <el-form-item label="配送ID:" prop="TransportId">
            <el-input v-model.number="formData.TransportId" :clearable="true" placeholder="请输入配送ID" />
          </el-form-item>
          <el-form-item label="结束时间:" prop="Endtime">
            <el-date-picker v-model="formData.Endtime" type="date" style="width:100%" placeholder="选择日期"
              :clearable="true" />
          </el-form-item>
          <el-form-item label="商品标题:" prop="GoodsTitle">
            <el-input v-model="formData.GoodsTitle" :clearable="true" placeholder="请输入商品标题" />
          </el-form-item>
          <el-form-item label="商品价格:" prop="GoodsPrice">
            <el-input v-model.number="formData.GoodsPrice" :clearable="true" placeholder="请输入商品价格" />
          </el-form-item>
          <el-form-item label="商品邮费:" prop="PostFee">
            <el-input v-model.number="formData.PostFee" :clearable="true" placeholder="请输入商品邮费" />
          </el-form-item>
          <el-form-item label="回调地址:" prop="CallUrl">
            <el-input v-model="formData.CallUrl" :clearable="true" placeholder="请输入回调地址" />
          </el-form-item>
          <el-form-item label="支付方式:" prop="PayMent">
            <el-select v-model="formData.PayMent" placeholder="请选择支付方式" style="width:100%" :clearable="true">
              <el-option v-for="(item, key) in PayMentOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
          </el-form-item>
          <el-form-item label="设备ID:" prop="DeviceId">
            <el-input v-model="formData.DeviceId" :clearable="true" placeholder="请输入设备ID" />
          </el-form-item>
        </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情"
      destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
          <el-descriptions-item label="APPID">
            {{ formData.AppId }}
          </el-descriptions-item>
          <el-descriptions-item label="商户号">
            {{ formData.MchId }}
          </el-descriptions-item>
          <el-descriptions-item label="商户订单">
            {{ formData.OutTradeNo }}
          </el-descriptions-item>
          <el-descriptions-item label="支付订单">
            {{ formData.TransactionId }}
          </el-descriptions-item>
          <el-descriptions-item label="交易类型">
            {{ formData.TradeType }}
          </el-descriptions-item>
          <el-descriptions-item label="交易状态">
            {{ filterDict(formData.TradeState, TradeStateOptions) }}
          </el-descriptions-item>
          <el-descriptions-item label="银行类型">
            {{ formData.BankType }}
          </el-descriptions-item>
          <el-descriptions-item label="附加数据">
            {{ formData.Attach }}
          </el-descriptions-item>
          <el-descriptions-item label="支付时间">
            {{ formatDate(formData.SuccessTime) }}
          </el-descriptions-item>
          <el-descriptions-item label="OPENID">
            {{ formData.OpenId }}
          </el-descriptions-item>
          <el-descriptions-item label="sub_appid">
            {{ formData.SubOpenId }}
          </el-descriptions-item>
          <el-descriptions-item label="订单金额">
            {{ formData.Total }}
          </el-descriptions-item>
          <el-descriptions-item label="支付金额">
            {{ formData.PayerTotal }}
          </el-descriptions-item>
          <el-descriptions-item label="用户ID">
            {{ formData.UserId }}
          </el-descriptions-item>
          <el-descriptions-item label="地址ID">
            {{ formData.AddrId }}
          </el-descriptions-item>
          <el-descriptions-item label="商品ID">
            {{ formData.GoodsId }}
          </el-descriptions-item>
          <el-descriptions-item label="店铺ID">
            {{ formData.ShopID }}
          </el-descriptions-item>
          <el-descriptions-item label="配送ID">
            {{ formData.TransportId }}
          </el-descriptions-item>
          <el-descriptions-item label="结束时间">
            {{ formatDate(formData.Endtime) }}
          </el-descriptions-item>
          <el-descriptions-item label="商品标题">
            {{ formData.GoodsTitle }}
          </el-descriptions-item>
          <el-descriptions-item label="商品价格">
            {{ formData.GoodsPrice }}
          </el-descriptions-item>
          <el-descriptions-item label="商品邮费">
            {{ formData.PostFee }}
          </el-descriptions-item>
          <el-descriptions-item label="回调地址">
            {{ formData.CallUrl }}
          </el-descriptions-item>
          <el-descriptions-item label="支付方式">
            {{ filterDict(formData.PayMent, PayMentOptions) }}
          </el-descriptions-item>
          <el-descriptions-item label="设备ID">
            {{ formData.DeviceId }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
//按钮控制
import { useBtnAuth } from '@/utils/btnAuth'
    const btnAuth = useBtnAuth()
    
import {
  createShopOrders,
  deleteShopOrders,
  deleteShopOrdersByIds,
  updateShopOrders,
  findShopOrders,
  getShopOrdersList
} from '@/api/shopOrders'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'ShopOrders'
})

// 自动化生成的字典（可能为空）以及字段
const PayMentOptions = ref([])
const TradeStateOptions = ref([])
const formData = ref({
  AppId: '',
  MchId: '',
  OutTradeNo: '',
  TransactionId: '',
  TradeType: '',
  TradeState: undefined,
  BankType: '',
  Attach: '',
  SuccessTime: new Date(),
  OpenId: '',
  SubOpenId: '',
  Total: 0,
  PayerTotal: 0,
  UserId: 0,
  AddrId: 0,
  GoodsId: 0,
  ShopID: '',
  TransportId: 0,
  Endtime: new Date(),
  GoodsTitle: '',
  GoodsPrice: 0,
  PostFee: 0,
  CallUrl: '',
  PayMent: undefined,
  DeviceId: '',
})


// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
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
const getTableData = async () => {
  const table = await getShopOrdersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
  PayMentOptions.value = await getDictFunc('PayMent')
  TradeStateOptions.value = await getDictFunc('TradeState')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


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
    deleteShopOrdersFunc(row)
  })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
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
  const res = await deleteShopOrdersByIds({ ids })
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
const updateShopOrdersFunc = async (row) => {
  const res = await findShopOrders({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reshopOrders
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteShopOrdersFunc = async (row) => {
  const res = await deleteShopOrders({ ID: row.ID })
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


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findShopOrders({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reshopOrders
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    AppId: '',
    MchId: '',
    OutTradeNo: '',
    TransactionId: '',
    TradeType: '',
    TradeState: undefined,
    BankType: '',
    Attach: '',
    SuccessTime: new Date(),
    OpenId: '',
    SubOpenId: '',
    Total: 0,
    PayerTotal: 0,
    UserId: 0,
    AddrId: 0,
    GoodsId: 0,
    ShopID: '',
    TransportId: 0,
    Endtime: new Date(),
    GoodsTitle: '',
    GoodsPrice: 0,
    PostFee: 0,
    CallUrl: '',
    PayMent: undefined,
    DeviceId: '',
  }
}


// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    AppId: '',
    MchId: '',
    OutTradeNo: '',
    TransactionId: '',
    TradeType: '',
    TradeState: undefined,
    BankType: '',
    Attach: '',
    SuccessTime: new Date(),
    OpenId: '',
    SubOpenId: '',
    Total: 0,
    PayerTotal: 0,
    UserId: 0,
    AddrId: 0,
    GoodsId: 0,
    ShopID: '',
    TransportId: 0,
    Endtime: new Date(),
    GoodsTitle: '',
    GoodsPrice: 0,
    PostFee: 0,
    CallUrl: '',
    PayMent: undefined,
    DeviceId: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createShopOrders(formData.value)
        break
      case 'update':
        res = await updateShopOrders(formData.value)
        break
      default:
        res = await createShopOrders(formData.value)
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

<style></style>
