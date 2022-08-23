<template>
  <div>
    <el-form :inline="true">
      <el-form-item label="请输入关键词">
        <el-input v-model="list.param.keyword" placeholder="请输入关键词搜索" size="small"></el-input>
      </el-form-item>
      <el-form-item label="请选择请求方式">
        <el-select placeholder="请选择请求方式" v-model="list.param.method" size="small" clearable>
          <el-option value="GET">GET</el-option>
          <el-option value="POST">POST</el-option>
          <el-option value="PUT">PUT</el-option>
          <el-option value="DELETE">DELETE</el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="请选择分组">
        <el-select placeholder="请选择分组" v-model="list.param.group" size="small" clearable>
          <el-option v-for="v in group" :value="v" :key="v" :label="v"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="创建时间">
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
        <el-button type="primary" icon="el-icon-plus" @click="create" size="small">创建api</el-button>
      </el-form-item>
    </el-form>
    <el-table
        :data="list.data"
        border
        style="width: 100%" :header-cell-style="{backgroundColor:'#fafafa'}">
      <el-table-column
          prop="path"
          label="路径"
          min-width="130">
      </el-table-column>
      <el-table-column
          prop="group"
          label="分组"
          min-width="130">
      </el-table-column>
      <el-table-column
          prop="name"
          label="名称"
          min-width="130">
      </el-table-column>
      <el-table-column
          label="请求方式"
          min-width="80">
        <template v-slot="scope">
          <el-tag
              :type="methodType(scope.row.method)"
              effect="dark"
              size="mini"
          >{{ scope.row.method }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column
          prop="create_time"
          label="创建时间"
          min-width="160">
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
        background
        :current-page="list.page"
        :page-size="list.page_size"
        :page-sizes="[10, 30, 50, 100]"
        :style="{float:'right',padding:'20px'}"
        :total="list.total"
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="pageChange"
        @size-change="pageSizeChange"
    />
    <!--新增编辑弹窗-->
    <el-dialog :title="dialog.title" :visible.sync="dialog.visible" @closed="dialogClosed">
      <el-form ref="form" :model="form" :rules="rules" label-width="100px" @keyup.enter.native="formSubmit">
        <el-form-item label="路径" prop="path">
          <el-input v-model="form.path" size="small" placeholder="请输入路径"/>
        </el-form-item>
        <el-form-item label="请求方式" prop="method">
          <el-select placeholder="请选择请求方式" v-model="form.method" style="width: 100%;" size="small">
            <el-option value="GET">GET</el-option>
            <el-option value="POST">POST</el-option>
            <el-option value="PUT">PUT</el-option>
            <el-option value="DELETE">DELETE</el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="分组" prop="group" ref="formGroup">
          <el-autocomplete
              v-model="form.group"
              :fetch-suggestions="groupSelectSearch"
              placeholder="请输入分组"
              style="width: 100%;"
              size="small"
              @select="groupSelectChecked">
            <template v-slot="{ item }">
              {{ item }}
            </template>
          </el-autocomplete>
        </el-form-item>
        <el-form-item label="名称" prop="name">
          <el-input v-model="form.name" size="small" placeholder="请输入名称"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialog.visible = false" size="small">取 消</el-button>
        <el-button type="primary" @click="formSubmit" size="small">确 定</el-button>
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
        url: "/admin/api/list",
        param: {
          method: "",
          group:"",
        }
      },
      group:[],
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
        group:"",
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
    groupSelectChecked(item){
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
      this.$confirm('此操作将永久删除该api, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
          .then(async () => {
            let res = await service({
              url: '/admin/api/delete',
              method: 'delete',
              data: {
                id: row.id
              }
            })
            if (res.code === 0) {
              Message({
                type: 'success',
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
    async groupList(){
      let res = await service({
        url: '/admin/api/group',
        method: 'get'
      })
      if(res.code === 0){
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
            type: 'error',
            message: '请填写完整后提交',
            showClose: true
          })
          return false
        }
        let successMsg
        let res
        if (this.form.id === "") {
          res = await service({
            url: '/admin/api/create',
            method: 'post',
            data: this.form
          })
          successMsg = "创建成功"
        } else {
          res = await service({
            url: '/admin/api/update',
            method: 'put',
            data: this.form
          })
          successMsg = "修改成功"
        }
        if (res.code === 0) {
          Message({
            type: 'success',
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
