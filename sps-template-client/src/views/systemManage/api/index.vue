<template>
  <div>
    <div class="search-term">
      <a-form layout="inline" :model="searchInfo">
        <a-form-item label="路径">
          <a-input v-model:value="searchInfo.path" placeholder="路径"></a-input>
        </a-form-item>
        <a-form-item label="描述">
          <a-input v-model:value="searchInfo.description" placeholder="描述"></a-input>
        </a-form-item>
        <a-form-item label="api组">
          <a-input v-model:value="searchInfo.apiGroup" placeholder="api组"></a-input>
        </a-form-item>
        <a-form-item label="请求">
          <a-select class="search-select" v-model:value="searchInfo.method" allowClear placeholder="请选择">
            <a-select-option
              v-for="item of methodOptions"
              :key="item.value"
              :value="item.value">
              {{ item.label }} ({{ item.value }})
            </a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item>
          <a-button @click="searchApi"><SearchOutlined />查询</a-button>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" @click="openDialog('add')"><PlusCircleOutlined />新增api</a-button>
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
      <template #method="{ text }">
        {{ text }}
        <a-tag :color="tagTypeFiletr(text)">{{ methodFilter(text) }}</a-tag>
      </template>
      <template #action="{ record }">
        <a-button type="primary" @click="editApi(record)"><EditOutlined />编辑</a-button>
        <a-popconfirm title="此操作将永久删除所有角色下该api, 是否继续?" @confirm="deleteApiRecord(record)">
          <a-button class="table-button" type="danger"><DeleteOutlined />删除</a-button>
        </a-popconfirm>
      </template>
    </a-table>
    <a-modal
      v-model:visible="dialogVisible"
      :title="dialogTitle"
      width="50%"
      :afterClose="initForm"
      @ok="okDialog">
      <a-form
        ref="apiForm"
        :model="formModel"
        :rules="rules"
        :labelCol="{ span: 4 }"
        :wrapperCol="{ span: 16 }">
        <a-row>
          <a-col :span="12">
            <a-form-item label="路径" name="path">
              <a-input v-model:value="formModel.path" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="请求" name="method">
              <a-select v-model:value="formModel.method" placeholder="请选择">
                <a-select-option
                  v-for="item of methodOptions"
                  :key="item.value"
                  :value="item.value">
                  {{ item.label }} ({{ item.value }})
                </a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="api分组" name="apiGroup">
              <a-input v-model:value="formModel.apiGroup" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="api简介" name="description">
              <a-input v-model:value="formModel.description" />
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
    </a-modal>
  </div>
</template>

<script>
import { getApiList, getApiById, createApi, updateApi, deleteApi } from '@/api/api'
import { reactive, ref, toRefs, unref, getCurrentInstance } from 'vue'
import useInfoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'

export default {
  name: 'ApiManage',
  setup () {
    const { ctx } = getCurrentInstance()
    const state = reactive({
      columns: [
        { title: 'id', dataIndex: 'ID', sorter: true, width: 60 },
        { title: 'api路径', dataIndex: 'path', sorter: true, width: 200 },
        { title: 'api分组', dataIndex: 'apiGroup', sorter: true, width: 150 },
        { title: 'api简介', dataIndex: 'description', sorter: true, width: 150 },
        { title: '请求', dataIndex: 'method', sorter: true, width: 100, slots: { customRender: 'method' } },
        { title: '操作', dataIndex: 'action', width: 200, slots: { customRender: 'action' } }
      ],
      searchInfo: {},
      sortInfo: {},
      dialogVisible: false,
      dialogTitle: '',
      dialogType: '',
      formModel: {},
      rules: {
        path: [{ required: true, message: '请输入api路径', trigger: 'blur' }],
        apiGroup: [{ required: true, message: '请输入组名称', trigger: 'blur' }],
        method: [{ required: true, message: '请选择请求方式', trigger: 'blur' }],
        description: [{ required: true, message: '请输入api介绍', trigger: 'blur' }]
      }
    })

    const apiForm = ref(null)

    const { getTableData, handleTableChange, tableData, total } = useInfoList(getApiList)
    getTableData()

    const methodOptions = [
      {
        value: 'POST',
        label: '创建',
        type: 'success'
      },
      {
        value: 'GET',
        label: '查看',
        type: 'processing'
      },
      {
        value: 'PUT',
        label: '更新',
        type: 'warning'
      },
      {
        value: 'DELETE',
        label: '删除',
        type: 'error'
      }
    ]

    const methodFilter = (value) => {
      const target = methodOptions.filter(item => item.value === value)[0]
      return target && `${target.label}`
    }

    const tagTypeFiletr = (value) => {
      const target = methodOptions.filter(item => item.value === value)[0]
      return target && `${target.type}`
    }

    // 表格的分页 排序条件变化
    const tableChange = (pagination, _, sort) => {
      state.sortInfo = {
        orderKey: toSQLLine(sort.field),
        desc: sort.order === 'descend'
      }
      handleTableChange(pagination.current, pagination.pageSize, {
        ...state.sortInfo,
        ...state.searchInfo
      })
    }

    // 条件查询
    const searchApi = () => {
      handleTableChange(1, 10, {
        ...state.sortInfo,
        ...state.searchInfo
      })
    }

    // 打开对话框
    const openDialog = (type) => {
      state.dialogTitle = type === 'add' ? '新增Api' : '编辑Api'
      state.dialogType = type
      state.dialogVisible = true
    }

    // 编辑Api
    const editApi = async (record) => {
      const res = await getApiById({ id: record.ID })
      state.formModel = res.data.api
      openDialog('edit')
    }

    // 删除Api
    const deleteApiRecord = async (record) => {
      const res = await deleteApi(record)
      if (res.code === 0) {
        ctx.$success('删除成功!')
        searchApi()
      }
    }

    // 初始化表单
    const initForm = () => {
      const form = unref(apiForm)
      form.resetFields()
      state.formModel = {}
    }

    // 提交表单
    const okDialog = async () => {
      try {
        const form = unref(apiForm)
        await form.validate()
        const isEdit = state.dialogType === 'edit'
        const handler = isEdit ? updateApi : createApi
        const message = `${isEdit ? '编辑' : '添加'}成功`
        const res = await handler(state.formModel)
        if (res.code === 0) {
          ctx.$success(message)
          searchApi()
          state.dialogVisible = false
        }
      } catch (err) {}
    }

    return {
      ...toRefs(state),
      apiForm,
      tableData,
      total,
      methodOptions,
      methodFilter,
      tagTypeFiletr,
      tableChange,
      searchApi,
      editApi,
      deleteApiRecord,
      openDialog,
      initForm,
      okDialog
    }
  }
}
</script>

<style scoped>
.search-select {
  width: 180px;
}
</style>
