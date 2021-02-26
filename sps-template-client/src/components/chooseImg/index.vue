<template>
  <a-drawer title="媒体库" v-model:visible="drawerVisible" width="30%">
    <div class="image-list">
      <a-image
        class="image-item"
        v-for="item in picList"
        :key="item.ID"
        :src="convertPath(item.url)"
        :preview="false"
        @click="chooseImg(item.url)">
      </a-image>
    </div>
  </a-drawer>
</template>

<script>
import { reactive, toRefs } from 'vue'
import { getFileList } from '@/api/fileUploadAndDownload'
import { convertPath } from '@/utils/download'

export default {
  name: 'ChooseImg',
  props: {
    target: [Object],
    targetKey: [String]
  },
  setup (props) {
    const state = reactive({
      drawerVisible: false,
      picList: []
    })

    const openDrawer = async () => {
      const res = await getFileList({ page: 1, pageSize: 999 })
      state.picList = res.data.list
      state.drawerVisible = true
    }

    // 选中图片
    const chooseImg = (url) => {
      const { target, targetKey } = props
      if (target && targetKey) {
        target[targetKey] = url
      }
      state.drawerVisible = false
    }

    return {
      ...toRefs(state),
      openDrawer,
      convertPath,
      chooseImg
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
