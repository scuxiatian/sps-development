<template>
  <div class="logo"></div>
  <a-menu
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
import { useRouter } from 'vue-router'
import { useStore, mapGetters } from 'vuex'
import { computed } from 'vue'
import MenuItem from './menuItem'

export default {
  name: 'LayoutAside',
  components: {
    MenuItem
  },
  setup () {
    const router = useRouter()
    const store = useStore()
    const getters = mapGetters('router', ['asyncRouters'])
    const asyncRouters = computed(getters.asyncRouters.bind({ $store: store }))

    const handleMenuClick = ({ key }) => {
      router.push({ name: key })
    }

    return {
      asyncRouters,
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
