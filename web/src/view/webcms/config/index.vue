<template>
  <div>
    <div class="gva-form-box">
      <el-tabs
        v-model="editableTabsValue"
        type="card"
        editable
        class="demo-tabs"
        @edit="handleTabsEdit"
        @tab-click="changeTabs"
      >
        <el-tab-pane
          v-for="item in editableTabs"
          :key="item.id"
          :label="item.company"
          :name="item.id"
        >
        <el-form ref="webConfigForm" label-position="right"  label-width="90px" :model="form">
          <el-form-item label="网站logo" prop="logo">
            <div style="display:inline-block;margin-right:10px;" @click="openHeaderChange('logo')">
                <img v-if="form.logo" alt="logo" class="header-img-box" :src="(form.logo && form.logo.slice(0, 4) !== 'http')?path+form.logo:form.logo">
                <div v-else class="header-img-box">从媒体库选择</div>
            </div>
            <el-text type="primary" >logo</el-text>
          </el-form-item>
          <el-form-item label="网站域名">
            <el-input v-model="form.siteUrl" class="keywords" />
            <el-text type="primary" >siteUrl</el-text>
          </el-form-item>
          <el-form-item label="网站名称">
            <el-input v-model="form.company" class="keywords" />
            <el-text type="primary" >company</el-text>
          </el-form-item>
          <el-form-item label="网站关键字">
            <el-input v-model="form.keywords" class="keywords" />
            <el-text type="primary" >keywords</el-text>
          </el-form-item>
          <el-form-item label="网站描述">
            <el-input v-model="form.description" class="keywords" />
            <el-text type="primary" >description</el-text>
          </el-form-item>
          <el-form-item label="热线电话">
            <el-input v-model="form.hotLine" class="keywords" />
            <el-text type="primary" >hotLine</el-text>
          </el-form-item>
          <el-form-item label="手机号码">
            <el-input v-model="form.phone" class="keywords" />
            <el-text type="primary" >phone</el-text>
          </el-form-item>
          <el-form-item label="wechat">
            <div style="display:inline-block;margin-right:10px;" @click="openHeaderChange('wechat')">
                <img v-if="form.wechat" alt="wechat" class="header-img-box" :src="(form.wechat && form.wechat.slice(0, 4) !== 'http')?path+form.wechat:form.wechat">
                <div v-else class="header-img-box">从媒体库选择</div>
            </div>
            <el-text type="primary" >wechat</el-text>
          </el-form-item>
          <el-form-item label="Email">
            <el-input v-model="form.email" class="keywords" />
            <el-text type="primary" >email</el-text>
          </el-form-item>
          <el-form-item label="网站备案">
            <el-input v-model="form.beian" class="keywords" />
            <el-text type="primary" >beian</el-text>
          </el-form-item>
          <el-form-item label="公司地址">
            <el-input v-model="form.address" class="keywords" />
            <el-text type="primary" >address</el-text>
          </el-form-item>
          <el-form-item>
            <el-button @click="updateConfig">保存</el-button>
          </el-form-item>
        </el-form>
        </el-tab-pane>
      </el-tabs>


      
    </div>
    <ChooseImg ref="chooseImg" :target="form" :target-key="targetkey" />
  </div>

</template>

<script>
export default {
  name: 'Config',
}
</script>

<script  setup>
import {
  createWebconfig,
  getWebconfig,
  setWebconfig,
  delWebconfig
} from '@/api/webconfig'
import ChooseImg from '@/components/chooseImg/index.vue'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
import { emitter } from '@/utils/bus.js'
let tabIndex = 1
const editableTabsValue = ref(1)
const editableTabs = ref([])
const path = ref(import.meta.env.VITE_BASE_API + '/')
const webConfigForm = ref(null)
const form = ref({
  ID:1,
  logo: '',
  hotLine:'',
  phone:'',
  wechat: '',
  email:'',
  beian:'',
  address: '',
  company: '',
  keywords: '',
  description: '',
  siteUrl:''
})
const handleTabsEdit = (
  targetName,
  action
) => {
  if (action === 'add') {
    createConfig()
  } else if (action === 'remove') {
    const tabs = editableTabs.value
    let activeName = editableTabsValue.value
    if (activeName === targetName) {
      tabs.forEach((tab, index) => {
        if (tab.id === targetName) {
          const nextTab = tabs[index + 1] || tabs[index - 1]
          if (nextTab) {
            activeName = nextTab.id
          }
        }
      })
    }
    editableTabsValue.value = activeName
    editableTabs.value = tabs.filter((tab) => tab.id !== targetName)
    // 删除数据
    delConfig(targetName)
  }
}

const changeTabs = (tab) => {
  form.value = editableTabs.value.find((tabval) => tabval.id === tab.props.name)
}



let targetkey = ref('')
const chooseImg = ref(null)
const openHeaderChange = (key) => {
  targetkey.value = key
  chooseImg.value.open()
}

const getTableData = async() => {
  const data = await getWebconfig()
  editableTabs.value = data.data.list
  form.value = editableTabs.value.find((tab) => tab.id === editableTabsValue.value)
}

getTableData()

const delConfig = async (id) => {
  const res = await delWebconfig(id)
  if (res.code === 0) {
    ElMessage.success(res.msg)
  }
}

const createConfig = async () => {
  let temp = ref({
      logo: '',
      hotLine:'',
      phone:'',
      wechat: '',
      email:'',
      beian:'',
      address: '',
      company: '新站点',
      keywords: '',
      description: '',
      siteUrl:''
    })
  const res = await createWebconfig(temp.value)
  if (res.code === 0) {
    ElMessage.success(res.msg)
    emitter.emit('reload')
  }
}

const updateConfig = async() => {
  const res = await setWebconfig(form.value)
  if (res.code === 0) {
    ElMessage.success('保存成功')
  }
}
</script>
<style>
.header-img-box {
  min-width: 200px;
  width: 100%;
  /* height: 200px; */
  border: 1px dashed #ccc;
  /* border-radius: 20px; */
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
.keywords{
  width:90%;
  margin-right:10px;
}
</style>
