<template>
  <div class="container">
    <div class="pay" v-for="v in data" :style="{background:v.background}">
      <div class="qrcode">
        <img :src="v.image" alt="">
        <div class="username">{{ v.title }}</div>
      </div>
      <div class="title">
        <span v-if="v.type === 'wxpay'">微信支付</span>
        <span v-if="v.type === 'alipay'">支付宝</span>
        <span v-if="v.type === 'yunsf'">云闪付</span>
        <span v-if="v.type === 'yinsheng'">银盛支付</span>
        <span v-if="v.type === 'pptong'">碰碰通</span>
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
      list: []
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
          title: this.list[k].title,
          qrcode: this.list[k].qrcode,
          type: this.list[k].type,
          image: qrcode.getQrBase64(this.list[k].qrcode)
        }
        let background = {
          wxpay: "#07c060",
          alipay: "#1577fe",
          yinsheng: "#c0a21e",
          yunsf: "#d4514f",
          pptong: "#3b82de"
        }
        data[k].background = background[this.list[k].type]
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
