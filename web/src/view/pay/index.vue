<template>
  <div class="container">
    <div class="pay" v-for="v in data" :style="{background: config[v.type] && config[v.type].background ? config[v.type].background : '#e7f8ff' }">
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
</template>

<script>
import qrcode from "jr-qrcode"
import service from "@/core/request";

export default {
  data() {
    return {
      list: [],
      config:{
        wxpay:{
          name:"微信支付",
          background:"#07c060",
        },
        alipay:{
          name:"支付宝",
          background:"#1577fe",
        },
        yinsheng:{
          name:"银盛支付",
          background:"#c0a21e",
        },
        yunshanfushouyintai:{
          name:"云闪付",
          background:"#d4514f",
        },
        pengpengtong:{
          name:"碰碰通",
          background:"#3b82de",
        },
        lakala:{
          name:"拉卡拉",
          background:"#0989f7",
        },
        miyishou:{
          name:"米易收",
          background:"#d02422",
        },
        duxiaoman:{
          name:"度小满",
          background:"#d73c38",
        },
        shuzirenminbi:{
          name:"数字人民币",
          background:"#ec0424",
        },
        haoshengyi:{
          name:"好生意",
          background:"#af6131",
        },
        alipayRedPacket:{
          name:"支付宝红包码",
          background:"#dc565b",
        },
      },
      qrcode:qrcode,
    }
  },
  async created() {
    let listRes = await service({
      url: "/api/pay",
      method: "get",
    })
    if (listRes.code === 0) {
      this.list = listRes.data
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
</style>
