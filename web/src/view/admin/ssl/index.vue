<template>
  <div>
    <el-form :inline="true">
      <el-form-item>
        <el-button type="primary" class="copy" icon="el-icon-document-copy" size="small" :data-clipboard-text="ssl.key" @click="copy">复制私钥</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" class="copy" icon="el-icon-document-copy" size="small" :data-clipboard-text="ssl.pem" @click="copy">复制公钥</el-button>
      </el-form-item>
      <el-form-item>
        <el-button icon="el-icon-download" size="small" @click="download('key')">下载私钥</el-button>
      </el-form-item>
      <el-form-item>
        <el-button icon="el-icon-download" size="small" @click="download('pem')">下载公钥</el-button>
      </el-form-item>
      <el-form-item>
        <span style="color: #3d763e;">证书到期时间: {{expire_time}}</span>
      </el-form-item>
    </el-form>
    <prism-editor class="my-editor" :value="ssl.key" readonly :highlight="highlighter" line-numbers style="width: 48%;float: left;"></prism-editor>
    <prism-editor class="my-editor" :value="ssl.pem" readonly :highlight="highlighter" line-numbers style="width: 48%;float: right;"></prism-editor>
  </div>
</template>

<script>
import {PrismEditor} from 'vue-prism-editor';
import 'vue-prism-editor/dist/prismeditor.min.css';
import {highlight, languages} from 'prismjs/components/prism-core';
import 'prismjs/components/prism-clike';
import 'prismjs/components/prism-javascript';
import 'prismjs/themes/prism-tomorrow.css';
import Clipboard from "clipboard";
import {Message} from "element-ui";
import service from "@/core/request";

export default {
  components: {
    PrismEditor,
  },
  data() {
    return {
      ssl:{
        key: "",
        pem: "",
        expire_time:"",
      },
    }
  },
  async created() {
    let res = await service({
      url: "/admin/ssl",
      method: "get",
    })
    if (res.code === 0) {
      this.ssl.key = res.data.key
      this.ssl.pem = res.data.pem
      this.ssl.expire_time = res.data.expire_time
    }
  },
  methods: {
    download(type){
      let export_blob = new Blob([this.ssl[type]])
      let save_link = document.createElement("a")
      save_link.href = window.URL.createObjectURL(export_blob)
      save_link.download = type +'.txt'
      save_link.click()
    },
    copy(){
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
</style>
