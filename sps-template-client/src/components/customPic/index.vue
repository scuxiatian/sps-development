<template>
  <span class="header-avatar">
    <template v-if="picType === 'avatar'">
      <a-avatar v-if="picSrc" :size="30" :src="file"></a-avatar>
      <a-avatar v-else :size="30">
        <template #icon>
          <UserOutlined />
        </template>
      </a-avatar>
    </template>
    <template v-else-if="picType === 'file'">
      <a-image :src="file"></a-image>
    </template>
  </span>
</template>

<script>
import { computed } from 'vue'
const path = process.env.VUE_APP_BASE_API

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
  },
  setup (props) {
    const file = computed(() => {
      const { picSrc } = props
      if (picSrc && picSrc.slice(0, 4) !== 'http') {
        return path + '/' + picSrc
      }
      return picSrc
    })

    return {
      file
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
