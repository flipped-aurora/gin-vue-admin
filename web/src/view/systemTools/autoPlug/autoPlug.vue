<template>
  <div>
    <div class="gva-table-box">
      <el-form label-width="140px" class="plug-form">
        <el-form-item label="插件名">
          <el-input v-model="form.plugName" />
        </el-form-item>
        <el-form-item label="路由组">
          <el-input v-model="form.routerGloup" />
        </el-form-item>
        <el-form-item label="使用全局属性">
          <el-checkbox v-model="form.hasGloabl" />
        </el-form-item>
        <el-form-item v-if="form.hasGloabl" label="全局属性">
          <div v-for="(i,k) in form.global" :key="k" class="plug-row">
            <span>
              <el-input v-model="i.key" placeholder="key" />
            </span>
            <span>
              <el-input v-model="i.type" placeholder="type" />
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
              <el-input v-model="i.key" placeholder="key" />
            </span>
            <span>
              <el-input v-model="i.type" placeholder="type" />
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
              <el-input v-model="i.key" placeholder="key" />
            </span>
            <span>
              <el-input v-model="i.type" placeholder="type" />
            </span>
            <span>
              <el-button :icon="Plus" circle @click="addkv(form.response)" />
            </span>
            <span>
              <el-button :icon="Minus" circle @click="minkv(form.response,k)" />
            </span>
          </div>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  Plus,
  Minus
} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

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
  }],
  request: [{
    key: '',
    type: '',
  }],
  response: [{
    key: '',
    type: '',
  }]
})

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
