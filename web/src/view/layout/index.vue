<template>
  <el-container class="layout-cont">
    <el-container :class="[isSider?'openside':'hideside',isMobile ? 'mobile': '']">
      <el-row :class="[isShadowBg?'shadowBg':'']" @click.native="changeShadow()"></el-row>
      <el-aside class="main-cont main-left">
        <Aside class="aside" />
      </el-aside>
      <!-- 分块滑动功能 -->
      <el-main class="main-cont main-right">
        <transition mode="out-in" name="el-fade-in-linear">
          <div
            class="topfix"
            :style="{width: `calc(100% - ${isMobile?'0px':isCollapse?'54px':'220px'})`}"
          >
            <el-header class="header-cont">
              <div @click="totalCollapse" class="menu-total">
                <i class="el-icon-s-unfold" v-if="isCollapse"></i>
                <i class="el-icon-s-fold" v-else></i>
              </div>
              <el-breadcrumb class="breadcrumb" separator-class="el-icon-arrow-right">
                <el-breadcrumb-item
                  :key="item.path"
                  v-for="item in matched.slice(1,matched.length)"
                >{{item.meta.title}}</el-breadcrumb-item>
              </el-breadcrumb>
              <div class="fl-right right-box">
                <el-dropdown>
                  <span class="el-dropdown-link">
                    <img :src="userInfo.headerImg" height="30" width="30" />
                    {{userInfo.title}}
                    <i class="el-icon-arrow-down"></i>
                  </span>
                  <el-dropdown-menu class="dropdown-group" slot="dropdown">
                    <el-dropdown-item>
                      <span>
                        更多信息
                        <el-badge is-dot />
                      </span>
                    </el-dropdown-item>
                    <el-dropdown-item @click.native="showPassword=true" icon="el-icon-s-custom">修改密码</el-dropdown-item>
                    <el-dropdown-item @click.native="toPerson" icon="el-icon-s-custom">个人信息</el-dropdown-item>
                    <el-dropdown-item @click.native="LoginOut" icon="el-icon-table-lamp">登 出</el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
              </div>
            </el-header>
            <!-- 当前面包屑用路由自动生成可根据需求修改 -->
            <!--
            :to="{ path: item.path }" 暂时注释不用-->
            <HistoryComponent />
          </div>
        </transition>
        <transition mode="out-in" name="el-fade-in-linear">
          <keep-alive>
            <router-view v-if="$route.meta.keepAlive" class="admin-box"></router-view>
          </keep-alive>
        </transition>
        <transition mode="out-in" name="el-fade-in-linear">
          <router-view v-if="!$route.meta.keepAlive" class="admin-box"></router-view>
        </transition>
      </el-main>
    </el-container>
    <el-dialog title="修改密码" :visible.sync="showPassword" @close="clearPassword" width="360px">
      <el-form ref="modifyPwdForm" :model="pwdModify" :rules="rules" label-width="80px">
        <el-form-item prop="password" :minlength="6" label="原密码">
          <el-input v-model="pwdModify.password" show-password></el-input>
        </el-form-item>
        <el-form-item prop="newPassword" :minlength="6" label="新密码">
          <el-input v-model="pwdModify.newPassword" show-password></el-input>
        </el-form-item>
        <el-form-item prop="confirmPassword" :minlength="6" label="确认密码">
          <el-input v-model="pwdModify.confirmPassword" show-password></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="showPassword=false">取 消</el-button>
        <el-button type="primary" @click="savePassword">确 定</el-button>
      </div>
    </el-dialog>
  </el-container>
</template>

<script>
import Aside from "@/view/layout/aside";
import HistoryComponent from "@/view/layout/aside/historyComponent/history";

import { mapGetters, mapActions } from "vuex";
import { changePassword } from "@/api/user";
export default {
  name: "Layout",
  data() {
    return {
      isCollapse: false,
      isSider: true,
      isMobile: false,
      isShadowBg: false,
      showPassword: false,
      pwdModify: {},
      rules: {
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          { min: 6, message: "最少6个字符", trigger: "blur" }
        ],
        newPassword: [
          { required: true, message: "请输入新密码", trigger: "blur" },
          { min: 6, message: "最少6个字符", trigger: "blur" }
        ],
        confirmPassword: [
          { required: true, message: "请输入确认密码", trigger: "blur" },
          { min: 6, message: "最少6个字符", trigger: "blur" },
          {
            validator: (rule, value, callback) => {
              if (value !== this.pwdModify.newPassword) {
                callback(new Error("两次密码不一致"));
              } else {
                callback();
              }
            },
            trigger: "blur"
          }
        ]
      }
    };
  },
  components: {
    Aside,
    HistoryComponent
  },
  methods: {
    ...mapActions("user", ["LoginOut"]),
    totalCollapse() {
      this.isCollapse = !this.isCollapse;
      this.isSider = !this.isCollapse;
      this.isShadowBg = !this.isCollapse;
      this.$bus.emit("collapse", this.isCollapse);
    },
    toPerson() {
      this.$router.push({ name: "person" });
    },
    changeShadow() {
      this.isShadowBg = !this.isShadowBg;
      this.isSider = !!this.isCollapse;
      this.totalCollapse();
    },
    savePassword() {
      this.$refs.modifyPwdForm.validate(valid => {
        if (valid) {
          changePassword({
            username: this.userInfo.userName,
            password: this.pwdModify.password,
            newPassword: this.pwdModify.newPassword
          }).then(() => {
            this.$message.success("修改密码成功！");
            this.showPassword = false;
          });
        } else {
          return false;
        }
      });
    },
    clearPassword() {
      this.pwdModify = {
        password: "",
        newPassword: "",
        confirmPassword: ""
      };
      this.$refs.modifyPwdForm.clearValidate();
    }
  },
  computed: {
    ...mapGetters("user", ["userInfo"]),
    ...mapGetters("history", ["historys", "activeValue"]),
    title() {
      return this.$route.meta.title || "当前页面";
    },
    matched() {
      return this.$route.matched;
    }
  },
  mounted() {
    let screenWidth = document.body.clientWidth;
    if (screenWidth < 1000) {
      this.isMobile = true;
      this.isSider = false;
      this.isCollapse = true;
    } else if (screenWidth >= 1000 && screenWidth < 1200) {
      this.isMobile = false;
      this.isSider = false;
      this.isCollapse = true;
    } else {
      this.isMobile = false;
      this.isSider = true;
      this.isCollapse = false;
    }
    this.$bus.emit("collapse", this.isCollapse);
    this.$bus.emit("mobile", this.isMobile);
    window.onresize = () => {
      return (() => {
        let screenWidth = document.body.clientWidth;
        if (screenWidth < 1000) {
          this.isMobile = true;
          this.isSider = false;
          this.isCollapse = true;
        } else if (screenWidth >= 1000 && screenWidth < 1200) {
          this.isMobile = false;
          this.isSider = false;
          this.isCollapse = true;
        } else {
          this.isMobile = false;
          this.isSider = true;
          this.isCollapse = false;
        }
        this.$bus.emit("collapse", this.isCollapse);
        this.$bus.emit("mobile", this.isMobile);
      })();
    };
  }
};
</script>

<style lang="scss">
$headerHigh: 52px;
$mainHight: 100vh;
.el-dropdown-link {
  cursor: pointer;
}
.dropdown-group {
  min-width: 100px;
}
.topfix {
  position: fixed;
  top: 0;
  box-sizing: border-box;
  z-index: 999;
}
.admin-box {
  background-color: rgb(255, 255, 255);
  margin-top: 100px;
}
.el-scrollbar__wrap {
  padding-bottom: 17px;
}
.layout-cont {
  .right-box {
    text-align: center;
    vertical-align: middle;
    img {
      vertical-align: middle;
      border: 1px solid #ccc;
      border-radius: 6px;
    }
  }
  .menu-contorl {
    line-height: 52px;
    font-size: 20px;
    color: #eee;
    display: table-cell;
    vertical-align: middle;
  }
  .header-cont {
    height: $headerHigh !important;
    background: #fff;
    box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
    line-height: $headerHigh;
  }
  .main-cont {
    .breadcrumb {
      line-height: 48px;
      display: inline-block;
      padding: 0 24px;
      // padding: 6px;
      // border-bottom: 1px solid #eee;
    }
    .router-history {
      background: #fff;
      padding: 0 6px;
      border-top: 1px solid #dcdcdc;
    }
    &.el-main {
      overflow: auto;
      background: #fff;
      // padding: 0px 10px;
      // background: #fff;
    }
    height: $mainHight !important;
    overflow: visible;
    position: relative;
    .menu-total {
      // z-index: 5;
      // position: absolute;
      // top: 10px;
      // right: -35px;
      margin-left: -10px;
      float: left;
      margin-top: 10px;
      width: 30px;
      height: 30px;
      line-height: 30px;
      font-size: 30px;
      // border: 0 solid #ffffff;
      // border-radius: 50%;
      // background: #fff;
    }
    .aside {
      overflow: auto;
      // background: #fff;
      &::-webkit-scrollbar {
        display: none;
      }
    }

    .el-menu-vertical {
      height: 100vh !important;
      visibility: auto;
      &:not(.el-menu--collapse) {
        width: 220px;
      }
    }

    &::-webkit-scrollbar {
      display: none;
    }
    &.main-left {
      width: auto !important;
    }
    &.main-right {
      .admin-title {
        float: left;
        font-size: 16px;
        vertical-align: middle;
        margin-left: 20px;
        img {
          vertical-align: middle;
        }
        &.collapse {
          width: 53px;
        }
      }
    }
  }
}
</style>
