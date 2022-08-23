import Vue from 'vue'
import Router from 'vue-router'
import {adminDefaultRouterName, adminRouterInit, isAdminRouterInit} from "@/core/adminMenu";

Vue.use(Router)

// 解决多次跳转路由报错
const originalPush = Router.prototype.push
const originalReplace = Router.prototype.replace
Router.prototype.push = function push(location, onResolve, onReject) {
    if (onResolve || onReject) return originalPush.call(this, location, onResolve, onReject)
    return originalPush.call(this, location).catch(err => err)
}
Router.prototype.replace = function push(location, onResolve, onReject) {
    if (onResolve || onReject) return originalReplace.call(this, location, onResolve, onReject)
    return originalReplace.call(this, location).catch(err => err)
}

const baseRouters = [
    {
        path: '/',
        redirect: {
            name: 'adminLogin'
        }
    },
    {
        path: '/admin/login',
        name: 'adminLogin',
        meta: {
            title: "后台登录"
        },
        component: () => import('@/view/admin/login/index')
    },
    {
        name: "admin",
        path: "/admin",
        component: () => import('@/view/admin/index.vue'),
        meta: {
            title: '后台管理'
        },
        children: [],
    },
    {
        path: '*',
        name: '404',
        meta: {
            title: "404"
        },
        component: () => import('@/view/404/index')
    },
]

// 需要通过后台数据来生成的组件
const createRouter = () => new Router({
    mode: process.env.NODE_ENV === "development" ? "hash" : "history",
    base:"/",
    routes: baseRouters
})

const router = createRouter()

const whiteList = ["adminLogin"]

router.beforeEach(async (to, from, next) => {
    let isWhite = false
    for (let k in whiteList) {
        if (to.name === whiteList[k]) {
            isWhite = true
        }
    }
    //路由在白名单
    if (isWhite) {
        //路由不是adminLogin 直接放行 渲染页面
        //如果是adminLogin 则进行后台路由初始化
        //如果初始化成功 则说明已登录 直接跳转默认后台页面
        //初始化失败 则继续放行 渲染登录页面
        if (to.name === "adminLogin") {
            let isSuccess = await adminRouterInit()
            if (isSuccess) {
                next({name: adminDefaultRouterName, replace: true})
                return
            }
        }
        next()
    } else {
        //不在白名单的时候
        //检查是否初始化 已初始化正常渲染
        //未初始化 去做初始化操作 成功则正常渲染 失败则跳转adminLogin页面
        if (isAdminRouterInit) {
            next()
        } else {
            let isSuccess = await adminRouterInit()
            if (isSuccess) {
                next({path: to.fullPath, replace: true})
            } else {
                if (to.name === "404") {
                    next()
                } else {
                    next({name: "adminLogin"})
                }
            }
        }
    }
})
//全局后置路由守卫 每次路由切换完成以后调用
router.afterEach((to) => {
    if (to.meta.title) {
        document.title = to.meta.title
    } else {
        document.title = "后台管理系统"
    }
})

export default router
