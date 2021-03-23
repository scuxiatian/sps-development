<template>
  <a-drawer title="签章库" v-model:visible="drawerVisible" width="30%">
    <a-form
      ref="signatureFormRef"
      :model="formModel"
      :rules="rules"
      :labelCol="{ span: 4 }"
      :wrapperCol="{ span: 12 }">
      <a-form-item label="签章" name="name">
        <a-input v-model:value="formModel.name" disabled></a-input>
      </a-form-item>
      <a-form-item label="密码" name="password">
        <a-input-password v-model:value="formModel.password"></a-input-password>
      </a-form-item>
      <a-form-item label="说明" name="description">
        <a-input v-model:value="formModel.description"></a-input>
      </a-form-item>
      <a-form-item>
        <a-button type="primary" @click="useSignature">使用签章</a-button>
      </a-form-item>
    </a-form>
    <div class="image-list">
      <a-image
        class="image-item"
        v-for="item in tableData"
        :key="item.ID"
        :src="convertPath(item.url)"
        :alt="item.name"
        :preview="false"
        @click="chooseSignature(item)">
      </a-image>
    </div>
  </a-drawer>
</template>

<script>
import { reactive, ref, toRefs, unref } from 'vue'
import useInfoList from '@/mixins/infoList'
import { getSignatureList, validateSignature } from '@/api/signature'
import { convertPath } from '@/utils/download'
import lodash from 'lodash'

export default {
  name: 'SignatureChoose',
  setup (_, { emit }) {
    const state = reactive({
      drawerVisible: false,
      formModel: {}
    })

    const rules = {
      name: { required: true, message: '请选择签章', trigger: 'blur' },
      password: { required: true, message: '请输入密码', trigger: 'change' }
    }

    const { tableData, getTableData } = useInfoList(getSignatureList)

    // 打开选择器
    const openDrawer = async () => {
      await getTableData()
      state.formModel = {}
      state.drawerVisible = true
    }

    // 关闭选择器
    const closeDrawer = () => {
      state.drawerVisible = false
    }

    // 选择签章
    const chooseSignature = (item) => {
      state.formModel = lodash.pick(item, ['ID', 'name', 'url'])
    }

    const signatureFormRef = ref(null)
    const useSignature = async () => {
      const signatureForm = unref(signatureFormRef)
      try {
        await signatureForm.validate()
        const res = await validateSignature(state.formModel)
        if (res.code === 0) {
          closeDrawer()
          emit('confirm', { ...state.formModel })
        }
      } catch (err) {
        console.log(err)
      }
    }

    return {
      ...toRefs(state),
      rules,
      openDrawer,
      closeDrawer,
      tableData,
      convertPath,
      chooseSignature,
      useSignature,
      signatureFormRef
    }
  }
}
</script>

<style lang="less" scoped>
.image-list {
  display: flex;
  justify-content: space-around;
  flex-wrap: wrap;
  padding-top: 40px;
  /deep/ .image-item{
    width: 180px;
    height: 180px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 180px;
    cursor: pointer;
  }
}
</style>
