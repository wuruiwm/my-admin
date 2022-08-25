<template>
  <div>
    <el-form :inline="true">
      <el-form-item label="用户名">
        <el-input v-model="list.param.keyword" placeholder="请输入用户名" size="small"></el-input>
      </el-form-item>
      <el-form-item label="最后登录时间">
        <el-date-picker
            v-model="list.date"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="yyyy-MM-dd HH:mm:ss"
            size="small">
        </el-date-picker>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" @click="this.getList" size="small">搜索</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-plus" @click="create" size="small">创建用户</el-button>
      </el-form-item>
    </el-form>
    <el-table
        :data="list.data"
        border
        style="width: 100%" :header-cell-style="{backgroundColor:'#fafafa'}">
      <el-table-column
          prop="username"
          label="用户名"
          min-width="100">
      </el-table-column>
      <el-table-column
          prop="nickname"
          label="昵称"
          min-width="130">
      </el-table-column>
      <el-table-column
          label="用户角色"
          min-width="150">
        <template v-slot="scope">
          <el-select placeholder="请选择角色" v-model="scope.row.admin_role_id" @change="roleUpdate(scope.row)">
            <el-option v-for="v in role" :label="v.name" :value="v.id" :key="v.id"></el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column
          prop="create_time"
          label="创建时间"
          min-width="160">
      </el-table-column>
      <el-table-column
          prop="last_login_time"
          label="最后登录时间"
          min-width="160">
      </el-table-column>
      <el-table-column
          fixed="right"
          label="操作"
          min-width="300">
        <template v-slot="scope">
          <el-button icon="el-icon-edit-outline" size="small" type="primary" @click="passwordUpdate(scope.row)">修改密码
          </el-button>
          <el-button icon="el-icon-edit" size="small" type="primary" @click="update(scope.row)">编辑</el-button>
          <el-button icon="el-icon-delete" size="small" type="danger" @click="del(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
        background
        :current-page="list.page"
        :page-size="list.page_size"
        :page-sizes="[10, 30, 50, 100]"
        style="float:right;padding:20px"
        :total="list.total"
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="pageChange"
        @size-change="pageSizeChange"
    />
    <!--新增编辑弹窗-->
    <el-dialog :title="dialog.title" :visible.sync="dialog.visible" @closed="dialogClosed">
      <el-form ref="form" :model="form" :rules="rules" label-width="100px" @keyup.enter.native="formSubmit">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" size="small" placeholder="请输入用户名"/>
        </el-form-item>
        <el-form-item label="昵称" prop="username">
          <el-input v-model="form.nickname" size="small" placeholder="请输入昵称"/>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="form.password" size="small"
                    :placeholder="isEdit ? '留空则不修改' : '请输入密码'"/>
        </el-form-item>
        <el-form-item label="请选择角色" prop="admin_role_id">
          <el-select placeholder="请选择角色" v-model="form.admin_role_id" style="width: 100%;">
            <el-option v-for="v in role" :label="v.name" :value="v.id" :key="v.id"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialog.visible = false" size="small">取 消</el-button>
        <el-button type="primary" @click="formSubmit" size="small">确 定</el-button>
      </div>
    </el-dialog>
    <!--修改密码弹窗-->
    <el-dialog title="修改密码" :visible.sync="passwordDialogVisible" @closed="passwordUpdateDialogClosed">
      <el-form ref="passwordUpdateForm" :model="form" :rules="passwordUpdateRules" label-width="100px" @keyup.enter.native="userPasswordUpdate">
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="form.password" size="small" placeholder="请输入密码"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="passwordDialogVisible = false" size="small">取 消</el-button>
        <el-button type="primary" @click="userPasswordUpdate" size="small">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import service from "@/core/request";
import {Message} from "element-ui";

export default {
  data() {
    return {
      list: {
        url: "/admin/user/list",
      },
      role: [],
      dialog: {
        title: "",
        visible: false,
      },
      isEdit: false,
      passwordDialogVisible: false,
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
          }
        ],
        nickname: [
          {required: true, message: "请填写昵称", trigger: "blur"},
        ],
        password: [],
        admin_role_id: [
          {required: true, message: "请选择角色", trigger: "blur"}
        ],
      },
      passwordUpdateRules: {
        password: [],
      },
      createPasswordRule: [
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
      form: {
        id: "",
        username: "",
        nickname: "",
        password: "",
        admin_role_id: "",
      },
    }
  },
  async created() {
    await this.getRoleList()
    this.getList()
  },
  methods: {
    async getRoleList() {
      let res = await service({
        url: '/admin/user/role',
        method: 'get',
      })
      if (res.code === 0) {
        this.role = res.data
      }
    },
    async roleUpdate(row) {
      let res = await service({
        url: "/admin/user/roleUpdate",
        method: "put",
        data: {
          id: row.id,
          admin_role_id: row.admin_role_id,
        }
      })
      if (res.code === 0) {
        Message({
          type: "success",
          message: "角色设置成功",
          showClose: true
        })
      }
    },
    create() {
      this.dialog.visible = true
      this.rules.password = this.createPasswordRule
      this.dialog.title = "创建用户"
      this.isEdit = false
    },
    update(row) {
      this.dialog.visible = true
      this.rules.password = []
      this.dialog.title = "编辑用户"
      this.isEdit = true
      this.form.id = row.id
      this.form.nickname = row.nickname
      this.form.username = row.username
      this.form.admin_role_id = row.admin_role_id
    },
    passwordUpdate(row) {
      this.passwordDialogVisible = true
      this.passwordUpdateRules.password = this.createPasswordRule
      this.form.id = row.id
    },
    del(row) {
      this.$confirm("此操作将永久删除该用户, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
          .then(async () => {
            let res = await service({
              url: "/admin/user/delete",
              method: "delete",
              data: {
                id: row.id
              }
            })
            if (res.code === 0) {
              Message({
                type: "success",
                message: "删除成功",
                showClose: true
              })
              this.getList()
            }
          })
          .catch(() => {
          })
    },
    dialogClosed() {
      this.$refs.form.resetFields()
      this.form.id = ""
      this.form.username = ""
      this.form.nickname = ""
      this.form.password = ""
      this.form.admin_role_id = ""
    },
    passwordUpdateDialogClosed() {
      this.$refs.passwordUpdateForm.resetFields()
      this.form.password = ""
    },
    async userPasswordUpdate() {
      await this.$refs.passwordUpdateForm.validate(async (valid) => {
        if (!valid) {
          Message({
            type: "error",
            message: "请填写完整后提交",
            showClose: true
          })
          return false
        }
        let res = await service({
          url: "/admin/user/passwordUpdate",
          method: "put",
          data: this.form
        })
        if (res.code === 0) {
          Message({
            type: "success",
            message: "修改成功",
            showClose: true
          })
          this.passwordDialogVisible = false
          this.getList()
        }
      })
    },
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
        let successMsg
        let res
        if (this.form.id === "") {
          res = await service({
            url: "/admin/user/create",
            method: "post",
            data: this.form
          })
          successMsg = "创建成功"
        } else {
          res = await service({
            url: '/admin/user/update',
            method: 'put',
            data: this.form
          })
          successMsg = "修改成功"
        }
        if (res.code === 0) {
          Message({
            type: "success",
            message: successMsg,
            showClose: true
          })
          this.dialog.visible = false
          this.getList()
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
