<template>
  <div>
    <!-- 折叠按钮 -->
    <component
      class="trigger"
      :is="collapsed ? 'MenuUnfoldOutlined' : 'MenuFoldOutlined'"
      @click="updateCollapsed"
    ></component>
    <!-- 面包屑导航 -->
    <a-breadcrumb class="header-path">
      <a-breadcrumb-item
        v-for="path of paths"
        :key="path.path"
      >
        {{ path.meta.title }}
      </a-breadcrumb-item>
    </a-breadcrumb>
  </div>
  <div>
    <!-- 用户头像 -->
    <a-avatar :src="userInfo.headerImg" />
    <!-- 下拉菜单 -->
    <a-dropdown class="user-dropdown">
      <span @click="e => e.preventDefault()">{{ userInfo.nickName }}<DownOutlined /></span>
      <template #overlay>
        <a-menu>
          <a-menu-item key="logout" @click="logout"><LogoutOutlined />退出登录</a-menu-item>
        </a-menu>
      </template>
    </a-dropdown>
  </div>
</template>

<script>
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
import { computed } from 'vue'
import CustomPic from '@/components/customPic'

export default {
  name: 'LayoutHeader',
  components: {
    CustomPic
  },
  props: {
    collapsed: {
      type: Boolean,
      default: false
    }
  },
  setup (props, context) {
    const route = useRoute()
    const store = useStore()

    // 更新侧边栏折叠属性
    const updateCollapsed = () => {
      context.emit('update:collapsed', !props.collapsed)
    }

    // 面包屑当前路由
    const paths = computed(() => {
      const arr = route.matched
      return arr.length > 0 ? arr.slice(1) : []
    })

    // 下拉菜单
    const userInfo = computed(() => store.getters['user/userInfo'])
    const logout = () => {
      store.dispatch('user/Logout')
    }

    return {
      userInfo,
      paths,
      updateCollapsed,
      logout
    }
  }
}
</script>

<style lang="less" scoped>
.trigger {
  font-size: 18px;
  line-height: 64px;
  padding: 0 24px;
  cursor: pointer;
  transition: color 0.3s;
  :hover {
    color: #1890ff;
  }
}
.header-path {
  display: inline-block;
}
.user-dropdown {
  margin-right: 30px;
  margin-left: 5px;
  cursor: pointer;
}
</style>
