<template>
  <div class="search-term">
    <a-button type="primary" @click="addSignature"><PlusOutlined />新增签章</a-button>
  </div>
  <a-table
    size="small"
    bordered
    :columns="columns"
    :dataSource="tableData"
    :pagination="false"
    rowKey="ID">
    <template #url="{ text }">
      <CustomPic :picSrc="text" />
    </template>
    <template #action="{ record }">
      <a-button type="primary" @click="editSignature(record)"><EditOutlined />编辑</a-button>
      <a-button class="table-button" @click="changePassword(record)"><LockOutlined />修改密码</a-button>
      <a-popconfirm title="此操作将永久删除签章, 是否继续?" @confirm="deleteSignatureRecord(record)">
        <a-button class="table-button" type="danger"><DeleteOutlined />删除</a-button>
      </a-popconfirm>
    </template>
  </a-table>
  <!-- 编辑签章信息 -->
  <a-modal
    v-model:visible="dialogVisible"
    width="50%"
    title="新增签章"
    :afterClose="initForm"
    @ok="okDialog">
    <a-form
      ref="signatureForm"
      :model="formModel"
      :rules="rules"
      :labelCol="{ span: 3 }"
      :wrapperCol="{ span: 21 }">
      <a-form-item label="签章名" name="name">
        <a-input v-model:value="formModel.name"></a-input>
      </a-form-item>
      <a-form-item v-if="!isEdit" label="密码" name="password">
        <a-input v-model:value="formModel.password"></a-input>
      </a-form-item>
      <a-form-item label="签章图片" name="url">
        <div style="display:inline-block" @click="openUrlChange">
          <img class="sd-img-box" v-if="formModel.url" :src="convertPath(formModel.url)" />
          <div v-else class="sd-img-box">从媒体库选择</div>
        </div>
      </a-form-item>
      <a-form-item label="描述" name="description">
        <a-textarea v-model:value="formModel.description"></a-textarea>
      </a-form-item>
    </a-form>
  </a-modal>
  <ChooseImg ref="chooseImg" :target="formModel" targetKey="url" />
  <!-- 修改签章密码 -->
  <a-modal
    v-model:visible="passwordDialogVisible"
    width="50%"
    title="修改密码"
    @ok="okPasswordDialog">
    <a-form
      ref="passwordForm"
      :model="passwordFormModel"
      :rules="passwordRules"
      :labelCol="{ span: 3 }"
      :wrapperCol="{ span: 21 }">
      <a-form-item label="原密码" name="password">
        <a-input v-model:value="passwordFormModel.password"></a-input>
      </a-form-item>
      <a-form-item label="新密码" name="newPassword">
        <a-input v-model:value="passwordFormModel.newPassword"></a-input>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script>
import { reactive, toRefs, ref, unref, getCurrentInstance } from 'vue'
import { convertPath } from '@/utils/download'
import { createSignature, getSignatureById, getSignatureList, changeSignaturePassword, updateSignature, deleteSignature } from '@/api/signature'
import CustomPic from '@/components/customPic'
import ChooseImg from '@/components/chooseImg'
import useInfoList from '@/mixins/infoList'
import _ from 'lodash'

export default {
  name: 'PublicSignature',
  components: {
    CustomPic,
    ChooseImg
  },
  setup () {
    const { ctx } = getCurrentInstance()
    const state = reactive({
      columns: [
        { title: '签章图片', dataIndex: 'url', width: 50, slots: { customRender: 'url' } },
        { title: 'id', dataIndex: 'ID', width: 50 },
        { title: '签章名', dataIndex: 'name', width: 150 },
        { title: '描述', dataIndex: 'description', width: 300 },
        { title: '操作', dataIndex: 'action', width: 250, slots: { customRender: 'action' } }
      ],
      formModel: {},
      rules: {
        name: [
          { required: true, message: '请输入签章名', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入签章密码', trigger: 'blur' }
        ],
        url: [
          { required: true, message: '请选择签章图片', trigger: 'change' }
        ]
      },
      dialogVisible: false,
      isEdit: false,
      passwordFormModel: {},
      passwordRules: {
        password: [
          { required: true, message: '请输入原密码', trigger: 'blur' }
        ],
        newPassword: [
          { required: true, message: '请输入新密码', trigger: 'blur' }
        ]
      },
      passwordDialogVisible: false
    })
    const { tableData, handleSizeChange, getTableData } = useInfoList(getSignatureList)
    handleSizeChange(999)

    const signatureForm = ref(null)
    const passwordForm = ref(null)

    const initForm = () => {
      state.formModel = {}
    }

    const addSignature = () => {
      state.isEdit = false
      state.dialogVisible = true
    }

    const editSignature = async (record) => {
      const res = await getSignatureById({ id: record.ID })
      if (res.code === 0) {
        state.formModel = _.pick(res.data, ['ID', 'name', 'url', 'description'])
        state.isEdit = true
        state.dialogVisible = true
      }
    }

    const changePassword = async (record) => {
      state.passwordFormModel = {
        id: record.ID
      }
      state.passwordDialogVisible = true
    }

    const deleteSignatureRecord = async (record) => {
      const res = await deleteSignature({ id: record.ID })
      if (res.code === 0) {
        ctx.$success('删除成功')
        getTableData()
      }
    }

    const chooseImg = ref(null)
    const openUrlChange = () => {
      const drawer = unref(chooseImg)
      drawer.openDrawer()
    }

    // 提交表单
    const okDialog = async () => {
      const form = unref(signatureForm)
      try {
        await form.validate()
        if (state.isEdit) {
          const res = await updateSignature(state.formModel)
          if (res.code === 0) {
            ctx.$success('更新成功')
            state.dialogVisible = false
            getTableData()
          }
        } else {
          const res = await createSignature({
            ...state.formModel,
            isPublic: true
          })
          if (res.code === 0) {
            ctx.$success('创建成功')
            state.dialogVisible = false
            getTableData()
          }
        }
      } catch (err) {}
    }

    // 提交密码表单
    const okPasswordDialog = async () => {
      const form = unref(passwordForm)
      try {
        await form.validate()
        const res = await changeSignaturePassword(state.passwordFormModel)
        if (res.code === 0) {
          ctx.$success('修改密码成功')
          state.passwordDialogVisible = false
        }
      } catch (err) {}
    }

    return {
      ...toRefs(state),
      signatureForm,
      passwordForm,
      chooseImg,
      initForm,
      addSignature,
      editSignature,
      deleteSignatureRecord,
      changePassword,
      convertPath,
      openUrlChange,
      okDialog,
      okPasswordDialog,
      tableData
    }
  }
}
</script>
