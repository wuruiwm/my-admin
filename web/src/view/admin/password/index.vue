<template>
  <div>
    <el-form :inline="true" @keyup.enter.native="getList">
      <el-form-item label="请输入关键词">
        <el-input v-model="list.param.keyword" placeholder="请输入关键词搜索" size="small"></el-input>
      </el-form-item>
      <el-form-item label="创建时间">
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
        <el-button icon="el-icon-search" size="small" type="primary" @click="getList">搜索</el-button>
      </el-form-item>
      <el-form-item>
        <el-button icon="el-icon-plus" size="small" type="primary" @click="create">创建密码</el-button>
      </el-form-item>
    </el-form>
    <el-table :data="list.data" :header-cell-style="{backgroundColor:'#fafafa'}" border
              style="width: 100%">
      <el-table-column
          label="名称"
          min-width="160"
          prop="name">
      </el-table-column>
      <el-table-column
          label="用户名"
          min-width="160">
        <template v-slot="scope">
          {{ scope.row.username }}
          <i class="el-icon-document-copy copy" :data-clipboard-text="scope.row.username" @click="copy"></i>
        </template>
      </el-table-column>
      <el-table-column
          label="密码"
          min-width="160">
        <template v-slot="scope">
          <span v-if="scope.row.is_show">
            {{ scope.row.password }}
            <img src="@/assets/eye.png" style="height: 16px;" alt="eye" @click="scope.row.is_show = false">
          </span>
          <span v-else>
            **********
            <img src="@/assets/eye-close.png" style="height: 16px;" alt="eye-close" @click="scope.row.is_show = true">
          </span>
          <i class="el-icon-document-copy copy" :data-clipboard-text="scope.row.password" @click="copy"></i>
        </template>
      </el-table-column>
      <el-table-column
          label="备注"
          min-width="160"
          prop="remark"
          show-overflow-tooltip>
        <template v-slot="scope">
          <span v-if="scope.row.remark">{{ scope.row.remark }}</span>
          <span v-else>无备注</span>
        </template>
      </el-table-column>
      <el-table-column
          label="创建时间"
          min-width="160"
          prop="create_time">
      </el-table-column>
      <el-table-column
          fixed="right"
          label="操作"
          min-width="180">
        <template v-slot="scope">
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
        <el-form-item label="名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入名称" size="small"/>
        </el-form-item>
        <el-form-item label="账号" prop="username">
          <el-input v-model="form.username" placeholder="请输入账号" size="small"/>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password" placeholder="请输入密码" size="small"/>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark" placeholder="请输入备注" size="small"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button size="small" @click="dialog.visible = false">取 消</el-button>
        <el-button size="small" type="primary" @click="formSubmit">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import service from "@/core/request"
import {Message} from "element-ui"
import Clipboard from "clipboard"

export default {
  data() {
    return {
      list: {
        url: "/admin/password/list",
      },
      dialog: {
        title: "",
        visible: false,
      },
      rules: {
        name: [
          {required: true, message: "请输入名称", trigger: "blur"}
        ],
        username: [
          {required: true, message: "请输入用户名", trigger: "blur"}
        ],
        password: [
          {required: true, message: "请输入密码", trigger: "blur"}
        ],
      },
      form: {
        id: "",
        name: "",
        username: "",
        password: "",
        remark: ""
      },
    }
  },
  async created() {
    this.getList()
  },
  methods: {
    copy() {
      let clipboard = new Clipboard(".copy")
      clipboard.on('success', e => {
        Message({
          type: 'success',
          message: '复制成功',
          showClose: true
        })
        clipboard.destroy()
      })
      clipboard.on('error', e => {
        Message({
          type: 'error',
          message: '复制失败,',
          showClose: true,
        })
        clipboard.destroy()
      })
    },
    watchListData(data) {
      data.forEach((v) => {
        v.is_show = false
      })
      return data
    },
    create() {
      this.dialog.visible = true
      this.dialog.title = "创建密码"
    },
    update(row) {
      this.dialog.visible = true
      this.dialog.title = "编辑密码"
      this.form.id = row.id
      this.form.name = row.name
      this.form.username = row.username
      this.form.password = row.password
      this.form.remark = row.remark
    },
    del(row) {
      this.$confirm("此操作将永久删除该密码, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
          .then(async () => {
            let res = await service({
              url: "/admin/password/delete",
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
      this.form.name = ""
      this.form.username = ""
      this.form.password = ""
      this.form.remark = ""
      this.form.id = ""
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
            url: "/admin/password/create",
            method: "post",
            data: this.form
          })
          successMsg = "创建成功"
        } else {
          res = await service({
            url: "/admin/password/update",
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
          this.getList()
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
