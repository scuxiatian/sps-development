<template>
  <a-sub-menu v-if="routerInfo.children && routerInfo.children.length">
    <template #title>
      <component :is="routerInfo.meta.icon"/>
      <span>{{ routerInfo.meta.title }}</span>
    </template>
    <template v-for="item in routerInfo.children">
      <MenuItem
        v-if="!item.meta.hidden"
        :key="item.name"
        :routerInfo="item"
      />
    </template>
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
