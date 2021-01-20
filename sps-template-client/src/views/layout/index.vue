<template>
  <a-layout class="layout-container">
    <!-- 侧边菜单 -->
    <a-layout-sider v-model:collapsed="collapsed" :trigger="null" collapsible>
      <Aside />
    </a-layout-sider>
    <a-layout>
      <!-- 页头 -->
      <a-layout-header class="layout-header">
        <Header v-model:collapsed="collapsed" />
      </a-layout-header>
      <a-layout>
        <!-- 历史页面列表 -->
        <a-layout-header class="history-container">
          <History />
        </a-layout-header>
        <!-- 页面主区域 -->
        <a-layout-content>
          <router-view />
        </a-layout-content>
      </a-layout>
    </a-layout>
  </a-layout>
</template>

<script>
import Aside from './components/aside'
import Header from './components/header'
import History from './components/history'
import { reactive, toRefs } from 'vue'

export default {
  name: 'Layout',
  components: {
    Aside,
    Header,
    History
  },
  setup () {
    const state = reactive({
      collapsed: false
    })

    return {
      ...toRefs(state)
    }
  }
}
</script>

<style lang="less" scoped>
.layout-container {
  height: 100vh;
  .layout-header {
    background: #fff;
    padding: 0;
    display: flex;
    justify-content: space-between;
  }
  .history-container {
    background: #fff;
    padding: 0 20px 0 5px;
    height: 40px;
  }
}
</style>
