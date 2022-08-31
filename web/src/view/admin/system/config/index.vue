<template>
  <div>
    <el-tabs v-model="group">
      <el-tab-pane label="通知配置" name="notice">
        <el-form label-width="150px" style="width: 50%" @keyup.enter.native="submit">
          <el-form-item label="类型">
            <el-radio-group v-model="notice.type">
              <el-radio label="email">邮件推送</el-radio>
              <el-radio label="gotify">gotify推送</el-radio>
            </el-radio-group>
          </el-form-item>
          <div v-show="notice.type === 'email'">
            <el-form-item label="主机">
              <el-input v-model="notice.email_server_host"></el-input>
            </el-form-item>
            <el-form-item label="用户名">
              <el-input v-model="notice.email_username"></el-input>
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="notice.email_password"></el-input>
            </el-form-item>
            <el-form-item label="ssl加密">
              <el-radio-group v-model="notice.email_is_encrypt">
                <el-radio label="0">不加密</el-radio>
                <el-radio label="1">加密</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="端口号">
              <el-input v-model="notice.email_port"></el-input>
            </el-form-item>
            <el-form-item label="接收人">
              <el-input v-model="notice.email_receive_user"></el-input>
            </el-form-item>
          </div>
          <div v-show="notice.type === 'gotify'">
            <el-form-item label="主机">
              <el-input v-model="notice.gotify_server_url"></el-input>
            </el-form-item>
            <el-form-item label="token">
              <el-input v-model="notice.gotify_server_token"></el-input>
            </el-form-item>
          </div>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="ssl证书配置" name="ssl">
        <el-form label-width="150px" style="width: 50%" @keyup.enter.native="submit">
          <el-form-item label="KEY(私钥)">
            <el-input v-model="ssl.key"></el-input>
          </el-form-item>
          <el-form-item label="PEM(公钥)">
            <el-input v-model="ssl.pem"></el-input>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="k8s配置" name="k8s">
        <el-form label-width="150px" style="width: 50%" @keyup.enter.native="submit">
          <el-form-item label="允许登录的用户">
            <el-input v-model="k8s.admin_username"></el-input>
          </el-form-item>
          <el-form-item label="主机">
            <el-input v-model="k8s.host"></el-input>
          </el-form-item>
          <el-form-item label="端口">
            <el-input v-model="k8s.port"></el-input>
          </el-form-item>
          <el-form-item label="账号">
            <el-input v-model="k8s.username"></el-input>
          </el-form-item>
          <el-form-item label="密码">
            <el-input v-model="k8s.password"></el-input>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="脚本配置" name="script">
        <el-form label-width="150px" style="width: 50%" @submit.native.prevent @keyup.enter.native="submit">
          <el-form-item label="台服lol幸运抽奖sk">
            <el-input v-model="script.tw_lol_luck_draw_sk"></el-input>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>
    <el-button style="margin-left: 150px" type="primary" @click="submit">保存修改</el-button>
  </div>
</template>

<script>
import service from "@/core/request"
import {Message} from "element-ui"

export default {
  data() {
    return {
      group: "notice",
      notice: {
        type: "email",//可选参数 email gotify
        //email配置
        email_server_host: "",
        email_username: "",
        email_password: "",
        email_is_encrypt: 0,
        email_port: "",
        email_receive_user: "",
        //gotify配置
        gotify_server_url: "",
        gotify_server_token: "",
      },
      ssl: {
        key: "",
        pem: "",
      },
      k8s: {
        admin_username: "",
        host: "",
        port: "",
        username: "",
        password: "",
      },
      script: {
        tw_lol_luck_draw_sk: "",
      },
    }
  },
  methods: {
    async submit() {
      let res = await service({
        url: "/admin/config/update",
        method: "put",
        data: {
          group: this.group,
          data: this.configObjToArray(this[this.group]),
        }
      })
      if (res.code === 0) {
        Message({
          type: "success",
          message: "保存成功",
          showClose: true
        })
      }
    },
    configObjToArray(obj) {
      let arr = []
      for (let k in obj) {
        arr.push({
          key: k,
          value: obj[k]
        })
      }
      return arr
    }
  },
  watch: {
    group: {
      immediate: true,
      async handler(val) {
        let res = await service({
          url: "/admin/config/list",
          method: "get",
          params: {
            group: val,
          }
        })
        if (res.code === 0) {
          res.data.forEach((v) => {
            this[val][v.key] = v.value
          })
        }
      }
    }
  }
}
</script>

<style scoped>

</style>
