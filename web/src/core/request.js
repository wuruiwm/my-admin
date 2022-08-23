import axios from "axios"
import {Message} from "element-ui"
import router from "@/router";
import vue from "@/main"

const service = axios.create({
    baseURL: process.env.NODE_ENV === "development" ? "/" : process.env.VUE_APP_PROD_API_URL,
    timeout: 30000
})
let acitveRequest = 0

function showLoading(){
    acitveRequest++
    vue.$bus.$emit("showLoading")
}
function closeLoading(){
    acitveRequest--
    if (acitveRequest <= 0) {
        vue.$bus.$emit("closeLoading")
    }
}
// http request 拦截器
service.interceptors.request.use(
    config => {
        showLoading()
        const token = localStorage.getItem("token")
        config.data = JSON.stringify(config.data)
        config.headers["Content-Type"] = "application/json"
        if (token) {
            config.headers["X-Admin-Token"] = token
        }
        return config
    },error => {
        closeLoading()
    }
)

// http response 拦截器
service.interceptors.response.use(
    response => {
        closeLoading()
        //内部处理掉
        //用户鉴权失败 跳转的场景
        //常规错误 需要弹窗提示的场景
        //成功场景 在外面处理
        if (response.data.code) {
            if(response.data.code === -1){
                Message({
                    type: "error",
                    message: response.data.msg,
                    showClose: true
                })
                localStorage.setItem("token", "")
                router.push({name: "adminLogin", replace: true})
            }else if(response.data.code === 1){
                Message({
                    type: "error",
                    message: response.data.msg,
                    showClose: true
                })
            }
        }
        return response.data
    },
    error => {
        closeLoading()
        let data = {
            code:1,
            msg:"请求失败,请重试",
            data:null
        }
        Message({
            type: "error",
            message: data.msg,
            showClose: true
        })
        return data
    }
)

export default service
