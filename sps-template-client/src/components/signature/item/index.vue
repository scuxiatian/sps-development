<template>
  <div ref="itemRef" :style="style" :draggable="!disabled" @dragstart="handleDragStart" @mouseenter="isHover=true" @mouseleave="isHover=false">
    <div class="image-container">
      <a-image :preview="false" :src="convertPath(signature.signature.url)" />
      <div class="signature-time">{{ formatTime(signature.CreatedAt, 'YYYY 年 MM 月 DD 日') }}</div>
      <div v-show="!disabled && isHover" class="hover-content">
        <DeleteOutlined @click="deleteSignature" class="hover-icon" />
        <!-- <FieldTimeOutlined class="hover-icon" /> -->
      </div>
    </div>
  </div>
</template>

<script>
import { computed, reactive, ref, toRefs, unref } from 'vue'
import { convertPath } from '@/utils/download'
import { formatTime } from '@/utils/date'

export default {
  name: 'SignatureItem',
  props: {
    signature: {
      type: Object
    },
    disabled: {
      type: Boolean,
      default: false
    }
  },
  setup (props, { emit }) {
    const itemRef = ref(null)

    const state = reactive({
      isHover: false
    })

    const handleDragStart = (e) => {
      if (props.disabled) return
      const item = unref(itemRef)
      e.dataTransfer.setData('targetId', props.signature.ID)
      e.dataTransfer.setData('pageX', e.pageX)
      e.dataTransfer.setData('pageY', e.pageY)
      e.dataTransfer.setData('width', item.offsetWidth)
      e.dataTransfer.setData('height', item.offsetHeight)
    }

    const style = computed(() => {
      return {
        left: props.signature.x + 'px',
        top: props.signature.y + 'px',
        position: 'absolute',
        cursor: props.disabled ? 'not-allowed' : 'move'
      }
    })

    const deleteSignature = () => {
      emit('delete', props.signature)
    }

    return {
      ...toRefs(state),
      itemRef,
      style,
      handleDragStart,
      deleteSignature,
      convertPath,
      formatTime
    }
  }
}
</script>

<style lang="less" scoped>
.image-container {
  position: relative;
  .signature-time {
    text-align: center;
    font-size: 18px;
  }
  .hover-content {
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(200, 200, 200, 0.5);
    position: absolute;
    .hover-icon {
      margin: 5px 5px;
      font-size: 25px;
      cursor: pointer;
    }
  }
}
</style>
