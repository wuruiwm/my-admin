import service from '@/core/request'
import router from "@/router";


export let adminMenu = []
export let adminDefaultRouterName = ''
export let isAdminRouterInit = false

function itemToMenu(item) {
    return {
        name: item.name,
        path: item.path,
        component: () => import('@/' + item.component),
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
        url: '/admin/user/menu',
        method: 'get'
    })
    if (menuRes.code !== 0) {
        return false
    }
    //key为父级路由id  value为该路由id下的所有子路由构成的数组
    let menuTmp = {}
    for (let k in menuRes.data) {
        menuRes.data[k].children = []
        if (menuRes.data[k].parent_id === "") {
            adminMenu.push(menuRes.data[k])
        } else {
            if (!menuTmp[menuRes.data[k].parent_id]) {
                menuTmp[menuRes.data[k].parent_id] = []
            }
            menuTmp[menuRes.data[k].parent_id].push(menuRes.data[k])
        }
        if (adminDefaultRouterName === "") {
            adminDefaultRouterName = menuRes.data[k].name
        }
    }
    //将子路由数组放到父级的children
    for (let k in adminMenu) {
        if (menuTmp[adminMenu[k].id]) {
            adminMenu[k]['children'] = menuTmp[adminMenu[k].id]
        }
    }
    //将后端返回的结构转换为路由需要的结构
    for (let k1 in adminMenu) {
        for (let k2 in adminMenu[k1].children) {
            adminMenu[k1].children[k2] = itemToMenu(adminMenu[k1].children[k2])
        }
        adminMenu[k1] = itemToMenu(adminMenu[k1])
    }
    adminMenu.forEach(route => {
        router.addRoute("admin",route)
    })
    router.addRoute("admin",{
        name: "admin404",
        path: '/admin*',
        component: () => import('@/view/admin/404/index.vue'),
        meta: {
            title: '404'
        },
    })
    isAdminRouterInit = true
    return true
}
