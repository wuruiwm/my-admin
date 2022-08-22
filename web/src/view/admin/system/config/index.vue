<template>
  <div>
    <el-tabs v-model="group" type="card">
      <el-tab-pane label="测试1" name="test1">
        <el-form label-width="80px" style="width: 50%">
          <el-form-item label="t1">
            <el-input v-model="test1.t1"></el-input>
          </el-form-item>
          <el-form-item label="t2">
            <el-input v-model="test1.t2"></el-input>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="测试2" name="test2">
        <el-form label-width="80px" style="width: 50%" @submit.native.prevent>
          <el-form-item label="t1">
            <el-input v-model="test2.t1"></el-input>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>
    <el-button type="primary" @click="submit" style="margin-left: 80px">保存修改</el-button>
  </div>
</template>

<script>
import service from "@/core/request";
import {Message} from "element-ui";

export default {
  data(){
    return {
      group:"test1",
      test1:{
        t1:"",
        t2:"",
      },
      test2:{
        t1:"",
      },
    }
  },
  methods:{
    async submit(){
      let res = await service({
        url: '/admin/config/update',
        method: 'put',
        data: {
          group: this.group,
          data: this.configObjToArray(this[this.group]),
        }
      })
      if (res.code === 0) {
        Message({
          type: 'success',
          message: "保存成功",
          showClose: true
        })
      }
    },
    configObjToArray(obj){
      let arr = []
      for(let k in obj){
        arr.push({
          key:k,
          value:obj[k]
        })
      }
      return arr
    }
  },
  watch:{
    group:{
      immediate:true,
      async handler(val){
        let res = await service({
          url: '/admin/config/list',
          method: 'get',
          params:{
            group:val,
          }
        })
        if(res.code === 0){
          res.data.forEach((v)=>{
            this[val][v.key] = v.value
          })
        }
      }
    }
  }
}
</script>

<style scoped>

</style>
