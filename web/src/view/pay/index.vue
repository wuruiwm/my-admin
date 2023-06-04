<template>
  <div>
    <div class="pay" v-for="v in data" :style="{background:v.background}">
      <div class="qrcode">
        <img :src="v.image" alt="">
        <div class="username">{{v.title}}</div>
      </div>
      <div class="title">
        <span v-if="v.type === 'wxpay'">微信支付</span>
        <span v-if="v.type === 'alipay'">支付宝</span>
        <span v-if="v.type === 'yunsf'">云闪付收银台</span>
        <span v-if="v.type === 'yinsheng'">银盛支付</span>
      </div>
    </div>
  </div>

</template>

<script>
import qrcode from "jr-qrcode"

export default {
  data() {
    return {
      list:[
        {
          title:"银盛支付",
          qrcode:"https://s.ysepay.com/516E7769544A2B55515A5065744456342B52434848684F506E7A6C6C44624130704C7065315A48336257513D.do",
          type: "yinsheng"
        },{
          title: "云闪付收银台",
          qrcode: "https://qr.95516.com/00010048/0cb6a8f06c7826abde476d816ea3fcc7a2467e16e1",
          type:"yunsf"
        },
        {
          title: "傍晚升起的太阳",
          qrcode: "wxp://f2f0sgISRWlMF8t5eVub3lgomowu9Dl8m_vY33t90sxuWnM",
          type:"wxpay"
        },
        {
          title: "小号1",
          qrcode: "wxp://f2f0W4wsfTfqul8zt0GVlOXtChQYJqnMKe2x_SD5GoIq7Os",
          type:"wxpay"
        },
        {
          title: "小号2",
          qrcode: "wxp://f2f0x4ivLzBCU2cxZpMPpBg5cqjfDG0uDVnhQkl46EpkYik",
          type:"wxpay"
        },
        {
          title: "傍晚升起的太阳",
          qrcode: "https://qr.alipay.com/fkx11920t7ya5soxztchv6b",
          type:"alipay"
        },
      ]
    }
  },
  computed:{
    data(){
      let data = []
      for (const k in this.list) {
        data[k] = {
          title:this.list[k].title,
          qrcode:this.list[k].qrcode,
          type:this.list[k].type,
          image:qrcode.getQrBase64(this.list[k].qrcode)
        }
        let background = {
          wxpay:"#07c060",
          alipay:"#1577fe",
          yinsheng:"#c0a21e",
          yunsf:"#d4514f"
        }
        data[k].background = background[this.list[k].type]
      }
      console.log(data)
      return data
    }
  }
}
</script>

<style scoped>
@media screen and (orientation: landscape) {
  .pay{
    height: 350px;
    width: 340px;
    position: relative;
    float: left;
    margin-right: 5px;
    margin-bottom: 5px;
  }
  .pay .qrcode{
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
  .pay .qrcode img{
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
  .pay .username{
    position: absolute;
    bottom: 20px;
    left: 0;
    right: 0;
    font-weight: bold;
  }
  .pay .title{
    position: absolute;
    color: #f8fbfe;
    font-size: 27px;
    left: 50%;
    top: 30px;
    transform: translateX(-50%);
  }
}

@media screen and (orientation: portrait) {
  .pay{
    background: black;
    width: 100px;
  }
}
</style>
