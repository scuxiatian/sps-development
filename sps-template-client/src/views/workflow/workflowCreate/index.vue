<template>
  <div>
    <a-button
      size="small"
      style="float: right; margin-top: 6px; margin-right: 6px"
      @click="saveJSON"
    >
      导出JSON
    </a-button>
    <a-button
      size="small"
      style="float: right; margin-top: 6px; margin-right: 6px"
      @click="saveImg"
    >
      导出图片
    </a-button>
    <a-popconfirm title="确认流程无误并保存吗" @confirm="saveFlow">
      <a-button
        size="small"
        style="float: right; margin-top: 6px; margin-right: 6px"
        :disabled="disabled"
      >
        保存流程
      </a-button>
    </a-popconfirm>
    <WorkflowDesign
      ref="wfdRef"
      v-if="done"
      :readOnly="disabled"
      :data="data"
      :height="590"
      :users="users"
      :authorities="authorities"
      :propProcessModel="processModel" />
  </div>
</template>

<script>
import WorkflowDesign from '@/components/workflowDesign'
import { getCurrentInstance, reactive, ref, toRefs, unref } from 'vue'
import { getUserList } from '@/api/user'
import { getAuthorityList } from '@/api/authority'
import { createWorkflowProcess, findWorkflowProcess, updateWorkflowProcess } from '@/api/workflowProcess'
import { useRoute } from 'vue-router'
import _ from 'lodash'

export default {
  name: 'workflowCreate',
  components: {
    WorkflowDesign
  },
  setup () {
    const { ctx } = getCurrentInstance()
    const route = useRoute()
    const state = reactive({
      data: {
        nodes: [],
        edges: []
      },
      processModel: {
        id: '',
        name: '',
        category: '',
        clazz: 'process',
        dataObjs: [],
        signalDefs: [],
        messageDefs: []
      },
      users: [],
      authorities: [],
      disabled: false,
      done: false,
      type: null
    })

    const wfdRef = ref(null)

    const fmtAuthority = (authorityList, list) => {
      authorityList.map(item => {
        list.push({
          id: item.authorityId,
          name: item.authorityName
        })
        if (item.children) {
          fmtAuthority(item.children, list)
        }
      })
    }

    const initData = async () => {
      const userRes = await getUserList({ page: 1, pageSize: 9999999 })
      if (userRes.code === 0) {
        userRes.data.list.map(item => {
          state.users.push({ id: item.ID, name: item.nickName })
        })
      }
      const authorityRes = await getAuthorityList({ page: 1, pageSize: 9999999 })
      if (authorityRes.code === 0) {
        fmtAuthority(authorityRes.data.list, state.authorities)
      }
      if (route.query.id) {
        const res = await findWorkflowProcess({ id: route.query.id })
        state.disabled = route.query.type === 'view'
        if (res.code === 0) {
          const reworkflowProcess = res.data.reworkflowProcess
          reworkflowProcess.nodes.map(item => {
            if (item.assignValue) {
              const waitUseArr = item.assignValue.substr(1, item.assignValue.length - 2).split(',')
              if (item.assignType === 'user') {
                item.assignValue = []
                waitUseArr.map(i => {
                  item.assignValue.push(Number(i))
                })
              } else {
                item.assignValue = waitUseArr
              }
            }
          })
          state.data.nodes = reworkflowProcess.nodes
          state.data.edges = reworkflowProcess.edges
          state.processModel = _.omit(reworkflowProcess, ['nodes', 'edges'])
        }
        state.type = route.query.type
      }
      state.done = true
    }
    initData()

    const saveImg = () => {
      unref(wfdRef).saveImg()
    }

    const saveJSON = () => {
      unref(wfdRef).saveJSON()
    }

    const saveFlow = async () => {
      const obj = unref(wfdRef).saveFlow()
      const processModel = JSON.parse(JSON.stringify(unref(wfdRef).processModel))
      processModel.edges = obj.edges
      processModel.nodes = JSON.parse(JSON.stringify(obj.nodes))
      processModel.nodes.map(node => {
        if (node.assignValue) {
          node.assignValue = ',' + String(node.assignValue) + ','
        }
      })
      if (!processModel.id) {
        ctx.$error('流程id必填，点击空白处填写流程基本信息')
        return
      }
      if (route.query.type === 'edit') {
        const res = await updateWorkflowProcess(processModel)
        if (res.code === 0) {
          ctx.$success('编辑成功')
        }
      } else {
        const res = await createWorkflowProcess(processModel)
        if (res.code === 0) {
          ctx.$success('创建成功')
        }
      }
    }

    return {
      ...toRefs(state),
      wfdRef,
      saveImg,
      saveJSON,
      saveFlow
    }
  }
}
</script>
