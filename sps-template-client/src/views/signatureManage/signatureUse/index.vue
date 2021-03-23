<template>
  <div class="search-term">
    <a-switch v-model:checked="disabled"></a-switch> 禁用编辑
    <a-button class="form-button" type="primary" @click="openDrawer"><AuditOutlined />签章</a-button>
    <a-button class="form-button" @click="saveSignature"><SaveOutlined />保存签章</a-button>
  </div>
  <SignatureChoose ref="signatureChooseRef" @confirm="handleSignature" />
  <div class="audit-area">
    <div class="audit-text">
      <p>测试印章</p>
      <p>当前签章流水号：20</p>
      <p>签章</p>
    </div>
    <SignatureContainer :id="20" ref="signatureContainerRef" @saveSuccess="handleSaveSuccess" :disabled="disabled" />
  </div>
</template>

<script>
import SignatureContainer from '@/components/signature/container'
import SignatureChoose from '@/components/signature/choose'
import { getCurrentInstance, ref, unref } from 'vue'

export default {
  name: 'SignatureUse',
  components: {
    SignatureContainer,
    SignatureChoose
  },
  setup () {
    const { ctx } = getCurrentInstance()
    const signatureChooseRef = ref(null)
    const signatureContainerRef = ref(null)
    const disabled = ref(false)

    const openDrawer = () => {
      const signatureChoose = unref(signatureChooseRef)
      signatureChoose.openDrawer()
    }

    const handleSignature = (signature) => {
      const signatureContainer = unref(signatureContainerRef)
      signatureContainer.addSignature(signature)
    }

    const saveSignature = () => {
      const signatureContainer = unref(signatureContainerRef)
      signatureContainer.saveSignature()
    }

    const handleSaveSuccess = () => {
      ctx.$success('保存签章位置成功')
    }

    return {
      signatureChooseRef,
      signatureContainerRef,
      openDrawer,
      handleSignature,
      saveSignature,
      handleSaveSuccess,
      disabled
    }
  }
}
</script>

<style lang="less" scoped>
.audit-area {
  width: 100%;
  height: 600px;
  border: 1px dashed #aaa;
  position: relative;
  .audit-text {
    font-size: 20px;
    text-align: center;
    margin: 20px
  }
}
</style>
