"use strict"

const path = require("path")

function resolve(dir) {
    return path.join(__dirname, dir)
}

// https://cli.vuejs.org/zh/config/ 配置文档
module.exports = {
    publicPath: "/",
    //编译生成的文件存放路径
    outputDir: "dist",
    //编译生成的静态文件存放路径 在outputDir配置下
    assetsDir: "static",
    //关闭eslint代码检查
    lintOnSave: false,
    //去除map文件 减少编译后的文件大小
    productionSourceMap: false,
    //run serve配置
    devServer: {
        port: process.env.SERVER_PORT,
        //run serve后是否自动打开页面
        open: false,
        //是否展示warning和error报错
        overlay: {
            warnings: true,
            errors: true
        },
        //反代配置 https://github.com/chimurai/http-proxy-middleware
        //vue代理默认先从自己本地找文件 如果找不到才反代
        proxy: {
            "/": { //匹配的路径前缀
                target: `${process.env.VUE_APP_DEVELOP_API_URL}`,//反代地址
                ws: true, //ws支持
            }
        }
    },
    configureWebpack: {
        // @路径走src文件夹
        resolve: {
            alias: {
                "@": resolve("src")
            }
        }
    }
}
