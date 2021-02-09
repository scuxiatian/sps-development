<template>
  <div>
    <div class="button-box">
      <a-button type="primary" @click="relation"><SaveOutlined />保存</a-button>
    </div>
    <a-tree
      ref="menuTree"
      :tree-data="menuTreeData"
      :replaceFields="{ key: 'ID' }"
      v-model:checkedKeys="checkedKeys"
      checkable
      checkStrictly
      @check="checkChange">
      <template #title="scope">
        {{ scope.meta.title }}
      </template>
    </a-tree>
  </div>
</template>

<script>
import { reactive, ref, toRefs, getCurrentInstance } from 'vue'
import { getBaseMenuTree, getMenuAuthority, addMenuAuthority } from '@/api/menu'

export default {
  name: 'MenuConfig',
  props: {
    record: {
      type: Object,
      default () { return {} }
    }
  },
  setup (props) {
    const { ctx } = getCurrentInstance()
    const state = reactive({
      menuTreeData: [],
      checkedKeys: [],
      checkedNodes: [],
      needConfirm: false
    })

    const menuTree = ref(null)

    // 初始化菜单树
    const initialTree = async () => {
      // 获取所有菜单树
      let res = await getBaseMenuTree()
      state.menuTreeData = res.data.menus

      res = await getMenuAuthority({ authorityId: props.record.authorityId })
      state.checkedKeys = res.data.menus.map(item => Number(item.menuId))
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
      const res = await addMenuAuthority({
        menus: state.checkedNodes,
        authorityId: props.record.authorityId
      })
      if (res.code === 0) {
        ctx.$success('菜单设置成功!')
        state.needConfirm = false
      }
    }

    return {
      ...toRefs(state),
      menuTree,
      checkChange,
      relation
    }
  }
}
</script>
