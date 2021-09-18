<template>
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card dashbord1">
        <div class="dashbord1-left">
          <div class="dashbord1-left-title">早安，管理员，请开始一天的工作吧</div>
          <div class="dashbord1-left-dot">今日晴，0℃ - 10℃，天气寒冷，注意添加衣物。</div>
          <div class="dashbord1-left-rows">
            <el-row :gutter="20">
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <i class="el-icon-sort icon" />今日流量 (1231231)
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <i class="el-icon-s-custom icon" />总用户数 (24001)
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <i class="el-icon-s-comment icon" />好评率 (99%)
                </div>
              </el-col>
            </el-row>
          </div>
        </div>
        <img src="@/assets/dashbord.png" class="dashbord1-right" alt>
      </div>
    </div>
    <div class="gva-card-box">
      <div class="gva-card dashbord2">
        <div>
          <div class="dashbord2-item">
            使用教学：
            <a
              style="color:#409EFF"
              target="view_window"
              href="https://www.bilibili.com/video/BV1Rg411u7xH/"
            >https://www.bilibili.com/video/BV1Rg411u7xH</a>
          </div>
          <div class="dashbord2-item">
            插件仓库：
            <a
              style="color:#409EFF"
              target="view_window"
              href="https://github.com/flipped-aurora/gva-plugins"
            >https://github.com/flipped-aurora/gva-plugins</a>
          </div>
        </div>
      </div>
    </div>
    <div class="gva-card-box">
      <el-card class="gva-card quick-entrance">
        <template #header>
          <div class="card-header">
            <span>快捷入口</span>
          </div>
        </template>
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
              <div class="quick-entrance-item-icon" :style="{ backgroundColor: card.bg }">
                <i :class="card.icon" :style="{ color: card.color }" />
              </div>
              <p>{{ card.label }}</p>
            </div>
          </el-col>
        </el-row>
      </el-card>
    <!-- <div class="quick-entrance-title"></div> -->
    </div>
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>数据统计</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <el-col :xs="24" :sm="18">
              <echarts-line />
            </el-col>
            <el-col :xs="24" :sm="6">
              <dashbord-table />
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import echartsLine from '@/view/dashboard/dashbordCharts/echartsLine.vue'
import dashbordTable from '@/view/dashboard/dashbordTable/dashbordTable.vue'
export default {
  components: {
    echartsLine,
    dashbordTable
  },
  data() {
    return {
      toolCards: [
        {
          label: '用户管理',
          icon: 'el-icon el-icon-monitor',
          name: 'user',
          color: '#ff9c6e',
          bg: 'rgba(255, 156, 110,.3)'
        },
        {
          label: '角色管理',
          icon: 'el-icon el-icon-setting',
          name: 'authority',
          color: '#69c0ff',
          bg: 'rgba(105, 192, 255,.3)'
        },
        {
          label: '菜单管理',
          icon: 'el-icon el-icon-menu',
          name: 'menu',
          color: '#b37feb',
          bg: 'rgba(179, 127, 235,.3)'
        },
        {
          label: '代码生成器',
          icon: ' el-icon-cpu',
          name: 'autoCode',
          color: '#ffd666',
          bg: 'rgba(255, 214, 102,.3)'
        },
        {
          label: '表单生成器',
          icon: 'el-icon-document-checked',
          name: 'formCreate',
          color: '#ff85c0',
          bg: 'rgba(255, 133, 192,.3)'
        },
        {
          label: '关于我们',
          icon: ' el-icon-user',
          name: 'about',
          color: '#5cdbd3',
          bg: 'rgba(92, 219, 211,.3)'
        }
      ]
    }
  },
  methods: {
    toTarget(name) {
      this.$router.push({ name })
    }
  }
}
</script>

<style lang="scss" scoped>
@mixin flex-center {
    display: flex;
    align-items: center;
}
.page {
    background: #f0f2f5;
    padding: 0;
    .gva-card-box{
      padding: 12px;
      &+.gva-card-box{
        padding-top: 0px;
      }
    }
    .gva-card {
      box-sizing: border-box;
        background-color: #fff;
        border-radius: 2px;
        height: auto;
        padding: 26px 30px;
        overflow: hidden;
        box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
    }
    .dashbord1 {
        height: 196px;
        @include flex-center;
        justify-content: space-between;
        color: #777;
        &-left {
          height: 100%;
          display: flex;
          flex-direction: column;
          justify-content: space-between;
            &-title {
                font-size: 22px;
                color: #343844;
            }
            &-dot {
                font-size: 14px;
                color: #6B7687;
            }
            &-rows {
                // margin-top: 15px;
                color: #6B7687;
                width: 600px;
                align-items: center;
            }
        }
        &-right {
            height: 600px;
            width: 600px;
            margin-top: 28px;
        }
    }
     ::v-deep(.el-card__header){
          padding:0;
          border-bottom: none;
        }
        .card-header{
          padding-bottom: 20px;
          border-bottom: 1px solid #e8e8e8;
        }
    .dashbord2 {
        @include flex-center;
        justify-content: flex-start;
        &-item {
            &+.dashbord2-item{
              margin-top:24px;
            }
            line-height: 25px;
        }
    }
    .quick-entrance-title {
        height: 30px;
        font-size: 22px;
        color: #333;
        width: 100%;
        border-bottom: 1px solid #eee;
    }
    .quick-entrance-items {
        @include flex-center;
        justify-content: center;
        text-align: center;
        color: #333;
        .quick-entrance-item {
            height: auto;
            text-align: center;
            // align-items: center;
            &-icon {
                width: 50px;
                height: 50px !important;
                border-radius: 8px;
                @include flex-center;
                justify-content: center;
                margin: 0 auto;
                i {
                    font-size: 24px;
                }
            }
            p {
                margin-top: 10px;
            }
        }
    }
    .echart-box{
      padding: 14px;
    }
}
.icon {
    font-size: 20px;
    color: rgb(85, 160, 248);
    width: 30px;
    height: 30px;
    margin-right: 10px;
    @include flex-center;
}
.flex-center {
    @include flex-center;
}

//小屏幕不显示右侧，将登陆框居中
@media (max-width: 750px) {
    .gva-card {
        padding: 20px 10px !important;
        .dashbord1 {
            height: auto;
            &-left {
                &-title {
                    font-size: 20px !important;
                }
                &-rows {
                    margin-top: 15px;
                    align-items: center;
                }
            }
            &-right {
                display: none;
            }
        }
        .dashbord2 {
            &-item {
                line-height: 20px;
            }
        }
        .icon {
            font-size: 18px;
        }
    }
}
</style>
