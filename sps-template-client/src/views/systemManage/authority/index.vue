<template>
  <div class="button-box">
    <a-button @click="addAuthority('0')" type="primary"><PlusOutlined />新增角色</a-button>
  </div>
  <!-- 角色列表 -->
  <a-table
    size="small"
    bordered
    :columns="columns"
    :dataSource="tableData"
    rowKey="authorityId"
    :pagination="false">
    <template #action="{ record }">
      <a-button type="primary" @click="openDrawer(record)"><LockOutlined />设置权限</a-button>
      <a-button class="table-button" @click="addAuthority(record.authorityId)"><PlusCircleOutlined />新建子角色</a-button>
      <a-button class="table-button" type="primary"><CopyOutlined />拷贝</a-button>
      <a-button class="table-button" @click="editAuthority(record)"><EditOutlined />编辑</a-button>
      <a-popconfirm title="此操作将永久删除该角色, 是否继续?" @confirm="deleteAuth(record)">
        <a-button class="table-button" type="danger"><DeleteOutlined />删除</a-button>
      </a-popconfirm>
    </template>
  </a-table>
  <!-- 新增 / 编辑角色弹窗 -->
  <a-modal
    v-model:visible="dialogVisible"
    :title="dialogTitle"
    @ok="okDialog">
    <a-form
      ref="authorityForm"
      :model="formModel"
      :rules="rules"
      :labelCol="{ span: 4 }"
      :wrapperCol="{ span: 14 }">
      <a-form-item label="父级角色" name="parentId">
        <a-tree-select
          v-model:value="formModel.parentId"
          :disabled="dialogType === 'add'"
          :tree-data="authorityOption">
        </a-tree-select>
      </a-form-item>
      <a-form-item label="角色ID" name="authorityId">
        <a-input v-model:value="formModel.authorityId"></a-input>
      </a-form-item>
      <a-form-item label="角色名称" name="authorityName">
        <a-input v-model:value="formModel.authorityName"></a-input>
      </a-form-item>
    </a-form>
  </a-modal>
  <!-- 权限配置 -->
  <a-drawer
    v-model:visible="drawerVisible"
    title="角色配置"
    width="40%"
    destroyOnClose>
    <a-tabs type="card">
      <a-tab-pane key="menu" tab="角色菜单">
        <MenuConfig :record="actionRecord" />
      </a-tab-pane>
      <a-tab-pane key="api" tab="角色api"></a-tab-pane>
      <a-tab-pane key="resource" tab="资源权限"></a-tab-pane>
    </a-tabs>
  </a-drawer>
</template>

<script>
import { reactive, toRefs, ref, getCurrentInstance } from 'vue'
import { getAuthorityList, createAuthority, updateAuthority, deleteAuthority } from '@/api/authority'
import useInfoList from '@/mixins/infoList'
import MenuConfig from './components/menuConfig'

export default {
  name: 'AuthorityMenu',
  components: {
    MenuConfig
  },
  setup () {
    const { ctx } = getCurrentInstance()
    const state = reactive({
      columns: [
        { title: '角色id', dataIndex: 'authorityId' },
        { title: '角色名称', dataIndex: 'authorityName' },
        { title: '操作', dataIndex: 'action', width: 600, slots: { customRender: 'action' } }
      ],
      dialogVisible: false,
      dialogTitle: '',
      dialogType: '',
      formModel: {},
      rules: {
        authorityId: [
          { required: true, message: '请输入角色ID', trigger: 'blur' },
          {
            validator (rule, value) {
              if (!(/^[0-9]*[1-9][0-9]*$/).test(value)) {
                // eslint-disable-next-line prefer-promise-reject-errors
                return Promise.reject('请输入正整数')
              }
              return Promise.resolve()
            },
            trigger: 'blur'
          }
        ],
        authorityName: [
          { required: true, message: '请输入角色名', trigger: 'blur' }
        ],
        parentId: [
          { required: true, message: '请选择父级角色', trigger: 'blur' }
        ]
      },
      authorityOption: [],
      drawerVisible: false,
      actionRecord: {}
    })

    const {
      handleSizeChange,
      tableData,
      getTableData
    } = useInfoList(getAuthorityList)

    handleSizeChange(999)

    const authorityForm = ref(null)

    // 初始化表单
    const initForm = () => {
      if (authorityForm.value) {
        authorityForm.value.resetFields()
      }
      state.formModel = {
        authorityId: '',
        authorityName: ''
      }
    }

    // 递归生成角色选项
    const setAuthorityOptions = (AuthorityData, optionsData, disabled) => {
      AuthorityData && AuthorityData.map(item => {
        const itemDisabled = disabled || item.authorityId === state.formModel.authorityId
        const hasChidlren = item.children && item.children.length
        const option = {
          value: item.authorityId,
          label: item.authorityName,
          disabled: itemDisabled,
          children: hasChidlren ? [] : undefined
        }
        hasChidlren && setAuthorityOptions(item.children, option.children, itemDisabled)
        optionsData.push(option)
      })
    }

    // 初始化角色选项
    const setOptions = () => {
      state.authorityOption = [
        {
          value: '0',
          label: '根角色'
        }
      ]
      state.formModel.authorityId = String(state.formModel.authorityId)
      setAuthorityOptions(tableData.value, state.authorityOption, false)
    }

    // 增加角色
    const addAuthority = (parentId) => {
      initForm()
      state.dialogTitle = '新增角色'
      state.dialogType = 'add'
      state.formModel.parentId = parentId
      setOptions()
      state.dialogVisible = true
    }

    // 编辑角色
    const editAuthority = (record) => {
      state.dialogTitle = '编辑角色'
      state.dialogType = 'edit'
      state.formModel = { ...record }
      setOptions()
      state.dialogVisible = true
    }

    // 删除角色
    const deleteAuth = async (record) => {
      const res = await deleteAuthority(record)
      if (res.code === 0) {
        ctx.$success('删除成功!')
      }
      getTableData()
    }

    // 打开配置权限面板
    const openDrawer = (record) => {
      state.drawerVisible = true
      state.actionRecord = record
    }

    // 关闭对话框
    const closeDialog = () => {
      initForm()
      state.dialogVisible = false
    }

    // 提交表单
    const okDialog = async () => {
      try {
        await authorityForm.value.validate()
        switch (state.dialogType) {
          case 'add': {
            const res = await createAuthority(state.formModel)
            if (res.code === 0) {
              ctx.$success('添加成功!')
              getTableData()
              closeDialog()
            }
            break
          }
          case 'edit': {
            const res = await updateAuthority(state.formModel)
            if (res.code === 0) {
              ctx.$success('编辑成功!')
              getTableData()
              closeDialog()
            }
            break
          }
        }
      } catch (err) {
        console.log(err)
      }
    }

    return {
      ...toRefs(state),
      authorityForm,
      tableData,
      addAuthority,
      editAuthority,
      deleteAuth,
      openDrawer,
      okDialog
    }
  }
}
</script>
