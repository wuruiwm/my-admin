<template>
  <div id="admin">
    <div id="userLayout">
      <div class="login_panle">
        <div class="login_panle_form">
          <div class="login_panle_form_title">
            <img alt="后台管理系统" class="login_panle_form_title_logo" src="@/assets/logo.png">
            <p class="login_panle_form_title_p" v-html="'&nbsp;&nbsp;后台管理系统&nbsp;&nbsp;'">后台管理系统</p>
          </div>
          <el-form
              ref="loginForm"
              :model="loginForm"
              :rules="rules"
              @keyup.enter.native="submitForm"
          >
            <el-form-item prop="username">
              <el-input v-model="loginForm.username" placeholder="请输入用户名">
                <i slot="suffix" class="el-input__icon el-icon-user"/>
              </el-input>
            </el-form-item>
            <el-form-item prop="password">
              <el-input
                  v-model="loginForm.password"
                  :type="lock === 'lock' ? 'password' : 'text'"
                  placeholder="请输入密码"
              >
                <i
                    slot="suffix"
                    :class="'el-input__icon el-icon-' + lock"
                    @click="changeLock"
                />
              </el-input>
            </el-form-item>
            <div/>
            <el-form-item>
              <el-button
                  style="width: 100%;"
                  type="primary"
                  @click="submitForm"
              >登 录
              </el-button>
            </el-form-item>
          </el-form>
        </div>
        <div class="login_panle_right"/>
        <div class="login_panle_foot">
          <div class="copyright">Copyright &copy; {{ new Date().getFullYear() }} 💖 wuruiwm</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {Message} from "element-ui";
import service from "@/core/request";
import {adminDefaultRouterName, adminRouterInit} from "@/core/adminMenu";
import router from "@/router";

export default {
  data() {
    return {
      lock: 'lock',
      loginForm: {
        username: 'admin',
        password: '123456',
      },
      rules: {
        username: [
          {required: true, message: '请填写用户名', trigger: 'blur'},
          {
            validator: (rule, value, callback) => {
              if (value.length < 3) {
                return callback(new Error('用户名不能少于三位'))
              } else {
                callback()
              }
            }, trigger: 'blur'
          }],
        password: [
          {required: true, message: '请填写密码', trigger: 'blur'},
          {
            validator: (rule, value, callback) => {
              if (value.length < 6) {
                return callback(new Error('密码不能少于六位'))
              } else {
                callback()
              }
            }, trigger: 'blur'
          }]
      },
      logVerify: '',
      picPath: ''
    }
  },
  methods: {
    async submitForm() {
      await this.$refs.loginForm.validate(async (valid) => {
        if (!valid) {
          Message({
            type: 'error',
            message: '请正确填写登录信息',
            showClose: true
          })
          return false
        }
        let loginRes = await service({
          url: '/admin/user/login',
          method: 'post',
          data: this.loginForm
        })
        if (loginRes.code !== 0) {
          return
        }
        localStorage.setItem("token", loginRes.data.token)
        let isSuccess = await adminRouterInit()
        if (isSuccess) {
          Message({
            type: 'success',
            message: '登录成功',
            showClose: true
          })
          router.push({name: adminDefaultRouterName})
        } else {
          Message({
            type: 'success',
            message: '登录成功 路由初始化失败',
            showClose: true
          })
        }
      })
    },
    changeLock() {
      this.lock = this.lock === 'lock' ? 'unlock' : 'lock'
    },
  }
}
</script>

<style lang="scss" scoped>
@import '@/style/base.scss';
@import '@/style/mobile.scss';
#admin {
  background: #eee;
  height: 100vh;
  overflow: hidden;
  font-weight: 400 !important;
}

.el-button {
  font-weight: 400 !important;
}

#userLayout {
  margin: 0;
  padding: 0;
  background-image: url("~@/assets/login_background.svg");
  background-size: cover;
  width: 100%;
  height: 100%;
  position: relative;

  .login_panle {
    position: absolute;
    top: 3vh;
    left: 2vw;
    width: 96vw;
    height: 94vh;
    background-color: rgba(255, 255, 255, .8);
    backdrop-filter: blur(5px);
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: space-evenly;

    .login_panle_right {
      background-image: url("~@/assets/login_left.svg");
      background-size: cover;
      width: 40%;
      height: 60%;
      float: right !important;
    }

    .login_panle_form {
      width: 420px;
      background-color: #fff;
      padding: 40px 40px 40px 40px;
      border-radius: 10px;
      box-shadow: 2px 3px 7px rgba(0, 0, 0, .2);

      .login_panle_form_title {

        display: flex;
        align-items: center;
        margin: 30px 0;

        .login_panle_form_title_logo {
          width: 90px;
          height: 72px;
        }

        .login_panle_form_title_p {
          font-size: 40px;

          padding-left: 20px;
        }
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
    }

    .login_panle_foot {
      position: absolute;
      bottom: 20px;

      .links {
        display: flex;
        align-items: center;
        justify-content: space-between;

        .link-icon {
          width: 30px;
          height: 30px;
        }
      }

      .copyright {
        color: #777777;
        margin-top: 5px;
      }
    }
  }
}

//小屏幕不显示右侧，将登陆框居中
@media (max-width: 750px) {
  .login_panle_right {
    display: none;
  }
  .login_panle {
    width: 100vw;
    height: 100vh;
    top: 0;
    left: 0;
  }
  .login_panle_form {
    width: 100%;
  }
}
</style>
