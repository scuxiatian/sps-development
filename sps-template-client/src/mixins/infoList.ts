import { reactive, toRefs } from 'vue'

export default function (listApi: Function) {
  const state = reactive({
    page: 1,
    total: 10,
    pageSize: 10,
    tableData: [],
    searchInfo: {}
  })

  const getTableData = async (page = state.page, pageSize = state.pageSize) => {
    const res = await listApi({ page, pageSize, ...state.searchInfo })
    if (res.code === 0) {
      state.tableData = res.data.list
      state.total = res.data.total
      state.page = res.data.page
      state.pageSize = res.data.pageSize
    }
  }

  const handleSizeChange = (val: number) => {
    state.pageSize = val
    getTableData()
  }

  return {
    ...toRefs(state),
    handleSizeChange,
    getTableData
  }
}
