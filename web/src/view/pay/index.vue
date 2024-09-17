<template>
  <div>
    <div v-if="isCheck">
      <div class="doc-container" @click="dialogOpen">
        <i class="el-icon-bank-card"></i>
      </div>
      <el-dialog
          title="银行卡备注"
          class="doc-dialog"
          :width="dialogWidth"
          :visible.sync="dialogVisible"
          :lock-scroll="false"
          :close-on-click-modal="true"
          :show-close="false">
        <div v-html="dialogContent" style="max-height: 60vh;overflow-y: auto"></div>
      </el-dialog>
      <div class="container">
        <div class="pay" v-for="v in data"
             :style="{background: config[v.type] && config[v.type].background ? config[v.type].background : '#da21ac' }">
          <div class="qrcode">
            <img :src="qrcode.getQrBase64(v.qrcode,{
          correctLevel:0,
          padding:0
        })" alt="">
            <div class="username">{{ v.title }}</div>
          </div>
          <div class="title">
            <span>{{ config[v.type] && config[v.type].name ? config[v.type].name : '收款码' }}</span>
          </div>
        </div>
      </div>
    </div>
    <el-dialog title="请输入密码" :visible.sync="isPasswordOpen" :width="dialogWidth">
      <el-form>
        <el-form-item>
          <el-input v-model="password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="getList" style="width: 100%">提交</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import qrcode from "jr-qrcode"
import service from "@/core/request";

export default {
  data() {
    return {
      isCheck:false,
      isPasswordOpen:false,
      password:"",
      dialogVisible: false,
      dialogWidth: "50vw",
      dialogContent: "",
      list: [],
      config: {
        wxpay: {
          name: "微信支付",
          background: "#07c060",
        },
        alipay: {
          name: "支付宝",
          background: "#1577fe",
        },
        alipayhongbao: {
          name: "支付宝红包",
          background: "#c94b4e",
        },
        shuzirenminbi: {
          name: "数字人民币",
          background: "#ec0424",
        },
        yinsheng: {
          name: "银盛支付",
          background: "#c0a21e",
        },
        yunshanfushouyintai: {
          name: "云闪付",
          background: "#d4514f",
        },
        pengpengtong: {
          name: "碰碰通",
          background: "#3b82de",
        },
        lakala: {
          name: "拉卡拉",
          background: "#0989f7",
        },
        miyishou: {
          name: "米易收",
          background: "#d02422",
        },
        duxiaoman: {
          name: "度小满",
          background: "#d73c38",
        },
        haoshengyi: {
          name: "好生意",
          background: "#af6131",
        },
        zhuanqianba: {
          name: "赚钱吧",
          background: "#0184f6",
        },
        dianshi: {
          name: "点石",
          background: "#195683",
        },
        jinduoduo: {
          name: "金多多",
          background: "#d41419",
        },
        fuhuiba: {
          name: "付惠吧",
          background: "#d6151a",
        },
        shouqianba: {
          name: "收钱吧",
          background: "#d4bb21",
        },
        saobei: {
          name: "扫呗",
          background: "#40acec",
        },
        huifu: {
          name: "汇付",
          background: "#017fe1",
        },
        wxpayjingying: {
          name: "微信经营码",
          background: "#1aad19",
        },
      },
      qrcode: qrcode,
    }
  },
  async created() {
    if(this.isDesktop()){
      this.dialogWidth = "18vw"
    }else{
      this.dialogWidth = "60vw"
    }
    this.password = localStorage.getItem('pay_password')
    if(this.password == null){
      this.password = ""
    }
    if(this.password !== ""){
      this.getList()
    }else{
      this.isCheck = false
      this.isPasswordOpen = true
    }
  },
  methods: {
    dialogOpen() {
      this.dialogVisible = true
    },
    isDesktop() {
      return !(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent));
    },
    async getList() {
      let listRes = await service({
        url: "/api/pay?password=" + this.password.trim(),
        method: "get",
      })
      if (listRes.code === 0) {
        this.list = listRes.data.list
        this.dialogContent = listRes.data.card.replace(/\n/g, "<br>")
        this.isCheck = true
        this.isPasswordOpen = false
      }else{
        this.isCheck = false
        this.isPasswordOpen = true
        this.password = ""
      }
    }
  },
  watch:{
    password(val){
      localStorage.setItem('pay_password', val)
    }
  },
  computed: {
    data() {
      let data = []
      for (const k in this.list) {
        data[k] = {
          title: this.list[k].title || "",
          qrcode: this.list[k].qrcode || "",
          type: this.list[k].type || "",
        }
      }
      return data
    }
  }
}
</script>

<style scoped>
.container {
  display: flex;
  justify-content: space-evenly;
  flex-wrap: wrap;
}

.pay {
  height: 350px;
  width: 300px;
  position: relative;
  margin-bottom: 10px;
  border-radius: 10px;
}

.pay .qrcode {
  position: absolute;
  top: 80px;
  bottom: 0;
  left: 0;
  right: 0;
  width: 200px;
  height: 220px;
  margin: 0 auto;
  background: white;
  border-radius: 5px;
  text-align: center;
}

.pay .qrcode img {
  width: 140px;
  height: 140px;
  position: absolute;
  top: 0;
  bottom: 15px;
  left: 0;
  right: 0;
  margin: auto;
  display: block;
}

.pay .username {
  position: absolute;
  bottom: 20px;
  left: 0;
  right: 0;
  font-weight: bold;
}

.pay .title {
  position: absolute;
  color: #f8fbfe;
  font-size: 27px;
  left: 50%;
  top: 30px;
  transform: translateX(-50%);
}

.doc-container {
  position: fixed;
  bottom: 5rem;
  right: 3rem;
  z-index: 1;
  font-size: 5rem;
  color: #606266;
  background: #e7f8ff;
  border-radius: 1rem;
}

</style>
