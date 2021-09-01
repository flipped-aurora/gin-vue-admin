<template>
    <div class="page">
        <div class="mycard dashbord1">
            <div class="dashbord1-left">
                <div class="dashbord1-left-title">早安，管理员，请开始一天的工作吧</div>
                <div class="dashbord1-left-dot">今日晴，0℃ - 10℃，天气寒冷，注意添加衣物。</div>
                <div class="dashbord1-left-rows">
                    <el-row :gutter="20">
                        <el-col :span="8" :xs="24" :sm="8">
                            <div class>
                                <i class="el-icon-sort icon"></i>今日流量 (1231231)
                            </div>
                        </el-col>
                        <el-col :span="8" :xs="24" :sm="8">
                            <div class>
                                <i class="el-icon-s-custom icon"></i>总用户数 (24001)
                            </div>
                        </el-col>
                        <el-col :span="8" :xs="24" :sm="8">
                            <div class>
                                <i class="el-icon-s-comment icon"></i>好评率 (99%)
                            </div>
                        </el-col>
                    </el-row>
                </div>
            </div>
            <img src="@/assets/dashbord.png" class="dashbord1-right" alt />
        </div>
        <div class="mycard dashbord2">
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

        <div class="mycard quick-entrance">
            <el-row :gutter="20">
                <el-col
                    v-for="(card, key) in toolCards"
                    :key="key"
                    :span="4"
                    :xs="8"
                    @click="toTarget(card.name)"
                    class="quick-entrance-items"
                >
                    <div class="quick-entrance-item">
                        <div
                            class="quick-entrance-item-icon"
                            :style="{ backgroundColor: card.bg }"
                        >
                            <i :class="card.icon" :style="{ color: card.color }" />
                        </div>
                        <p>{{ card.label }}</p>
                    </div>
                </el-col>
            </el-row>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            toolCards: [
                {
                    label: '用户管理',
                    icon: 'el-icon el-icon-monitor',
                    name: 'user',
                    color: '#ff9c6e',
                    bg :'rgba(255, 156, 110,.3)'
                },
                {
                    label: '角色管理',
                    icon: 'el-icon el-icon-setting',
                    name: 'authority',
                    color: '#69c0ff',
                    bg :'rgba(105, 192, 255,.3)'
                },
                {
                    label: '菜单管理',
                    icon: 'el-icon el-icon-menu',
                    name: 'menu',
                    color: '#b37feb',
                    bg :'rgba(179, 127, 235,.3)'
                },
                {
                    label: '代码生成器',
                    icon: ' el-icon-cpu',
                    name: 'autoCode',
                    color: '#ffd666',
                    bg :'rgba(255, 214, 102,.3)'
                },
                {
                    label: '表单生成器',
                    icon: 'el-icon-document-checked',
                    name: 'formCreate',
                    color: '#ff85c0',
                    bg :'rgba(255, 133, 192,.3)'
                },
                {
                    label: '关于我们',
                    icon: ' el-icon-user',
                    name: 'about',
                    color: '#5cdbd3',
                    bg :'rgba(92, 219, 211,.3)'
                }
            ]
        }
    }
}
</script>

<style lang="scss" scoped>
.page {
    background: #f8f8f8;
    padding: 10px;

    .mycard {
        background-color: #fff;
        border-radius: 10px;
        height: auto;
        padding: 10px 30px;
        overflow: hidden;
        margin-bottom: 20px;
    }
    .dashbord1 {
        height: 150px;
        display: flex;
        align-items: center;
        justify-content: space-between;
        color: #777;
        &-left {
            &-title {
                font-size: 22px;
                color: #000;
            }
            &-dot {
                font-size: 14px;
                margin-top: 10px;
            }
            &-rows {
                margin-top: 15px;
                width: 600px;
                align-items: center;
            }
        }
        &-right {
            height: 600px;
            width: 600px;
            margin-top: -20px;
        }
    }
    .dashbord2 {
        display: flex;
        align-items: center;
        justify-content: flex-start;
        height: 60px;
        &-item {
            line-height: 25px;
        }
    }

    .quick-entrance-items {
        display: flex;
        align-items: center;
        justify-content: center;
        text-align: center;
        color: #333;
        .quick-entrance-item {
            height: auto;
            &-icon {
                width: 50px;
                height: 50px !important;
                border-radius: 8px;
                display: flex;
                align-items: center;
                justify-content: center;
                i{
                    font-size: 24px;
                }
            }
            p{
                margin-top: 10px;
            }
        }
    }
}
.icon {
    font-size: 22px;
    color: rgb(85, 160, 248);
    width: 30px;
    height: 30px;
    margin-right: 10px;
}

//小屏幕不显示右侧，将登陆框居中
@media (max-width: 750px) {
    .mycard {
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