import Vue from "vue"
import ElementUI from "element-ui"
import "element-ui/lib/theme-chalk/index.css"
import App from "@/App.vue"
import router from "@/router/index"
import listMixin from "@/core/listMixin"

//关闭vue生成环境提示
Vue.config.productionTip = false
//调整element 点击空白处不关闭弹窗
ElementUI.Dialog.props.closeOnClickModal.default = false
//使用element插件
Vue.use(ElementUI)
//全局list mixin 方便后台列表调用
Vue.mixin(listMixin)

export default new Vue({
    el: "#app",//容器id
    render: h => h(App),
    router,
    //安装全局事件总线 实际上就是把当前的vue实例 挂到vue原型上
    beforeCreate() {
        Vue.prototype.$bus = this
    }
})

