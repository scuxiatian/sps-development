<template>
  <div class="logo"></div>
  <a-menu
    v-model:selectedKeys="selectedKeys"
    v-model:openKeys="openKeys"
    mode="inline"
    theme="dark"
    @click="handleMenuClick">
    <MenuItem
      v-for="item of asyncRouters[0].children"
      :key="item.path"
      :routerInfo="item" />
  </a-menu>
</template>

<script>
import { useRoute, useRouter, onBeforeRouteUpdate } from 'vue-router'
import { useStore, mapGetters } from 'vuex'
import { reactive, toRefs, computed } from 'vue'
import MenuItem from './menuItem'

export default {
  name: 'LayoutAside',
  components: {
    MenuItem
  },
  setup () {
    const route = useRoute()
    const router = useRouter()
    const store = useStore()
    const getters = mapGetters('router', ['asyncRouters'])
    const asyncRouters = computed(getters.asyncRouters.bind({ $store: store }))
    const state = reactive({
      openKeys: [],
      selectedKeys: [route.name]
    })

    const handleMenuClick = ({ key }) => {
      router.push({ name: key })
    }

    onBeforeRouteUpdate((to, _, next) => {
      state.selectedKeys = [to.name]
      const paths = to.fullPath.split('/')
      if (paths.length < 4) {
        state.openKeys = []
        return next()
      }
      state.openKeys = paths.slice(2, -1)
      next()
    })

    return {
      ...toRefs(state),
      asyncRouters,
      // selectedKeys,
      handleMenuClick
    }
  }
}
</script>

<style lang="less" scoped>
.logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.2);
  margin: 16px;
}
</style>
