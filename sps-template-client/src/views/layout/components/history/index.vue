<template>
  <a-tabs
    v-model:activeKey="activeKey"
    :type="closeable ? 'editable-card' : 'card'"
    hideAdd
    @change="changeTab"
    @edit="removeTab">
    <a-tab-pane
      v-for="item of historys"
      :key="item.name"
      :tab="item.meta.title">
    </a-tab-pane>
    <!-- Tab控制下拉菜单 -->
    <template #tabBarExtraContent>
      <a-dropdown :trigger="['click']">
        <a-button @click="e => e.preventDefault()">关闭标签 <DownOutlined /></a-button>
        <template #overlay>
          <a-menu>
            <a-menu-item @click="closeAll">关闭所有</a-menu-item>
            <a-menu-item @click="closeLeft">关闭左侧</a-menu-item>
            <a-menu-item @click="closeRight">关闭右侧</a-menu-item>
            <a-menu-item @click="closeOther">关闭其他</a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </template>
  </a-tabs>
</template>

<script>
import { reactive, toRefs, computed } from 'vue'
import { useRoute, useRouter, onBeforeRouteUpdate } from 'vue-router'

export default {
  name: 'LayoutHistory',
  setup () {
    const route = useRoute()
    const router = useRouter()

    const state = reactive({
      historys: [],
      activeKey: 'dashboard'
    })

    // Tab可关闭
    const closeable = computed(() => {
      return !(state.historys.length === 1 && route.name === 'dashboard')
    })

    // 更新Tab
    const setTab = (route) => {
      if (!state.historys.some((item) => item.name === route.name)) {
        const obj = {
          name: route.name,
          meta: route.meta,
          query: route.query,
          params: route.params
        }
        state.historys.push(obj)
      }
      state.activeKey = route.name
    }

    // 更新Session存储
    const updateSession = () => {
      sessionStorage.setItem('historys', JSON.stringify(state.historys))
    }

    // 初始化历史列表
    const initHistorys = [
      {
        name: 'dashboard',
        meta: {
          title: '仪表盘'
        }
      }
    ]
    state.historys = JSON.parse(sessionStorage.getItem('historys')) || initHistorys
    setTab(route)

    // 切换路由时更新历史列表
    onBeforeRouteUpdate((to) => {
      state.historys = state.historys.filter(item => !item.meta.hidden)
      setTab(to)
      updateSession()
    })

    // 切换Tab
    const changeTab = (targetKey) => {
      router.push({ name: targetKey })
    }

    // 删除Tab
    const removeTab = (targetKey) => {
      const index = state.historys.findIndex(item => item.name === targetKey)
      if (route.name === targetKey) {
        if (state.historys.length === 1) {
          router.push({ name: 'dashboard' })
        } else {
          const toRoute = index < state.historys.length - 1
            ? state.historys[index + 1]
            : state.historys[index - 1]
          router.push({ name: toRoute.name, query: toRoute.query, params: toRoute.params })
        }
      }
      state.historys.splice(index, 1)
      updateSession()
    }

    // 关闭所有标签
    const closeAll = () => {
      router.push({ name: 'dashboard' })
      state.historys = initHistorys
      updateSession()
    }

    // 关闭左侧标签
    const closeLeft = () => {
      const index = state.historys.findIndex(item => item.name === state.activeKey)
      state.historys = state.historys.slice(index)
      updateSession()
    }

    // 关闭右侧标签
    const closeRight = () => {
      const index = state.historys.findIndex(item => item.name === state.activeKey)
      state.historys = state.historys.slice(0, index + 1)
      updateSession()
    }

    // 关闭其他标签
    const closeOther = () => {
      const index = state.historys.findIndex(item => item.name === state.activeKey)
      state.historys = [state.historys[index]]
      updateSession()
    }

    return {
      ...toRefs(state),
      closeable,
      changeTab,
      removeTab,
      closeAll,
      closeLeft,
      closeRight,
      closeOther
    }
  }
}
</script>

<style lang="less" scoped>
</style>
