<template>
    <div>
        <el-tabs v-model="group">
            <el-tab-pane label="通知配置" name="notice">
                <el-form label-width="150px" style="width: 50%" @keyup.enter.native="submit">
                    <el-form-item label="类型">
                        <el-radio-group v-model="notice.type">
                            <el-radio label="email">邮件推送</el-radio>
                            <el-radio label="gotify">gotify推送</el-radio>
                        </el-radio-group>
                    </el-form-item>
                    <div v-show="notice.type === 'email'">
                        <el-form-item label="主机">
                            <el-input v-model="notice.email_server_host"></el-input>
                        </el-form-item>
                        <el-form-item label="用户名">
                            <el-input v-model="notice.email_username"></el-input>
                        </el-form-item>
                        <el-form-item label="密码">
                            <el-input v-model="notice.email_password"></el-input>
                        </el-form-item>
                        <el-form-item label="ssl加密">
                            <el-radio-group v-model="notice.email_is_encrypt">
                                <el-radio label="0">不加密</el-radio>
                                <el-radio label="1">加密</el-radio>
                            </el-radio-group>
                        </el-form-item>
                        <el-form-item label="端口号">
                            <el-input v-model="notice.email_port"></el-input>
                        </el-form-item>
                        <el-form-item label="接收人">
                            <el-input v-model="notice.email_receive_user"></el-input>
                        </el-form-item>
                    </div>
                    <div v-show="notice.type === 'gotify'">
                        <el-form-item label="主机">
                            <el-input v-model="notice.gotify_server_url"></el-input>
                        </el-form-item>
                        <el-form-item label="token">
                            <el-input v-model="notice.gotify_server_token"></el-input>
                        </el-form-item>
                    </div>
                </el-form>
            </el-tab-pane>
            <el-tab-pane label="ssl证书配置" name="ssl">
                <el-form label-width="150px" style="width: 50%" @keyup.enter.native="submit">
                    <el-form-item label="域名配置">
                      <el-input v-model="ssl.domain" type="textarea" placeholder="请填写域名配置" autosize></el-input>
                    </el-form-item>
                </el-form>
            </el-tab-pane>
            <el-tab-pane label="k8s配置" name="k8s">
                <el-form label-width="150px" style="width: 50%" @keyup.enter.native="submit">
                    <el-form-item label="允许登录的用户">
                        <el-input v-model="k8s.admin_username"></el-input>
                    </el-form-item>
                    <el-form-item label="主机">
                        <el-input v-model="k8s.host"></el-input>
                    </el-form-item>
                    <el-form-item label="端口">
                        <el-input v-model="k8s.port"></el-input>
                    </el-form-item>
                    <el-form-item label="账号">
                        <el-input v-model="k8s.username"></el-input>
                    </el-form-item>
                    <el-form-item label="密码">
                        <el-input v-model="k8s.password"></el-input>
                    </el-form-item>
                </el-form>
            </el-tab-pane>
            <el-tab-pane label="脚本配置" name="script">
                <el-form label-width="150px" style="width: 50%" @submit.native.prevent @keyup.enter.native="submit">
                    <el-form-item label="台服lol幸运抽奖sk">
                        <el-input v-model="script.tw_lol_luck_draw_sk"></el-input>
                    </el-form-item>
                    <el-form-item label="youtube音乐保存目录">
                        <el-input v-model="script.youtube_save_dir"></el-input>
                    </el-form-item>
                  <el-form-item label="m3u8视频保存目录">
                    <el-input v-model="script.m3u8_save_dir"></el-input>
                  </el-form-item>
                  <el-form-item label="youtube cookie">
                    <el-input v-model="script.youtube_cookie" type="textarea" placeholder="请输入youtube cookie" autosize></el-input>
                  </el-form-item>
                </el-form>
            </el-tab-pane>
            <el-tab-pane label="cloudflare配置" name="cloudflare">
                <el-form label-width="150px" style="width: 50%" @submit.native.prevent @keyup.enter.native="submit">
                    <el-form-item label="CloudflareKey">
                        <el-input v-model="cloudflare.key"></el-input>
                    </el-form-item>
                    <el-form-item label="CloudflareEmail">
                        <el-input v-model="cloudflare.email"></el-input>
                    </el-form-item>
                    <el-form-item label="CloudflareZoneId">
                        <el-input v-model="cloudflare.zone_id"></el-input>
                    </el-form-item>
                    <el-form-item label="Cloudflare修改域名">
                        <el-input v-model="cloudflare.dns" type="textarea" placeholder="请输入需要修改域名的json" autosize></el-input>
                    </el-form-item>
                </el-form>
            </el-tab-pane>
            <el-tab-pane label="nas配置" name="nas">
                <el-form label-width="150px" style="width: 50%" @submit.native.prevent @keyup.enter.native="submit">
                    <el-form-item label="主机">
                        <el-input v-model="nas.host"></el-input>
                    </el-form-item>
                    <el-form-item label="用户名">
                        <el-input v-model="nas.username"></el-input>
                    </el-form-item>
                    <el-form-item label="密码">
                        <el-input v-model="nas.password"></el-input>
                    </el-form-item>
                </el-form>
            </el-tab-pane>
            <el-tab-pane label="收款码配置" name="pay">
              <el-form label-width="150px" style="width: 50%" @submit.native.prevent @keyup.enter.native="submit">
                <el-form-item label="银行卡">
                  <el-input v-model="pay.card" type="textarea" placeholder="请输入银行卡备注" autosize></el-input>
                </el-form-item>
                <el-form-item label="密码">
                  <el-input v-model="pay.password"></el-input>
                </el-form-item>
                <el-form-item label="收款码配置">
                  <el-input v-model="pay.config" type="textarea" placeholder="请输入需要收款码配置json" autosize></el-input>
                </el-form-item>
              </el-form>
            </el-tab-pane>
            <el-tab-pane label="openwrt配置" name="openwrt">
              <el-form label-width="150px" style="width: 50%" @keyup.enter.native="submit">
                <el-form-item label="主机">
                  <el-input v-model="openwrt.host"></el-input>
                </el-form-item>
                <el-form-item label="端口">
                  <el-input v-model="openwrt.port"></el-input>
                </el-form-item>
                <el-form-item label="账号">
                  <el-input v-model="openwrt.username"></el-input>
                </el-form-item>
                <el-form-item label="密码">
                  <el-input v-model="openwrt.password"></el-input>
                </el-form-item>
                <el-form-item label="验证域名">
                  <el-input v-model="openwrt.domain"></el-input>
                </el-form-item>
                <el-form-item label="验证域名ip">
                  <el-input v-model="openwrt.domain_ip"></el-input>
                </el-form-item>
              </el-form>
            </el-tab-pane>
          <el-tab-pane label="阿里云配置" name="aliyun">
            <el-form label-width="150px" style="width: 50%" @keyup.enter.native="submit">
              <el-form-item label="AccessKeyId">
                <el-input v-model="aliyun.access_key_id"></el-input>
              </el-form-item>
              <el-form-item label="AccessKeySecret">
                <el-input v-model="aliyun.access_key_secret"></el-input>
              </el-form-item>
              <el-form-item label="区域">
                <el-input v-model="aliyun.region_id"></el-input>
              </el-form-item>
              <el-form-item label="安全组ID">
                <el-input v-model="aliyun.security_group_id"></el-input>
              </el-form-item>
              <el-form-item label="限制流量(G)">
                <el-input v-model="aliyun.limit"></el-input>
              </el-form-item>
              <el-form-item label="已用流量(G)">
                <el-input v-model="aliyunCdtFlow" disabled="disabled"></el-input>
              </el-form-item>
            </el-form>
          </el-tab-pane>
          <el-tab-pane label="抖音配置" name="douyin">
            <el-form label-width="150px" style="width: 50%" @keyup.enter.native="submit">
              <el-form-item label="url">
                <el-input v-model="douyin.url"></el-input>
              </el-form-item>
              <el-form-item label="用户id">
                <el-input v-model="douyin.user_id"></el-input>
              </el-form-item>
              <el-form-item label="Cookie">
                <el-input v-model="douyin.cookie" type="textarea" placeholder="请输入Cookie" autosize></el-input>
              </el-form-item>
            </el-form>
          </el-tab-pane>
        </el-tabs>
        <el-button style="margin-left: 150px" type="primary" @click="submit">保存修改</el-button>
    </div>
</template>

<script>
import service from "@/core/request"
import {Message} from "element-ui"

export default {
    data() {
        return {
            group: "notice",
            notice: {
                type: "email",//可选参数 email gotify
                //email配置
                email_server_host: "",
                email_username: "",
                email_password: "",
                email_is_encrypt: 0,
                email_port: "",
                email_receive_user: "",
                //gotify配置
                gotify_server_url: "",
                gotify_server_token: "",
            },
            ssl: {
                domain: "",
            },
            k8s: {
                admin_username: "",
                host: "",
                port: "",
                username: "",
                password: "",
            },
            script: {
                tw_lol_luck_draw_sk: "",
                youtube_save_dir: "",
                m3u8_save_dir: "",
                youtube_cookie: "",
            },
            cloudflare:{
                email: "",
                key: "",
                zone_id: "",
                dns: "",
            },
            nas:{
                host: "",
                username: "",
                password: "",
            },
            pay:{
                config:"",
                card:"",
                password:"",
            },
            openwrt:{
              host:"",
              port:"",
              username:"",
              password:"",
              domain:"",
              domain_ip:""
            },
            aliyun:{
              access_key_id:"",
              access_key_secret:"",
              region_id:"",
              security_group_id:"",
              limit:0,
            },
            douyin:{
              url:"",
              cookie:"",
            },
            aliyunCdtFlow:"",
        }
    },
    created() {
      this.GetAliyunCdtFlow()
    },
    methods: {
        async submit() {
            let res = await service({
                url: "/admin/config/update",
                method: "put",
                data: {
                    group: this.group,
                    data: this.configObjToArray(this[this.group]),
                }
            })
            if (res.code === 0) {
                Message({
                    type: "success",
                    message: "保存成功",
                    showClose: true
                })
            }
        },
        configObjToArray(obj) {
            let arr = []
            for (let k in obj) {
                arr.push({
                    key: k,
                    value: obj[k]
                })
            }
            return arr
        },
        async GetAliyunCdtFlow() {
          let res = await service({
            url: "/admin/aliyunCdt/flow",
            method: "get"
          })
          if (res.code === 0) {
            this.aliyunCdtFlow = res.data.flow
          }
        }
    },
    watch: {
        group: {
            immediate: true,
            async handler(val) {
                let res = await service({
                    url: "/admin/config/list",
                    method: "get",
                    params: {
                        group: val,
                    }
                })
                if (res.code === 0) {
                    res.data.forEach((v) => {
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
