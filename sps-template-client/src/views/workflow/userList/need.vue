<template>
  <div>
    <a-table
      size="small"
      bordered
      :columns="columns"
      :dataSource="tableData"
      rowKey="ID">
      <template #CreatedAt="{ text }">
        {{ formatTime(text) }}
      </template>
      <template #isActive="{ text }">
        {{ text ? '进行中' : '已结束' }}
      </template>
      <template #action="{ record }">
        <a-button type="primary" @click="handle(record)"><EyeOutlined />处理</a-button>
      </template>
    </a-table>
  </div>
</template>

<script>
import { getMyNeed } from '@/api/workflowProcess'
import { formatTime } from '@/utils/date'
import { reactive, toRefs } from 'vue'
import { useRouter } from 'vue-router'

export default {
  name: 'MyNeed',
  setup () {
    const router = useRouter()
    const state = reactive({
      tableData: [],
      columns: [
        { title: 'id', dataIndex: 'ID', width: 60 },
        { title: '流程名称', dataIndex: 'workflowProcess.label', width: 150 },
        { title: '发起人', dataIndex: 'promoter.nickName', width: 120 },
        { title: '节点日期', dataIndex: 'CreatedAt', width: 180, slots: { customRender: 'CreatedAt' } },
        { title: '业务代码', dataIndex: 'businessType', width: 120 },
        { title: '当前节点', dataIndex: 'workflowNode.label', width: 120 },
        { title: '流程状态', dataIndex: 'isActive', width: 120, slots: { customRender: 'isActive' } },
        { title: '详细介绍', dataIndex: 'workflowProcess.description', width: 200 },
        { title: '操作', dataIndex: 'action', width: 200, slots: { customRender: 'action' } }
      ]
    })

    const init = async () => {
      const res = await getMyNeed()
      if (res.code === 0) {
        state.tableData = res.data.wfms
      }
    }

    init()

    const handle = (record) => {
      router.push({
        name: 'workflowUse',
        query: {
          workflowMoveID: record.ID
        }
      })
    }

    return {
      ...toRefs(state),
      formatTime,
      handle
    }
  }
}
</script>
