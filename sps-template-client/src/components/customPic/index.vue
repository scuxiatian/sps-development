<template>
  <span class="header-avatar">
    <template v-if="picType === 'avatar'">
      <a-avatar v-if="picSrc" :size="30" :src="picSrc"></a-avatar>
      <a-avatar v-else :size="30">
        <template #icon>
          <UserOutlined />
        </template>
      </a-avatar>
    </template>
  </span>
</template>

<script>
export default {
  name: 'CustomPic',
  props: {
    picType: {
      type: String,
      default: 'avatar'
    },
    picSrc: {
      type: String,
      default: ''
    }
  }
}
</script>

<style lang="less" scoped>
.header-avatar {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
