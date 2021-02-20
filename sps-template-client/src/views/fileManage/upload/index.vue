<template>
  <div>
    <a-row class="button-box">
      <a-col :span="12">
        <a-upload
          :customRequest="handleUpload"
          :showUploadList="false">
          <a-button type="primary">
            <UploadOutlined />点击上传
          </a-button>
          <div>只能上传jpg/png文件，且不超过2MB</div>
        </a-upload>
      </a-col>
      <a-col :span="12"></a-col>
    </a-row>

    <a-table
      size="small"
      borderd
      :columns="columns"
      :dataSource="tableData">
      <template #url="{ text }">
        <CustomPic picType="file" :picSrc="text" />
      </template>
      <template #UpdatedAt="{ text }">
        {{ formatTime(text, 'yyyy-MM-DD hh:mm:ss') }}
      </template>
      <template #action="{ record }">
        <a-button type="link" @click="downloadImage(record.url, record.name)">下载</a-button>
        <a-popconfirm title="此操作将永久文件, 是否继续?" @confirm="deleteFileRecord(record)">
          <a-button type="link">删除</a-button>
        </a-popconfirm>
      </template>
    </a-table>
  </div>
</template>

<script>
import { useStore } from 'vuex'
import { deleteFile, getFileList, uploadFile } from '@/api/fileUploadAndDownload'
import { getCurrentInstance, reactive, toRefs } from 'vue'
import useInfoList from '@/mixins/infoList'
import CustomPic from '@/components/customPic'
import { formatTime } from '@/utils/date'
import { downloadImage } from '@/utils/download'
const path = process.env.VUE_APP_BASE_API

export default {
  name: 'UploadAndDownload',
  components: {
    CustomPic
  },
  setup () {
    const { ctx } = getCurrentInstance()
    const store = useStore()
    const state = reactive({})

    const columns = [
      { title: '预览', dataIndex: 'url', width: 100, slots: { customRender: 'url' } },
      { title: '日期', dataIndex: 'UpdatedAt', width: 180, slots: { customRender: 'UpdatedAt' } },
      { title: '文件名', dataIndex: 'name', width: 180 },
      { title: '链接', dataIndex: 'url', width: 300 },
      { title: '标签', dataIndex: 'tag', width: 100 },
      { title: '操作', dataIndex: 'action', width: 160, slots: { customRender: 'action' } }
    ]

    const checkFile = (file) => {
      const isImage = /image/.test(file.type)
      const isLt2M = file.size / 1024 / 1024 / 2
      if (!isImage) {
        ctx.$error('上传头像图片只能是图片格式!')
      }
      if (!isLt2M) {
        ctx.$error('上传头像图片大小不能超过 2MB!')
      }
      return isImage && isLt2M
    }

    const { tableData, getTableData } = useInfoList(getFileList)
    getTableData()

    const handleUpload = async (data) => {
      if (!checkFile(data.file)) return
      const formData = new FormData()
      formData.append('file', data.file)
      const res = await uploadFile(formData, store.getters['user/token'])
      if (res.code === 0) {
        ctx.$success('上传成功')
        getTableData()
      }
    }

    const deleteFileRecord = async (record) => {
      const res = await deleteFile(record)
      if (res.code === 0) {
        ctx.$success('删除成功')
        getTableData()
      }
    }

    return {
      ...toRefs(state),
      path,
      columns,
      handleUpload,
      tableData,
      formatTime,
      downloadImage,
      deleteFileRecord
    }
  }
}
</script>
