<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="name" label="插件标题" show-overflow-tooltip width="200">
        <template #default="{ row }">
          <a
            class="text-black dark:text-white decoration-black/20 dark:decoration-white/20 hover:text-active"
            :href="`https://plugin.gin-vue-admin.com/details/${row.ID}`"
            target="_blank"
          >{{ row.name }}</a>
        </template>
      </el-table-column>
      <el-table-column prop="resume" label="简介" show-overflow-tooltip></el-table-column>
      <el-table-column prop="money" label="价格" width="100">
        <template #default="{ row }">
          <span v-if="row.money === 0">免费</span>
          <span v-else>￥{{ row.money }}</span>
        </template>
      </el-table-column>
    </el-table>
     <div class="gva-pagination">
      <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[5, 10, 20]"
        :total="total"
        layout="total, prev, pager, next"
        size="small"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>
  </div>
</template>

<script setup>
  import { getShopPluginList } from '@/api/plugin/api'
  import { ref } from 'vue'

  const tableData = ref([])
  const page = ref(1)
  const pageSize = ref(5)
  const total = ref(0)

  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  const getTableData = async() => {
    const res = await getShopPluginList({ page: page.value, pageSize: pageSize.value ,updatedAt: 1})
    if (res.code === 0) {
      tableData.value = res.data.list
      total.value = res.data.total
    }
  }

  getTableData()
</script>

<style scoped lang="scss"></style>
