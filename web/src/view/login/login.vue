<template>
  <div id="userLayout" class="user-layout-wrapper">
    <div class="container">
      <div class="top">
        <div class="desc">
          <img class="logo_login" src="@/assets/logo_login.png" alt="" />
        </div>
        <div class="header">
          <a href="/">
            <!-- <img src="~@/assets/logo.png" class="logo" alt="logo" /> -->
            <span class="title">Gin-Vue-Admin</span>
          </a>
        </div>
      </div>
      <div class="main">
        <el-form
          :model="loginForm"
          :rules="rules"
          ref="loginForm"
          @keyup.enter.native="submitForm"
        >
          <el-form-item prop="username">
            <el-input
              placeholder="请输入用户名"
              v-model="loginForm.username"
            >
            <i
                class="el-input__icon el-icon-user"
                slot="suffix"
              ></i></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
              v-model="loginForm.password"
            >
              <i
                :class="'el-input__icon el-icon-' + lock"
                @click="changeLock"
                slot="suffix"
              ></i>
            </el-input>
          </el-form-item>
          <el-form-item style="position:relative">
            <el-input
              v-model="loginForm.captcha"
              name="logVerify"
              placeholder="请输入验证码"
              style="width:60%"
            />
            <div class="vPic">
              <img
                v-if="picPath"
                :src="path + picPath"
                alt="请输入验证码"
                @click="loginVefify()"
              />
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm" style="width:100%"
              >登 录</el-button
            >
          </el-form-item>
        </el-form>
      </div>

      <div class="footer">
        <div class="links">
          <a href="http://doc.henrongyi.top/"
            ><img src="@/assets/docs.png" class="link-icon"
          /></a>
          <a href="https://www.yuque.com/flipped-aurora/"
            ><img src="@/assets/yuque.png" class="link-icon"
          /></a>
          <a href="https://github.com/flipped-aurora/gin-vue-admin"
            ><img src="@/assets/github.png" class="link-icon"
          /></a>
          <a href="https://space.bilibili.com/322210472"
            ><img src="@/assets/video.png" class="link-icon"
          /></a>
        </div>
        <div class="copyright">
          Copyright &copy; 2020 💖flipped-aurora
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import { captcha } from "@/api/user";
const path = process.env.VUE_APP_BASE_API;
export default {
  name: "Login",
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5 || value.length > 12) {
        return callback(new Error("请输入正确的用户名"));
      } else {
        callback();
      }
    };
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6 || value.length > 12) {
        return callback(new Error("请输入正确的密码"));
      } else {
        callback();
      }
    };

    return {
      lock: "lock",
      loginForm: {
        username: "admin",
        password: "123456",
        captcha: "",
        captchaId: "",
      },
      rules: {
        username: [{ validator: checkUsername, trigger: "blur" }],
        password: [{ validator: checkPassword, trigger: "blur" }],
      },
      path: path,
      logVerify: "",
      picPath: "",
    };
  },
  created() {
    this.loginVefify();
  },
  methods: {
    ...mapActions("user", ["LoginIn"]),
    async login() {
      await this.LoginIn(this.loginForm);
    },
    async submitForm() {
      this.$refs.loginForm.validate(async (v) => {
        if (v) {
          this.login();
          this.loginVefify();
        } else {
          this.$message({
            type: "error",
            message: "请正确填写登录信息",
            showClose: true,
          });
          this.loginVefify();
          return false;
        }
      });
    },
    changeLock() {
      this.lock === "lock" ? (this.lock = "unlock") : (this.lock = "lock");
    },
    loginVefify() {
      captcha({}).then((ele) => {
        this.picPath = ele.data.picPath;
        this.loginForm.captchaId = ele.data.captchaId;
      });
    },
  },
};
</script>

<style scoped lang="scss">
.login-register-box {
  height: 100vh;
  .login-box {
    width: 40vw;
    position: absolute;
    left: 50%;
    margin-left: -22vw;
    top: 5vh;
    .logo {
      height: 35vh;
      width: 35vh;
    }
  }
}

.link-icon {
  width: 20px;
  min-width: 20px;
  height: 20px;
  border-radius: 10px;
}

.vPic {
  width: 33%;
  height: 38px;
  float: right !important;
  background: #ccc;
  img {
    cursor: pointer;
    vertical-align: middle;
  }
}

.logo_login {
  width: 100px;
}

#userLayout.user-layout-wrapper {
  height: 100%;
  position: relative;
  &.mobile {
    .container {
      .main {
        max-width: 368px;
        width: 98%;
      }
    }
  }

  .container {
    width: 100%;
    min-height: 100%;
    background: #f0f2f5 url(~@/assets/background.svg) no-repeat 50%;
    background-size: 100%;
    padding: 110px 0 144px;
    a {
      text-decoration: none;
    }

    .top {
      text-align: center;
      .header {
        height: 44px;
        line-height: 44px;
        margin-bottom: 30px;
        .badge {
          position: absolute;
          display: inline-block;
          line-height: 1;
          vertical-align: middle;
          margin-left: -12px;
          margin-top: -10px;
          opacity: 0.8;
        }

        .logo {
          height: 44px;
          vertical-align: top;
          margin-right: 16px;
          border-style: none;
        }

        .title {
          font-size: 33px;
          color: rgba(0, 0, 0, 0.85);
          font-family: Avenir, "Helvetica Neue", Arial, Helvetica, sans-serif;
          font-weight: 600;
          position: relative;
          top: 2px;
        }
      }
      .desc {
        font-size: 14px;
        color: rgba(0, 0, 0, 0.45);
        margin-top: 12px;
      }
    }

    .main {
      min-width: 260px;
      width: 368px;
      margin: 0 auto;
    }

    .footer {
      position: absolute;
      width: 100%;
      padding: 0 16px;
      margin: 120px 0 24px;
      text-align: center;
      bottom: 100px;
      .links {
        margin-bottom: 8px;
        font-size: 14px;
        width: 330px;
        display: inline-flex;
        flex-direction: row;
        justify-content: space-between;
        padding-right: 40px;
        a {
          color: rgba(0, 0, 0, 0.45);
          transition: all 0.3s;
        }
      }
      .copyright {
        color: rgba(0, 0, 0, 0.45);
        font-size: 14px;
        padding-right: 40px;
      }
    }
  }
}
</style>
