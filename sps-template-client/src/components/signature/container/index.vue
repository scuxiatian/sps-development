<template>
  <div id="containerRef" ref="containerRef" class="signature-container" @dragover="handleDragOver" @drop="handleDrop">
    <SignatureItem
      v-for="(signature) of record.signatures"
      :key="signature.ID"
      :signature="signature"
      :disabled="disabled"
      @delete="handleDelete"
    />
  </div>
</template>

<script>
import { onMounted, reactive, ref, toRefs, unref } from 'vue'
import SignatureItem from '../item'
import { CancelSignature, getSignatureRecordById, saveSignaturePosition, useSignature } from '@/api/signature'

export default {
  name: 'SignatureContainer',
  components: {
    SignatureItem
  },
  props: {
    id: {
      type: Number,
      default: 0
    },
    disabled: {
      type: Boolean,
      default: false
    }
  },
  setup (props, { emit }) {
    const state = reactive({
      record: {
        ID: 0,
        signatures: []
      }
    })

    const containerRef = ref(null)

    const handleDragOver = (e) => {
      e.preventDefault()
    }

    const handleDrop = (e) => {
      e.preventDefault()
      if (props.disabled) return
      const container = unref(containerRef)
      const targetId = e.dataTransfer.getData('targetId')
      const target = state.record.signatures.find(item => {
        return item.ID === parseInt(targetId)
      })
      const x = target.x + e.pageX - e.dataTransfer.getData('pageX')
      const y = target.y + e.pageY - e.dataTransfer.getData('pageY')
      const itemWidth = e.dataTransfer.getData('width')
      const itemHeight = e.dataTransfer.getData('height')

      // 标签不超出父容器范围
      target.x = x < 0 ? 0 : (x > container.offsetWidth - itemWidth ? container.offsetWidth - itemWidth : x)
      target.y = y < 0 ? 0 : (y > container.offsetHeight - itemHeight ? container.offsetHeight - itemHeight : y)
    }

    onMounted(async () => {
      if (props.id) {
        const res = await getSignatureRecordById({ id: props.id })
        if (res.code === 0) {
          state.record = res.data
        }
      } else {
        state.record = {
          signatures: []
        }
      }
    })

    // 新增签章
    const addSignature = async (signature) => {
      const res = await useSignature({ recordId: state.record.ID, signatureId: signature.ID, description: signature.description })
      if (res.code === 0) {
        const { record } = res.data
        state.record = {
          ...record
        }
      }
      const container = unref(containerRef)
      container.focus()
      container.dispatchEvent(new MouseEvent('click'))
    }

    // 保存签章位置
    const saveSignature = async () => {
      const res = await saveSignaturePosition(state.record)
      if (res.code === 0) {
        emit('saveSuccess', state.record)
      }
    }

    // 取消签章
    const handleDelete = async (signature) => {
      const res = await CancelSignature(signature)
      if (res.code === 0) {
        state.record = res.data
        emit('cancelSignature', signature)
      }
    }

    return {
      ...toRefs(state),
      containerRef,
      handleDrop,
      handleDragOver,
      addSignature,
      saveSignature,
      handleDelete
    }
  }
}
</script>

<style lang="less" scoped>
.signature-container {
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;
  position: absolute;
}
</style>
