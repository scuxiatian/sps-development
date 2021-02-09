<template>
  <div>
    <div class="search-term">
      <a-form layout="inline" :model="searchInfo">
        <a-form-item label="用户名">
          <a-input v-model:value="searchInfo.userName" placeholder="路径"></a-input>
        </a-form-item>
        <a-form-item label="昵称">
          <a-input v-model:value="searchInfo.nickName" placeholder="描述"></a-input>
        </a-form-item>
        <a-form-item label="用户角色">
          <a-tree-select
            placeholder="请选择角色"
            class="auth-select"
            allowClear
            v-model:value="searchInfo.authorityId"
            :tree-data="authOptions">
          </a-tree-select>
        </a-form-item>
        <a-form-item>
          <a-button @click="searchUser"><SearchOutlined />查询</a-button>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" @click="addUser"><PlusCircleOutlined />新增用户</a-button>
        </a-form-item>
      </a-form>
    </div>
    <a-table
      size="small"
      bordered
      :columns="columns"
      :dataSource="tableData"
      :pagination="{
        total,
        showTotal: total => `共 ${total} 条`,
        pageSizeOptions: ['10', '30', '50', '100'],
        showQuickJumper: true,
        showSizeChanger: true
      }"
      @change="tableChange"
      rowKey="ID">
      <template #headerImg="{ text }">
        <CustomPic :picSrc="text" />
      </template>
      <template #authority="{ record }">
        <a-tree-select
          class="auth-select"
          v-model:value="record.authority.authorityId"
          :tree-data="authOptions"
          @change="changeAuthority(record)">
        </a-tree-select>
      </template>
      <template #action="{ record }">
        <a-popconfirm title="此操作将永久删除用户, 是否继续?"  @confirm="deleteUserRecord(record)">
          <a-button type="danger"><DeleteOutlined />删除</a-button>
        </a-popconfirm>
      </template>
    </a-table>
    <a-modal
      v-model:visible="dialogVisible"
      width="50%"
      title="新增用户"
      class="user-dialog"
      :afterClose="initForm"
      @ok="okDialog">
      <a-form
        ref="userForm"
        :model="formModel"
        :rules="rules"
        :labelCol="{ span: 3 }"
        :wrapperCol="{ span: 21 }">
        <a-form-item label="用户名" name="userName">
          <a-input v-model:value="formModel.userName"></a-input>
        </a-form-item>
        <a-form-item label="密码" name="password">
          <a-input v-model:value="formModel.password"></a-input>
        </a-form-item>
        <a-form-item label="别名" name="nickName">
          <a-input v-model:value="formModel.nickName"></a-input>
        </a-form-item>
        <a-form-item label="头像" name="headerImg">
          <div style="display:inline-block">
            <img class="header-img-box" v-if="formModel.headerImg" :src="formModel.headerImg" />
            <div v-else class="header-img-box">从媒体库选择</div>
          </div>
        </a-form-item>
        <a-form-item label="用户角色" name="authorityId">
          <a-tree-select
            placeholder="请选择角色"
            class="auth-select"
            v-model:value="formModel.authorityId"
            :tree-data="authOptions">
          </a-tree-select>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script>
import { getUserList, setUserAuthority, register, deleteUser } from '@/api/user'
import { getAuthorityList } from '@/api/authority'
import useInfoList from '@/mixins/infoList'
import { reactive, toRefs, getCurrentInstance, ref, unref } from 'vue'
import CustomPic from '@/components/customPic'

export default {
  name: 'UserManage',
  components: {
    CustomPic
  },
  setup () {
    const { ctx } = getCurrentInstance()
    const state = reactive({
      columns: [
        { title: '头像', dataIndex: 'headerImg', width: 50, slots: { customRender: 'headerImg' } },
        { title: 'uuid', dataIndex: 'uuid', width: 250 },
        { title: '用户名', dataIndex: 'userName', width: 150 },
        { title: '昵称', dataIndex: 'nickName', width: 150 },
        { title: '用户角色', dataIndex: 'authority', width: 150, slots: { customRender: 'authority' } },
        { title: '操作', dataIndex: 'action', width: 150, slots: { customRender: 'action' } }
      ],
      searchInfo: {},
      authOptions: [],
      dialogVisible: false,
      formModel: {},
      rules: {
        userName: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 6, message: '最低6位字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入用户密码', trigger: 'blur' },
          { min: 6, message: '最低6位字符', trigger: 'blur' }
        ],
        nickName: [
          { required: true, message: '请输入用户昵称', trigger: 'blur' }
        ],
        authorityId: [
          { required: true, message: '请选择用户角色', trigger: 'blur' }
        ]
      }
    })

    const userForm = ref(null)
    const { getTableData, tableData, total, handleTableChange } = useInfoList(getUserList)
    getTableData()

    // 递归生成角色选项
    const setAuthorityOptions = (AuthorityData, optionsData) => {
      AuthorityData && AuthorityData.map(item => {
        const hasChidlren = item.children && item.children.length
        const option = {
          value: item.authorityId,
          label: item.authorityName,
          children: hasChidlren ? [] : undefined
        }
        hasChidlren && setAuthorityOptions(item.children, option.children)
        optionsData.push(option)
      })
    }

    const initAuthorityOptions = async () => {
      const res = await getAuthorityList({ page: 1, pageSize: 999 })
      state.authOptions = []
      setAuthorityOptions(res.data.list, state.authOptions)
    }

    initAuthorityOptions()

    const tableChange = (pagination) => {
      handleTableChange(pagination.current, pagination.pageSize, {
        ...state.searchInfo
      })
    }

    const searchUser = () => {
      handleTableChange(1, 10, {
        ...state.searchInfo
      })
    }

    // 修改用户角色
    const changeAuthority = async (record) => {
      const res = await setUserAuthority({
        uuid: record.uuid,
        authorityId: record.authority.authorityId
      })
      if (res.code === 0) {
        ctx.$success('角色设置成功')
      }
    }

    // 初始化表单
    const initForm = () => {
      const form = unref(userForm)
      form && form.resetFields()
      state.formModel = {
        userName: '',
        password: '',
        nickName: '',
        authorityId: ''
      }
    }

    // 新增用户
    const addUser = () => {
      initForm()
      state.dialogVisible = true
    }

    // 提交表单
    const okDialog = async () => {
      const form = unref(userForm)
      try {
        await form.validate()
        const res = await register(state.formModel)
        if (res.code === 0) {
          ctx.$success('创建成功')
          await searchUser()
          state.dialogVisible = false
        }
      } catch (err) {}
    }

    // 删除用户
    const deleteUserRecord = async (record) => {
      const res = await deleteUser({ id: record.ID })
      if (res.code === 0) {
        ctx.$success('删除成功')
        searchUser()
      }
    }

    return {
      ...toRefs(state),
      userForm,
      tableData,
      total,
      tableChange,
      searchUser,
      changeAuthority,
      initForm,
      addUser,
      okDialog,
      deleteUserRecord
    }
  }
}
</script>

<style lang="less" scoped>
.auth-select {
  width: 150px
}
.user-dialog {
  .auth-select {
    width: 200px
  }
  .header-img-box {
    width: 200px;
    height: 200px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 200px;
    cursor: pointer;
  }
}

</style>
