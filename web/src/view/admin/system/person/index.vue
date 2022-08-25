<template>
  <div style="width: 25rem">
    <el-form ref="form" :model="user" :rules="rules" label-width="100px" @keyup.enter.native="formSubmit">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="user.username" disabled placeholder="请输入用户名" size="small"/>
      </el-form-item>
      <el-form-item label="请选择角色" prop="admin_role_id">
        <el-select v-model="user.admin_role_id" disabled placeholder="请选择角色" style="width: 100%">
          <el-option :label="user.admin_role_name" :value="user.admin_role_id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="昵称" prop="username">
        <el-input v-model="user.nickname" placeholder="请输入昵称" size="small"/>
      </el-form-item>
      <el-form-item label="原密码" prop="oldPassword">
        <el-input v-model="user.oldPassword" placeholder="请输入原密码" size="small" type="password"/>
      </el-form-item>
      <el-form-item label="新密码" prop="newPassword">
        <el-input v-model="user.newPassword" placeholder="请输入新密码" size="small" type="password"/>
      </el-form-item>
      <el-form-item label="确认新密码" prop="confirmNewPassword">
        <el-input v-model="user.confirmNewPassword" placeholder="请确认新密码" size="small" type="password"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="formSubmit">保存修改</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import service from "@/core/request"
import {Message} from "element-ui"

export default {
  data() {
    return {
      user: {
        username: "",
        nickname: "",
        oldPassword: "",
        newPassword: "",
        confirmNewPassword: "",
        admin_role_id: "",
        admin_role_name: ""
      },
      rules: {
        username: [
          {required: true, message: "请填写用户名", trigger: "blur"},
        ],
        nickname: [
          {required: true, message: "请填写昵称", trigger: "blur"},
        ],
        oldPassword: [
          {required: true, message: "请填写密码", trigger: "blur"},
          {
            validator: (rule, value, callback) => {
              if (value.length < 6) {
                return callback(new Error("密码不能少于六位"))
              } else {
                callback()
              }
            }, trigger: "blur"
          }
        ],
        newPassword: [
          {required: true, message: "请填写新密码", trigger: "blur"},
          {
            validator: (rule, value, callback) => {
              if (value.length < 6) {
                return callback(new Error("新密码不能少于六位"))
              } else {
                callback()
              }
            }, trigger: "blur"
          }
        ],
        confirmNewPassword: [
          {required: true, message: "请确认新密码", trigger: "blur"},
          {
            validator: (rule, value, callback) => {
              if (this.user.newPassword !== this.user.confirmNewPassword) {
                return callback(new Error("密码不一致"))
              } else {
                callback()
              }
            }, trigger: "blur"
          }
        ],
        admin_role_id: [
          {required: true, message: "请选择角色", trigger: "blur"}
        ],
      }
    }
  },
  methods: {
    async formSubmit() {
      await this.$refs.form.validate(async (valid) => {
        if (!valid) {
          Message({
            type: "error",
            message: "请填写完整后提交",
            showClose: true
          })
          return false
        }
        let res = await service({
          url: "/admin/user/info/update",
          method: "put",
          data: {
            nickname: this.user.nickname,
            old_password: this.user.oldPassword,
            new_password: this.user.newPassword,
          }
        })
        if (res.code === 0) {
          Message({
            type: "success",
            message: "保存成功",
            showClose: true
          })
        }
      })
    }
  },
  async created() {
    let res = await service({
      url: "/admin/user/info",
      method: "get",
    })
    if (res.code === 0) {
      this.user.username = res.data.username
      this.user.nickname = res.data.nickname
      this.user.admin_role_id = res.data.admin_role_id
      this.user.admin_role_name = res.data.admin_role_name
    }
  }
}
</script>

<style scoped>

</style>
