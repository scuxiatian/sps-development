<script lang="tsx">
import { useRoute } from 'vue-router'
import { findWorkflowStep, getWorkflowMoveByID } from '@/api/workflowProcess'
import { computed, defineComponent, markRaw, reactive, toRefs } from 'vue'

export default defineComponent({
  async setup () {
    const route = useRoute()
    const state = reactive({
      workflow: {
        view: ''
      },
      node: {
        view: ''
      },
      business: null,
      moves: [] as any[],
      move: {}
    })
    let componentObj: any

    const createDone = async () => {
      let path = ''
      if (state.node.view) {
        path = state.node.view
      } else {
        path = state.workflow.view
      }
      componentObj = (await (() => import('@/views/' + path))()).default
    }

    const init = async () => {
      const workflowId = route.query.workflowId as string
      const workflowMoveID = route.query.workflowMoveID as string
      if (workflowId) {
        const res = await findWorkflowStep({ id: workflowId }) as any
        if (res.code === 0) {
          state.workflow = res.data.workflow
          state.node = res.data.workflow.nodes[0]
          await createDone()
        }
      } else if (workflowMoveID) {
        const res = await getWorkflowMoveByID({ id: workflowMoveID }) as any
        if (res.code === 0) {
          state.business = res.data.business
          state.workflow = res.data.move.workflowProcess
          state.node = res.data.move.workflowNode
          state.move = res.data.move
          state.moves = res.data.moves
          await createDone()
        }
      }
    }

    await init()

    const flowComponent = markRaw(componentObj)

    const processStatus = computed(() => {
      const node = state.moves[state.moves.length - 1]
      if (node && node.workflowNode.clazz === 'end') {
        if (node.workflowNode.success) {
          return 'finish'
        }
        return 'error'
      }
      return 'process'
    })

    return {
      ...toRefs(state),
      flowComponent,
      processStatus
    }
  },
  render () {
    const { flowComponent, node, business, move, moves, processStatus, $route } = this

    return (
      <div>
        <div style="padding: 10px 20px;">
          <a-steps current={ moves.length - 1} status={ processStatus }>
            {
              moves.map(item => {
                const slots = {
                  description: () => {
                    const operator = item.operator.nickName ? <div>操作人: { item.operator.nickName }</div> : null
                    return (
                      <div>
                        <div>节点说明: { item.workflowNode.description }</div>
                        { operator }
                        <div>操作参数: { item.param || '无参数' }</div>
                      </div>
                    )
                  }
                }
                return (
                  <a-step
                    title={ item.workflowNode.label }
                    v-slots={ slots }>
                  </a-step>
                )
              })
            }
          </a-steps>
        </div>
        <flowComponent
          node={ node }
          business={ business }
          move={ move }
          workflowMoveID={ parseInt($route.query.workflowMoveID as string) }></flowComponent>
      </div>
    )
  }
})

</script>
