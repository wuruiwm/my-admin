<template>
  <div>
    <el-tabs v-model="domain">
      <template v-for="v in this.list">
      <el-tab-pane :label="v.domain" :name="v.domain">
        <el-form :inline="true">
          <el-form-item>
            <el-button :data-clipboard-text="v.key" :disabled="!v.key" class="copy" icon="el-icon-document-copy"
                       size="small"
                       type="primary" @click="copy">复制私钥
            </el-button>
          </el-form-item>
          <el-form-item>
            <el-button :data-clipboard-text="v.pem" :disabled="!v.pem" class="copy" icon="el-icon-document-copy"
                       size="small"
                       type="primary" @click="copy">复制公钥
            </el-button>
          </el-form-item>
          <el-form-item>
            <el-button :disabled="!v.key" icon="el-icon-download" size="small" @click="download(v.key,'key')">下载私钥
            </el-button>
          </el-form-item>
          <el-form-item>
            <el-button :disabled="!v.pem" icon="el-icon-download" size="small" @click="download(v.pem,'pem')">下载公钥
            </el-button>
          </el-form-item>
          <el-form-item v-if="v.expire_time">
            <span style="color: #3d763e;">证书到期时间: {{ v.expire_time }}</span>
          </el-form-item>
        </el-form>
        <prism-editor v-if="v.key" :highlight="highlighter" :value="v.key" class="my-editor" line-numbers readonly
                      style="width: 48%;float: left;"></prism-editor>
        <prism-editor v-if="v.pem" :highlight="highlighter" :value="v.pem" class="my-editor" line-numbers readonly
                      style="width: 48%;float: right;"></prism-editor>
      </el-tab-pane>
      </template>
    </el-tabs>
  </div>
</template>

<script>
import {PrismEditor} from "vue-prism-editor"
import "vue-prism-editor/dist/prismeditor.min.css"
import {highlight, languages} from "prismjs/components/prism-core"
import "prismjs/components/prism-clike"
import "prismjs/components/prism-javascript"
import "prismjs/themes/prism-tomorrow.css"
import Clipboard from "clipboard"
import {Message} from "element-ui"
import service from "@/core/request"

export default {
  components: {
    PrismEditor,
  },
  data() {
    return {
      domain:"",
      list:[],
    }
  },
  async created() {
    let res = await service({
      url: "/admin/ssl",
      method: "get",
    })
    if (res.code === 0) {
      this.list = res.data
      if(this.list.length > 0){
        this.domain = this.list[0].domain
      }
    }
  },
  methods: {
    download(content,type) {
      let export_blob = new Blob([content])
      let save_link = document.createElement("a")
      save_link.href = window.URL.createObjectURL(export_blob)
      save_link.download = type + ".txt"
      save_link.click()
    },
    copy() {
      let clipboard = new Clipboard(".copy")
      clipboard.on("success", e => {
        Message({
          type: "success",
          message: "复制成功",
          showClose: true
        })
        clipboard.destroy()
      })
      clipboard.on("error", e => {
        Message({
          type: "error",
          message: "复制失败",
          showClose: true,
        })
        clipboard.destroy()
      })
    },
    highlighter(code) {
      return highlight(code, languages.plaintext)
    },
  },
}
</script>

<style scoped>
.my-editor {
  background: #fafafa;
  font-family: Fira code, Fira Mono, Consolas, Menlo, Courier, monospace;
  font-size: 14px;
  line-height: 1.5;
  padding: 5px 5px 5px 10px;
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
