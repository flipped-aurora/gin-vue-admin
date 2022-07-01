<template>
  <div>
    <div class="gva-table-box">
      <el-form label-width="140px" class="plug-form">
        <el-form-item label="插件名">
          <el-input v-model="form.plugName" placeholder="必填（英文大写字母开头）" @blur="titleCase" />
        </el-form-item>
        <el-form-item label="路由组">
          <el-input v-model="form.routerGroup" placeholder="将会作为插件路由组使用" />
        </el-form-item>
        <el-form-item label="使用全局属性">
          <el-checkbox v-model="form.hasGlobal" />
        </el-form-item>
        <el-form-item v-if="form.hasGlobal" label="全局属性">
          <div v-for="(i,k) in form.global" :key="k" class="plug-row">
            <span>
              <el-input v-model="i.key" placeholder="key 必填" />
            </span>
            <span>
              <el-select v-model="i.type" placeholder="type 必填">
                <el-option label="string" value="string" />
                <el-option label="int" value="int" />
                <el-option label="float32" value="float32" />
                <el-option label="float64" value="float64" />
                <el-option label="bool" value="bool" />
              </el-select>
            </span>
            <span>
              <el-input v-model="i.desc" placeholder="备注" />
            </span>
            <span>
              <el-button :icon="Plus" circle @click="addkv(form.global)" />
            </span>
            <span>
              <el-button :icon="Minus" circle @click="minkv(form.global,k)" />
            </span>
          </div>
        </el-form-item>
        <el-form-item label="使用Request">
          <el-checkbox v-model="form.hasRequest" />
        </el-form-item>
        <el-form-item v-if="form.hasRequest" label="Request">
          <div v-for="(i,k) in form.request" :key="k" class="plug-row">
            <span>
              <el-input v-model="i.key" placeholder="key 必填" />
            </span>
            <span>
              <el-select v-model="i.type" placeholder="type 必填">
                <el-option label="string" value="string" />
                <el-option label="int" value="int" />
                <el-option label="float32" value="float32" />
                <el-option label="float64" value="float64" />
                <el-option label="bool" value="bool" />
              </el-select>
            </span>
            <span>
              <el-input v-model="i.desc" placeholder="备注" />
            </span>
            <span>
              <el-button :icon="Plus" circle @click="addkv(form.request)" />
            </span>
            <span>
              <el-button :icon="Minus" circle @click="minkv(form.request,k)" />
            </span>
          </div>
        </el-form-item>
        <el-form-item label="使用Response">
          <el-checkbox v-model="form.hasResponse" />
        </el-form-item>
        <el-form-item v-if="form.hasResponse" label="Response">
          <div v-for="(i,k) in form.response" :key="k" class="plug-row">
            <span>
              <el-input v-model="i.key" placeholder="key 必填" />
            </span>
            <span>
              <el-select v-model="i.type" placeholder="type 必填">
                <el-option label="string" value="string" />
                <el-option label="int" value="int" />
                <el-option label="float32" value="float32" />
                <el-option label="float64" value="float64" />
                <el-option label="bool" value="bool" />
              </el-select>
            </span>
            <span>
              <el-input v-model="i.desc" placeholder="备注" />
            </span>
            <span>
              <el-button :icon="Plus" circle @click="addkv(form.response)" />
            </span>
            <span>
              <el-button :icon="Minus" circle @click="minkv(form.response,k)" />
            </span>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="createPlug">创建</el-button>
        </el-form-item>
      </el-form>

    </div>
  </div>
</template>

<script setup>
import { toUpperCase } from '@/utils/stringFun'

import {
  Plus,
  Minus
} from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'

import { createPlugApi } from '@/api/autoCode.js'

import { reactive } from 'vue'

const form = reactive({
  plugName: '',
  routerGroup: '',
  hasGlobal: true,
  hasRequest: true,
  hasResponse: true,
  global: [{
    key: '',
    type: '',
    desc: '',
  }],
  request: [{
    key: '',
    type: '',
    desc: '',
  }],
  response: [{
    key: '',
    type: '',
    desc: '',
  }]
})

const titleCase = () => {
  form.plugName = toUpperCase(form.plugName)
}

const createPlug = async() => {
  if (!form.plugName || !form.routerGroup) {
    ElMessage.error('插件名称和插件路由组为必填项')
    return
  }
  if (form.hasGlobal) {
    const intercept = form.global.some(i => {
      if (!i.key || !i.type) {
        return true
      }
    })
    if (intercept) {
      ElMessage.error('全局属性的key和type为必填项')
      return
    }
  }
  if (form.hasRequest) {
    const intercept = form.request.some(i => {
      if (!i.key || !i.type) {
        return true
      }
    })
    if (intercept) {
      ElMessage.error('请求属性的key和type为必填项')
      return
    }
  }
  if (form.hasResponse) {
    const intercept = form.response.some(i => {
      if (!i.key || !i.type) {
        return true
      }
    })
    if (intercept) {
      ElMessage.error('响应属性的key和type为必填项')
      return
    }
  }
  const res = await createPlugApi(form)
  if (res.code === 0) {
    ElMessageBox('创建成功，插件已自动写入后端plugin目录下，请按照自己的逻辑进行创造')
  }
}

const addkv = (arr) => {
  arr.push({
    key: '',
    value: '',
  })
}

const minkv = (arr, key) => {
  if (arr.length === 1) {
    ElMessage.warning('至少有一个全局属性')
    return
  }
  arr.splice(key, 1)
}

</script>

<style lang="scss" scoped>
        .plug-form{
            width: 680px;
        }
    .plug-row{
        display: flex;
        align-items: center;
        width: 100%;
        &+&{
            margin-top: 12px;
        }
        &>span{
            margin-left: 8px;
        }
    }
</style>
