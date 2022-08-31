<template>
  <div>
    <el-form :inline="true" @keyup.enter.native="getList">
      <el-form-item label="角色名称">
        <el-input v-model="list.param.keyword" placeholder="请输入角色名称" size="small"></el-input>
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
        <el-button icon="el-icon-plus" size="small" type="primary" @click="create">创建角色</el-button>
      </el-form-item>
    </el-form>
    <el-table :data="list.data" :header-cell-style="{backgroundColor:'#fafafa'}" border style="width: 100%">
      <el-table-column
          label="角色名称"
          min-width="100"
          prop="name">
      </el-table-column>
      <el-table-column
          label="创建时间"
          min-width="160"
          prop="create_time">
      </el-table-column>
      <el-table-column
          fixed="right"
          label="操作"
          min-width="300">
        <template v-slot="scope">
          <el-button icon="el-icon-setting" size="small" type="primary" @click="authUpdate(scope.row)">设置权限
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
      <el-form ref="form" :model="form" :rules="rules" label-width="100px" @keyup.enter.native="formSubmit"
               @submit.native.prevent>
        <el-form-item label="角色名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入角色名称" size="small"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button size="small" @click="dialog.visible = false">取 消</el-button>
        <el-button size="small" type="primary" @click="formSubmit">确 定</el-button>
      </div>
    </el-dialog>
    <!--编辑权限侧边栏-->
    <el-drawer :visible.sync="drawer.visible" :with-header="false" size="40%" title="设置权限">
      <el-tabs class="role-box" type="border-card">
        <el-tab-pane label="菜单权限">
          <div class="clear-flex" style="margin-bottom: 18px;margin-left: 24px">
            <el-button size="small" type="primary" @click="selectAll('menu')">全选</el-button>
            <el-button size="small" type="primary" @click="notSelectAll('menu')">全不选</el-button>
          </div>
          <el-tree ref="menu" :data="auth.menu.data" :default-checked-keys="auth.menu.checked"
                   :props="{children: 'children',label: 'name'}" default-expand-all node-key="id" show-checkbox>
          </el-tree>
        </el-tab-pane>
        <el-tab-pane label="api权限">
          <div class="clear-flex" style="margin-bottom: 18px;margin-left: 24px">
            <el-button size="small" type="primary" @click="selectAll('api')">全选</el-button>
            <el-button size="small" type="primary" @click="notSelectAll('api')">全不选</el-button>
          </div>
          <el-tree ref="api" :data="auth.api.data" :default-checked-keys="auth.api.checked"
                   :props="{children: 'children',label: 'name'}" node-key="id" show-checkbox>
          </el-tree>
        </el-tab-pane>
      </el-tabs>
      <div class="clear-flex" style="margin-top: 18px;margin-left: 40px">
        <el-button size="small" type="primary" @click="authSubmit">确 定</el-button>
      </div>
    </el-drawer>
  </div>
</template>

<script>
import service from "@/core/request"
import {Message} from "element-ui"

export default {
  data() {
    return {
      list: {
        url: "/admin/role/list",
      },
      dialog: {
        title: "",
        visible: false,
      },
      drawer: {
        visible: false,
      },
      auth: {
        id: "",
        menu: {
          data: [],
          checked: [],
        },
        api: {
          data: [],
          checked: [],
        },
      },
      rules: {
        name: [
          {required: true, message: "请输入角色名称", trigger: "blur"}
        ],
      },
      form: {
        id: "",
        name: ""
      },
    }
  },
  created() {
    this.getList()
  },
  methods: {
    selectAll(ref) {
      this.$refs[ref].setCheckedNodes(this.auth[ref].data)
    },
    notSelectAll(ref) {
      this.$refs[ref].setCheckedNodes([])
    },
    create() {
      this.dialog.visible = true
      this.dialog.title = "创建角色"
    },
    update(row) {
      this.dialog.visible = true
      this.dialog.title = "编辑角色"
      this.form.id = row.id
      this.form.name = row.name
    },
    del(row) {
      this.$confirm("此操作将永久删除该角色, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
          .then(async () => {
            let res = await service({
              url: "/admin/role/delete",
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
    async authUpdate(row) {
      this.auth.id = row.id
      let res = await service({
        url: "/admin/role/auth",
        method: "get",
        params: {
          id: row.id
        }
      })
      if (res.code === 0) {
        let menuData = []
        let menuChecked = []
        let apiData = []
        let apiChecked = []
        //key为父级路由id  value为该路由id下的所有子路由构成的数组
        let menuTmp = {}
        let apiTmp = {}
        res.data.menu.forEach((v) => {
          v.children = []
          if (v.parent_id === "") {
            menuData.push(v)
          } else {
            if (v.is_checked === 1) {
              menuChecked.push(v.id)
            }
            if (!menuTmp[v.parent_id]) {
              menuTmp[v.parent_id] = []
            }
            menuTmp[v.parent_id].push(v)
          }
        })
        menuData.forEach((v) => {
          if (menuTmp[v.id]) {
            v.children = menuTmp[v.id]
          } else if (v.is_checked === 1) {
            menuChecked.push(v.id)
          }
        })
        res.data.api.forEach((v) => {
          v.children = []
          if (v.is_checked === 1) {
            apiChecked.push(v.id)
          }
          if (!apiTmp[v.parent_id]) {
            apiTmp[v.parent_id] = []
          }
          apiTmp[v.parent_id].push(v)
        })
        for (let k in apiTmp) {
          apiData.push({
            id: k,
            name: k,
            parent_id: "",
            children: apiTmp[k],
          })
        }
        this.auth.menu.data = menuData
        this.auth.menu.checked = menuChecked
        this.auth.api.data = apiData
        this.auth.api.checked = apiChecked
        this.drawer.visible = true
      }
    },
    async authSubmit() {
      let menu_id = []
      let api_id = []
      this.$refs.menu.getCheckedNodes(false, true).forEach((v) => {
        menu_id.push(v.id)
      })
      this.$refs.api.getCheckedNodes(false, true).forEach((v) => {
        if (v.parent_id) {
          api_id.push(v.id)
        }
      })
      let res = await service({
        url: "/admin/role/authUpdate",
        method: "put",
        data: {
          id: this.auth.id,
          menu_id: menu_id,
          api_id: api_id,
        }
      })
      if (res.code === 0) {
        Message({
          type: "success",
          message: "设置权限成功",
          showClose: true
        })
        this.drawer.visible = false
      }
    },
    dialogClosed() {
      this.$refs.form.resetFields()
      this.form.id = ""
      this.form.name = ""
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
            url: "/admin/role/create",
            method: "post",
            data: this.form
          })
          successMsg = "创建成功"
        } else {
          res = await service({
            url: "/admin/role/update",
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
