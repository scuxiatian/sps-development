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
        <a-layout-content class="main-content">
          <router-view v-if="keepAlive" v-slot="{ Component }">
            <keep-alive>
              <component :is="Component"></component>
            </keep-alive>
          </router-view>
          <router-view v-if="!keepAlive" />
        </a-layout-content>
      </a-layout>
    </a-layout>
  </a-layout>
</template>

<script lang="ts">
import Aside from './components/aside/index.vue'
import Header from './components/header/index.vue'
import History from './components/history/index.vue'
import { reactive, toRefs, computed } from 'vue'
import { useRoute } from 'vue-router'

export default {
  name: 'Layout',
  components: {
    Aside,
    Header,
    History
  },
  setup () {
    const route = useRoute()
    const state = reactive({
      collapsed: false
    })

    const keepAlive = computed(() => route.meta.keepAlive)

    return {
      ...toRefs(state),
      keepAlive
    }
  }
}
</script>

<style lang="less" scoped>
.layout-container {
  height: 100vh;
  .layout-header {
    background: @white-bg;
    padding: 0;
    height: 60px;
    display: flex;
    justify-content: space-between;
  }
  .history-container {
    background: @white-bg;
    padding: 0 20px 0 5px;
    height: 40px;
  }
  .main-content {
    height: calc(100vh - 130px);
    overflow: auto;
    padding: 15px;
    margin: 15px;
    background: @white-bg;
  }
}
</style>
