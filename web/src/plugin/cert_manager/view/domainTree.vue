<template>
  <div class="domain-tree-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>子域名探测</span>
          <div class="header-actions">
            <el-input
              v-model="rootDomain"
              placeholder="请输入根域名"
              style="width: 200px; margin-right: 10px"
              @keyup.enter="loadDomainTree"
            />
            <el-button type="primary" @click="loadDomainTree">加载</el-button>
            <el-checkbox v-model="deepScan" label="深度扫描" style="margin-right: 10px" />
            <el-button type="success" @click="discoverSubdomains">发现子域名</el-button>
            <el-button type="info" @click="exportReport">导出报告</el-button>
            <el-button type="warning" @click="reProbeAll">重新探测</el-button>
          </div>
        </div>
      </template>

      <el-table
        :data="domainTree"
        row-key="id"
        :tree-props="{ children: 'children' }"
        :expand-row-keys="expandedKeys"
        v-loading="loading"
        border
      >
        <el-table-column prop="domain" label="域名" min-width="200" />
        <el-table-column prop="ip" label="解析 IP" width="140" />
        <el-table-column prop="source" label="发现来源" width="120" />
        <el-table-column prop="certType" label="证书类型" width="120">
          <template #default="{ row }">
            <el-tag v-if="row.certType === '泛域名证书'" type="warning">泛域名证书</el-tag>
            <el-tag v-else-if="row.certType === '独立证书'" type="success">独立证书</el-tag>
            <el-tag v-else type="info">未探测</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="domainStatus" label="域名状态" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.domainStatus === 1" type="success">正常</el-tag>
            <el-tag v-else type="danger">异常</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="certStatus" label="证书状态" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.certStatus === 1" type="success">正常</el-tag>
            <el-tag v-else type="danger">异常</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="category" label="项目分类" width="120" />
        <el-table-column prop="isIgnored" label="是否忽略" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.isIgnored" type="info">已忽略</el-tag>
            <el-tag v-else type="success">监控中</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="lastProbedAt" label="最后探测时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.lastProbedAt) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button
              link
              type="primary"
              size="small"
              @click="toggleIgnore(row)"
            >
              {{ row.isIgnored ? '取消忽略' : '忽略' }}
            </el-button>
            <el-button
              link
              type="warning"
              size="small"
              @click="reProbeDomain(row)"
            >
              探测
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage, ElLoading } from 'element-plus'
import { getDomainTree, discoverSubdomains as discoverSubdomainsApi, reProbeDomainTree, ignoreDomain, exportSubdomainReport, batchReprobe } from '@/plugin/cert_manager/api/certAdvanced'

const rootDomain = ref('')
const deepScan = ref(true)
const domainTree = ref([])
const expandedKeys = ref([])
const loading = ref(false)

const loadDomainTree = async () => {
  if (!rootDomain.value) {
    ElMessage.warning('请输入根域名')
    return
  }

  loading.value = true
  try {
    const res = await getDomainTree({ rootDomain: rootDomain.value })
    if (res.code === 0) {
      domainTree.value = res.data || []
      // 默认展开所有第一级节点
      expandedKeys.value = domainTree.value.map(item => String(item.id))
    } else {
      ElMessage.error(res.msg || '加载失败')
    }
  } catch (error) {
    ElMessage.error('加载失败: ' + error.message)
  } finally {
    loading.value = false
  }
}

const discoverSubdomains = async () => {
  if (!rootDomain.value) {
    ElMessage.warning('请输入根域名')
    return
  }

  const loadingInstance = ElLoading.service({
    lock: true,
    text: '正在执行子域名探测，请稍候...',
    background: 'rgba(0, 0, 0, 0.7)',
  })

  try {
    const res = await discoverSubdomainsApi({ 
      rootDomain: rootDomain.value,
      deepScan: deepScan.value
    })
    if (res.code === 0) {
      ElMessage.success('子域名发现成功')
      await loadDomainTree()
    } else {
      ElMessage.error(res.msg || '发现失败')
    }
  } catch (error) {
    ElMessage.error('发现失败: ' + error.message)
  } finally {
    loadingInstance.close()
  }
}

const exportReport = async () => {
  if (!rootDomain.value) {
    ElMessage.warning('请输入根域名')
    return
  }

  try {
    const response = await exportSubdomainReport({ rootDomain: rootDomain.value })
    const url = window.URL.createObjectURL(new Blob([response]))
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', `DomainReport_${rootDomain.value}.xlsx`)
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    ElMessage.success('报告下载开始')
  } catch (error) {
    ElMessage.error('导出失败: ' + error.message)
  }
}

const reProbeAll = async () => {
  if (!rootDomain.value) {
    ElMessage.warning('请输入根域名')
    return
  }

  loading.value = true
  try {
    const res = await reProbeDomainTree({ rootDomain: rootDomain.value })
    if (res.code === 0) {
      ElMessage.success('重新探测成功')
      await loadDomainTree()
    } else {
      ElMessage.error(res.msg || '探测失败')
    }
  } catch (error) {
    ElMessage.error('探测失败: ' + error.message)
  } finally {
    loading.value = false
  }
}

const toggleIgnore = async (row) => {
  try {
    const res = await ignoreDomain({ domainId: row.id, ignore: !row.isIgnored })
    if (res.code === 0) {
      ElMessage.success('更新成功')
      row.isIgnored = !row.isIgnored
    } else {
      ElMessage.error(res.msg || '更新失败')
    }
  } catch (error) {
    ElMessage.error('更新失败: ' + error.message)
  }
}

const reProbeDomain = async (row) => {
  const loadingInstance = ElLoading.service({
    lock: true,
    text: `正在探测: ${row.domain}...`,
    background: 'rgba(0, 0, 0, 0.7)',
  })
  try {
    const res = await batchReprobe({ ids: [row.id] })
    if (res.code === 0) {
      ElMessage.success('探测成功')
      await loadDomainTree()
    } else {
      ElMessage.error(res.msg || '探测失败')
    }
  } catch (error) {
    ElMessage.error('请求失败: ' + error.message)
  } finally {
    loadingInstance.close()
  }
}

const formatDate = (dateStr) => {
  if (!dateStr) return '-'
  return new Date(dateStr).toLocaleString('zh-CN')
}
</script>

<style scoped>
.domain-tree-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-actions {
  display: flex;
  align-items: center;
}
</style>
