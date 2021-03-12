<template>
  <div>
    <div class="search-term">
      <a-form layout="inline" :model="searchInfo">
        <a-form-item label="流程名称">
          <a-input v-model:value="searchInfo.name" placeholder="流程名称"></a-input>
        </a-form-item>
        <a-form-item label="流程标题">
          <a-input v-model:value="searchInfo.label" placeholder="流程标题"></a-input>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" @click="searchProcess"><SearchOutlined />查询</a-button>
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
      rowKey="id">
      <template #CreatedAt="{ text }">
        {{ formatTime(text) }}
      </template>
      <template #hideIcon="{ text }">
        {{ text ? '是' : '否' }}
      </template>
      <template #action="{ record }">
        <a-button type="primary" @click="useWorkflowProcess(record)"><CaretRightOutlined />启动</a-button>
        <a-button class="table-button" @click="updateWorkflowProcess(record)"><EditOutlined />变更</a-button>
        <a-button class="table-button" type="primary" @click="viewWorkflowProcess(record)"><EyeOutlined />查看</a-button>
        <a-popconfirm title="此操作将永久删除流程, 是否继续?" @confirm="deleteWorkProcess(record)">
          <a-button class="table-button" type="danger"><DeleteOutlined />删除</a-button>
        </a-popconfirm>
      </template>
    </a-table>
  </div>
</template>

<script>
import { deleteWorkflowProcess, getWorkflowProcessList } from '@/api/workflowProcess'
import useInfoList from '@/mixins/infoList'
import { getCurrentInstance, reactive, toRefs } from 'vue'
import { formatTime } from '@/utils/date'
import { useRouter } from 'vue-router'

export default {
  name: 'WorkflowProcess',
  setup () {
    const { ctx } = getCurrentInstance()
    const router = useRouter()
    const state = reactive({
      columns: [
        { title: '创建时间', dataIndex: 'CreatedAt', width: 180, slots: { customRender: 'CreatedAt' } },
        { title: '流程名称', dataIndex: 'name', width: 120 },
        // { title: '分类', dataIndex: 'category', width: 120 },
        { title: '类型', dataIndex: 'clazz', width: 120 },
        { title: '流程标题', dataIndex: 'label', width: 120 },
        { title: '是否隐藏图标', dataIndex: 'hideIcon', width: 120, slots: { customRender: 'hideIcon' } },
        { title: '详细介绍', dataIndex: 'description', width: 180 },
        { title: '操作', key: 'action', slots: { customRender: 'action' } }
      ],
      searchInfo: {}
    })

    const { tableData, getTableData, total, handleTableChange } = useInfoList(getWorkflowProcessList)
    getTableData()

    // 表格的分页 排序条件变化
    const tableChange = (pagination) => {
      handleTableChange(pagination.current, pagination.pageSize, {
        ...state.searchInfo
      })
    }

    // 条件查询
    const searchProcess = () => {
      handleTableChange(1, 10, {
        ...state.sortInfo,
        ...state.searchInfo
      })
    }

    // 启动流程
    const useWorkflowProcess = (record) => {
      router.push({
        name: 'workflowUse',
        query: {
          workflowId: record.id
        }
      })
    }

    // 查看流程
    const viewWorkflowProcess = async (record) => {
      router.push({
        name: 'workflowCreate',
        query: {
          id: record.id,
          type: 'view'
        }
      })
    }

    // 更新流程
    const updateWorkflowProcess = async (record) => {
      router.push({
        name: 'workflowCreate',
        query: {
          id: record.id,
          type: 'edit'
        }
      })
    }

    // 删除流程
    const deleteWorkProcess = async (record) => {
      const res = await deleteWorkflowProcess(record)
      if (res.code === 0) {
        ctx.$success('删除成功')
        getTableData()
      }
    }

    return {
      ...toRefs(state),
      tableData,
      formatTime,
      total,
      tableChange,
      searchProcess,
      useWorkflowProcess,
      viewWorkflowProcess,
      updateWorkflowProcess,
      deleteWorkProcess
    }
  }
}
</script>
