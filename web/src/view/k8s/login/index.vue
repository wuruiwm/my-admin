<template>
  <div class="box">
    <div style="background-color: #ffffff;width:60%;margin: 0 auto;border-radius: 4px;">
      <div class="title">
        Kubernetes Dashboard
      </div>
      <el-form ref="form" :model="form" :rules="rules" label-width="40px" style="margin-top: 20px;"
               @keyup.enter.native="login">
        <el-form-item prop="username">
          <el-input v-model="form.username" placeholder="请输入用户名">
            <i slot="prefix" class="el-input__icon el-icon-user"/>
          </el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input v-model="form.password" placeholder="请输入密码" type="password">
            <i slot="prefix" class="el-input__icon el-icon-lock"/>
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="medium" style="background: #326de6;border-color: #326de6;" type="primary" @click="login">
            登录
          </el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import service from "@/core/request";
import cookie from "vue-cookie"
import {Message} from "element-ui";

export default {
  data() {
    return {
      form: {
        username: "",
        password: "",
      },
      rules: {
        username: [
          {required: true, message: "请填写用户名", trigger: "blur"},
          {
            validator: (rule, value, callback) => {
              if (value.length < 3) {
                return callback(new Error("用户名不能少于三位"))
              } else {
                callback()
              }
            }, trigger: "blur"
          }],
        password: [
          {required: true, message: "请填写密码", trigger: "blur"},
          {
            validator: (rule, value, callback) => {
              if (value.length < 6) {
                return callback(new Error("密码不能少于六位"))
              } else {
                callback()
              }
            }, trigger: "blur"
          }]
      },
      loading: null,
      baseUrl: window.location.origin
    }
  },
  created() {
    this.checkLoginStatus()
  },
  mounted() {
    this.updateIco()
  },
  methods: {
    async checkLoginStatus() {
      let jweToken = cookie.get("jweToken")
      if (!jweToken) {
        return
      }
      let res = await service({
        baseURL: this.baseUrl,
        url: "/api/v1/csrftoken/token",
        method: "get",
      })
      if (!res.hasOwnProperty("token")) {
        return
      }
      let refreshRes = await service({
        baseURL: this.baseUrl,
        url: "/api/v1/token/refresh",
        method: "post",
        headers: {
          "x-csrf-token": res.token,
          "jwetoken": jweToken
        },
        data: {
          "jweToken": jweToken
        }
      })
      if (refreshRes.hasOwnProperty("jweToken")) {
        Message({
          type: "success",
          message: "已登录",
          showClose: true
        })
        location.href = "/"
      }
    },
    updateIco() {
      let icoUrl = "/assets/images/kubernetes-logo.png"
      let ico = document.querySelector('link[rel="icon"]')
      if (ico !== null) {
        ico.href = icoUrl
      } else {
        ico = document.createElement("link")
        ico.rel = "icon"
        ico.href = icoUrl
        document.head.appendChild(ico)
      }
    },
    openLoading() {
      this.loading = this.$loading({
        lock: true,
        text: "Loading",
        spinner: "el-icon-loading",
        background: "rgba(0, 0, 0, 0.7)"
      });
    },
    closeLoading() {
      this.loading.close()
    },
    async login() {
      await this.$refs.form.validate(async (valid) => {
        if (!valid) {
          Message({
            type: "error",
            message: "请正确填写登录信息",
            showClose: true
          })
          return false
        }
        this.openLoading()
        let getToken = service({
          url: "/api/k8s/login",
          method: "post",
          data: this.form
        })
        let getCsrfToken = service({
          baseURL: this.baseUrl,
          url: "/api/v1/csrftoken/login",
          method: "get",
        })
        Promise.all([getToken, getCsrfToken]).then(async (result) => {
          this.closeLoading()
          if (result[0].code !== 0) {
            return
          }
          if (!result[1].hasOwnProperty("token")) {
            Message({
              type: "error",
              message: "登录失败",
              showClose: true
            })
            return
          }
          this.openLoading()
          let loginRes = await service({
            baseURL: this.baseUrl,
            url: "/api/v1/login",
            method: "post",
            headers: {
              "x-csrf-token": result[1].token
            },
            data: {
              token: result[0].data.token
            }
          })
          this.closeLoading()
          if (loginRes.hasOwnProperty("jweToken")) {
            cookie.set("jweToken", loginRes.jweToken, 30)
            location.href = "/"
          } else {
            Message({
              type: "error",
              message: "登录失败",
              showClose: true
            })
          }
        }).catch((error) => {
          this.closeLoading()
        })
      })
    }
  }
}
</script>

<style scoped>
.el-form-item {
  margin-right: 40px
}

.title {
  background: #326de6;
  width: 100%;
  color: #ffffff;
  font-size: 20px;
  height: 50px;
  line-height: 50px;
  text-indent: 22px;
  border-radius: 4px 4px 0 0;
}

.box {
  height: 100%;
  width: 100%;
  background: #edeff0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
