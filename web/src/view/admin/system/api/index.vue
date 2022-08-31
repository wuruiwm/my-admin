<template>
  <div>
    <el-form :inline="true" @keyup.enter.native="this.getList">
      <el-form-item label="请输入关键词">
        <el-input v-model="list.param.keyword" placeholder="请输入关键词搜索" size="small"></el-input>
      </el-form-item>
      <el-form-item label="请选择请求方式">
        <el-select v-model="list.param.method" clearable placeholder="请选择请求方式" size="small">
          <el-option value="GET">GET</el-option>
          <el-option value="POST">POST</el-option>
          <el-option value="PUT">PUT</el-option>
          <el-option value="DELETE">DELETE</el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="请选择分组">
        <el-select v-model="list.param.group" clearable placeholder="请选择分组" size="small">
          <el-option v-for="v in group" :key="v" :label="v" :value="v"></el-option>
        </el-select>
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
        <el-button icon="el-icon-search" size="small" type="primary" @click="this.getList">搜索</el-button>
      </el-form-item>
      <el-form-item>
        <el-button icon="el-icon-plus" size="small" type="primary" @click="create">创建api</el-button>
      </el-form-item>
    </el-form>
    <el-table :data="list.data" :header-cell-style="{backgroundColor:'#fafafa'}" border style="width: 100%">
      <el-table-column
          label="路径"
          min-width="130"
          prop="path">
      </el-table-column>
      <el-table-column
          label="分组"
          min-width="130"
          prop="group">
      </el-table-column>
      <el-table-column
          label="名称"
          min-width="130"
          prop="name">
      </el-table-column>
      <el-table-column
          label="请求方式"
          min-width="80">
        <template v-slot="scope">
          <el-tag :type="methodType(scope.row.method)" effect="dark" size="mini">
            {{ scope.row.method }}
          </el-tag>
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
        <el-form-item label="路径" prop="path">
          <el-input v-model="form.path" placeholder="请输入路径" size="small"/>
        </el-form-item>
        <el-form-item label="请求方式" prop="method">
          <el-select v-model="form.method" placeholder="请选择请求方式" size="small" style="width: 100%">
            <el-option value="GET">GET</el-option>
            <el-option value="POST">POST</el-option>
            <el-option value="PUT">PUT</el-option>
            <el-option value="DELETE">DELETE</el-option>
          </el-select>
        </el-form-item>
        <el-form-item ref="formGroup" label="分组" prop="group">
          <el-autocomplete
              v-model="form.group"
              :fetch-suggestions="groupSelectSearch"
              placeholder="请输入分组"
              size="small"
              style="width: 100%"
              @select="groupSelectChecked">
            <template v-slot="{ item }">
              {{ item }}
            </template>
          </el-autocomplete>
        </el-form-item>
        <el-form-item label="名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入名称" size="small"/>
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

export default {
  data() {
    return {
      list: {
        url: "/admin/api/list",
        param: {
          method: "",
          group: "",
        }
      },
      group: [],
      dialog: {
        title: "",
        visible: false,
      },
      rules: {
        path: [
          {required: true, message: "请输入路径", trigger: "blur"}
        ],
        method: [
          {required: true, message: "请选择请求方式", trigger: "blur"}
        ],
        group: [
          {required: true, message: "请输入分组", trigger: "blur"}
        ],
        name: [
          {required: true, message: "请输入名称", trigger: "blur"}
        ],
      },
      form: {
        id: "",
        path: "",
        group: "",
        method: "",
        name: ""
      },
    }
  },
  async created() {
    this.getList()
    this.groupList()
  },
  methods: {
    groupSelectSearch(queryString, cb) {
      let result = queryString ? this.group.filter((v) => {
        return v.indexOf(queryString) !== -1
      }) : this.group
      cb(result)
    },
    groupSelectChecked(item) {
      this.$refs.formGroup.resetField()
      this.form.group = item
    },
    create() {
      this.dialog.visible = true
      this.dialog.title = "创建api"
    },
    update(row) {
      this.dialog.visible = true
      this.dialog.title = "编辑api"
      this.form.id = row.id
      this.form.path = row.path
      this.form.group = row.group
      this.form.method = row.method
      this.form.name = row.name
    },
    del(row) {
      this.$confirm("此操作将永久删除该api, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
          .then(async () => {
            let res = await service({
              url: "/admin/api/delete",
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
              this.groupList()
            }
          })
          .catch(() => {
          })
    },
    async groupList() {
      let res = await service({
        url: "/admin/api/group",
        method: "get"
      })
      if (res.code === 0) {
        this.group = res.data
      }
    },
    dialogClosed() {
      this.$refs.form.resetFields()
      this.form.path = ""
      this.form.group = ""
      this.form.method = ""
      this.form.name = ""
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
            url: "/admin/api/create",
            method: "post",
            data: this.form
          })
          successMsg = "创建成功"
        } else {
          res = await service({
            url: "/admin/api/update",
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
          this.groupList()
        }
      })
    },
    methodType(method) {
      let methodType = {
        GET: "",
        POST: "success",
        PUT: "warning",
        DELETE: "danger",
      }
      if (methodType.hasOwnProperty(method)) {
        return methodType[method]
      } else {
        return ""
      }
    },
  }
}
</script>

<style scoped>

</style>
