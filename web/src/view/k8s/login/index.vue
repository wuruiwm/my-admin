<template>
  <div>
    <el-form ref="form" label-width="80px">
      <el-form-item label="用户名">
        <el-input v-model="form.username"></el-input>
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="form.password" type="password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="login">登录</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import service from "@/core/request";
import cookie from "vue-cookie"
import router from "@/router";

export default {
  data(){
    return {
      form:{
        username:"",
        password:"",
      }
    }
  },
  methods:{
    async login(){
      let k8sLoginRes = await service({
        url: "/api/k8s/login",
        method: "post",
        data: this.form
      })
      if(k8sLoginRes.code === 0){
        let csrfTokenRes = await service({
          baseURL:"/",
          url: "/api/v1/csrftoken/login",
          method: "get",
        })
        if(csrfTokenRes.hasOwnProperty("token") && csrfTokenRes.token){
          let loginRes = await service({
            baseURL:"/",
            url: "/api/v1/login",
            method: "post",
            headers:{
              "x-csrf-token":csrfTokenRes.token
            },
            data:{
              token:k8sLoginRes.data.token
            }
          })
          console.log(loginRes)
          if(loginRes.hasOwnProperty("jweToken") && loginRes.jweToken){
            cookie.set("jweToken",loginRes.jweToken,30)
            location.href = "/"
          }
        }
      }
    }
  }
}
</script>

<style scoped>

</style>
