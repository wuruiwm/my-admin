<template>
  <div id="admin">
    <el-container class="layout-cont">
      <el-container :class="[isSide ? 'open-side' : 'hide-side',isMobile ? 'mobile' : '']">
        <el-row :class="[isShadowBg ? 'shadowBg' : '']" @click.native="changeShadow()"/>
        <el-aside class="main-cont main-left">
          <div class="admin-title" style="background: #191a23">
            <img alt class="logo-img" src="@/assets/logo.png">
            <h2 v-if="isSide" class="tit-text" style="color:#fff" v-html="'&nbsp;&nbsp;后台管理系统&nbsp;&nbsp;'">
              后台管理系统</h2>
          </div>
          <div class="aside">
            <el-scrollbar style="height:calc(100vh - 64px)">
              <transition :duration="{ enter: 800, leave: 100 }" mode="out-in" name="el-fade-in-linear">
                <el-menu
                    :collapse="isCollapse"
                    :collapse-transition="true"
                    :default-active="active"
                    active-text-color="#1890ff"
                    background-color="#191a23"
                    class="el-menu-vertical"
                    text-color="#fff"
                    unique-opened
                    @select="selectMenuItem"
                >
                  <template v-for="item in adminMenu">
                    <el-submenu v-if="item.children && item.children.length" :key="item.name" ref="subMenu"
                                :index="item.name" :popper-append-to-body="false">
                      <template v-if="!item.meta.is_hidden" slot="title">
                        <i :class="'el-icon-'+item.meta.icon"/>
                        <span slot="title">{{ item.meta.title }}</span>
                      </template>
                      <template v-for="v in item.children">
                        <el-menu-item v-if="!v.meta.is_hidden" :key="v.name" :index="v.name" :route="{parameters:[]}">
                          <i :class="'el-icon-'+v.meta.icon"/>
                          <span slot="title">{{ v.meta.title }}</span>
                        </el-menu-item>
                      </template>
                    </el-submenu>
                    <el-menu-item v-else :key="item.name" :index="item.name" :route="{parameters:[]}">
                      <template v-if="!item.meta.is_hidden">
                        <i :class="'el-icon-'+item.meta.icon"/>
                        <span slot="title">{{ item.meta.title }}</span>
                      </template>
                    </el-menu-item>
                  </template>
                </el-menu>
              </transition>
            </el-scrollbar>
          </div>
        </el-aside>
        <!-- 分块滑动功能 -->
        <el-main class="main-cont main-right">
          <transition :duration="{ enter: 800, leave: 100 }" mode="out-in" name="el-fade-in-linear">
            <div
                :style="{width: `calc(100% - ${isMobile ? '0px' : isCollapse ? '54px' : '220px'})`}"
                class="top-fix"
            >
              <el-row>
                <el-header class="header-cont">
                  <el-col :lg="1" :md="1" :sm="1" :xl="1" :xs="2">
                    <div class="menu-total" @click="totalCollapse">
                      <i v-if="isCollapse" class="el-icon-s-unfold"/>
                      <i v-else class="el-icon-s-fold"/>
                    </div>
                  </el-col>
                  <el-col :lg="14" :md="14" :sm="9" :xl="14" :xs="10">
                    <el-breadcrumb class="breadcrumb" separator-class="el-icon-arrow-right">
                      <el-breadcrumb-item
                          v-for="item in $route.matched.slice(1,$route.matched.length)"
                          :key="item.path"
                      >{{ item.meta.title }}
                      </el-breadcrumb-item>
                    </el-breadcrumb>
                  </el-col>
                  <el-col :lg="9" :md="9" :sm="14" :xl="9" :xs="12">
                    <div class="fl-right right-box">
                      <div class="search-component">
                        <div
                            class="user-box"
                            style="display:inline-block;float:right;width:31px;text-align:left;font-size:16px;padding-top:2px"
                        >
                          <i :class="[reloading ? 'reloading' : '']" class="el-icon-refresh reload"
                             style="cursor:pointer;padding-left:1px" @click="reload"/>
                        </div>
                      </div>
                      <el-dropdown>
                      <span class="header-avatar" style="cursor: pointer">
                        <span style="margin-left: 5px">{{ userInfo.nickname ? userInfo.nickname : "加载中" }}</span>
                        <i class="el-icon-arrow-down"/>
                      </span>
                        <el-dropdown-menu slot="dropdown" class="dropdown-group">
                          <el-dropdown-item icon="el-icon-s-custom"
                                            @click.native="$router.push({name: 'systemPerson'})">个人信息
                          </el-dropdown-item>
                          <el-dropdown-item icon="el-icon-table-lamp" @click.native="loginOut">退出登陆
                          </el-dropdown-item>
                        </el-dropdown-menu>
                      </el-dropdown>
                    </div>
                  </el-col>
                </el-header>
              </el-row>
              <div class="router-history">
                <el-tabs
                    v-model="activeValue"
                    :closable="!(history.length === 1 && this.$route.name === adminDefaultRouterName)"
                    type="card"
                    @tab-click="changeTab"
                    @tab-remove="removeTab"
                >
                  <el-tab-pane
                      v-for="item in history"
                      :key="name(item)"
                      :label="item.meta.title"
                      :name="name(item)"
                      :tab="item"
                      class="gva-tab"
                  >
                  <span slot="label" :style="{color: activeValue===name(item)?'#1890ff':'#333'}">
                    <i :style="{backgroundColor:activeValue===name(item)?'#1890ff':'#ddd'}" class="dot"/>
                    {{ item.meta.title }}
                  </span>
                  </el-tab-pane>
                </el-tabs>
              </div>
            </div>
          </transition>
          <transition mode="out-in" name="el-fade-in-linear">
            <router-view v-if="reloadFlag" v-loading="isLoading" class="admin-box" element-loading-text="正在加载中"/>
          </transition>
          <div class="bottom-info">
            <div>
              Copyright © 2022 💖 wuruiwm
            </div>
          </div>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import router from "@/router"
import {adminDefaultRouterName, adminMenu} from "@/core/adminMenu"
import service from "@/core/request"

export default {
  data() {
    return {
      isCollapse: false,
      isSide: true,
      isMobile: false,
      isShadowBg: false,
      reloadFlag: true,
      reloading: false,
      isLoading: false,
      active: "",
      history: [],
      activeValue: "",
      userInfo: {
        nickname: ""
      },
    }
  },
  computed: {
    adminMenu() {
      return adminMenu
    },
    adminDefaultRouterName() {
      return adminDefaultRouterName
    }
  },
  created() {
    this.getUserInfo()
  },
  mounted() {
    const screenWidth = document.body.clientWidth
    if (screenWidth < 1000) {
      this.isMobile = true
      this.isSide = false
      this.isCollapse = true
    } else if (screenWidth >= 1000 && screenWidth < 1200) {
      this.isMobile = false
      this.isSide = false
      this.isCollapse = true
    } else {
      this.isMobile = false
      this.isSide = true
      this.isCollapse = false
    }
    this.$bus.$on("reload", this.reload)
    this.$bus.$on("showLoading", () => {
      this.isLoading = true
    })
    this.$bus.$on("closeLoading", () => {
      this.isLoading = false
    })
    window.onresize = () => {
      return (() => {
        const screenWidth = document.body.clientWidth
        if (screenWidth < 1000) {
          this.isMobile = true
          this.isSide = false
          this.isCollapse = true
        } else if (screenWidth >= 1000 && screenWidth < 1200) {
          this.isMobile = false
          this.isSide = false
          this.isCollapse = true
        } else {
          this.isMobile = false
          this.isSide = true
          this.isCollapse = false
        }
      })()
    }
  },
  watch: {
    $route: {
      immediate: true,
      handler(to, now) {
        this.active = this.$route.name
        this.activeValue = this.getFmtString(this.$route)
        this.setTab(to)
        if (now && to && now.name === to.name) {
          this.reload()
        }
      }
    }
  },
  methods: {
    reload() {
      this.reloading = true
      this.reloadFlag = false
      setTimeout(() => {
        this.reloading = false
      }, 500)
      this.$nextTick(() => {
        this.reloadFlag = true
      })
    },
    totalCollapse() {
      this.isCollapse = !this.isCollapse
      this.isSide = !this.isCollapse
      this.isShadowBg = !this.isCollapse
    },
    changeShadow() {
      this.isShadowBg = !this.isShadowBg
      this.isSide = !this.isCollapse
      this.totalCollapse()
    },
    loginOut() {
      localStorage.setItem("token", "")
      router.push({name: "adminLogin"})
    },
    selectMenuItem(index) {
      if (index === this.$route.name) {
        return
      }
      this.$router.push({name: index})
    },
    getFmtString(item) {
      return item.name + Object.entries(item.query).toString() + Object.entries(item.params).toString()
    },
    name(item) {
      return item.name + Object.entries(item.query).toString() + Object.entries(item.params).toString()
    },
    isSame(route1, route2) {
      if (route1.name !== route2.name) {
        return false
      }
      if (Object.entries(route1.query).toString() !== Object.entries(route2.query).toString()) {
        return false
      }
      return Object.entries(route1.params).toString() === Object.entries(route2.params).toString()
    },
    setTab(route) {
      if (!this.history.some(item => this.isSame(item, route))) {
        const obj = {}
        obj.name = route.name
        obj.meta = route.meta
        obj.query = route.query
        obj.params = route.params
        this.history.push(obj)
      }
    },
    changeTab(component) {
      const tab = component.$attrs.tab
      this.$router.push({
        name: tab.name,
        query: tab.query,
        params: tab.params
      })
    },
    removeTab(tab) {
      const index = this.history.findIndex(
          item => this.getFmtString(item) === tab
      )
      if (this.getFmtString(this.$route) === tab) {
        if (this.history.length === 1) {
          this.$router.push({name: this.adminDefaultRouterName})
        } else {
          let newIndex
          if (index < this.history.length - 1) {
            newIndex = index + 1
          } else {
            newIndex = index - 1
          }
          this.$router.push({
            name: this.history[newIndex].name,
            query: this.history[newIndex].query,
            params: this.history[newIndex].params
          })
        }
      }
      this.history.splice(index, 1)
    },
    async getUserInfo() {
      let res = await service({
        url: "/admin/user/info",
        method: "get",
      })
      if (res.code === 0) {
        this.userInfo.nickname = res.data.nickname
      }
    },
  }
}
</script>

<style lang="scss">
@import "@/style/base.scss";
@import "@/style/mobile.scss";

#admin {
  background: #eee;
  height: 100vh;
  overflow: hidden;
  font-weight: 400 !important;
}

.el-button {
  font-weight: 400 !important;
}

.dark {
  background-color: #191a23 !important;
  color: #fff !important;
}

.light {
  background-color: #fff !important;
  color: #000 !important;
}

.reload {
  font-size: 17px;

  &:hover {
    transform: scale(1.02)
  }
}

.reloading {
  animation: turn 0.5s linear infinite;
}

@keyframes turn {
  0% {
    -webkit-transform: rotate(0deg);
  }
  25% {
    -webkit-transform: rotate(90deg);
  }
  50% {
    -webkit-transform: rotate(180deg);
  }
  75% {
    -webkit-transform: rotate(270deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
  }
}

.bottom-info {
  color: #888;
  height: 30px;
  line-height: 12px;

  div {
    display: flex;
    justify-content: center;
  }
}

.el-submenu__title, .el-menu-item {
  i {
    color: inherit !important;
  }
}

.el-submenu__title:hover, .el-menu-item:hover {
  i {
    color: inherit !important;
  }

  span {
    color: inherit !important;
  }
}

.el-scrollbar {
  .el-scrollbar__view {
    height: 100%;
  }
}

.contextmenu {
  width: 100px;
  margin: 0;
  border: 1px solid #ccc;
  background: #fff;
  z-index: 3000;
  position: absolute;
  list-style-type: none;
  padding: 5px 0;
  border-radius: 4px;
  font-size: 14px;
  color: #333;
  box-shadow: 2px 2px 3px 0 rgba(0, 0, 0, 0.2);
}

.el-tabs__item .el-icon-close {
  color: initial !important;
}

.el-tabs__item .dot {
  content: "";
  width: 9px;
  height: 9px;
  margin-right: 8px;
  display: inline-block;
  border-radius: 50%;
  transition: background-color .2s;
}
</style>
