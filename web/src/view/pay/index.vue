<template>
  <div>
    <a :href="v.qrcode" v-for="v in data">
      <img :src="v.image">
      <br>
      <div style="color: black;text-align: center">{{v.title}}</div>
    </a>
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
        },{
          title: "云闪付收银台",
          qrcode: "https://qr.95516.com/00010048/0cb6a8f06c7826abde476d816ea3fcc7a2467e16e1",
        },
        {
          title: "微信-傍晚升起的太阳",
          qrcode: "wxp://f2f0sgISRWlMF8t5eVub3lgomowu9Dl8m_vY33t90sxuWnM",
        },
        {
          title: "微信-小号1",
          qrcode: "wxp://f2f0W4wsfTfqul8zt0GVlOXtChQYJqnMKe2x_SD5GoIq7Os",
        },
        {
          title: "微信-小号2",
          qrcode: "wxp://f2f0x4ivLzBCU2cxZpMPpBg5cqjfDG0uDVnhQkl46EpkYik",
        },
        {
          title: "支付宝-傍晚升起的太阳",
          qrcode: "https://qr.alipay.com/fkx11920t7ya5soxztchv6b",
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
          image:qrcode.getQrBase64(this.list[k].qrcode)
        }
      }
      console.log(data)
      return data
    }
  },
  methods:{
    t(){
      qrcode.jrQrcode.getQrBase64()
    }
  }
}
</script>

<style scoped>
div > img{
  width: 340px;
}
a {
  display: block;
  width: 256px;
  float: left;
}
</style>
