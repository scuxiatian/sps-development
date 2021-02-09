<template>
  <div>
    <div class="button-box">
      <a-button type="primary" @click="relation"><SaveOutlined />保存</a-button>
    </div>
    <a-tree
      ref="apiTree"
      :tree-data="apiTreeData"
      :replaceFields="{ key: 'onlyId' }"
      v-model:checkedKeys="checkedKeys"
      checkable
      @check="checkChange">
      <template #title="scope">
        {{ scope.description }}
      </template>
    </a-tree>
  </div>
</template>

<script>
import { reactive, ref, toRefs, getCurrentInstance } from 'vue'
import { getAllApis } from '@/api/api'
import { getPolicyPathByAuthorityId, updateCasbin } from '@/api/casbin'

export default {
  name: 'ApiConfig',
  props: {
    record: {
      type: Object,
      default () { return {} }
    }
  },
  setup (props) {
    const { ctx } = getCurrentInstance()
    const state = reactive({
      apiTreeData: [],
      checkedKeys: [],
      checkedNodes: [],
      needConfirm: false
    })

    const apiTree = ref(null)

    // 创建api树方法
    const buildApiTree = (apis) => {
      const apiObj = {}
      apis && apis.map(item => {
        item.onlyId = 'p:' + item.path + 'm:' + item.method
        if (Object.prototype.hasOwnProperty.call(apiObj, item.apiGroup)) {
          apiObj[item.apiGroup].push(item)
        } else {
          Object.assign(apiObj, { [item.apiGroup]: [item] })
        }
      })
      const apiTree = []
      for (const key in apiObj) {
        const treeNode = {
          ID: key,
          description: key + '组',
          children: apiObj[key]
        }
        apiTree.push(treeNode)
      }
      return apiTree
    }

    // 初始化菜单树
    const initialTree = async () => {
      // 获取所有菜单树
      let res = await getAllApis()
      state.apiTreeData = buildApiTree(res.data.apis)

      res = await getPolicyPathByAuthorityId({ authorityId: props.record.authorityId })
      state.checkedKeys = res.data.paths.map(item => 'p:' + item.path + 'm:' + item.method)
    }
    initialTree()

    // 选中状态变化
    const checkChange = (_, e) => {
      state.needConfirm = true
      state.checkedNodes = e.checkedNodes.map(item => item.props.dataRef)
    }

    // 关联角色菜单
    const relation = async () => {
      if (!state.needConfirm) {
        ctx.$warning('未发生更改')
        return
      }
      const casbinInfos = []
      state.checkedNodes && state.checkedNodes
        .filter(item => item.path)
        .forEach(item => {
          casbinInfos.push({
            path: item.path,
            method: item.method
          })
        })
      const res = await updateCasbin({
        authorityId: props.record.authorityId,
        casbinInfos
      })
      if (res.code === 0) {
        ctx.$success('api设置成功!')
        state.needConfirm = false
      }
    }

    return {
      ...toRefs(state),
      apiTree,
      checkChange,
      relation
    }
  }
}
</script>
