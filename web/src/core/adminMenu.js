import service from "@/core/request"
import router from "@/router";


export let adminMenu = []
export let adminDefaultRouterName = ""
export let isAdminRouterInit = false

function itemToMenu(item) {
    return {
        name: item.name,
        path: item.path,
        component: () => import("@/" + item.component),
        meta: {
            title: item.title,
            icon: item.icon,
            is_hidden: item.is_hidden
        },
        children: item.children
    }
}

export const adminRouterInit = async () => {
    //未登录不请求 直接返回空数组
    const token = localStorage.getItem("token")
    if (!token) {
        return false
    }
    //已初始化过后台路由 不重新初始化
    if (isAdminRouterInit) {
        return true
    }
    //请求后台路由接口
    let menuRes = await service({
        url: "/admin/user/menu",
        method: "get"
    })
    if (menuRes.code !== 0) {
        return false
    }
    //key为父级路由id  value为该路由id下的所有子路由构成的数组
    let menuTmp = {}
    menuRes.data.forEach((v)=>{
        v.children = []
        if(v.parent_id === ""){
            adminMenu.push(v)
        }else{
            if (!menuTmp[v.parent_id]) {
                menuTmp[v.parent_id] = []
            }
            menuTmp[v.parent_id].push(v)
        }
        if (adminDefaultRouterName === "") {
            adminDefaultRouterName = v.name
        }
    })
    //将子路由数组放到父级的children
    adminMenu.forEach((v)=>{
        if (menuTmp[v.id]) {
            v.children = menuTmp[v.id]
        }
    })
    //将后端返回的结构转换为路由需要的结构
    adminMenu.forEach((v1,k1)=>{
        v1.children.forEach((v2,k2)=>{
            v1.children[k2] = itemToMenu(v2)
        })
        adminMenu[k1] = itemToMenu(v1)
    })
    adminMenu.forEach(route => {
        router.addRoute("admin",route)
    })
    router.addRoute("admin",{
        name: "admin404",
        path: "/admin*",
        component: () => import("@/view/404/index"),
        meta: {
            title: "404"
        },
    })
    isAdminRouterInit = true
    return true
}
