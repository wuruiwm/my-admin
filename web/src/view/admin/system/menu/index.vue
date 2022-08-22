<template>
  <div>
    <el-form :inline="true">
      <el-form-item>
        <el-button type="primary"
                   :icon="isSort ?'el-icon-finished' :'el-icon-sort'"
                   @click="isSort = !isSort"
                   size="small">{{ isSort ? "退出排序" : "开始排序" }}
        </el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-plus" @click="create" size="small">创建菜单</el-button>
      </el-form-item>
    </el-form>
    <el-table
        :data="isSort ? sortList : list.data"
        row-key="id"
        border
        style="width: 100%" :header-cell-style="{backgroundColor:'#fafafa'}">
      <el-table-column
          prop="title"
          label="菜单名称"
          min-width="120">
      </el-table-column>
      <el-table-column
          prop="path"
          label="路由path"
          min-width="100">
      </el-table-column>
      <el-table-column
          prop="name"
          label="路由name"
          min-width="100">
      </el-table-column>
      <el-table-column
          prop="component"
          label="文件路径"
          min-width="260">
      </el-table-column>
      <el-table-column
          label="图标"
          min-width="120">
        <template v-slot="scope">
          <div v-if="scope.row.icon">
            <span :class="'el-icon-'+scope.row.icon"></span>&nbsp;
            <span>{{ scope.row.icon }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column
          label="是否隐藏"
          min-width="100">
        <template v-slot="scope">
          <el-tag
              :type="scope.row.is_hidden ? 'danger' : 'success'"
              effect="dark"
              size="mini"
          >{{ scope.row.is_hidden ? "隐藏" : "显示" }}
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
          label="排序"
          min-width="180"
          v-if="isSort">
        <template v-slot="scope">
          <el-button icon="el-icon-top" size="small" :type="scope.row.parent_id ? 'success' : 'primary'"
                     :disabled="!scope.row.top" @click="sort(scope.row,0)"></el-button>
          <el-button icon="el-icon-bottom" size="small" :type="scope.row.parent_id ? 'success' : 'primary'"
                     :disabled="!scope.row.bottom" @click="sort(scope.row,1)"></el-button>
        </template>
      </el-table-column>
      <el-table-column
          fixed="right"
          label="操作"
          min-width="180"
          v-if="!isSort">
        <template v-slot="scope">
          <el-button icon="el-icon-edit" size="small" type="primary" @click="update(scope.row)">编辑</el-button>
          <el-button icon="el-icon-delete" size="small" type="danger" @click="del(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!--新增编辑弹窗-->
    <el-dialog :title="dialog.title" :visible.sync="dialog.visible" @closed="dialogClosed">
      <el-form ref="form" :model="form" :rules="rules" label-width="100px" @keyup.enter.native="formSubmit">
        <el-form-item label="上级菜单" prop="parent_id">
          <el-select placeholder="请选择上级菜单" size="small" v-model="form.parent_id" style="width: 100%;">
            <el-option value="" label="顶级菜单"></el-option>
            <el-option v-for="v in list.data" :value="v.id" :key="v.id" :label="v.title"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="路由path" prop="path">
          <el-input v-model="form.path" size="small" placeholder="请输入路由path"/>
        </el-form-item>
        <el-form-item label="路由name" prop="name">
          <el-input v-model="form.name" size="small" placeholder="请输入路由name"/>
        </el-form-item>
        <el-form-item label="菜单名称" prop="title">
          <el-input v-model="form.title" size="small" placeholder="请输入菜单名称"/>
        </el-form-item>
        <el-form-item label="文件路径" prop="component">
          <el-input v-model="form.component" size="small" placeholder="请输入文件路径"/>
        </el-form-item>
        <el-form-item label="是否隐藏" prop="is_hidden">
          <el-select placeholder="请选择是否隐藏" size="small" v-model="form.is_hidden" style="width: 100%;">
            <el-option :value="0" label="显示"></el-option>
            <el-option :value="1" label="隐藏"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="图标" prop="icon" v-show="!form.is_hidden">
          <Icon :form="form"></Icon>
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
import Icon from '@/components/icon/icon'

export default {
  components: {
    Icon
  },
  data() {
    return {
      list: {
        url: "/admin/menu/list",
      },
      isSort: false,
      dialog: {
        title: "",
        visible: false,
      },
      rules: {
        name: [
          {required: true, message: '请输入路由name', trigger: 'blur'}
        ],
        path: [
          {required: true, message: '请输入路由path', trigger: 'blur'}
        ],
        component: [
          {required: true, message: '请输入文件路径', trigger: 'blur'}
        ],
        title: [
          {required: true, message: '请输入菜单名称', trigger: 'blur'}
        ],
        is_hidden: [
          {required: true, message: '请选择是否隐藏', trigger: 'blur'}
        ],
        icon: [],
      },
      iconRule:[
        {required: true, message: '请选择icon', trigger: 'blur'}
      ],
      form: {
        id: "",
        parent_id: "",
        path:"",
        name: "",
        component: "",
        icon: "",
        title: "",
        is_hidden: 0,
      },
    }
  },
  async created() {
    this.getList()
  },
  computed: {
    sortList() {
      let data = JSON.parse(JSON.stringify(this.list.data))
      data = data.filter((v) => {
        return v.is_hidden === 0
      })
      data.forEach((v) => {
        v.children = v.children.filter((v2) => {
          return v2.is_hidden === 0
        })
      })
      this.arrSortButton(data)
      return data
    }
  },
  watch:{
    'form.is_hidden':{
      immediate:true,
      handler(val){
        if(val === 1){
          this.form.icon = ""
          this.rules.icon = []
        }else{
          this.rules.icon = this.iconRule
        }
      }
    }
  },
  methods: {
    create() {
      this.dialog.visible = true
      this.dialog.title = "创建菜单"
    },
    update(row) {
      this.dialog.visible = true
      this.dialog.title = "编辑菜单"
      this.form.id = row.id
      this.form.parent_id = row.parent_id
      this.form.path = row.path
      this.form.name = row.name
      this.form.component = row.component
      this.form.icon = row.icon
      this.form.title = row.title
      this.form.is_hidden = row.is_hidden
    },
    del(row) {
      this.$confirm('此操作将永久删除该菜单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
          .then(async () => {
            let res = await service({
              url: '/admin/menu/delete',
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
            }
          })
          .catch(() => {
          })
    },
    async getList() {
      let listRes = await service({
        url: this.list.url,
        method: 'get',
      })
      if (listRes.code === 0) {
        this.list.data = []
        //key为父级路由id  value为该路由id下的所有子路由构成的数组
        let menuTmp = {}
        listRes.data.forEach((v) => {
          v.children = []
          v.top = true
          v.bottom = true
          if (v.parent_id === "") {
            this.list.data.push(v)
          } else {
            if (!menuTmp[v.parent_id]) {
              menuTmp[v.parent_id] = []
            }
            menuTmp[v.parent_id].push(v)
          }
        })
        this.list.data.forEach((v) => {
          if (menuTmp[v.id]) {
            v.children = menuTmp[v.id]
          }
        })
        this.arrSortButton(this.list.data)
      }
    },
    arrSortButton(arr) {
      arr.forEach((v, k) => {
        if (k === 0) {
          v.top = false
        }
        if (arr.length === (k + 1)) {
          v.bottom = false
        }
        if (v.children.length) {
          this.arrSortButton(v.children)
        }
      })
    },
    dialogClosed() {
      this.$refs.form.resetFields()
      this.form.id = ""
      this.form.parent_id = ""
      this.form.name = ""
      this.form.component = ""
      this.form.icon = ""
      this.form.title = ""
      this.form.is_hidden = 0
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
            url: '/admin/menu/create',
            method: 'post',
            data: this.form
          })
          successMsg = "创建成功"
        } else {
          res = await service({
            url: '/admin/menu/update',
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
        }
      })
    },
    //type 0上移 1下移
    async sort(row, type) {
      let res = await service({
        url: "/admin/menu/sort",
        method: "put",
        data: {
          id: row.id,
          type: type
        }
      })
      if (res.code === 0) {
        Message({
          type: "success",
          message: "操作成功",
          showClose: true
        })
        this.getList()
      }
    }
  }
}
</script>

<style scoped>

</style>
