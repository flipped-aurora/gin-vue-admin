<template>
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card gva-top-card">
        <div class="gva-top-card-left">
          <div class="gva-top-card-left-title">{{ t('view.dashboard.title') }}</div>
          <div class="gva-top-card-left-dot">{{ weatherInfo }}</div>
          <el-row class="my-8 w-[500px]">
            <el-col
              :span="8"
              :xs="24"
              :sm="8"
            >
              <div class="flex items-center">
                <el-icon class="dashboard-icon">
                  <sort />
                </el-icon>
                {{ t('view.dashboard.todaysTraffic') }} (1231231)
              </div>
            </el-col>
            <el-col
              :span="8"
              :xs="24"
              :sm="8"
            >
              <div class="flex items-center">
                <el-icon class="dashboard-icon">
                  <avatar />
                </el-icon>
                {{ t('view.dashboard.totalNumberOfUsers') }} (24001)
              </div>
            </el-col>
            <el-col
              :span="8"
              :xs="24"
              :sm="8"
            >
              <div class="flex items-center">
                <el-icon class="dashboard-icon">
                  <comment />
                </el-icon>
                {{ t('view.dashboard.positiveRatings') }} (99%)
              </div>
            </el-col>
          </el-row>
          <div>
            <div class="gva-top-card-left-item">
              {{ t('view.dashboard.instructionalUse') }}
              <a
                style="color:#409EFF"
                target="view_window"
                href="https://www.bilibili.com/video/BV1Rg411u7xH/"
              >https://www.bilibili.com/video/BV1Rg411u7xH</a>
            </div>
            <div class="gva-top-card-left-item">
              {{ t('view.dashboard.pluginRepo') }}
              <a
                style="color:#409EFF"
                target="view_window"
                href="https://plugin.gin-vue-admin.com/#/layout/home"
              >https://plugin.gin-vue-admin.com</a>
            </div>
          </div>
        </div>
        <img
          src="@/assets/dashboard.png"
          class="gva-top-card-right"
          alt
        >
      </div>
    </div>
    <div class="gva-card-box">
      <div class="gva-card quick-entrance">
        <div class="gva-card-title">{{ t('view.dashboard.quickEntry') }}</div>
        <el-row :gutter="20">
          <el-col
            v-for="(card, key) in toolCards"
            :key="key"
            :span="4"
            :xs="8"
            class="quick-entrance-items"
            @click="toTarget(card.name)"
          >
            <div class="quick-entrance-item">
              <div
                class="quick-entrance-item-icon"
                :style="{ backgroundColor: card.bg }"
              >
                <el-icon>
                  <component
                    :is="card.icon"
                    :style="{ color: card.color }"
                  />
                </el-icon>
              </div>
              <p>{{ card.label }}</p>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="gva-card-title">{{ t('view.dashboard.statistics') }}</div>
        <div class="p-4">
          <el-row :gutter="20">
            <el-col
              :xs="24"
              :sm="18"
            >
              <echarts-line />
            </el-col>
            <el-col
              :xs="24"
              :sm="6"
            >
              <dashboard-table />
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import EchartsLine from '@/view/dashboard/dashboardCharts/echartsLine.vue'
import DashboardTable from '@/view/dashboard/dashboardTable/dashboardTable.vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useWeatherInfo } from '@/view/dashboard/weather.js'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

defineOptions({
  name: 'Dashboard'
})

const weatherInfo = useWeatherInfo()

const toolCards = ref([
  {
    label: t('view.dashboard.userManage'),
    icon: 'monitor',
    name: 'user',
    color: '#ff9c6e',
    bg: 'rgba(255, 156, 110,.3)'
  },
  {
    label: t('view.dashboard.roleManage'),
    icon: 'setting',
    name: 'authority',
    color: '#69c0ff',
    bg: 'rgba(105, 192, 255,.3)'
  },
  {
    label: t('view.dashboard.menuManage'),
    icon: 'menu',
    name: 'menu',
    color: '#b37feb',
    bg: 'rgba(179, 127, 235,.3)'
  },
  {
    label: t('view.dashboard.codeGen'),
    icon: 'cpu',
    name: 'autoCode',
    color: '#ffd666',
    bg: 'rgba(255, 214, 102,.3)'
  },
  {
    label: t('view.dashboard.formCreator'),
    icon: 'document-checked',
    name: 'formCreate',
    color: '#ff85c0',
    bg: 'rgba(255, 133, 192,.3)'
  },
  {
    label: t('view.dashboard.about'),
    icon: 'user',
    name: 'about',
    color: '#5cdbd3',
    bg: 'rgba(92, 219, 211,.3)'
  }
])

const router = useRouter()

const toTarget = (name) => {
  router.push({ name })
}

</script>

<style lang="scss" scoped>
.page {
    @apply p-0;
    .gva-card-box{
      @apply p-4;
      &+.gva-card-box{
        @apply pt-0;
      }
    }
    .gva-card {
      @apply box-border bg-white rounded h-auto px-6 py-8 overflow-hidden shadow-sm;
      .gva-card-title{
        @apply pb-5 border-t-0 border-l-0 border-r-0 border-b border-solid border-gray-100;
      }
    }
    .gva-top-card {
        @apply h-72 flex items-center justify-between text-gray-500;
        &-left {
          @apply h-full flex flex-col w-auto;
            &-title {
              @apply text-3xl text-gray-600;
            }
            &-dot {
              @apply mt-4 text-gray-600 text-lg;
            }
            &-item{
              +.gva-top-card-left-item{
                margin-top: 24px;
              }
              margin-top: 14px;
            }
        }
        &-right {
            height: 600px;
            width: 600px;
            margin-top: 28px;
        }
    }
     ::v-deep(.el-card__header){
          @apply p-0  border-gray-200;
        }
        .card-header{
          @apply pb-5 border-b border-solid border-gray-200 border-t-0 border-l-0 border-r-0;
        }
    .quick-entrance-items {
      @apply flex items-center justify-center text-center text-gray-800;
        .quick-entrance-item {
          @apply px-8 py-6 flex items-center flex-col transition-all duration-100 ease-in-out rounded-lg cursor-pointer;
          &:hover{
            @apply shadow-lg;
          }
            &-icon {
              @apply flex items-center h-16 w-16 rounded-lg justify-center mx-0 my-auto text-2xl;
            }
            p {
                @apply mt-2.5;
            }
        }
    }
}
.dashboard-icon {
  @apply flex items-center text-xl mr-2 text-blue-400;
}

</style>
