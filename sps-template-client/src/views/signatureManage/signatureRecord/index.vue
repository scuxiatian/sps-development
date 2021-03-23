<template>
  <div class="search-term">
    <a-form
      layout="inline"
      :model="searchInfo">
      <a-form-item label="流水号">
        <a-input-number v-model:value="searchInfo.ID" placeholder="流水号"></a-input-number>
      </a-form-item>
      <a-form-item>
        <a-button type="primary" @click="searchRecord"><SearchOutlined />查询</a-button>
      </a-form-item>
    </a-form>
  </div>
  <a-table
    size="small"
    bordered
    :columns="columns"
    :dataSource="tableData"
    :pagination="{
        current,
        total,
        showTotal: total => `共 ${total} 条`,
        pageSizeOptions: ['10', '30', '50', '100'],
        showQuickJumper: true,
        showSizeChanger: true
    }"
    @change="tableChange"
    rowKey="ID">
    <template #CreatedAt="{ text }">
      {{ formatTime(text) }}
    </template>
    <template #expandedRowRender="{ record }">
      <a-table
        size="small"
        bordered
        :columns="signatureColumns"
        :dataSource="record.signatures"
        :pagination="false"
        rowKey="ID">
        <template #url="{ text: subText }">
          <CustomPic :picSrc="subText" />
        </template>
        <template #position="{ record: subRecord }">
          ({{ subRecord.x }}, {{ subRecord.y }})
        </template>
        <template #CreatedAt="{ text: subText }">
          {{ formatTime(subText) }}
        </template>
        <template #UpdatedAt="{ text: subText }">
          {{ formatTime(subText) }}
        </template>
      </a-table>
    </template>
  </a-table>
</template>

<script>
import useInfoList from '@/mixins/infoList'
import { getSignatureRecordList } from '@/api/signature'
import { reactive, toRefs } from 'vue'
import { formatTime } from '@/utils/date'
import CustomPic from '@/components/customPic'

export default {
  name: 'SignatureRecord',
  components: {
    CustomPic
  },
  setup () {
    const state = reactive({
      columns: [
        { title: '流水号', dataIndex: 'ID' },
        { title: '创建时间', dataIndex: 'CreatedAt', slots: { customRender: 'CreatedAt' } }
      ],
      signatureColumns: [
        { title: '签章图片', dataIndex: 'signature.url', width: 100, slots: { customRender: 'url' } },
        { title: '签章名', dataIndex: 'signature.name' },
        { title: '签章时间', dataIndex: 'CreatedAt', slots: { customRender: 'CreatedAt' } },
        { title: '更新时间', dataIndex: 'UpdatedAt', slots: { customRender: 'UpdatedAt' } },
        { title: '签章位置', dataIndex: 'position', slots: { customRender: 'position' } },
        { title: '说明', dataIndex: 'description', width: 300 },
        { title: '操作者', dataIndex: 'operator.nickName' }
      ],
      searchInfo: {},
      current: 1
    })
    const { tableData, total, getTableData, handleTableChange } = useInfoList(getSignatureRecordList)
    getTableData()

    // 表格的分页 排序条件变化
    const tableChange = (pagination) => {
      state.current = pagination.current
      handleTableChange(pagination.current, pagination.pageSize, {
        ...state.searchInfo
      })
    }

    // 条件查询
    const searchRecord = () => {
      state.current = 1
      handleTableChange(1, undefined, {
        ...state.searchInfo
      })
    }

    return {
      ...toRefs(state),
      tableData,
      total,
      formatTime,
      tableChange,
      searchRecord
    }
  }
}
</script>
