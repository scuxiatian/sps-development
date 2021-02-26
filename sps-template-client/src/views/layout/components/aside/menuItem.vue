<template>
  <a-sub-menu v-if="routerInfo.children && routerInfo.children.length">
    <template #title>
      <component :is="routerInfo.meta.icon"/>
      <span>{{ routerInfo.meta.title }}</span>
    </template>
    <MenuItem
      v-for="item in routerInfo.children"
      :key="item.name"
      :routerInfo="item"
    />
  </a-sub-menu>
  <a-menu-item
    v-else
    :key="routerInfo.name">
    <component :is="routerInfo.meta.icon"/>
    <span>{{ routerInfo.meta.title }}</span>
  </a-menu-item>
</template>

<script>
export default {
  name: 'MenuItem',
  props: {
    routerInfo: {
      default: () => ({}),
      type: Object
    }
  }
}
</script>
