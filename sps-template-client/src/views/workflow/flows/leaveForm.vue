<template>
  <div>
    <a-form :model="formData">
      <a-form-item label="开始时间">
        <a-date-picker
          v-model:value="formData.startTime"
          allowClear></a-date-picker>
      </a-form-item>
      <a-form-item label="结束时间">
        <a-date-picker
          v-model:value="formData.endTime"
          allowClear></a-date-picker>
      </a-form-item>
      <a-form-item label="请假原因">
        <a-textarea
          v-model:value="formData.cause"
          allowClear
          placeholder="请输入">
        </a-textarea>
      </a-form-item>
      <a-form-item>
        <template v-if="node.clazz === 'start'">
          <a-button type="primary" @click="start"><CaretRightOutlined />启动</a-button>
        </template>
        <template v-else-if="node.clazz === 'userTask'">
          <a-button type="primary" @click="complete('yes')"><CheckOutlined />同意</a-button>
          <a-button class="form-button" type="danger" @click="complete('no')"><CloseOutlined />拒绝</a-button>
        </template>
        <a-button class="form-button" @click="back"><RollbackOutlined />返回</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script>
import { computed, getCurrentInstance, reactive, toRefs } from 'vue'
import { useRouter } from 'vue-router'
import { completeWorkflowMove, startWorkflow } from '@/api/workflowProcess'
import { useStore } from 'vuex'
import moment from 'moment'

export default {
  name: 'leaveForm',
  props: {
    business: {
      type: Object,
      default: () => (null)
    },
    node: {
      type: Object,
      default: () => ({})
    },
    workflowMoveID: {
      type: Number,
      default: 0
    }
  },
  setup (props) {
    const { ctx } = getCurrentInstance()
    const router = useRouter()
    const store = useStore()

    const state = reactive({
      formData: props.business ? {
        ...props.business,
        startTime: moment(props.business.startTime),
        endTime: moment(props.business.endTime)
      } : {}
    })

    const userInfo = computed(() => store.getters['user/userInfo'])

    const back = () => {
      router.go(-1)
    }

    const start = async () => {
      const res = await startWorkflow({
        business: state.formData,
        wf: {
          workflowMoveID: props.workflowMoveID,
          businessId: 0,
          businessType: 'leave',
          workflowProcessID: props.node.workflowProcessID,
          workflowNodeID: props.node.id,
          promoterID: userInfo.value.ID,
          operatorID: userInfo.value.ID,
          action: 'create',
          param: ''
        }
      })
      if (res.code === 0) {
        ctx.$success('启动成功')
        back()
      }
    }

    const complete = async (param) => {
      const res = await completeWorkflowMove({
        business: state.formData,
        wf: {
          workflowMoveID: props.workflowMoveID,
          businessID: state.formData.ID,
          businessType: 'leave',
          workflowProcessID: props.node.workflowProcessID,
          workflowNodeID: props.node.id,
          promoterID: userInfo.value.ID,
          operatorID: userInfo.value.ID,
          action: 'complete',
          param: param
        }
      })
      if (res.code === 0) {
        ctx.$success('提交成功')
        back()
      }
    }

    return {
      ...toRefs(state),
      start,
      complete,
      back
    }
  }
}
</script>
