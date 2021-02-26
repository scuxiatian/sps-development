<template>
  <div class="detail-panel" :style="{ height: height + 'px' }">
    <ProcessDetail
      v-if="model.clazz === 'process'"
      :model="model"
      :onChange="onChange"
      :readOnly="readOnly"
    />
    <StartEventDetail
      v-else-if="model.clazz === 'start'"
      :model="model"
      :onChange="onChange"
      :readOnly="readOnly"
    />
    <EndEventDetail
      v-else-if="model.clazz === 'end'"
      :model="model"
      :onChange="onChange"
      :readOnly="readOnly"
    />
    <UserTaskDetail
      v-else-if="model.clazz === 'userTask'"
      :model="model"
      :onChange="onChange"
      :readOnly="readOnly"
      :users="users"
      :authorities="authorities"
      :groups="groups"
    />
    <GatewayDetail
      v-else-if="
        model.clazz === 'gateway' ||
        model.clazz === 'exclusiveGateway' ||
        model.clazz === 'parallelGateway' ||
        model.clazz === 'inclusiveGateway'
      "
      :model="model"
      :onChange="onChange"
      :readOnly="readOnly"
    />
    <FlowDetail
      v-else-if="model.clazz === 'flow'"
      :model="model"
      :onChange="onChange"
      :readOnly="readOnly"
    />
  </div>
</template>

<script>
import ProcessDetail from './ProcessDetail'
import StartEventDetail from './StartEventDetail'
import EndEventDetail from './EndEventDetail'
import UserTaskDetail from './UserTaskDetail'
import GatewayDetail from './GatewayDetail'
import FlowDetail from './FlowDetail'

export default {
  name: 'DetailPanel',
  components: {
    ProcessDetail,
    StartEventDetail,
    EndEventDetail,
    UserTaskDetail,
    GatewayDetail,
    FlowDetail
  },
  props: {
    height: {
      type: Number,
      default: 800
    },
    model: {
      type: Object,
      default: () => ({})
    },
    onChange: {
      type: Function,
      default: () => {}
    },
    readOnly: {
      type: Boolean,
      default: false
    },
    users: {
      type: Array,
      default: () => ([])
    },
    authorities: {
      type: Array,
      default: () => ([])
    },
    groups: {
      type: Array,
      default: () => ([])
    }
  }
}
</script>

<style lang="less">
.detail-panel {
  height: 100%;
  background: #f0f2f5;
  flex: 0 0 auto;
  float: left;
  width: 20%;
  border-right: 1px solid #e9e9e9;
  border-bottom: 1px solid #e9e9e9;

  .panel-title {
    text-align: left;
    height: 32px;
    padding-left: 12px;
    color: #000;
    line-height: 28px;
    background: #ebeef2;
    border-bottom: 1px solid #dce3e8;
  }

  .panel-body {
    .panel-row {
      text-align: left;
      display: inline-block;
      font-size: 12px;
      width: 100%;
      padding: 5px 12px;
    }
  }
}
</style>
