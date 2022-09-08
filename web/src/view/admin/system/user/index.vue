<template>
  <div>
    <el-form :inline="true" @keyup.enter.native="refreshList">
      <el-form-item label="用户名">
        <el-input v-model="list.param.keyword" placeholder="请输入用户名" size="small"></el-input>
      </el-form-item>
      <el-form-item label="最后登录时间">
        <el-date-picker
            v-model="list.date"
            end-placeholder="结束日期"
            range-separator="至"
            size="small"
            start-placeholder="开始日期"
            type="datetimerange"
            value-format="yyyy-MM-dd HH:mm:ss">
        </el-date-picker>
      </el-form-item>
      <el-form-item>
        <el-button icon="el-icon-search" size="small" type="primary" @click="refreshList">搜索</el-button>
      </el-form-item>
      <el-form-item>
        <el-button icon="el-icon-plus" size="small" type="primary" @click="create">创建用户</el-button>
      </el-form-item>
    </el-form>
    <el-table
        :data="list.data"
        :header-cell-style="{backgroundColor:'#fafafa'}"
        border style="width: 100%">
      <el-table-column
          label="用户名"
          min-width="100"
          prop="username">
      </el-table-column>
      <el-table-column
          label="昵称"
          min-width="130"
          prop="nickname">
      </el-table-column>
      <el-table-column
          label="用户角色"
          min-width="150">
        <template v-slot="scope">
          <el-select v-model="scope.row.admin_role_id" placeholder="请选择角色" @change="roleUpdate(scope.row)">
            <el-option v-for="v in role" :key="v.id" :label="v.name" :value="v.id"></el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column
          label="创建时间"
          min-width="160"
          prop="create_time">
      </el-table-column>
      <el-table-column
          label="最后登录时间"
          min-width="160"
          prop="last_login_time">
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
        :current-page="list.page"
        :page-size="list.page_size"
        :page-sizes="[10, 30, 50, 100]"
        :total="list.total"
        background
        layout="total, sizes, prev, pager, next, jumper"
        style="float:right;padding:20px"
        @current-change="pageChange"
        @size-change="pageSizeChange"
    />
    <!--新增编辑弹窗-->
    <el-dialog :title="dialog.title" :visible.sync="dialog.visible" @closed="dialogClosed">
      <el-form ref="form" :model="form" :rules="rules" label-width="100px" @keyup.enter.native="formSubmit">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" placeholder="请输入用户名" size="small"/>
        </el-form-item>
        <el-form-item label="昵称" prop="username">
          <el-input v-model="form.nickname" placeholder="请输入昵称" size="small"/>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password" :placeholder="isEdit ? '留空则不修改' : '请输入密码'" size="small"
                    type="password"/>
        </el-form-item>
        <el-form-item label="请选择角色" prop="admin_role_id">
          <el-select v-model="form.admin_role_id" placeholder="请选择角色" style="width: 100%">
            <el-option v-for="v in role" :key="v.id" :label="v.name" :value="v.id"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button size="small" @click="dialog.visible = false">取 消</el-button>
        <el-button size="small" type="primary" @click="formSubmit">确 定</el-button>
      </div>
    </el-dialog>
    <!--修改密码弹窗-->
    <el-dialog :visible.sync="passwordDialogVisible" title="修改密码" @closed="passwordUpdateDialogClosed">
      <el-form ref="passwordUpdateForm" :model="form" :rules="passwordUpdateRules" label-width="100px"
               @keyup.enter.native="userPasswordUpdate">
        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password" placeholder="请输入密码" size="small" type="password"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button size="small" @click="passwordDialogVisible = false">取 消</el-button>
        <el-button size="small" type="primary" @click="userPasswordUpdate">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import service from "@/core/request"
import {Message} from "element-ui"

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
        url: "/admin/user/role",
        method: "get",
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
              this.refreshList()
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
          this.refreshList()
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
            url: "/admin/user/update",
            method: "put",
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
          this.refreshList()
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
