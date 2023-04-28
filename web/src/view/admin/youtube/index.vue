<template>
    <div>
        <el-form :inline="true" @keyup.enter.native="refreshList">
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
                <el-button icon="el-icon-search" size="small" type="primary" @click="refreshList">搜索</el-button>
            </el-form-item>
            <el-form-item>
                <el-button icon="el-icon-plus" size="small" type="primary" @click="create">创建视频下载</el-button>
            </el-form-item>
        </el-form>
        <el-table :data="list.data" :header-cell-style="{backgroundColor:'#fafafa'}" border
                  style="width: 100%">
            <el-table-column
                    label="音乐名"
                    min-width="80"
                    prop="name">
            </el-table-column>
            <el-table-column
                    label="视频地址"
                    min-width="260"
                    prop="url">
            </el-table-column>
            <el-table-column
                label="执行命令"
                min-width="400"
                prop="command">
            </el-table-column>
            <el-table-column
                label="创建时间"
                min-width="110"
                prop="create_time">
            </el-table-column>
            <el-table-column
                    label="状态"
                    min-width="50"
                    prop="status">
                <template v-slot="scope">
                  <span v-if="scope.row.status === 0">
                      <el-tag type="info" effect="dark">
                        等待
                      </el-tag>
                  </span>
                    <span v-else-if="scope.row.status === 1">
                      <el-tag type="success" effect="dark">
                        成功
                      </el-tag>
                  </span>
                    <span v-else-if="scope.row.status === 2">
                      <el-tag type="warning" effect="dark">
                        失败
                      </el-tag>
                  </span>
                </template>
            </el-table-column>
            <el-table-column
                    fixed="right"
                    label="操作"
                    min-width="290">
                <template v-slot="scope">
                    <el-button v-if="scope.row.status !== 0" icon="el-icon-document" size="small" type="info" @click="log(scope.row)">执行日志</el-button>
                    <el-button icon="el-icon-refresh" size="small" type="primary" @click="retry(scope.row)">重试</el-button>
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
                <el-form-item label="音乐名" prop="name">
                    <el-input v-model="form.name" placeholder="请输入音乐名" size="small"/>
                </el-form-item>
                <el-form-item label="视频地址" prop="url">
                    <el-input v-model="form.url" placeholder="请输入视频地址" size="small"/>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button size="small" @click="dialog.visible = false">取 消</el-button>
                <el-button size="small" type="primary" @click="formSubmit">确 定</el-button>
            </div>
        </el-dialog>
        <el-dialog
            title="查看执行日志"
            :visible.sync="logDialog.visible">
            <prism-editor :highlight="highlighter" :value="logDialog.content" class="my-editor" line-numbers readonly style="width: 100%"></prism-editor>
        </el-dialog>
    </div>
</template>

<script>
import {PrismEditor} from "vue-prism-editor"
import "vue-prism-editor/dist/prismeditor.min.css"
import {highlight, languages} from "prismjs/components/prism-core"
import "prismjs/components/prism-clike"
import "prismjs/components/prism-javascript"
import "prismjs/themes/prism-tomorrow.css"
import service from "@/core/request"
import {Message} from "element-ui"

export default {
    components: {PrismEditor},
    data() {
        return {
            list: {
                url: "/admin/youtube/list",
            },
            dialog: {
                title: "",
                visible: false,
            },
            logDialog :{
                visible: false,
                content:"",
            },
            rules: {
                name: [
                    {required: true, message: "请输入音乐名", trigger: "blur"}
                ],
                url: [
                    {required: true, message: "请输入视频地址", trigger: "blur"}
                ],
            },
            form: {
                id: "",
                name: "",
                url: "",
            },
        }
    },
    async created() {
        this.getList()
    },
    methods: {
        create() {
            this.dialog.visible = true
            this.dialog.title = "创建视频下载"
        },
        log(row){
            this.logDialog.visible = true
            this.logDialog.content = row.content
        },
        retry(row) {
            this.$confirm("此操作将重新下载视频, 是否继续?", "提示", {
                confirmButtonText: "确定",
                cancelButtonText: "取消",
                type: "warning"
            })
                .then(async () => {
                    let res = await service({
                        url: "/admin/youtube/retry",
                        method: "put",
                        data: {
                            id: row.id
                        }
                    })
                    if (res.code === 0) {
                        Message({
                            type: "success",
                            message: "重试成功",
                            showClose: true
                        })
                        this.refreshList()
                    }
                })
                .catch(() => {
                })
        },
        del(row) {
            this.$confirm("此操作将永久删除该视频下载记录, 是否继续?", "提示", {
                confirmButtonText: "确定",
                cancelButtonText: "取消",
                type: "warning"
            })
                .then(async () => {
                    let res = await service({
                        url: "/admin/youtube/delete",
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
            this.form.name = ""
            this.form.url = ""
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
                res = await service({
                    url: "/admin/youtube/create",
                    method: "post",
                    data: this.form
                })
                successMsg = "创建成功"
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
        },
        highlighter(code) {
            return highlight(code, languages.plaintext)
        },
    }
}
</script>

<style scoped>
.my-editor {
    background: #fafafa;
    font-family: Fira code, Fira Mono, Consolas, Menlo, Courier, monospace;
    font-size: 14px;
    line-height: 1.5;
    padding-top: 5px;
    border: 1px solid #e2e2e2;
}
</style>
<style>
.my-editor textarea:focus {
    outline: none;
}

.my-editor textarea::selection {
    background-color: yellow;
}
</style>

