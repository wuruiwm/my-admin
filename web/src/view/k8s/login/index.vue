<template>
  <div style="height: 100%;width: 100%;background: #edeff0;display: flex;justify-content: space-between;align-items: center;">
    <div style="background-color: #ffffff;width:60%;margin: 0 auto;border-radius: 4px;">
      <div style="background: #326de6;width: 100%;color:#ffffff;font-size: 20px;height: 50px;line-height: 50px;text-indent: 22px;border-radius: 4px 4px 0 0;">
        Kubernetes Dashboard
      </div>
      <el-form ref="form" label-width="40px" style="margin-top: 20px;" @keyup.enter.native="login">
        <el-form-item>
          <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item>
          <el-input v-model="form.password" type="password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="medium" @click="login" style="background: #326de6;border-color: #326de6;">登录</el-button>
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
      loading:null,
    }
  },
  methods: {
    openLoading(){
      this.loading = this.$loading({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      });
    },
    closeLoading(){
      this.loading.close()
    },
    login() {
      this.openLoading()
      Promise.all([service({
        url: "/api/k8s/login",
        method: "post",
        data: this.form
      }), service({
        baseURL: "/",
        url: "/api/v1/csrftoken/login",
        method: "get",
      })]).then(async (result) => {
        this.closeLoading()
        if (result[0].code !== 0) {
          return
        }
        if(!result[1].hasOwnProperty("token")){
          Message({
            type: "error",
            message: "登录失败",
            showClose: true
          })
          return
        }
        this.openLoading()
        let loginRes = await service({
          baseURL: "/",
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
        }else{
          Message({
            type: "error",
            message: "登录失败",
            showClose: true
          })
        }
      }).catch((error) => {
        this.closeLoading()
      })
    }
  }
}
</script>

<style scoped>
.el-form-item{
  margin-right: 40px
}
</style>
